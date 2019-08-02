from graphqlclient import GraphQLClient
from ast import literal_eval


def infoDummy(s, *args):
    pass

class DBClient:
    def __init__(self, url, log):
        self.url = url
        self.client = None
        if log != None:
            self.info = log.info
        else:
            self.info = infoDummy

    def prepare_client(self):
        if self.client == None:
            self.client = GraphQLClient(self.url)
            self.info("Connected to %s", self.url)

    def fetch(self, query):
        result = self.client.execute(query)
        return literal_eval(result)

    def fetch_products(self):
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
