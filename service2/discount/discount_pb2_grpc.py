# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from discount import discount_pb2 as discount_dot_discount__pb2


class DiscountStub(object):
  """Discount will contain ...
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.VerifyDiscount = channel.unary_unary(
        '/protocols.Discount/VerifyDiscount',
        request_serializer=discount_dot_discount__pb2.VerifyDiscountForm.SerializeToString,
        response_deserializer=discount_dot_discount__pb2.DiscountInfo.FromString,
        )


class DiscountServicer(object):
  """Discount will contain ...
  """

  def VerifyDiscount(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_DiscountServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'VerifyDiscount': grpc.unary_unary_rpc_method_handler(
          servicer.VerifyDiscount,
          request_deserializer=discount_dot_discount__pb2.VerifyDiscountForm.FromString,
          response_serializer=discount_dot_discount__pb2.DiscountInfo.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'protocols.Discount', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
