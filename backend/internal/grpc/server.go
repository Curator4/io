package grpc

import (
	"context"
	"log"

	"github.com/curator4/io/backend/internal/core"
	"github.com/curator4/io/backend/internal/domain"
	pb "github.com/curator4/io/backend/internal/proto"
)

// Server implements IOService grpc service,
type Server struct {
	pb.UnimplementedIOServiceServer
	core *core.Core
}

// NewServer creates a new grpc server instance
func NewServer(core *core.Core) *Server {
	return &Server{
		core: core,
	}
}

// SendMessage handles the type conversion from pb to and from domain types,
// calls the core handler, HandleSendMessage, to do the actual logic
func (s *Server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	// convert from protobuf to domain format
	content := domain.MessageContentFromPb(req.Content)

	// core handler (now returns LLMResponse and lifecycle)
	llmResponse, lifecycle, err := s.core.HandleSendMessage(ctx, content, req.Username)
	if err != nil {
		log.Printf("SendMessage error for user %s: %v", req.Username, err)
		return nil, toGRPCError(err)
	}

	// convert response content to protobuf
	responseContentPb := domain.MessageContentToPb(llmResponse.Content)

	// convert actions to protobuf
	actionsPb := make([]*pb.Action, len(llmResponse.Actions))
	for i, action := range llmResponse.Actions {
		actionsPb[i] = domain.ActionToPb(action)
	}

	// convert lifecycle to protobuf
	lifecyclePb := domain.ConversationLifecycleToPb(lifecycle)

	// respond
	return &pb.SendMessageResponse{
		Content:   responseContentPb,
		Actions:   actionsPb,
		Lifecycle: lifecyclePb,
	}, nil
}

// StoreMessage handles the type conversions to and from domain types,
// calls the core handler, HandleStoreMessage
func (s *Server) StoreMessage(ctx context.Context, req *pb.StoreMessageRequest) (*pb.StoreMessageResponse, error) {
	content := domain.MessageContentFromPb(req.Content)

	lifecycle, err := s.core.HandleStoreMessage(ctx, content, req.Username)
	if err != nil {
		log.Printf("StoreMessage error for user %s: %v", req.Username, err)
		return nil, toGRPCError(err)
	}

	lifecyclePb := domain.ConversationLifecycleToPb(*lifecycle)

	return &pb.StoreMessageResponse{
		Success:   true,
		Lifecycle: lifecyclePb,
	}, nil
}
