# ADQL To SQL Parser (Python)

This is a simple parser that converts ADQL queries to SQL queries. It is written in Python and uses the `ADQL` library to parse the ADQL queries.

It is exposed as a gRPC service, consumed by the main TAP service when a user submits an ADQL query.

## Development

Install with poetry:

```bash
poetry install
```

Run the tests:

```bash
poetry run pytest
```

Generate python code from the proto file:

```bash
poetry run python -m grpc_tools.protoc -Igrpc_server=./grpc_server/protos --python_out=. --pyi_out=. --grpc_python_out=. ./grpc_server/protos/adqlparser.proto
```
