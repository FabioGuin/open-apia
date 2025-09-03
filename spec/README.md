# OpenAPIA Specification Files

This directory contains the OpenAPIA specification in different formats and related files.

## Structure

```
spec/
├── openapia-0.1.yaml              # Main specification (human-readable)
├── examples/                       # Example specifications
│   ├── openapia-0.1-example.json  # Complete example in JSON format
│   ├── content-moderator.yaml     # Content moderation example
│   ├── customer-support.yaml      # Customer support example
│   ├── multilingual-chatbot.yaml  # Multilingual chatbot example
│   ├── mcp-integration.yaml       # Model Context Protocol integration
│   ├── ecommerce-automation.yaml  # E-commerce with n8n automation
│   └── zapier-automation.yaml     # Customer support with Zapier
└── schemas/                        # Validation schemas
    └── openapia-0.1-schema.json   # JSON Schema for validation
```

## File Types

### 1. Main Specification
- **`openapia-0.1.yaml`** - The official OpenAPIA 0.1 specification in YAML format
  - Human-readable format
  - Defines the structure and requirements
  - Reference for implementers

### 2. Example Specifications
- **`examples/openapia-0.1-example.json`** - Complete working example in JSON format
  - Machine-readable format for tools and SDKs
  - Contains realistic data for all sections
  - Demonstrates proper usage

- **`examples/*.yaml`** - Domain-specific examples
  - Real-world use cases
  - Different complexity levels
  - Reference implementations

#### Core AI Examples
- **`content-moderator.yaml`** - AI-powered content filtering system
- **`customer-support.yaml`** - E-commerce customer support assistant
- **`multilingual-chatbot.yaml`** - Multi-language chatbot implementation

#### Integration Examples
- **`mcp-integration.yaml`** - Model Context Protocol server integration
- **`ecommerce-automation.yaml`** - Complete e-commerce system with n8n automation workflows
- **`zapier-automation.yaml`** - Customer support with Zapier webhook automation

### 3. Validation Schemas
- **`schemas/openapia-0.1-schema.json`** - JSON Schema for validation
  - Defines validation rules
  - Used by validators to check specification compliance
  - Machine-readable validation rules

## Usage

### For Developers
- Use `openapia-0.1.yaml` to understand the specification
- Use `examples/openapia-0.1-example.json` as a template for your implementations
- Use `schemas/openapia-0.1-schema.json` for programmatic validation

### For Tools and SDKs
- Parse `examples/openapia-0.1-example.json` for structure understanding
- Use `schemas/openapia-0.1-schema.json` for validation
- Reference `openapia-0.1.yaml` for documentation

### For Validators
- Use `schemas/openapia-0.1-schema.json` as the validation schema
- Test against `examples/*.yaml` and `examples/*.json` files
