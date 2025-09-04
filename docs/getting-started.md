# Getting Started with OpenAPIA

This guide will help you get started with OpenAPIA, the open standard for describing AI systems.

## Table of Contents

1. [What is OpenAPIA?](#what-is-openapia)
   - [Key Features](#key-features)

2. [Quick Start](#quick-start)
   - [1. Clone the Repository](#1-clone-the-repository)
   - [2. Create Your First Specification](#2-create-your-first-specification)
   - [3. Validate Your Specification](#3-validate-your-specification)

3. [Understanding the Specification Structure](#understanding-the-specification-structure)
   - [Basic Structure](#basic-structure)
   - [Required Fields](#required-fields)
   - [Optional Fields](#optional-fields)

4. [Core Concepts](#core-concepts)
   - [Models](#models)
   - [Prompts](#prompts)
   - [Constraints](#constraints)
   - [Tasks](#tasks)
   - [Context](#context)
   - [Evaluation](#evaluation)

5. [Advanced Features](#advanced-features)
   - [Hierarchical Composition](#hierarchical-composition)
   - [Multi-Agent Systems](#multi-agent-systems)
   - [Automation Integration](#automation-integration)

6. [Best Practices](#best-practices)
   - [Security](#security)
   - [Performance](#performance)
   - [Maintainability](#maintainability)

7. [Next Steps](#next-steps)

---

## What is OpenAPIA?

OpenAPIA is an open, vendor-agnostic standard for describing, documenting, and validating artificial intelligence systems. It provides a machine-readable format to specify AI models, prompts, constraints, workflows, and evaluation metrics.

### Key Features

- **AI-Native Design**: Models, prompts, and constraints as first-class entities
- **Vendor Agnostic**: Works with any AI provider (OpenAI, Anthropic, Google, etc.)
- **Multi-Modal Support**: LLM, Vision, Audio, and Multimodal AI systems
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

### 2. Create Your First Specification

Create a file called `my-ai-system.yaml`:

```yaml
openapia: "0.1.0"

info:
  title: "My AI Assistant"
  version: "1.0.0"
  description: "A helpful AI assistant"

models:
  - id: "main_model"
    type: "LLM"
    provider: "openai"
    name: "gpt-4"
    purpose: "conversation"

prompts:
  - id: "system_prompt"
    role: "system"
    template: "You are a helpful AI assistant"

constraints:
  - id: "safety"
    rule: "output NOT contains harmful_content"
    severity: "critical"

tasks:
  - id: "handle_query"
    description: "Process user queries"
    steps:
      - action: "generate"
        model: "main_model"
        prompt: "system_prompt"

context:
  memory:
    type: "session"
    retention: "7d"

evaluation:
  metrics:
    - name: "accuracy"
      target: 0.9
```

### 3. Validate Your Specification

```bash
# Simple validation
python validators/python/openapia_validator.py my-ai-system.yaml

# Hierarchical validation (with inheritance)
python validators/python/openapia_validator.py validate my-ai-system.yaml --hierarchical

# Using JavaScript validator
node validators/javascript/cli.js validate my-ai-system.yaml

# Validate entire hierarchy
python validators/python/openapia_validator.py validate . --recursive
```

## Core Concepts

### Models

Models define the AI systems you're using:

```yaml
models:
  - id: "language_model"
    type: "LLM"
    provider: "openai"
    name: "gpt-4"
    purpose: "conversation"
    parameters:
      temperature: 0.7
      max_tokens: 500
```

### Prompts

Prompts define how to interact with your models:

```yaml
prompts:
  - id: "system_prompt"
    role: "system"
    template: "You are a helpful assistant for {{company_name}}"
    variables:
      company_name:
        type: "string"
        required: true
        default: "My Company"
```

### Constraints

Constraints ensure your AI system behaves safely and ethically:

```yaml
constraints:
  - id: "safety"
    rule: "output NOT contains harmful_content"
    severity: "critical"
    enforcement: "automatic"
```

### Tasks

Tasks define the workflows and business logic:

```yaml
tasks:
  - id: "handle_query"
    description: "Process user queries safely"
    steps:
      - name: "analyze"
        action: "analyze"
        model: "language_model"
      - name: "generate"
        action: "generate"
        model: "language_model"
        prompt: "system_prompt"
```

### Context

Context manages state and memory:

```yaml
context:
  memory:
    type: "session"
    retention: "7d"
    scope: "per_user"
```

### Evaluation

Evaluation defines metrics and tests:

```yaml
evaluation:
  metrics:
    - name: "accuracy"
      target: 0.9
      measurement:
        method: "human_evaluation"
        frequency: "weekly"
```

## Examples

Check out our [examples directory](../examples/) for complete implementations:

- [Customer Support AI](../examples/core/customer-support.yaml) - E-commerce assistant
- [Content Moderator](../examples/core/content-moderator.yaml) - AI-powered content filtering
- [Multilingual Chatbot](../examples/core/multilingual-chatbot.yaml) - Multi-language support

## Tools

### Validators

- **Python**: Full-featured validation library
- **PHP**: PHP validation with Symfony YAML
- **JavaScript**: Node.js and browser support
- **Go**: High-performance validation

### CLI Usage

Each validator includes a command-line interface:

```bash
# Python
python validators/python/openapia_validator.py spec.yaml

# PHP
php validators/php/cli.php validate spec.yaml

# JavaScript
node validators/javascript/cli.js validate spec.yaml

# Go
go run validators/go/cli.go validate spec.yaml
```

## Best Practices

### 1. Start Simple

Begin with a basic specification and add complexity gradually.

### 2. Use Clear Names

Choose descriptive IDs and names for your models, prompts, and tasks.

### 3. Define Constraints Early

Always include safety and ethical constraints from the start.

### 4. Test Regularly

Use the evaluation section to define and run tests.

### 5. Document Everything

Include clear descriptions for all components.

## Common Patterns

### Multi-Model Systems

```yaml
models:
  - id: "main_llm"
    type: "LLM"
    provider: "openai"
    name: "gpt-4"
  - id: "sentiment_analyzer"
    type: "Classification"
    provider: "huggingface"
    name: "sentiment-model"
```

### Conditional Workflows

```yaml
tasks:
  - id: "process_message"
    steps:
      - name: "analyze_sentiment"
        action: "analyze"
        model: "sentiment_analyzer"
        conditions:
          - if: "sentiment == 'negative'"
            then: "escalate"
      - name: "generate_response"
        action: "generate"
        model: "main_llm"
```

### Multi-Language Support

```yaml
prompts:
  - id: "multilingual_prompt"
    role: "system"
    language: "auto"
    template: "Respond in {{language}}"
    variables:
      language:
        type: "string"
        enum: ["en", "es", "fr", "de"]
```

## Troubleshooting

### Common Issues

1. **Missing Required Fields**: Check that all required sections are present
2. **Invalid References**: Ensure model and prompt IDs are correctly referenced
3. **Constraint Conflicts**: Verify that constraints don't conflict with each other
4. **Performance Issues**: Monitor response times and token usage

### Getting Help

- **Documentation**: Check the [specification](../specification.md)
- **Examples**: See [examples](../examples/)
- **Issues**: Search [GitHub Issues](https://github.com/FabioGuin/OpenAPIA/issues)
- **Discussions**: Use [GitHub Discussions](https://github.com/FabioGuin/OpenAPIA/discussions)

## Hierarchical Composition

For complex organizational structures, OpenAPIA supports hierarchical composition:

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

### Hierarchical Commands

```bash
# Show hierarchy tree
python validators/python/openapia_validator.py tree openapia.yaml

# Merge specifications
python validators/python/openapia_validator.py merge output.yaml spec1.yaml spec2.yaml

# Using other validators
node validators/javascript/cli.js tree openapia.yaml
php validators/php/cli.php tree openapia.yaml
go run validators/go/cli.go tree openapia.yaml
```

For detailed information, see the [Hierarchical Composition Guide](hierarchical-composition.md).

## Next Steps

1. **Explore Examples**: Look at the example specifications
2. **Build Your System**: Create your own OpenAPIA specification
3. **Try Hierarchical**: Experiment with inheritance and composition
4. **Validate**: Use the validators to check your work
5. **Contribute**: Help improve OpenAPIA by contributing

## Resources

- [OpenAPIA Specification](../spec/openapia-0.1.yaml)
- [Examples](../examples/)
- [Validators](../validators/)
- [Contributing Guide](../CONTRIBUTING.md)
