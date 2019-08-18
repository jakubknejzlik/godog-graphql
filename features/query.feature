Feature: It should be possible to send query

    Scenario: Simple query
        When I send query:
            """
            query {
            hello
            }
            """
        Then the response should be:
            """
            {
                "hello": "world"
            }
            """
    Scenario: Simple query with error
        When I send query:
            """
            query {
            blah
            }
            """
        Then the error should be:
            """
            graphql: Cannot query field "blah" on type "Query".
            """

    Scenario: Query with variables
        When I have variables:
            """
            {
                "blah": "xxx"
            }
            """
        And I send query:
            """
            query ($blah: String!) {
            foo(blah: $blah)
            }
            """
        Then the response should be:
            """
            {
                "foo": "this is blah: xxx"
            }
            """
