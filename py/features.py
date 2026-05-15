# MercadoBitcoin SDK feature factory

from feature.base_feature import MercadoBitcoinBaseFeature
from feature.test_feature import MercadoBitcoinTestFeature


def _make_feature(name):
    features = {
        "base": lambda: MercadoBitcoinBaseFeature(),
        "test": lambda: MercadoBitcoinTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
