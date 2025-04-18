package main

import (
	"context"
	"flag"
	"log"
	"obsidianOptimizeMCP/pkg"
	"obsidianOptimizeMCP/types"

	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	var baseURL string
	var token string
	flag.StringVar(&baseURL, "obsidian-url",
		getEnv("OBSIDIAN_URL", "http://127.0.0.1:27123"), "Obsidian API base URL")
	flag.StringVar(&token, "obsidian-token",
		getEnv("OBSIDIAN_TOKEN", "xxxxxxxxxx"), "Obsidian Bearer token")
	flag.Parse()

	config := &types.Config{
		ObsidianURL:   baseURL,
		ObsidianToken: token,
	}

	messageEndpointURL := "/message"

	sseTransport, mcpHandler, err := transport.NewSSEServerTransportAndHandler(messageEndpointURL)
	if err != nil {
		log.Panicf("new sse transport and handler with error: %v", err)
	}

	// new mcp server
	mcpServer, err := server.NewServer(sseTransport)
	if err != nil {
		log.Panicf("new mcp server error: %v", err)
	}
	pkg.Register(mcpServer, config)

	// start mcp Server
	go func() {
		mcpServer.Run()
	}()

	defer mcpServer.Shutdown(context.Background())

	gin.DefaultWriter = logrusWriter{}
	gin.DefaultErrorWriter = logrusWriter{}
	r := gin.New()
	r.GET("/sse", func(ctx *gin.Context) {
		mcpHandler.HandleSSE().ServeHTTP(ctx.Writer, ctx.Request)
	})
	r.POST(messageEndpointURL, func(ctx *gin.Context) {
		mcpHandler.HandleMessage().ServeHTTP(ctx.Writer, ctx.Request)
	})

	if err = r.Run(":8080"); err != nil {
		return
	}
}
