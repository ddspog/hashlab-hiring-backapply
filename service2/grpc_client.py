import grpc
from  discount import DiscountStub, VerifyDiscountForm


def infoDummy(s, *args):
    pass

class DiscountClient:
    def __init__(self, url, log):
        self.url = url
        self.client = None
        if log != None:
            self.info = log.info
        else:
            self.info = infoDummy

    def prepare_client(self):
        if self.client == None:
            channel = grpc.insecure_channel(self.url)
            self.client = DiscountStub(channel)
            self.info("Connected to %s", self.url)

    def verify(self, uid, pid):
        self.prepare_client()

        form = VerifyDiscountForm(userid=uid, prodid=pid)
        info = self.client.VerifyDiscount(form)

        return {
            "pct": info.pct,
            "value_in_cents": info.value_in_cents
        }