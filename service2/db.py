"""DB Module

The DB Module creates a client to access GraphQL database engines,
running simple queries.

This module contains the following functions and classes:

    * DBClient - client to access GraphQL database engines.
"""

from graphqlclient import GraphQLClient
from ast import literal_eval
from util import infoDummy


class DBClient:
    """Client to access GraphQL database engines.

    Attributes:
        url (str): Address to access the database engine.
        client (graphqlclient.GraphQLClient): A GraphQL client to access
    engines connected to databases.
        info (function): A logging function to print info about execution.

    Methods:
        prepare_client(): If not defined, it prepares the GraphQL client.
        fetch(query): Executes query and return JSON as dict object.
        fetch_products(): Fetch all products on database.
    """
    
    def __init__(self, url, log):
        """
        Parameters:
            url (str): Address to access the GraphQL engine.
            log (function): A logger to obtain the info function.
        """

        self.url = url
        self.client = None
        if log != None:
            self.info = log.info
        else:
            self.info = infoDummy

    def prepare_client(self):
        """If not defined, it prepares the GraphQL client."""

        if self.client == None:
            self.client = GraphQLClient(self.url)
            self.info("Connected to %s", self.url)

    def fetch(self, query):
        """Executes a query and return JSON result.

        Parameters:
            query (str): The GraphQL query to execute.

        Returns:
            dict: A JSON value as a dict object.
        """

        result = self.client.execute(query)
        return literal_eval(result)

    def fetch_products(self):
        """Fetch all products on database.

        Returns:
            list: A list of all products, each in dict objects.
        """

        self.prepare_client()
        json = self.fetch("""{
            Product {
                id
                price_in_cents
                title
                description
            }
        }""")

        if json != None and json['data'] != None:
            return json['data']['Product']
        else:
            return json
