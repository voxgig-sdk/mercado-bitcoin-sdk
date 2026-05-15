# ProjectName SDK exists test

import pytest
from mercadobitcoin_sdk import MercadoBitcoinSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = MercadoBitcoinSDK.test(None, None)
        assert testsdk is not None
