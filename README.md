# OpenAPIA - Open Artificial Intelligence Architecture

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/badge/version-0.1.0-green.svg)](https://github.com/FabioGuin/OpenAPIA/releases)

## Overview

OpenAPIA is an open, vendor-agnostic standard for describing, documenting, and validating artificial intelligence systems. Inspired by the success of OpenAPI for REST APIs, OpenAPIA provides a machine-readable format to specify AI models, prompts, constraints, workflows, and evaluation metrics.

## Key Features

- **AI-Native Design**: Models, prompts, and constraints as first-class entities
- **Vendor Agnostic**: Works with any AI provider (OpenAI, Anthropic, Google, etc.)
- **Multi-Modal Support**: LLM, Vision, Audio, and Multimodal AI systems
- **Built-in Ethics**: Mandatory fields for safety, bias prevention, and explainability
- **Comprehensive Evaluation**: Metrics for accuracy, performance, and safety
- **Extensible**: Support for custom use cases and domain-specific requirements

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/FabioGuin/OpenAPIA.git
cd OpenAPIA
```

### 2. Create Your First OpenAPIA Specification

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

### 3. Validate Your Specification

```bash
# Using Python validator
python validators/python/openapia_validator.py spec/my-ai-system.yaml

# Using JavaScript validator
node validators/javascript/cli.js -f spec/my-ai-system.yaml
```

## Examples

Check out our [examples directory](spec/examples/) for complete implementations:

- [Customer Support AI](spec/examples/customer-support.yaml) - E-commerce assistant
- [Content Moderator](spec/examples/content-moderator.yaml) - AI-powered content filtering
- [Multilingual Chatbot](spec/examples/multilingual-chatbot.yaml) - Multi-language support

## Specification

The OpenAPIA specification is defined in YAML format and includes these core sections:

- **`openapia`** - Specification version
- **`info`** - System metadata and AI-specific information
- **`models`** - AI models with capabilities, limits, and costs
- **`prompts`** - Structured prompts with variables and configuration
- **`constraints`** - Safety, ethical, and operational constraints
- **`tasks`** - Declarative workflows and business logic
- **`context`** - State management and memory configuration
- **`evaluation`** - Metrics, tests, and performance monitoring

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

- [x] **v0.1.0** - Core specification and validators (Python, PHP, JavaScript, Go)
- [ ] **v0.2.0** - Multi-agent support and advanced features
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
