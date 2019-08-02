"""Discount Module

The Discount Module aggregates all protocol objects from discount.proto
defined at this project.
"""

from .discount_pb2_grpc import DiscountStub
from .discount_pb2 import DiscountInfo
from .discount_pb2 import VerifyDiscountForm