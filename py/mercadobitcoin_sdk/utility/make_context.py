# MercadoBitcoin SDK utility: make_context

from mercadobitcoin_sdk.core.context import MercadoBitcoinContext


def make_context_util(ctxmap, basectx):
    return MercadoBitcoinContext(ctxmap, basectx)
