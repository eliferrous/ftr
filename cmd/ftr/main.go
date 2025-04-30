package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/eliferrous/ftr/execmtr"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	srv := server.NewMCPServer("FTR", "0.1.0")

	runMTR := mcp.NewTool("run_mtr",
		mcp.WithDescription("Run MTR (IPv4) and return raw JSON"),
		mcp.WithString("target", mcp.Required()),
	)

	srv.AddTool(runMTR, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return callMTR(ctx, req.Params.Arguments["target"].(string), false)
	})

	if err := serveMCP(srv); err != nil {
		fmt.Fprintln(os.Stderr, "server error:", err)
		os.Exit(1)
	}
}

func callMTR(ctx context.Context, target string, v6 bool) (*mcp.CallToolResult, error) {
	runner := execmtr.New()

	rep, err := runner.Run(ctx, target, 5)
	if err != nil {
		return nil, err
	}

	// marshal back to JSON for the MCP payload
	pretty, _ := json.MarshalIndent(rep, "", "  ")
	return mcp.NewToolResultText(string(pretty)), nil
}

func serveMCP(srv *server.MCPServer) error {
	return server.ServeStdio(srv)
}
