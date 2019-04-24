Feature: It should be possible to run select query

    Scenario: Basic select
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