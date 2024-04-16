import time

from behave import given, when, then

from netboxlabs.diode.sdk import DiodeClient
from netboxlabs.diode.sdk.diode.v1.ingester_pb2 import Entity
from netboxlabs.diode.sdk.diode.v1.site_pb2 import Site

from steps.config import configs

from steps.utils import get_site, send_delete_request


api_key = str(configs["api_key"])


@given('a new site "{site_name}" object')
def step_create_new_site_object(context, site_name):
    """Set the body of the request to create a new site."""
    context.site_name = site_name


@when("the site object is ingested")
def ingest_site_object(context):
    """Ingest the site object using the Diode SDK"""
    with DiodeClient(
        target="localhost:8081",
        app_name="my-test-app",
        app_version="0.0.1",
        api_key=api_key,
    ) as client:
        entities = [
            Entity(site=Site(name=context.site_name)),
        ]

        context.response = client.ingest(entities=entities)
        return context.response


@then("the site object is created in the database")
@then("do nothing")
def check_site_object(context):
    """Check if the response status code is 200 and the result is success"""
    assert context.response is not None
    # Wait for the site object to be added to the cache
    time.sleep(3)
    site = get_site(context.site_name)
    assert site.get("name") == context.site_name
    assert site.get("status").get("value") == "active"


@given("site object already exists in the database")
def retrieve_existing_site_object(context):
    """Retrieve the site object from the database"""
    context.site_name = "Site A"
    context.site = get_site(context.site_name)
    context.site_name = context.site.get("name")


@given('site object with status "{status}" and description "{description}"')
def create_site_object_to_update(context, status, description):
    """Create a site object with a status and description to update"""
    context.site_name = "Site A"
    context.status = status
    context.description = description


@when("the site object is ingested with the updates")
def update_site_object(context):
    """Update the site object using the Diode SDK"""
    with DiodeClient(
        target="localhost:8081",
        app_name="my-test-app",
        app_version="0.0.1",
        api_key=api_key,
    ) as client:
        entities = [
            Entity(
                site=Site(
                    name=context.site_name,
                    status=context.status,
                    description=context.description,
                )
            ),
        ]

        context.response = client.ingest(entities=entities)
        return context.response


@then("the site object is updated in the database")
def check_site_object_updated(context):
    """Check if the response status code is 200 and the result is success"""
    assert context.response is not None
    site = get_site(context.site_name)
    assert site.get("name") == context.site_name
    assert site.get("status").get("value") == context.status
    assert site.get("description") == context.description
