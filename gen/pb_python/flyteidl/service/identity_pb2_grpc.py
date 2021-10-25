# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from flyteidl.service import identity_pb2 as flyteidl_dot_service_dot_identity__pb2


class IdentityServiceStub(object):
  """IdentityService defines an RPC Service that interacts with user/app identities.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.UserInfo = channel.unary_unary(
        '/flyteidl.service.IdentityService/UserInfo',
        request_serializer=flyteidl_dot_service_dot_identity__pb2.UserInfoRequest.SerializeToString,
        response_deserializer=flyteidl_dot_service_dot_identity__pb2.UserInfoResponse.FromString,
        )


class IdentityServiceServicer(object):
  """IdentityService defines an RPC Service that interacts with user/app identities.
  """

  def UserInfo(self, request, context):
    """Retrieves user information about the currently logged in user.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_IdentityServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'UserInfo': grpc.unary_unary_rpc_method_handler(
          servicer.UserInfo,
          request_deserializer=flyteidl_dot_service_dot_identity__pb2.UserInfoRequest.FromString,
          response_serializer=flyteidl_dot_service_dot_identity__pb2.UserInfoResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'flyteidl.service.IdentityService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
