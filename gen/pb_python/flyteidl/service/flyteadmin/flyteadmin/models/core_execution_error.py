# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.execution_error_error_kind import ExecutionErrorErrorKind  # noqa: F401,E501


class CoreExecutionError(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'code': 'str',
        'message': 'str',
        'error_uri': 'str',
        'kind': 'ExecutionErrorErrorKind'
    }

    attribute_map = {
        'code': 'code',
        'message': 'message',
        'error_uri': 'error_uri',
        'kind': 'kind'
    }

    def __init__(self, code=None, message=None, error_uri=None, kind=None):  # noqa: E501
        """CoreExecutionError - a model defined in Swagger"""  # noqa: E501

        self._code = None
        self._message = None
        self._error_uri = None
        self._kind = None
        self.discriminator = None

        if code is not None:
            self.code = code
        if message is not None:
            self.message = message
        if error_uri is not None:
            self.error_uri = error_uri
        if kind is not None:
            self.kind = kind

    @property
    def code(self):
        """Gets the code of this CoreExecutionError.  # noqa: E501


        :return: The code of this CoreExecutionError.  # noqa: E501
        :rtype: str
        """
        return self._code

    @code.setter
    def code(self, code):
        """Sets the code of this CoreExecutionError.


        :param code: The code of this CoreExecutionError.  # noqa: E501
        :type: str
        """

        self._code = code

    @property
    def message(self):
        """Gets the message of this CoreExecutionError.  # noqa: E501

        Detailed description of the error - including stack trace.  # noqa: E501

        :return: The message of this CoreExecutionError.  # noqa: E501
        :rtype: str
        """
        return self._message

    @message.setter
    def message(self, message):
        """Sets the message of this CoreExecutionError.

        Detailed description of the error - including stack trace.  # noqa: E501

        :param message: The message of this CoreExecutionError.  # noqa: E501
        :type: str
        """

        self._message = message

    @property
    def error_uri(self):
        """Gets the error_uri of this CoreExecutionError.  # noqa: E501


        :return: The error_uri of this CoreExecutionError.  # noqa: E501
        :rtype: str
        """
        return self._error_uri

    @error_uri.setter
    def error_uri(self, error_uri):
        """Sets the error_uri of this CoreExecutionError.


        :param error_uri: The error_uri of this CoreExecutionError.  # noqa: E501
        :type: str
        """

        self._error_uri = error_uri

    @property
    def kind(self):
        """Gets the kind of this CoreExecutionError.  # noqa: E501


        :return: The kind of this CoreExecutionError.  # noqa: E501
        :rtype: ExecutionErrorErrorKind
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this CoreExecutionError.


        :param kind: The kind of this CoreExecutionError.  # noqa: E501
        :type: ExecutionErrorErrorKind
        """

        self._kind = kind

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreExecutionError, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreExecutionError):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
