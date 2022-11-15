#!/usr/bin/env python3

import pytest


pytest_plugins = ["docker_compose"]

@pytest.fixture(scope="function")
def wait_for_api(function_scoped_container_getter):
    service = function_scoped_container_getter.get("my_api_service").network_info[0]
    return service
