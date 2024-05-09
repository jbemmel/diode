#!/usr/bin/env python
# Copyright 2024 NetBox Labs Inc
"""NetBox Labs, Diode - SDK - Client."""

import logging
import os
import platform
import uuid
from typing import Iterable, Optional
from urllib.parse import urlparse

import certifi
import grpc

from netboxlabs.diode.sdk.diode.v1 import ingester_pb2, ingester_pb2_grpc
from netboxlabs.diode.sdk.exceptions import DiodeClientError, DiodeConfigError

_DIODE_API_KEY_ENVVAR_NAME = "DIODE_API_KEY"
_DIODE_API_TLS_VERIFY_ENVVAR_NAME = "DIODE_API_TLS_VERIFY"
_DIODE_SDK_LOG_LEVEL_ENVVAR_NAME = "DIODE_SDK_LOG_LEVEL"
_DEFAULT_STREAM = "latest"
_LOGGER = logging.getLogger(__name__)


def _certs() -> bytes:
    with open(certifi.where(), "rb") as f:
        return f.read()


def _api_key(api_key: Optional[str] = None) -> str:
    if api_key is None:
        api_key = os.getenv(_DIODE_API_KEY_ENVVAR_NAME)
    if api_key is None:
        raise DiodeConfigError(
            f"api_key param or {_DIODE_API_KEY_ENVVAR_NAME} environment variable required"
        )
    return api_key


def _tls_verify(tls_verify: Optional[bool]) -> bool:
    if tls_verify is None:
        tls_verify_env_var = os.getenv(_DIODE_API_TLS_VERIFY_ENVVAR_NAME, "false")
        return tls_verify_env_var.lower() in ["true", "1", "yes"]
    if not isinstance(tls_verify, bool):
        raise DiodeConfigError("tls_verify must be a boolean")

    return tls_verify


class DiodeClient:
    """Diode Client."""

    _name = "diode-sdk-python"
    _version = "0.0.1"
    _app_name = None
    _app_version = None
    _channel = None
    _stub = None

    def __init__(
        self,
        target: str,
        app_name: str,
        app_version: str,
        api_key: Optional[str] = None,
        tls_verify: bool = None,
    ):
        """Initiate a new client."""
        log_level = os.getenv(_DIODE_SDK_LOG_LEVEL_ENVVAR_NAME, "INFO").upper()
        logging.basicConfig(level=log_level)

        # TODO: validate target
        self._target = target

        self._app_name = app_name
        self._app_version = app_version

        api_key = _api_key(api_key)
        self._metadata = (
            ("diode-api-key", api_key),
            ("platform", platform.platform()),
            ("python-version", platform.python_version()),
        )

        if _tls_verify(tls_verify):
            self._channel = grpc.secure_channel(
                self._target,
                grpc.ssl_channel_credentials(
                    root_certificates=_certs(),
                ),
            )
        else:
            self._channel = grpc.insecure_channel(self._target)

        self._stub = ingester_pb2_grpc.IngesterServiceStub(self._channel)

    @property
    def name(self) -> str:
        """Retrieve the name."""
        return self._name

    @property
    def version(self) -> str:
        """Retrieve the version."""
        return self._version

    @property
    def target(self) -> str:
        """Retrieve the target."""
        return self._target

    @property
    def app_name(self) -> str:
        """Retrieve the app name."""
        return self._app_name

    @property
    def app_version(self) -> str:
        """Retrieve the app version."""
        return self._app_version

    @property
    def channel(self) -> grpc.Channel:
        """Retrieve the channel."""
        return self._channel

    def __enter__(self):
        """Enters the runtime context related to the channel object."""
        return self

    def __exit__(self, exc_type, exc_value, exc_traceback):
        """Exits the runtime context related to the channel object."""
        self.close()

    def close(self):
        """Close the channel."""
        self._channel.close()

    def ingest(
        self,
        entities: Iterable[ingester_pb2.Entity],
        stream: Optional[str] = _DEFAULT_STREAM,
    ) -> ingester_pb2.IngestResponse:
        """Ingest entities."""
        try:
            request = ingester_pb2.IngestRequest(
                stream=stream,
                id=str(uuid.uuid4()),
                entities=entities,
                sdk_name=self.name,
                sdk_version=self.version,
                producer_app_name=self.app_name,
                producer_app_version=self.app_version,
            )

            return self._stub.Ingest(request, metadata=self._metadata)
        except grpc.RpcError as err:
            raise DiodeClientError(err) from err
