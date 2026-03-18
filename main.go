package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	URL = "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m,apparent_temperature,relative_humidity_2m,dew_point_2m,precipitation_probability,precipitation,rain,showers,snowfall,snow_depth,weather_code,pressure_msl,surface_pressure,visibility,wind_speed_10m,wind_direction_10m,wind_gusts_10m,soil_temperature_0cm,soil_moisture_0_to_1cm,cloud_cover"
)

type args struct {
	Latitude  float64 `json:"latitude" jsonschema:"The float value of the latitude"`
	Longitude float64 `json:"longitude" jsonschema:"The float value of the longitude"`
}

func get_weather(latitude, longitude float64) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://wttr.in/%f,%f", latitude, longitude))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "Weather-MCP-Server", Title: "Weather MCP Server", Version: "0.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "get_weather", Description: "Supply latitude and longitude to this tool, to receive data about the weather around that location."}, func(ctx context.Context, req *mcp.CallToolRequest, args args) (*mcp.CallToolResult, any, error) {
		text, err := get_weather(args.Latitude, args.Longitude)
		if err != nil {
			log.Printf("Failed to get weather data from website: %v", err)
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: text},
			},
		}, nil, nil
	})

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Printf("Server failed %v", err)
	}
}
