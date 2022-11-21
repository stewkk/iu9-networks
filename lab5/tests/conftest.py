#!/usr/bin/env python3

import pytest

from websocket import create_connection
import json


pytest_plugins = ["docker_compose"]

@pytest.fixture(scope="function")
def wait_for_api(function_scoped_container_getter):
    service = function_scoped_container_getter.get("app").network_info[0]

    return service

@pytest.fixture
def service(wait_for_api):
    return Service(wait_for_api)


class Service:
    def __init__(self, service):
        self.ws = create_connection("ws://localhost:%s" % service.host_port)
    def make_request(self, payload):
        self.ws.send(json.dumps(payload))
        return self.ws.recv_data()
