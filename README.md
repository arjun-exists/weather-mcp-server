# Weather MCP Server

A simple MCP server written in Go that fetches weather data and supplies it to an AI model. You give it a latitude and longitude, and it returns the current weather for that location. The server communicates with the AI via the Model Context Protocol over stdio.

## Running it in Claude Desktop

First build the binary:

```bash
go build -o weather-mcp-server.exe .
```

Then open your Claude Desktop config file at `%APPDATA%\Claude\claude_desktop_config.json` and add the following:

```json
{
  "mcpServers": {
    "weather": {
      "command": "C:\\path\\to\\weather-mcp-server.exe"
    }
  }
}
```

Replace the path with wherever you built the binary. Restart Claude Desktop and the weather tool will be available in your conversations.
