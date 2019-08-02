Feature: Acessing GraphQL DB
    In order to access products values
    As db module
    It should connect through GQL

    Scenario: List of Products
        Given I have connected to https://hash-ddspog-apply-dbhasura.herokuapp.com/v1/graphql engine
        When fetching all products
        Then these products must return
            """
                [{  "id": "1a184013-d405-4da1-b956-4781e5e4d256", "price_in_cents": 100,
                    "title": "Simple Pencil A2", "description": "The most simple pencil that can be used by anyone at anytime!"
                }, {
                    "id": "84b21225-9be4-4d9c-9bf0-6e5b321e2ca4", "price_in_cents": 2660,
                    "title": "Spiral Notebook Neon (TILIBRA) 96 pgs No-Lines", "description": "Simple and colorful notebook, for drawing to making various notes. Comes with an elastic to mark where you at."
                }, {
                    "id": "71bbc322-ffef-47af-8d87-c4bc596900af", "price_in_cents": 3000,
                    "title": "University Notebook with Cotton 100 pgs", "description": "Simple and delicate Notebook with lines to use on class and further purposes."
                }, {
                    "id": "1cd9ab74-56ba-444e-8666-c00eb9597e69", "price_in_cents": 2390,
                    "title": "University Notebook 300 pgs (15 subj.)", "description": "Simple Notebook with subject division for any student requiring more organization." }]
            """