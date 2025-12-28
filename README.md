# Time MCP Server

A simple MCP server that provides the current time in any timezone.

## Building

```bash
go build
```

## Running the Server

```bash
./timemcp
```

## Connecting from Claude Web UI

When adding this as a custom MCP connector in Claude's web UI, use the `/mcp` path:

```
https://your-hostname:8000/mcp
```

## Configuration

### Environment Variables

| Variable | Default                 | Description        |
| -------- | ----------------------- | ------------------ |
| `HOST`   | (empty, all interfaces) | Host/IP to bind to |
| `PORT`   | `8000`                  | Port to listen on  |

Example:

```bash
HOST=localhost PORT=3000 ./timemcp
```
