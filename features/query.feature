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
        And the error should not be empty

    Scenario: Query with variables
        When I have variables:
            """
            {
                "blah": "xxx"
            }
            """
        And I have headers:
            """
            {
                "foo":"blah header"
            }
            """
        And I send query:
            """
            query ($blah: String!) {
            foo(blah: $blah)
            foo2: foo(blah: $blah)
            fooHeader: header(name: "foo")
            }
            """
        Then the response should be:
            """
            {
                "foo2": "this is blah: xxx",
                "foo": "this is blah: xxx",
                "fooHeader": "blah header"
            }
            """
        And the error should be empty
