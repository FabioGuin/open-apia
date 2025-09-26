# APAI Specification Files

This directory contains the Architecture Protocol for Artificial Intelligence specification in different formats and related files.

## Structure

```
spec/
├── apai-0.1.yaml              # Main specification (human-readable)
└── schemas/                        # Validation schemas
    └── apai-0.1-schema.json   # JSON Schema for validation
```

## File Types

### 1. Main Specification
- **`apai-0.1.yaml`** - The official APAI 0.1 specification in YAML format
  - Human-readable format
  - Defines the structure and requirements
  - Reference for implementers

### 2. Example Specifications
Examples are now located in the root `examples/` directory. See the [Examples README](../examples/README.md) for complete documentation.

- **`../examples/apai-0.1-example.json`** - Complete working example in JSON format
  - Machine-readable format for tools and SDKs
  - Contains realistic data for all sections
  - Demonstrates proper usage

- **`../examples/*.yaml`** - Domain-specific examples
  - Real-world use cases
  - Different complexity levels
  - Reference implementations



### 3. Validation Schemas
- **`schemas/apai-0.1-schema.json`** - JSON Schema for validation
  - Defines validation rules
  - Used by validators to check specification compliance
  - Machine-readable validation rules

## Usage

### For Developers
- Use `apai-0.1.yaml` to understand the specification
- Start with `../examples/templates/basic-template.yaml` for new projects
- Use `../examples/apai-0.1-example.json` as a machine-readable reference
- Use `schemas/apai-0.1-schema.json` for programmatic validation
- Explore `../examples/README.md` for detailed examples and learning paths

### For Tools and SDKs
- Parse `../examples/apai-0.1-example.json` for structure understanding
- Use `schemas/apai-0.1-schema.json` for validation
- Reference `apai-0.1.yaml` for documentation
- Test against examples in `../examples/` directory

### For Validators
- Use `schemas/apai-0.1-schema.json` as the validation schema
- Test against all `../examples/**/*.yaml` and `../examples/*.json` files
- Validate hierarchical compositions using multi-agent examples
