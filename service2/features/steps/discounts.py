from behave import given, when, then
from grpc_client import DiscountClient
from ast import literal_eval
import re

@given(u'I have connected to {url} discount server')
def step_impl(context, url):
    context.client = DiscountClient(url, None)

@when(u'fetching a discount for user \'{userid}\'')
def step_impl(context, userid):
    context.actual_user = userid


@then(u'for \'{prodid}\' product it will receive')
def step_impl(context, prodid):
    discount = context.client.verify(context.actual_user, prodid)
    expected = literal_eval(re.sub(r'^[\s]*', '', context.text))
    assert discount == expected