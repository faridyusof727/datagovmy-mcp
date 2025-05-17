package main

import (
	"fmt"

	"github.com/faridyusof727/datagovmy-mcp/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"data.gov.my",
		"1.0.0",
	)

	for tool, reg := range tools.LoadTools() {
		s.AddTool(*tool, reg)
	}

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
