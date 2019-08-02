"""Service 2

The script initializes a Flask REST server with only one route: product. 
This route returns all products listed at database

This module contains the following functions and classes:

    * DiscountClient - client to access gRPC Discount service.
"""

from cmd import service_cmd
from db import DBClient
from grpc_client import DiscountClient
from flask import Flask, request


app = Flask(__name__)

@app.route('/product')
def list_products():
    product_list = app.db.fetch_products()

    if product_list != None:
        for i in range(len(product_list)):
            try:
                d = app.discount.verify(
                    request.headers['X-USER-ID'],
                    product_list[i]['id']
                )
                product_list[i]['discount'] = d
            except Exception as err:
                app.logger.exception(f'An exception ocurred: {inst}')
                pass

        return {
            "products": product_list,
        }
    else:
        return {
            "data": product_list
        }

if __name__ == '__main__':
    args = service_cmd().parse_args()

    app.db = DBClient(args.graphql_server, app.logger)
    app.discount = DiscountClient(args.discount_server, app.logger)
    app.run(host=args.host, port=args.port, debug=args.debug)
    