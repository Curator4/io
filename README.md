# Io

Personal AI assistant with multi-frontend support and centralized Go backend.

## Roadmap

### Backend (Go)

**Infrastructure**
- [x] CI/CD pipeline
- [x] Docker containerization
- [x] PostgreSQL database
  - [x] Schema design
  - [x] Goose migrations
  - [x] sqlc integration

**MCP Host**
- [ ] MCP server implementation
- [ ] MCP clients/session management

**AI Providers**
- [ ] Provider orchestration
- [x] OpenAI integration
- [ ] Claude integration
- [ ] Grok integration
- [ ] Gemini integration
- [ ] Streaming support

**Interactive Tools**
- [ ] `ask_user` tool — AI presents choices/questions to user mid-conversation
- [ ] MCP elicitation — MCP servers request structured user input via protocol

**Advanced Features**
- [ ] Personalities system
- [ ] Autonomy features
- [ ] Notifications

### Discord Frontend (discord.js)

**Core Functionality**
- [x] Basic message flow
- [ ] Conversation UI indicators
- [ ] Reasoning level dynamic + UI
- [ ] Advanced simulated streaming/typing indicators
- [ ] Tool call/orchestration indicators
- [ ] Interactive components (buttons/select menus) for `ask_user` + elicitation

**Commands**
- [ ] `/status` - Display system info
  - Container status
  - API/model info
  - Mood/autonomy state
  - Conversation/message stats
  - MCP sessions
  - Streaming status
- [ ] `/model` - Switch AI model
- [ ] `/api` - Switch AI provider
- [ ] `/mcp` - MCP management
- [ ] `/clear` - Clear conversation
- [ ] `/resume` - Resume conversation
- [ ] `/help` - Command help

**Integrations**
- [ ] Discord MCP client
- [ ] Reaction handling
- [ ] Discord.js utilities
- [ ] Memory management

cool orchestrator pattern from anthropic here: https://www.anthropic.com/engineering/building-effective-agents?ref=chris.sotherden.io
