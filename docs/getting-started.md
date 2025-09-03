# Getting Started with OpenAPIA

This guide will help you get started with OpenAPIA, the open standard for describing AI systems.

## What is OpenAPIA?

OpenAPIA is an open, vendor-agnostic standard for describing, documenting, and validating artificial intelligence systems. It provides a machine-readable format to specify AI models, prompts, constraints, workflows, and evaluation metrics.

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
# Using Python validator
python validators/python/openapia_validator.py my-ai-system.yaml

# Using JavaScript validator
node validators/javascript/cli.js -f my-ai-system.yaml
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

Check out our [examples directory](../spec/examples/) for complete implementations:

- [Customer Support AI](../spec/examples/customer-support.yaml) - E-commerce assistant
- [Content Moderator](../spec/examples/content-moderator.yaml) - AI-powered content filtering
- [Multilingual Chatbot](../spec/examples/multilingual-chatbot.yaml) - Multi-language support

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
php validators/php/cli.php -f spec.yaml

# JavaScript
node validators/javascript/cli.js -f spec.yaml

# Go
go run validators/go/. -f spec.yaml
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
- **Examples**: See [examples](../spec/examples/)
- **Issues**: Search [GitHub Issues](https://github.com/FabioGuin/OpenAPIA/issues)
- **Discussions**: Use [GitHub Discussions](https://github.com/FabioGuin/OpenAPIA/discussions)

## Next Steps

1. **Explore Examples**: Look at the example specifications
2. **Build Your System**: Create your own OpenAPIA specification
3. **Validate**: Use the validators to check your work
4. **Contribute**: Help improve OpenAPIA by contributing

## Resources

- [OpenAPIA Specification](../spec/openapia-0.1.yaml)
- [Examples](../spec/examples/)
- [Validators](../validators/)
- [Contributing Guide](../CONTRIBUTING.md)
