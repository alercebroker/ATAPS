from grpc_server.adqlparser_pb2 import ADQLRequest, SQLResponse
from grpc_server.adqlparser_pb2_grpc import (
    ADQLParserServicer,
    add_ADQLParserServicer_to_server,
)
import grpc
from concurrent import futures
from ADQL.adql import ADQL
from spatial_index import SpatialIndex


class ADQLParserServer(ADQLParserServicer):
    def __init__(self):
        super().__init__()
        self.adql = ADQL(
            dbms="oracle",
            level=20,
            racol="ra",
            deccol="dec",
            xcol="X",
            ycol="Y",
            zcol="Z",
            indxcol="HTM20",
            mode=SpatialIndex.HTM,
            encoding=SpatialIndex.BASE10,
        )

    def Parse(self, request: ADQLRequest, context):
        try:
            self.adql.sql(request.query)
            return SQLResponse(error="", parsed=request.query)
        except Exception as e:
            return SQLResponse(error=str(e), parsed="")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_ADQLParserServicer_to_server(ADQLParserServer(), server)
    server.add_insecure_port("[::]:50051")
    print("Starting server. Listening on port 50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
