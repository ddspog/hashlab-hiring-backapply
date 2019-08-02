"""Util Module

The Util Module contains useful functions, needed on the other modules
logic.

This module contains the following functions:

    * envor - returns an environment var value, or default.
    * infoDummy - a dummy log info method.
"""

import os


def envor(envkey: str, default: str) -> str:
    """Returns an environment var value, or default.

    Parameters:
        envkey (str): An environment variable name.
        default (str): A default value to return, in case the environment
    variable ain't defined or have a value.

    Returns:
        str: If defined, the environment var value, else default value.
    """
    if envkey in os.environ and os.environ[envkey] != '':
        return os.environ[envkey]
    else:
        return default

def infoDummy(s, *args):
    """A dummy log info method."""
    pass