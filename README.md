# Time MCP Server

A simple MCP server that provides the current time in any timezone.

## Running the Server

```bash
uv run python main.py
```

**Important:** Do NOT use `uv run mcp run main.py` - that command seemingly only supports `stdio` and `sse` transports,
not `streamable-http`.

## Connecting from Claude Web UI

When adding this as a custom MCP connector in Claude's web UI, use the `/mcp` path:

```
https://your-hostname:8000/mcp
```

## Configuration

### Environment Variables

| Variable        | Default   | Description                               |
| --------------- | --------- | ----------------------------------------- |
| `HOST`          | `0.0.0.0` | Host/IP to bind to                        |
| `PORT`          | `8000`    | Port to listen on                         |
| `ALLOWED_HOSTS` | (none)    | Comma-separated list of allowed hostnames |

Example:

```bash
PORT=3000 ALLOWED_HOSTS=my-server.example.com,other-host.local uv run python main.py
```

### Allowed Hosts (DNS Rebinding Protection)

The MCP SDK includes DNS rebinding protection that validates the `Host` header. If you're accessing the server through a
proxy or custom domain (e.g., Tailscale), you'll get `421 Misdirected Request` errors. Add any required hosts via the
`ALLOWED_HOSTS` environment variable.

## Tool

### `time`

Get the current time in a given timezone.

- **Parameter:** `tz` - Timezone string (default: "Europe/Amsterdam")
- **Returns:** Formatted time string like "Sunday, 28 December 2025 16:07:02 CET"
