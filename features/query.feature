Feature: It should be possible to run select query

    Scenario: Basic select
        When I send query:
            """
            query {
            hello
            }
            """
        Then the response should be:
            """
            {
            "hello" :"world"
            }
            """