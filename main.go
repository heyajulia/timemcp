package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Timezone string `json:"tz" jsonschema:"IANA timezone string. Defaults to Europe/Amsterdam if empty."`
}

type Output struct {
	Time string `json:"time" jsonschema:"Current time, formatted as 'Sunday, 28 December 2025 16:07:02 CET'."`
}

func Time(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
	tz := "Europe/Amsterdam"
	if input.Timezone != "" {
		tz = input.Timezone
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, Output{}, fmt.Errorf("load location %q: %w", tz, err)
	}

	return nil, Output{Time: time.Now().In(loc).Format("Monday, 2 January 2006 15:04:05 MST")}, nil
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("%s %s completed in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	handler := mcp.NewSSEHandler(func(*http.Request) *mcp.Server {
		server := mcp.NewServer(&mcp.Implementation{Name: "What's the time?", Version: "v1.0.0"}, nil)

		mcp.AddTool(server, &mcp.Tool{Name: "time", Description: "Get the current time in a given timezone"}, Time)

		return server
	}, nil)

	http.Handle("/mcp", withLogging(handler))

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	addr := host + ":" + port

	log.Printf("MCP server listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
