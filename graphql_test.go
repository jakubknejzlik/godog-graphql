package godogsql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	// "github.com/machinebox/graphql"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/go-test/deep"
	"github.com/jakubknejzlik/godog-graphql/graphql"
	"github.com/sergi/go-diff/diffmatchpatch"
)

type gqlFeature struct {
	client    *graphql.Client
	query     string
	variables map[string]interface{}
	response  interface{}
}

func (f *gqlFeature) iSendQuery(arg1 *gherkin.DocString) error {
	f.query = arg1.Content

	ctx := context.Background()
	return f.client.SendQuery(ctx, f.query, f.variables, &f.response)
}

func (f *gqlFeature) iHaveVariables(arg1 *gherkin.DocString) error {
	return json.Unmarshal([]byte(arg1.Content), &f.variables)
}

func (f *gqlFeature) theResponseShouldBe(arg1 *gherkin.DocString) (err error) {
	var expected interface{}
	err = json.Unmarshal([]byte(arg1.Content), &expected)
	if err != nil {
		return
	}

	if diff := deep.Equal(expected, f.response); diff != nil {
		dmp := diffmatchpatch.New()
		text1, _ := json.MarshalIndent(expected, "", " ")
		text2, _ := json.MarshalIndent(f.response, "", " ")
		diffs := dmp.DiffMain(string(text1), string(text2), true)
		err = errors.New(dmp.DiffPrettyText(diffs))
	}
	return
}

func FeatureContext(s *godog.Suite) {
	URL := os.Getenv("GRAPHQL_URL")
	if URL == "" {
		panic(fmt.Errorf("Missing required environment variable GRAPHQL_URL"))
	}

	c, err := graphql.NewClient(URL)
	if err != nil {
		panic(err)
	}
	feature := &gqlFeature{client: c}

	s.Step(`^I send query:$`, feature.iSendQuery)
	s.Step(`^I have variables:$`, feature.iHaveVariables)
	s.Step(`^the response should be:$`, feature.theResponseShouldBe)
}
