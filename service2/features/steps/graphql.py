from behave import given, when, then
from db import DBClient
from ast import literal_eval
import re


@given('I have connected to {url} engine')
def step_impl(context, url):
    context.client = DBClient(url, None)

@when('fetching all products')
def step_impl(context):
    context.all_products = context.client.fetch_products()

@then('these products must return')
def step_impl(context):
    for ra, rb in zip(literal_eval(re.sub(r'^[\s]*', '', context.text)), context.all_products):
        assert ra == rb