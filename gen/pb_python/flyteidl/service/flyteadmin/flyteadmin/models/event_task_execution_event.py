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

from flyteadmin.models.core_execution_error import CoreExecutionError  # noqa: F401,E501
from flyteadmin.models.core_identifier import CoreIdentifier  # noqa: F401,E501
from flyteadmin.models.core_literal_map import CoreLiteralMap  # noqa: F401,E501
from flyteadmin.models.core_node_execution_identifier import CoreNodeExecutionIdentifier  # noqa: F401,E501
from flyteadmin.models.core_task_execution_phase import CoreTaskExecutionPhase  # noqa: F401,E501
from flyteadmin.models.core_task_log import CoreTaskLog  # noqa: F401,E501
from flyteadmin.models.event_task_execution_metadata import EventTaskExecutionMetadata  # noqa: F401,E501
from flyteadmin.models.protobuf_struct import ProtobufStruct  # noqa: F401,E501


class EventTaskExecutionEvent(object):
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
        'task_id': 'CoreIdentifier',
        'parent_node_execution_id': 'CoreNodeExecutionIdentifier',
        'retry_attempt': 'int',
        'phase': 'CoreTaskExecutionPhase',
        'producer_id': 'str',
        'logs': 'list[CoreTaskLog]',
        'occurred_at': 'datetime',
        'input_uri': 'str',
        'output_uri': 'str',
        'error': 'CoreExecutionError',
        'output_data': 'CoreLiteralMap',
        'custom_info': 'ProtobufStruct',
        'phase_version': 'int',
        'reason': 'str',
        'task_type': 'str',
        'metadata': 'EventTaskExecutionMetadata'
    }

    attribute_map = {
        'task_id': 'task_id',
        'parent_node_execution_id': 'parent_node_execution_id',
        'retry_attempt': 'retry_attempt',
        'phase': 'phase',
        'producer_id': 'producer_id',
        'logs': 'logs',
        'occurred_at': 'occurred_at',
        'input_uri': 'input_uri',
        'output_uri': 'output_uri',
        'error': 'error',
        'output_data': 'output_data',
        'custom_info': 'custom_info',
        'phase_version': 'phase_version',
        'reason': 'reason',
        'task_type': 'task_type',
        'metadata': 'metadata'
    }

    def __init__(self, task_id=None, parent_node_execution_id=None, retry_attempt=None, phase=None, producer_id=None, logs=None, occurred_at=None, input_uri=None, output_uri=None, error=None, output_data=None, custom_info=None, phase_version=None, reason=None, task_type=None, metadata=None):  # noqa: E501
        """EventTaskExecutionEvent - a model defined in Swagger"""  # noqa: E501

        self._task_id = None
        self._parent_node_execution_id = None
        self._retry_attempt = None
        self._phase = None
        self._producer_id = None
        self._logs = None
        self._occurred_at = None
        self._input_uri = None
        self._output_uri = None
        self._error = None
        self._output_data = None
        self._custom_info = None
        self._phase_version = None
        self._reason = None
        self._task_type = None
        self._metadata = None
        self.discriminator = None

        if task_id is not None:
            self.task_id = task_id
        if parent_node_execution_id is not None:
            self.parent_node_execution_id = parent_node_execution_id
        if retry_attempt is not None:
            self.retry_attempt = retry_attempt
        if phase is not None:
            self.phase = phase
        if producer_id is not None:
            self.producer_id = producer_id
        if logs is not None:
            self.logs = logs
        if occurred_at is not None:
            self.occurred_at = occurred_at
        if input_uri is not None:
            self.input_uri = input_uri
        if output_uri is not None:
            self.output_uri = output_uri
        if error is not None:
            self.error = error
        if output_data is not None:
            self.output_data = output_data
        if custom_info is not None:
            self.custom_info = custom_info
        if phase_version is not None:
            self.phase_version = phase_version
        if reason is not None:
            self.reason = reason
        if task_type is not None:
            self.task_type = task_type
        if metadata is not None:
            self.metadata = metadata

    @property
    def task_id(self):
        """Gets the task_id of this EventTaskExecutionEvent.  # noqa: E501

        ID of the task. In combination with the retryAttempt this will indicate the task execution uniquely for a given parent node execution.  # noqa: E501

        :return: The task_id of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: CoreIdentifier
        """
        return self._task_id

    @task_id.setter
    def task_id(self, task_id):
        """Sets the task_id of this EventTaskExecutionEvent.

        ID of the task. In combination with the retryAttempt this will indicate the task execution uniquely for a given parent node execution.  # noqa: E501

        :param task_id: The task_id of this EventTaskExecutionEvent.  # noqa: E501
        :type: CoreIdentifier
        """

        self._task_id = task_id

    @property
    def parent_node_execution_id(self):
        """Gets the parent_node_execution_id of this EventTaskExecutionEvent.  # noqa: E501


        :return: The parent_node_execution_id of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: CoreNodeExecutionIdentifier
        """
        return self._parent_node_execution_id

    @parent_node_execution_id.setter
    def parent_node_execution_id(self, parent_node_execution_id):
        """Sets the parent_node_execution_id of this EventTaskExecutionEvent.


        :param parent_node_execution_id: The parent_node_execution_id of this EventTaskExecutionEvent.  # noqa: E501
        :type: CoreNodeExecutionIdentifier
        """

        self._parent_node_execution_id = parent_node_execution_id

    @property
    def retry_attempt(self):
        """Gets the retry_attempt of this EventTaskExecutionEvent.  # noqa: E501


        :return: The retry_attempt of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: int
        """
        return self._retry_attempt

    @retry_attempt.setter
    def retry_attempt(self, retry_attempt):
        """Sets the retry_attempt of this EventTaskExecutionEvent.


        :param retry_attempt: The retry_attempt of this EventTaskExecutionEvent.  # noqa: E501
        :type: int
        """

        self._retry_attempt = retry_attempt

    @property
    def phase(self):
        """Gets the phase of this EventTaskExecutionEvent.  # noqa: E501


        :return: The phase of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: CoreTaskExecutionPhase
        """
        return self._phase

    @phase.setter
    def phase(self, phase):
        """Sets the phase of this EventTaskExecutionEvent.


        :param phase: The phase of this EventTaskExecutionEvent.  # noqa: E501
        :type: CoreTaskExecutionPhase
        """

        self._phase = phase

    @property
    def producer_id(self):
        """Gets the producer_id of this EventTaskExecutionEvent.  # noqa: E501


        :return: The producer_id of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._producer_id

    @producer_id.setter
    def producer_id(self, producer_id):
        """Sets the producer_id of this EventTaskExecutionEvent.


        :param producer_id: The producer_id of this EventTaskExecutionEvent.  # noqa: E501
        :type: str
        """

        self._producer_id = producer_id

    @property
    def logs(self):
        """Gets the logs of this EventTaskExecutionEvent.  # noqa: E501


        :return: The logs of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: list[CoreTaskLog]
        """
        return self._logs

    @logs.setter
    def logs(self, logs):
        """Sets the logs of this EventTaskExecutionEvent.


        :param logs: The logs of this EventTaskExecutionEvent.  # noqa: E501
        :type: list[CoreTaskLog]
        """

        self._logs = logs

    @property
    def occurred_at(self):
        """Gets the occurred_at of this EventTaskExecutionEvent.  # noqa: E501

        This timestamp represents when the original event occurred, it is generated by the executor of the task.  # noqa: E501

        :return: The occurred_at of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: datetime
        """
        return self._occurred_at

    @occurred_at.setter
    def occurred_at(self, occurred_at):
        """Sets the occurred_at of this EventTaskExecutionEvent.

        This timestamp represents when the original event occurred, it is generated by the executor of the task.  # noqa: E501

        :param occurred_at: The occurred_at of this EventTaskExecutionEvent.  # noqa: E501
        :type: datetime
        """

        self._occurred_at = occurred_at

    @property
    def input_uri(self):
        """Gets the input_uri of this EventTaskExecutionEvent.  # noqa: E501

        URI of the input file, it encodes all the information including Cloud source provider. ie., s3://...  # noqa: E501

        :return: The input_uri of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._input_uri

    @input_uri.setter
    def input_uri(self, input_uri):
        """Sets the input_uri of this EventTaskExecutionEvent.

        URI of the input file, it encodes all the information including Cloud source provider. ie., s3://...  # noqa: E501

        :param input_uri: The input_uri of this EventTaskExecutionEvent.  # noqa: E501
        :type: str
        """

        self._input_uri = input_uri

    @property
    def output_uri(self):
        """Gets the output_uri of this EventTaskExecutionEvent.  # noqa: E501

        URI to the output of the execution, it will be in a format that encodes all the information including Cloud source provider. ie., s3://...  # noqa: E501

        :return: The output_uri of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._output_uri

    @output_uri.setter
    def output_uri(self, output_uri):
        """Sets the output_uri of this EventTaskExecutionEvent.

        URI to the output of the execution, it will be in a format that encodes all the information including Cloud source provider. ie., s3://...  # noqa: E501

        :param output_uri: The output_uri of this EventTaskExecutionEvent.  # noqa: E501
        :type: str
        """

        self._output_uri = output_uri

    @property
    def error(self):
        """Gets the error of this EventTaskExecutionEvent.  # noqa: E501


        :return: The error of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: CoreExecutionError
        """
        return self._error

    @error.setter
    def error(self, error):
        """Sets the error of this EventTaskExecutionEvent.


        :param error: The error of this EventTaskExecutionEvent.  # noqa: E501
        :type: CoreExecutionError
        """

        self._error = error

    @property
    def output_data(self):
        """Gets the output_data of this EventTaskExecutionEvent.  # noqa: E501

        Raw output data produced by this task execution.  # noqa: E501

        :return: The output_data of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: CoreLiteralMap
        """
        return self._output_data

    @output_data.setter
    def output_data(self, output_data):
        """Sets the output_data of this EventTaskExecutionEvent.

        Raw output data produced by this task execution.  # noqa: E501

        :param output_data: The output_data of this EventTaskExecutionEvent.  # noqa: E501
        :type: CoreLiteralMap
        """

        self._output_data = output_data

    @property
    def custom_info(self):
        """Gets the custom_info of this EventTaskExecutionEvent.  # noqa: E501

        Custom data that the task plugin sends back. This is extensible to allow various plugins in the system.  # noqa: E501

        :return: The custom_info of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: ProtobufStruct
        """
        return self._custom_info

    @custom_info.setter
    def custom_info(self, custom_info):
        """Sets the custom_info of this EventTaskExecutionEvent.

        Custom data that the task plugin sends back. This is extensible to allow various plugins in the system.  # noqa: E501

        :param custom_info: The custom_info of this EventTaskExecutionEvent.  # noqa: E501
        :type: ProtobufStruct
        """

        self._custom_info = custom_info

    @property
    def phase_version(self):
        """Gets the phase_version of this EventTaskExecutionEvent.  # noqa: E501

        Some phases, like RUNNING, can send multiple events with changed metadata (new logs, additional custom_info, etc) that should be recorded regardless of the lack of phase change. The version field should be incremented when metadata changes across the duration of an individual phase.  # noqa: E501

        :return: The phase_version of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: int
        """
        return self._phase_version

    @phase_version.setter
    def phase_version(self, phase_version):
        """Sets the phase_version of this EventTaskExecutionEvent.

        Some phases, like RUNNING, can send multiple events with changed metadata (new logs, additional custom_info, etc) that should be recorded regardless of the lack of phase change. The version field should be incremented when metadata changes across the duration of an individual phase.  # noqa: E501

        :param phase_version: The phase_version of this EventTaskExecutionEvent.  # noqa: E501
        :type: int
        """

        self._phase_version = phase_version

    @property
    def reason(self):
        """Gets the reason of this EventTaskExecutionEvent.  # noqa: E501

        An optional explanation for the phase transition.  # noqa: E501

        :return: The reason of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._reason

    @reason.setter
    def reason(self, reason):
        """Sets the reason of this EventTaskExecutionEvent.

        An optional explanation for the phase transition.  # noqa: E501

        :param reason: The reason of this EventTaskExecutionEvent.  # noqa: E501
        :type: str
        """

        self._reason = reason

    @property
    def task_type(self):
        """Gets the task_type of this EventTaskExecutionEvent.  # noqa: E501

        A predefined yet extensible Task type identifier. If the task definition is already registered in flyte admin this type will be identical, but not all task executions necessarily use pre-registered definitions and this type is useful to render the task in the UI, filter task executions, etc.  # noqa: E501

        :return: The task_type of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: str
        """
        return self._task_type

    @task_type.setter
    def task_type(self, task_type):
        """Sets the task_type of this EventTaskExecutionEvent.

        A predefined yet extensible Task type identifier. If the task definition is already registered in flyte admin this type will be identical, but not all task executions necessarily use pre-registered definitions and this type is useful to render the task in the UI, filter task executions, etc.  # noqa: E501

        :param task_type: The task_type of this EventTaskExecutionEvent.  # noqa: E501
        :type: str
        """

        self._task_type = task_type

    @property
    def metadata(self):
        """Gets the metadata of this EventTaskExecutionEvent.  # noqa: E501

        Metadata around how a task was executed.  # noqa: E501

        :return: The metadata of this EventTaskExecutionEvent.  # noqa: E501
        :rtype: EventTaskExecutionMetadata
        """
        return self._metadata

    @metadata.setter
    def metadata(self, metadata):
        """Sets the metadata of this EventTaskExecutionEvent.

        Metadata around how a task was executed.  # noqa: E501

        :param metadata: The metadata of this EventTaskExecutionEvent.  # noqa: E501
        :type: EventTaskExecutionMetadata
        """

        self._metadata = metadata

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
        if issubclass(EventTaskExecutionEvent, dict):
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
        if not isinstance(other, EventTaskExecutionEvent):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
