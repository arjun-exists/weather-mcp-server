package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type args struct {
	Latitude  float64 `json:"latitude" jsonschema:"The float value of the latitude"`
	Longitude float64 `json:"longitude" jsonschema:"The float value of the longitude"`
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "Weather-MCP-Server", Title: "Weather MCP Server", Version: "0.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "Get weather", Description: "Supply latitude and longitude to this tool, to receive data about the weather around that location."}, func(ctx *context.Context, req *mcp.CallToolRequest, args args) {

	})
}
