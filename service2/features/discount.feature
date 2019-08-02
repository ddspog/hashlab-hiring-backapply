Feature: Using Discount service
    In order to access products discounts
    As discount module
    It should connect through the discount service

    Scenario: Discount for one Product
        Given I have connected to localhost:5001 discount server
        When fetching a discount for user '543.004.756-11'
        Then for '84b21225-9be4-4d9c-9bf0-6e5b321e2ca4' product it will receive
            """
                { "pct": 0, "value_in_cents": 0 }
            """