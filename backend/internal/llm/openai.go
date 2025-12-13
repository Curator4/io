// llm is the package that manages llm providers. see Provider interface
package llm

import (
	"context"
	"fmt"

	"github.com/curator4/io/backend/internal/domain"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"github.com/openai/openai-go/v3/shared"
)

type OpenAIProvider struct {
	client *openai.Client
}

// NewOpenAIProvider creates a new OpenAI provider with the given API key
func NewOpenAIProvider(apiKey string) Provider {
	return &OpenAIProvider{
		client: NewOpenAIClient(apiKey),
	}
}

// supportedModels maps model names to OpenAI SDK constants
var supportedModels = map[string]shared.ResponsesModel{
	"gpt-5.1":    openai.ChatModelGPT5ChatLatest,
	"gpt-5-nano": openai.ChatModelGPT5Nano,
	"gpt-5-mini": openai.ChatModelGPT5Mini,
}

func (p OpenAIProvider) SendMessage(ctx context.Context, messages []domain.Message, config domain.AIConfig) (domain.MessageContent, error) {

	model, ok := supportedModels[config.Model.Name]
	if !ok {
		return domain.MessageContent{}, fmt.Errorf("unknown or unsupported model: %s", config.Model.Name)
	}

	input := responses.ResponseNewParamsInputUnion{
		OfInputItemList: messagesToOpenAIInput(messages),
	}

	params := responses.ResponseNewParams{
		Model:        model,
		Instructions: openai.String(config.SystemPrompt),
		Input:        input,
		Reasoning: openai.ReasoningParam{
			Effort: openai.ReasoningEffortLow,
		},
	}

	resp, err := p.client.Responses.New(ctx, params)
	if err != nil {
		return domain.MessageContent{}, fmt.Errorf("openai api error: %w", err)
	}

	return domain.MessageContent{
		Text: resp.OutputText(),
	}, nil
}

func NewOpenAIClient(apikey string) *openai.Client {
	client := openai.NewClient(
		option.WithAPIKey(apikey),
	)
	return &client
}

// buildInputMessage creates an input message (user/system/developer) for OpenAI
func buildInputMessage(msg domain.Message) responses.ResponseInputItemUnionParam {
	contentItems := make(responses.ResponseInputMessageContentListParam, 0)

	// Add text content (with user prefix if multi-user)
	text := msg.Content.Text
	if msg.User != nil && text != "" {
		text = fmt.Sprintf("%s: %s", msg.User.Name, text)
	}
	if text != "" {
		contentItems = append(contentItems, responses.ResponseInputContentUnionParam{
			OfInputText: &responses.ResponseInputTextParam{
				Text: text,
				Type: "input_text",
			},
		})
	}

	// Add all media items
	for _, media := range msg.Content.Media {
		switch media.Type {
		case "image":
			contentItems = append(contentItems, responses.ResponseInputContentUnionParam{
				OfInputImage: &responses.ResponseInputImageParam{
					ImageURL: openai.String(media.URL),
					Type:     "input_image",
				},
			})
		case "video", "file", "audio":
			contentItems = append(contentItems, responses.ResponseInputContentUnionParam{
				OfInputFile: &responses.ResponseInputFileParam{
					FileURL: openai.String(media.URL),
					Type:    "input_file",
				},
			})
		}
	}

	return responses.ResponseInputItemUnionParam{
		OfInputMessage: &responses.ResponseInputItemMessageParam{
			Content: contentItems,
			Role:    string(msg.Role),
			Type:    "message",
		},
	}
}

// buildOutputMessage creates an output message (assistant) for OpenAI
func buildOutputMessage(msg domain.Message) responses.ResponseInputItemUnionParam {
	outputContent := make([]responses.ResponseOutputMessageContentUnionParam, 0)

	if msg.Content.Text != "" {
		outputContent = append(outputContent, responses.ResponseOutputMessageContentUnionParam{
			OfOutputText: &responses.ResponseOutputTextParam{
				Text: msg.Content.Text,
				Type: "output_text",
			},
		})
	}

	// Note: Assistant-generated media (images from DALL-E, etc.) come through tool calls,
	// not as direct message content, so we don't include msg.Content.Media here

	return responses.ResponseInputItemUnionParam{
		OfOutputMessage: &responses.ResponseOutputMessageParam{
			ID:      fmt.Sprintf("msg_%s", msg.ID.String()),
			Content: outputContent,
			Status:  "completed",
			Role:    "assistant",
			Type:    "message",
		},
	}
}

// messagesToOpenAIInput takes a slice of messages as input, and converts them to a ResponseInputParam for OpenAI
// the type definitions for all these are a real jungle, rely heavily on the go docs
// ResponseInputParam is the input, it a slice of ResponseInputItemUnionParams
// those itemunions can be many things, here just input and output messages, each have their seperate subtypes
func messagesToOpenAIInput(messages []domain.Message) responses.ResponseInputParam {
	items := make([]responses.ResponseInputItemUnionParam, 0, len(messages))

	for _, msg := range messages {
		if msg.Role == domain.RoleAssistant {
			items = append(items, buildOutputMessage(msg))
		} else {
			items = append(items, buildInputMessage(msg))
		}
	}

	return items
}
