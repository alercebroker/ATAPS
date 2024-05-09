from grpc_server.adqlparser_pb2 import ADQLRequest
from grpc_server.adqlparser_server import ADQLParserServer


def test_parse():
    server = ADQLParserServer()
    # query = "select ra, dec from iraspsc where (contains(point('icrs', ra, dec), polygon('ECLIPTIC', 233.56, 34.567, 233.56, 33.567, 234.56, 33.567, 234.56, 34.567)) = 1 order by dec desc"
    query = """SELECT *, DISTANCE(81.28, -69.78, ra, dec) AS ang_sep
FROM gaiadr3.gaia_source
WHERE DISTANCE(81.28, -69.78, ra, dec) < 5./60.
AND phot_g_mean_mag < 20.5
AND parallax IS NOT NULL
ORDER BY ang_sep ASC
    """
    request = ADQLRequest(query=query)
    response = server.Parse(request, None)
    assert response.error == ""
    assert response.parsed == ""
