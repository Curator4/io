package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net"
	"os"

	"github.com/curator4/io/backend/internal/core"
	"github.com/curator4/io/backend/internal/database"
	grpcserver "github.com/curator4/io/backend/internal/grpc"
	"github.com/curator4/io/backend/internal/llm"
	pb "github.com/curator4/io/backend/internal/proto"
	"google.golang.org/grpc"
)

func main() {
	// load env config
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	// connect to db
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	queries := database.New(db)

	// validate required API keys
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is required")
	}

	// initialize LLM providers
	providers := map[string]llm.Provider{
		"openai": llm.NewOpenAIProvider(openaiKey),
	}

	// intialize backend core
	coreInstance := core.NewCore(queries, providers)

	// initialize gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterIOServiceServer(grpcServer, grpcserver.NewServer(coreInstance))

	// initialize tcp listerner
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	log.Printf("backend server succesfully initialized, listening on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
