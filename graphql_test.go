package godogsql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	// "github.com/machinebox/graphql"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/go-test/deep"
	"github.com/jakubknejzlik/godog-graphql/graphql"
)

type gqlFeature struct {
	client    *graphql.Client
	query     string
	variables map[string]interface{}
	response  interface{}
}

func (f *gqlFeature) iSendQuery(arg1 *gherkin.DocString) error {
	f.query = arg1.Content
	return nil
}

func (f *gqlFeature) iHaveVariables(arg1 *gherkin.DocString) error {
	return json.Unmarshal([]byte(arg1.Content), &f.variables)
}

func (f *gqlFeature) theResponseShouldBe(arg1 *gherkin.DocString) (err error) {
	ctx := context.Background()
	var res interface{}
	var expected interface{}
	err = json.Unmarshal([]byte(arg1.Content), &expected)
	if err != nil {
		return
	}
	err = f.client.SendQuery(ctx, f.query, f.variables, &res)
	if err != nil {
		return
	}
	if diff := deep.Equal(expected, res); diff != nil {
		err = errors.New(strings.Join(diff, "\n"))
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
