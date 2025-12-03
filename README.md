# Io

Personal AI assistant with multi-frontend support and centralized Go backend.

## Roadmap

### Backend (Go)

**Infrastructure**
- [x] CI/CD pipeline
- [ ] Docker containerization
- [x] PostgreSQL database
  - [x] Schema design
  - [x] Goose migrations
  - [x] sqlc integration

**MCP Host**
- [ ] MCP server implementation
- [ ] MCP clients/session management

**AI Providers**
- [ ] Provider orchestration
- [ ] OpenAI integration
- [ ] Claude integration
- [ ] Grok integration
- [ ] Gemini integration
- [ ] Streaming support

**Advanced Features**
- [ ] Personalities system
- [ ] Autonomy features
- [ ] Notifications

### Discord Frontend (discord.js)

**Core Functionality**
- [ ] Basic message flow
- [ ] Advanced simulated streaming/typing indicators
- [ ] Tool call/orchestration indicators

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
