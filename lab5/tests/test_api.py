#!/usr/bin/env python3

import pytest
import json
import struct

from conftest import Service


def test_calculates_integral(service):
    _, response = service.make_request({
        "polynom": {
            "a": 0.0,
            "b": 0.0,
            "c": 1.0,
        },
        "range": {
            "start": 0.0,
            "end": 1.0,
        }
    })

    assert json.loads(response)["sum"] == pytest.approx(1.0, 1e-6)

def test_returns_error_on_wrong_format(service):
    code, response = service.make_request([])

    assert code == 8
    assert struct.unpack("!H", response[0:2])[0] == 1007

def test_calculates_two_integrals(service):
    payload = {
        "polynom": {
            "a": 0.0,
            "b": 0.0,
            "c": 1.0,
        },
        "range": {
            "start": 0.0,
            "end": 1.0,
        }
    }
    _, response = service.make_request(payload)

    payload["range"]["end"] = 2.0
    _, response = service.make_request(payload)

    assert json.loads(response)["sum"] == pytest.approx(2.0, 1e-6)

