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
            "hello" :"world"
            }
            """

    Scenario: Query with variables
        When I send query:
            """
            query ($blah: String!) {
            foo(blah: $blah)
            }
            """
        And I have variables:
            """
            {"blah":"xxx"}
            """
        Then the response should be:
            """
            {
            "foo" :"this is blah: xxx"
            }
            """