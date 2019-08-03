"""CMD Module

The CMD Module contains functions needed to process the service as a cli
application. Its functions defined various needed properties.

This module contains the following functions:

    * service_cmd - returns the parser for the service args.
"""

import argparse
from util import envor


def service_cmd():
    """Returns the parser for the service args.

    Returns:
        argparse.ArgumentParser: A parser for the service args.
    """

    parser = argparse.ArgumentParser(
        prog='service2', 
        description='Service responsible of Service 002 micro service'
    )

    parser.add_argument('--verbose', '-v', action='store_true', default=False,
                        help='if set, will print more detailed information')
    parser.add_argument('--debug', '-d', action='store_true', default=False,
                        help='if set, will print debug information')
    parser.add_argument('--port', '-p', metavar='PORT', default=envor('SERVICE2_PORT', '5002'),
                        help='port to load the server')
    parser.add_argument('--host', metavar='HOST', default=envor('SERVICE2_HOST', '0.0.0.0'),
                        help='host to load the server')
    parser.add_argument('--graphql_server', '-g', metavar='SRV', default=envor('SERVICE2_GRAPHQLSRV', 'localhost:8080/v1/graphql'),
                        help='GraphQL engine server to load data')
    parser.add_argument('--discount_server', metavar='SRV', default=envor('SERVICE2_DISCOUNTSRV', 'localhost:5001'),
                        help='Discount gRPC server to fetch data')
    
    return parser