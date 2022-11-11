#!/usr/bin/env python3

import pytest

import json
import socket
import pathlib

import pytest_services


@pytest.fixture(scope='session')
def root_dir():
    """Path to root directory service."""
    return pathlib.Path(__file__).parent.parent

@pytest.fixture
def service(watcher_getter, request, root_dir):
    return Service(watcher_getter, request, root_dir)

class Service:
    def __init__(self, watcher_getter, request, root_dir):
        self.sock_ = socket.socket()
        self.service_ = watcher_getter(
            name=str(root_dir)+'/build_debug/beer-paint',
            checker = lambda: True,
            request=request,
        )
        retry = 0
        while self.sock_.connect_ex(('localhost', 2014)) != 0 or retry < 10:
            retry += 1
        if retry == 10:
            raise "Can't connect"
        self.sock_file_ = self.sock_.makefile()
    def make_request(self, payload):
        self.sock_.sendall(json.dumps(payload).encode('utf-8'))
    def get_response(self):
        line = self.sock_file_.readline()
        return json.loads(line)
