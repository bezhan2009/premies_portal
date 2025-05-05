import grpc

from gen.cards import (cards_pb2, cards_pb2_grpc)


# from google.protobuf.empty_pb2 import Empty


def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = cards_pb2_grpc.CardsServiceStub(channel)

    response_upload = stub.UploadCardsData(
        cards_pb2.CardsUploadRequest(file_path="./uploads/cards.xlsx"))
    print(response_upload.status)

    # response_clean = stub.CleanCardsTable(Empty())
    # print(response_clean.status)


if __name__ == '__main__':
    run()
