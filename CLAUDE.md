# Io - Personal AI Assistant

Multi-frontend AI assistant with centralized Go backend. Supports multiple AI providers and conversation management.

## Architecture

- **Backend** (`backend/`) - Go gRPC server, shared state/logic
- **Frontends** - Multiple clients (currently Discord bot), communicate via gRPC
- **Proto** (`proto/`) - Protobuf service definitions, shared between backend + frontends
- **Database** - PostgreSQL for persistence
- **Active State** - Current conversation/AI config lives in app memory, NOT database

## Backend (`backend/`)

**Tech Stack:**
- Go
- PostgreSQL with sqlc (type-safe queries) and goose (migrations)
- Docker containerized

**Database Schema:**
- `ai_configs` - AI provider configurations (api, model, system_prompt)
- `conversations` - Chat conversations with names and timestamps
- `messages` - Individual messages with role (user/assistant/system) and content
- Foreign keys: `messages.conversation_id → conversations.id` (CASCADE delete)

**Directory Structure:**
- `sql/schema/` - Goose migration files (numbered)
- `sql/queries/` - sqlc query definitions
- `internal/database/` - Generated sqlc code (don't edit manually)
- `internal/mcp/` - MCP integration
- `internal/llm/` - LLM provider abstractions

## Discord Frontend (`discord/`)

- Built with discord.js (TypeScript, ESM)
- Communicates with backend via gRPC (@grpc/grpc-js, ts-proto)
- Handles Discord-specific rendering (embeds, buttons, select menus)

## Current Feature Work

**`ask_user` tool** — allows the AI to present interactive questions/choices mid-conversation:
- Backend: OpenAI tool definition → returns structured question to frontend
- Proto: carries question + options between backend and Discord
- Discord: renders as buttons/select menus, collects response, sends back
- Backend: feeds answer back into conversation, AI continues

**MCP elicitation** (next) — same Discord UI, but triggered by MCP servers via protocol
