# Changelog

All notable changes to OpenAPIA will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **Security Best Practices**: Comprehensive security framework for managing sensitive data
  - New `.env.example` template with all variables used in OpenAPIA examples and documentation
  - Complete analysis of environment variables across all files (examples, docs, templates)
  - Fixed hardcoded webhook URLs in automation examples, replaced with environment variables
  - New `docs/ai-security-considerations.md` with comprehensive AI security framework
  - New `docs/security-vulnerabilities.md` with OWASP Top 10 for AI and vulnerability prevention
  - Updated `.gitignore` to properly handle environment files
  - New `docs/security-best-practices.md` with complete security guidelines
  - Environment variable usage examples in all YAML specifications
  - Updated documentation with security setup instructions
  - Template examples showing secure authentication patterns
- **Model Context Protocol (MCP) Support**: Full integration of MCP servers for external data and tool access
  - New `context.mcp_servers` section for configuring MCP server connections
  - Support for stdio, SSE, and WebSocket transport protocols
  - Authentication configuration (none, api_key, oauth, custom)
  - Security settings including rate limiting and operation restrictions
  - Health check configuration for MCP server monitoring
  - New task actions: `mcp_tool` and `mcp_resource` for MCP integration
  - MCP-specific task step fields: `mcp_server`, `mcp_tool`, `mcp_resource`, `mcp_parameters`
  - Advanced MCP configuration in `extensions.advanced.mcp`
  - Comprehensive MCP validation in all validators
  - Example MCP integration specification (`examples/automation/mcp-integration.yaml`)
  - Updated documentation with complete MCP reference

- **Automation Integration**: Declarative integration with external automation platforms
  - New `automations` section for defining external automation workflows
  - Support for multiple automation providers: n8n, Zapier, Microsoft Power Automate, custom webhooks
  - Trigger configuration: webhook, scheduled, event, and manual triggers
  - Data contracts for explicit input/output schemas
  - Monitoring and health check configuration
  - New task action: `automation` for triggering external workflows
  - Task step fields: `automation`, `automation_parameters`, `check_automation`
  - Advanced automation configuration in `extensions.advanced.automation`
  - Provider-specific settings for n8n, Zapier, and other platforms
  - Example specifications: e-commerce automation (`examples/automation/ecommerce-automation.yaml`) and Zapier integration (`examples/automation/zapier-automation.yaml`)
  - Complete automation integration guide (`docs/automation-integration.md`)

### Changed
- Extended task step actions to include MCP operations and automation triggers
- Enhanced context section with MCP server configuration
- Updated validators to support MCP validation and cross-references
- Extended task step actions to include automation operations
- Enhanced extensions section with automation support configuration
- Updated JSON schema to include automation validation

### Deprecated
- Nothing yet

### Removed
- Nothing yet

### Fixed
- Nothing yet

### Security
- Added MCP-specific security configurations including rate limiting and operation restrictions

## [0.1.0] - 2025-09-03

### Added
- Initial OpenAPIA v0.1.0 specification and validators
- Core sections: `openapia`, `info`, `models`, `prompts`, `constraints`, `tasks`, `context`, `evaluation`
- Example specifications for customer support AI, content moderation, and multilingual chatbot
- Multi-language validators (JavaScript, PHP, Python, Go)
- Hierarchical composition support for complex AI systems
- Machine-readable JSON specification format
- Basic governance structure
- Apache 2.0 license
- Contributing guidelines
- README with quick start guide

### Specification Features
- **Models**: Support for LLM, Vision, Audio, Multimodal, Classification, Embedding
- **Prompts**: Structured prompts with variables and configuration
- **Constraints**: Safety, ethical, and operational constraints with enforcement
- **Tasks**: Declarative workflows with conditional steps
- **Context**: State management and memory configuration
- **Evaluation**: Metrics, tests, and performance monitoring
- **Extensions**: Support for custom use cases
- **Hierarchical Composition**: Support for multi-agent systems and complex workflows

### Validators
- JavaScript validator with CLI support
- PHP validator with comprehensive testing
- Python validator for cross-platform validation
- Go validator for high-performance validation

### Documentation
- Comprehensive README with examples
- Formal OpenAPIA 0.1 specification documentation
- Hierarchical composition documentation
- Contributing guidelines
- Governance documentation
- License and legal information

---

## Release Notes Format

### Added
- New features and capabilities

### Changed
- Changes to existing functionality

### Deprecated
- Features that will be removed in future versions

### Removed
- Features removed in this version

### Fixed
- Bug fixes

### Security
- Security improvements and fixes

## Version History

- **0.1.0**: Initial release with core specification
- **Future versions**: Will be documented here as they are released

## Breaking Changes

Breaking changes will be clearly marked and documented with migration guides.

## Migration Guides

Migration guides will be provided for major version changes to help users upgrade their specifications.

---

*This changelog follows the [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) format.*
