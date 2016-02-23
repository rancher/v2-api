import pytest
import cattle


@pytest.fixture
def client():
    url = 'http://localhost:8899/v2/schemas'
    return cattle.from_env(url=url)


def test_service_list(client):
    services = client.list_service()
    assert len(services) > 0

    assert services[0].name is not None
