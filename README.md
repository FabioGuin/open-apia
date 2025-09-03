# OpenAPIA - Open Artificial Intelligence Architecture

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/badge/version-0.1.0-green.svg)](https://github.com/FabioGuin/OpenAPIA/releases)

## Overview

OpenAPIA is an open, vendor-agnostic standard for describing, documenting, and validating artificial intelligence systems. Inspired by the success of OpenAPI for REST APIs, OpenAPIA provides a machine-readable format to specify AI models, prompts, constraints, workflows, and evaluation metrics.

## Key Features

- **AI-Native Design**: Models, prompts, and constraints as first-class entities
- **Vendor Agnostic**: Works with any AI provider (OpenAI, Anthropic, Google, etc.)
- **Multi-Modal Support**: LLM, Vision, Audio, and Multimodal AI systems
- **Automation Integration**: Declarative integration with n8n, Zapier, and other automation platforms
- **Built-in Ethics**: Mandatory fields for safety, bias prevention, and explainability
- **Comprehensive Evaluation**: Metrics for accuracy, performance, and safety
- **Extensible**: Support for custom use cases and domain-specific requirements
- **Hierarchical Composition**: Inherit and compose specifications across organizational levels
- **Multi-Environment Support**: Different configurations for dev, staging, and production

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/FabioGuin/OpenAPIA.git
cd OpenAPIA
```

### 2. Create Your First OpenAPIA Specification

#### Simple Specification
```yaml
openapia: "0.1.0"
info:
  title: "My AI Assistant"
  version: "1.0.0"
  description: "A helpful AI assistant for customer support"

models:
  - id: "main_model"
    type: "LLM"
    provider: "openai"
    name: "gpt-4"
    purpose: "conversation"

prompts:
  - id: "system_prompt"
    role: "system"
    template: "You are a helpful AI assistant for {{company_name}}"

constraints:
  - id: "safety"
    rule: "output NOT contains harmful_content"
    severity: "critical"

tasks:
  - id: "handle_query"
    description: "Process user queries safely"
    steps:
      - action: "generate"
        model: "main_model"
        prompt: "system_prompt"
```

#### Hierarchical Specification
```yaml
openapia: "0.1.0"

# Inherit from parent specifications
inherits:
  - "../openapia-global.yaml"
  - "../../openapia-team.yaml"

info:
  title: "Feature-Specific AI Assistant"
  version: "1.0.0"
  description: "AI assistant for specific feature"
  ai_metadata:
    hierarchy_info:
      level: "feature"
      scope: "project"
      inheritance_mode: "merge"

# Additional models and constraints specific to this feature
models:
  - id: "feature_specific_model"
    type: "Classification"
    provider: "huggingface"
    name: "sentiment-analyzer"

constraints:
  - id: "feature_performance"
    type: "performance"
    rule: "response_time < 2s"
    severity: "high"
```

### 3. Validate Your Specification

```bash
# Simple validation
python validators/python/openapia_validator.py validate spec/my-ai-system.yaml

# Hierarchical validation (with inheritance)
python validators/python/openapia_validator.py validate spec/my-ai-system.yaml --hierarchical

# Using JavaScript validator
node validators/javascript/cli.js validate spec/my-ai-system.yaml

# Using PHP validator
php validators/php/cli.php validate spec/my-ai-system.yaml

# Using Go validator
go run validators/go/cli.go validate spec/my-ai-system.yaml
```

## Examples

Check out our [examples directory](spec/examples/) for complete implementations:

### Core AI Examples
- [Customer Support AI](spec/examples/customer-support.yaml) - E-commerce assistant
- [Content Moderator](spec/examples/content-moderator.yaml) - AI-powered content filtering
- [Multilingual Chatbot](spec/examples/multilingual-chatbot.yaml) - Multi-language support

### Automation Integration Examples
- [E-commerce with n8n](spec/examples/ecommerce-automation.yaml) - Complete order processing with automation workflows
- [Customer Support with Zapier](spec/examples/zapier-automation.yaml) - Simple webhook-based automation integration
- [MCP Integration](spec/examples/mcp-integration.yaml) - Model Context Protocol server integration

## Hierarchical Composition

OpenAPIA supports hierarchical composition for complex organizational structures:

- **[Hierarchical Composition Guide](docs/hierarchical-composition.md)** - Complete guide to inheritance and composition
- **Enterprise Use Cases** - Global, regional, and department-level specifications
- **Team Development** - Sprint and feature-level configurations
- **Environment Management** - Dev, staging, and production configurations

### Quick Hierarchical Example

```yaml
# Global specification
openapia: "0.1.0"
info:
  title: "Global AI Standards"
  ai_metadata:
    hierarchy_info:
      level: "global"
      scope: "organization"

# Feature specification inheriting from global
openapia: "0.1.0"
inherits:
  - "../openapia-global.yaml"
info:
  title: "Feature-Specific AI"
  ai_metadata:
    hierarchy_info:
      level: "feature"
      scope: "project"
```

## Specification

The OpenAPIA specification is available in multiple formats:

- **YAML**: `spec/openapia-0.1.yaml` - Official specification (human-readable)
- **JSON Example**: `spec/examples/openapia-0.1-example.json` - Complete working example for tools and SDKs
- **JSON Schema**: `spec/schemas/openapia-0.1-schema.json` - Validation schema
- **Examples**: `spec/examples/*.yaml` - Real-world use case examples

See `spec/README.md` for detailed file structure and usage guide.

The specification includes these core sections:

- **`openapia`** - Specification version
- **`info`** - System metadata and AI-specific information
- **`models`** - AI models with capabilities, limits, and costs
- **`prompts`** - Structured prompts with variables and configuration
- **`constraints`** - Safety, ethical, and operational constraints
- **`tasks`** - Declarative workflows and business logic
- **`context`** - State management and memory configuration
- **`evaluation`** - Metrics, tests, and performance monitoring

### Documentation

- **[OpenAPIA 0.1 Specification](docs/openapia-0.1-specification.md)** - Complete formal specification documentation
- **[Getting Started Guide](docs/getting-started.md)** - Step-by-step tutorial for new users
- **[Hierarchical Composition Guide](docs/hierarchical-composition.md)** - Advanced inheritance and composition patterns

## Tools and Libraries

### Validators
- [Python Validator](validators/python/) - Full-featured validation library
- [PHP Validator](validators/php/) - PHP validation with Symfony YAML
- [JavaScript Validator](validators/javascript/) - Node.js and browser support
- [Go Validator](validators/go/) - High-performance validation

### CLI Usage
Each validator includes a command-line interface for easy validation:

```bash
# Python
python validators/python/openapia_validator.py spec.yaml

# PHP
php validators/php/cli.php -f spec.yaml

# JavaScript
node validators/javascript/cli.js -f spec.yaml

# Go
go run validators/go/. -f spec.yaml
```

## Automation Integration

OpenAPIA 0.1 introduces declarative integration with external automation platforms, enabling AI systems to trigger and coordinate with automation workflows without tight coupling.

### Supported Platforms

- **n8n**: Complex business process automation
- **Zapier**: Simple integrations and notifications  
- **Microsoft Power Automate**: Enterprise workflows
- **Custom Webhooks**: Legacy system connections

### Example: E-commerce Order Processing

```yaml
automations:
  - id: "order_processing_workflow"
    name: "Order Processing Automation"
    provider: "n8n"
    
    trigger:
      type: "webhook"
      endpoint: "/webhooks/order-created"
      conditions:
        - "order_status == 'validated'"
    
    integration:
      type: "external_workflow"
      workflow_id: "n8n://workflows/order-processing"
      timeout: "5m"
    
    data_contract:
      input:
        order_id:
          type: "string"
          required: true
        customer_id:
          type: "string"
          required: true
      output:
        processing_status:
          type: "string"
          enum: ["processing", "shipped", "delivered", "failed"]

tasks:
  - id: "process_order"
    steps:
      - name: "trigger_processing"
        action: "automation"
        automation: "order_processing_workflow"
        automation_parameters:
          order_id: "${input.order_id}"
          customer_id: "${input.customer_id}"
```

### Key Benefits

- **Declarative**: Define what automations to trigger, not how
- **Vendor Agnostic**: Works with any automation platform
- **Monitored**: Built-in health checks and metrics
- **Secure**: Configurable authentication and validation

### Documentation

- **[Automation Integration Guide](docs/automation-integration.md)** - Complete guide with examples and best practices
- **[E-commerce Example](spec/examples/ecommerce-automation.yaml)** - Full n8n integration example
- **[Zapier Example](spec/examples/zapier-automation.yaml)** - Simple webhook integration example

## Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Process

1. **Issues**: Open an issue to discuss changes
2. **Pull Requests**: Submit PRs for implementation
3. **Review**: Maintainer review with community feedback
4. **Decision**: Transparent decision-making process

## Governance

OpenAPIA is currently in **Bootstrap Phase** with Fabio Guin as the initial maintainer. We plan to transition to a community-driven governance model as the project grows.

See [GOVERNANCE.md](GOVERNANCE.md) for details.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Roadmap

- [x] **v0.1.0** - Core specification, validators, and automation integration
- [ ] **v0.2.0** - Multi-agent support and advanced automation features
- [ ] **v0.3.0** - CLI tools and documentation generator
- [ ] **v1.0.0** - Community governance and ecosystem

## Community

- **Discussions**: [GitHub Discussions](https://github.com/FabioGuin/OpenAPIA/discussions)
- **Issues**: [GitHub Issues](https://github.com/FabioGuin/OpenAPIA/issues)
- **Email**: random@starzero.it

## Acknowledgments

OpenAPIA is inspired by:
- [OpenAPI Specification](https://github.com/OAI/OpenAPI-Specification) for API documentation
- [ISO/IEC 42001](https://www.iso.org/standard/81231.html) for AI governance
- [LangChain](https://github.com/langchain-ai/langchain) for AI orchestration
- [IEEE AI-ESTATE](https://standards.ieee.org/standard/1232-2010.html) for AI component formalization

---

**Note**: This project is in active development. The specification may change between versions. Please check the [CHANGELOG](CHANGELOG.md) for breaking changes.
