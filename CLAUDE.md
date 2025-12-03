# Io - Personal AI Assistant

Multi-frontend AI assistant with centralized Go backend. Supports multiple AI providers and conversation management.

## Architecture

- **Backend** (`backend/`) - Go REST API, shared state/logic
- **Frontends** - Multiple clients (currently Discord bot)
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

- Built with discord.js
- Communicates with backend API
