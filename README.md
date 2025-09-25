# APAI - Architecture Protocol for Artificial Intelligence

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/badge/version-0.1.0-green.svg)](https://github.com/FabioGuin/OpenAPIA/releases)

## Overview

APAI is an open protocol for describing, documenting, and validating artificial intelligence systems. It provides a structured format to specify AI models, prompts, constraints, workflows, and evaluation metrics.

**Generative Map Concept**: The structured YAML specification serves as a generative map - a machine-readable blueprint that enables automated code generation, documentation creation, system orchestration, and AI infrastructure deployment. The structured format makes it suitable for generative development tools and workflows.

## Quick Navigation

- **[Key Features](#key-features)** - Core capabilities and design principles
- **[Generative Map Potential](#generative-map-potential)** - Code generation and automation possibilities
- **[Quick Start](#quick-start)** - Get started with APAI in minutes
- **[Examples](#examples)** - Real-world implementations and templates
- **[Hierarchical Composition](#hierarchical-composition)** - Building complex systems through inheritance
- **[Specification](#specification)** - Official specification and documentation
- **[Tools and Libraries](#tools-and-libraries)** - Validators and community tools
- **[Automation Integration](#automation-integration)** - External platform integration
- **[Contributing](#contributing)** - How to contribute to APAI
- **[Governance](#governance)** - Project governance and community structure
- **[Whitepapers](#whitepapers)** - Comprehensive technical and business analysis

## Key Features

- **AI-Focused Design**: Models, prompts, and constraints as core components
- **Provider Independence**: Works with any AI provider (OpenAI, Anthropic, Google, etc.)
- **Multi-Modal Support**: LLM, Vision, Audio, and Multimodal AI systems
- **Automation Integration**: Integration with n8n, Zapier, and other automation platforms
- **Built-in Ethics**: Required fields for safety, bias prevention, and explainability
- **Evaluation Framework**: Metrics for accuracy, performance, and safety
- **Extensible**: Support for custom use cases and domain-specific requirements
- **Hierarchical Composition**: Inherit and compose specifications across organizational levels
- **Multi-Environment Support**: Different configurations for dev, staging, and production
- **Code Generation**: Structured format enables automated code generation, documentation, and system orchestration

## Generative Map Potential

The APAI YAML specification's structured format makes it suitable for generative development. The protocol's metadata and standardized format help developers build tools that can automatically create:

### Code Generation
- **API Clients**: Generate client libraries in multiple languages (Python, JavaScript, PHP, Go)
- **Server Implementations**: Create backend services that implement the AI system
- **SDK Generation**: Build software development kits for easy integration
- **Configuration Files**: Generate deployment configs for Docker, Kubernetes, cloud platforms

### Documentation Generation
- **Interactive Docs**: Create web-based documentation with live examples
- **API References**: Generate comprehensive API documentation
- **Integration Guides**: Auto-create step-by-step integration tutorials
- **Architecture Diagrams**: Visual representations of AI system architecture

### System Orchestration
- **Workflow Automation**: Deploy n8n, Zapier, or custom automation workflows
- **MCP Server Setup**: Configure Model Context Protocol servers automatically
- **Monitoring Dashboards**: Set up metrics collection and alerting systems
- **Testing Frameworks**: Generate test suites and validation scripts

### Infrastructure as Code
- **Cloud Deployments**: Generate Terraform, CloudFormation, or Pulumi configurations
- **Container Orchestration**: Create Docker Compose and Kubernetes manifests
- **CI/CD Pipelines**: Set up automated testing and deployment workflows
- **Environment Management**: Configure dev, staging, and production environments

### Example: Generative Development Possibilities
```yaml
# Your APAI specification
openapia: "0.1.0"
info:
  title: "Customer Support AI"
models:
  - id: "support_model"
    type: "LLM"
    provider: "openai"
    name: "gpt-4"
tasks:
  - id: "handle_support"
    steps:
      - action: "generate"
        model: "support_model"
```

**The structured format helps developers build tools that could generate:**
- API clients and server implementations
- Interactive documentation and guides
- Automation workflows and integrations
- Infrastructure configurations and deployments
- Monitoring dashboards and testing frameworks
- Development tools and SDKs

*The possibilities are limited only by the creativity of the developer community!*

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/FabioGuin/APAI.git
cd APAI
```

### 2. Set Up Environment Variables (Security)

```bash
# Copy the template and fill in your actual values
cp .env.example .env
# Edit .env with your API keys and credentials
# NEVER commit the .env file to version control
```

### 3. Create Your First APAI Specification

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
python validators/python/apai_validator.py validate spec/my-ai-system.yaml

# Hierarchical validation (with inheritance)
python validators/python/apai_validator.py validate spec/my-ai-system.yaml --hierarchical

# Using JavaScript validator
node validators/javascript/cli.js validate spec/my-ai-system.yaml

# Using PHP validator
php validators/php/cli.php validate spec/my-ai-system.yaml

# Using Go validator
go run validators/go/cli.go validate spec/my-ai-system.yaml
```

## Examples

Check out our [examples directory](examples/) for complete implementations organized by category and complexity:

### Core AI Examples
- [Customer Support AI](examples/core/customer-support.yaml) - E-commerce assistant with order management and multilingual support
- [Content Moderator](examples/core/content-moderator.yaml) - AI-powered content filtering and safety enforcement
- [Multilingual Chatbot](examples/core/multilingual-chatbot.yaml) - Multi-language conversation support

### Multi-Agent System Examples
- [Multi-Agent Customer Support](examples/multi-agent/multi-agent-customer-support.yaml) - Complete multi-agent system using hierarchical composition
- [Sentiment Analysis Agent](examples/agents/sentiment-analyzer.yaml) - Specialized agent for sentiment analysis

### Automation Integration Examples
- [E-commerce with n8n](examples/automation/ecommerce-automation.yaml) - Complete order processing with automation workflows
- [Customer Support with Zapier](examples/automation/zapier-automation.yaml) - Simple webhook-based automation integration
- [MCP Integration](examples/automation/mcp-integration.yaml) - Model Context Protocol server integration

### Templates
- [Basic Template](examples/templates/basic-template.yaml) - Minimal starting template for new specifications
- [Complete JSON Example](examples/openapia-0.1-example.json) - Machine-readable reference for tools and SDKs

See the [Examples README](examples/README.md) for detailed descriptions and learning paths.

## Hierarchical Composition

APAI supports hierarchical composition for complex organizational structures:

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

The APAI specification is available in multiple formats:

- **YAML**: `spec/openapia-0.1.yaml` - Official specification (human-readable)
- **JSON Example**: `examples/openapia-0.1-example.json` - Complete working example for tools and SDKs
- **JSON Schema**: `spec/schemas/openapia-0.1-schema.json` - Validation schema
- **Examples**: `examples/*.yaml` - Real-world use case examples

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

- **[Documentation Index](docs/README.md)** - Complete documentation overview and learning paths
- **[Getting Started Guide](docs/getting-started.md)** - Step-by-step tutorial for new users
- **[APAI 0.1 Specification](docs/apai-0.1-specification.md)** - Complete formal specification documentation
- **[Hierarchical Composition Guide](docs/hierarchical-composition.md)** - Advanced inheritance and composition patterns
- **[Multi-Agent Systems Guide](docs/multi-agent-systems.md)** - Building multi-agent systems using existing APAI features
- **[Automation Integration Guide](docs/automation-integration.md)** - Integrating with external automation platforms
- **[Security Best Practices](docs/security-best-practices.md)** - Managing secrets and environment variables securely
- **[AI Security Considerations](docs/ai-security-considerations.md)** - Comprehensive AI-specific security framework
- **[Security Vulnerabilities](docs/security-vulnerabilities.md)** - Common AI security vulnerabilities and prevention

## Tools and Libraries

### Validators
- [Python Validator](validators/python/) - Full-featured validation library
- [PHP Validator](validators/php/) - PHP validation with Symfony YAML
- [JavaScript Validator](validators/javascript/) - Node.js and browser support
- [Go Validator](validators/go/) - High-performance validation

### Community Tools
- *The APAI community is encouraged to build generative tools and integrations*
- *Examples: Code generators, documentation tools, deployment automation, testing frameworks*

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

APAI 0.1 includes integration with external automation platforms, allowing AI systems to trigger and coordinate with automation workflows without tight coupling.

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
- **Platform Independent**: Works with any automation platform
- **Monitored**: Built-in health checks and metrics
- **Secure**: Configurable authentication and validation

### Documentation

- **[Automation Integration Guide](docs/automation-integration.md)** - Complete guide with examples and best practices
- **[E-commerce Example](examples/automation/ecommerce-automation.yaml)** - Full n8n integration example
- **[Zapier Example](examples/automation/zapier-automation.yaml)** - Simple webhook integration example

## Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Process

1. **Issues**: Open an issue to discuss changes
2. **Pull Requests**: Submit PRs for implementation
3. **Review**: Maintainer review with community feedback
4. **Decision**: Transparent decision-making process

## Governance

APAI is currently in **Bootstrap Phase** with Fabio Guin as the initial maintainer. We plan to transition to a community-driven governance model as the project grows.

See [GOVERNANCE.md](GOVERNANCE.md) for details.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Whitepapers

Comprehensive whitepapers providing detailed analysis and insights into APAI:

- **[APAI Whitepaper (English)](APAI_Whitepaper_EN.md)** - Complete technical and business analysis of APAI's capabilities, architecture, and potential
- **[APAI Whitepaper (Italiano)](APAI_Whitepaper_IT.md)** - Analisi completa tecnica e business delle capacità, architettura e potenziale di APAI

### Whitepaper Contents

Both whitepapers cover:
- **Executive Summary** - Overview of APAI's value proposition
- **AI Standardization Challenge** - Current industry challenges and needs
- **Core Architecture** - Detailed technical specifications and features
- **Advanced Capabilities** - Multi-agent systems, automation integration, MCP support
- **Generative Development Potential** - Code generation, documentation, and tooling opportunities
- **Implementation Guide** - Getting started and adoption strategies
- **Governance & Community** - Project governance and contribution guidelines
- **Future Roadmap** - Long-term vision and development plans

Perfect for:
- **Technical Leaders** - Understanding APAI's technical capabilities
- **Business Stakeholders** - Evaluating APAI's business value
- **Developers** - Learning implementation approaches and best practices
- **Decision Makers** - Making informed adoption decisions

## Community

- **Discussions**: [GitHub Discussions](https://github.com/FabioGuin/APAI/discussions)
- **Issues**: [GitHub Issues](https://github.com/FabioGuin/APAI/issues)

## Acknowledgments

APAI is inspired by:
- [ISO/IEC 42001](https://www.iso.org/standard/81231.html) for AI governance
- [LangChain](https://github.com/langchain-ai/langchain) for AI orchestration
- [IEEE AI-ESTATE](https://standards.ieee.org/standard/1232-2010.html) for AI component formalization

---

**Note**: This project is in active development. The specification may change between versions. Please check the [CHANGELOG](CHANGELOG.md) for breaking changes.
