import os

import arrow
from mcp.server.fastmcp import FastMCP
from mcp.server.transport_security import TransportSecuritySettings


def get_allowed_hosts() -> list[str]:
    return [host.strip() for host in os.environ["ALLOWED_HOSTS"].split(",")]


mcp = FastMCP(
    "What time is it?",
    stateless_http=True,
    json_response=True,
    host=os.getenv("HOST", "0.0.0.0"),
    port=int(os.getenv("PORT", 8000)),
    transport_security=TransportSecuritySettings(allowed_hosts=get_allowed_hosts()),
)


@mcp.tool()
def time(tz: str = "Europe/Amsterdam") -> str:
    """
    Get the current time in a given timezone.

    :param tz: IANA timezone string. Defaults to "Europe/Amsterdam".
    :return: Current time formatted as "Sunday, 28 December 2025 16:07:02 CET"
    """
    return arrow.now(tz).format("dddd, DD MMMM YYYY HH:mm:ss ZZZ")


if __name__ == "__main__":
    mcp.run(transport="streamable-http")
