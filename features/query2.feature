Feature: Query in next steps

    Scenario: Query without specified headers
        When I send query:
            """
            query {
            fooHeader: header(name: "foo")
            }
            """
        Then the response should be:
            """
            {
                "fooHeader": ""
            }
            """
        And the error should be empty
