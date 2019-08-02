"""gRPC Client Module

The gRPC Client Module creates a client to access gRPC services, calling
its methods.

This module contains the following functions and classes:

    * DiscountClient - client to access gRPC Discount service.
"""

import grpc
from discount import DiscountStub, VerifyDiscountForm
from util import infoDummy


class DiscountClient:
    """Client to access gRPC Discount service.

    Attributes:
        url (str): Address to access the gRPC service.
        client (discount.DiscountStub): A gRPC client to access discount
    services.
        info (function): A logging function to print info about execution.

    Methods:
        prepare_client(): If not defined, it prepares the gRPC client.
        verify(query): Executes query and return JSON as dict object.
    """

    def __init__(self, url, log):
        """
        Parameters:
            url (str): Address to access the gRPC client.
            log (function): A logger to obtain the info function.
        """

        self.url = url
        self.client = None
        if log != None:
            self.info = log.info
        else:
            self.info = infoDummy

    def prepare_client(self):
        """If not defined, it prepares the gRPC client."""

        if self.client == None:
            channel = grpc.insecure_channel(self.url)
            self.client = DiscountStub(channel)
            self.info("Connected to %s", self.url)

    def verify(self, uid, pid):
        """Verify and return discount of a product for the user.

        Parameters:
            uid (str): User id interested at discount.
            pid (str): Product id to verify discount.

        Returns:
            dict: A dict value containing pct and value_in_cents.
        """

        self.prepare_client()

        form = VerifyDiscountForm(userid=uid, prodid=pid)
        info = self.client.VerifyDiscount(form)

        return {
            "pct": info.pct,
            "value_in_cents": info.value_in_cents
        }