# APAI Go Validator

A high-performance Go validator for APAI specifications.

## Features

- **High Performance**: Fast validation with minimal memory usage
- **Full Validation**: Validates all APAI specification sections
- **Hierarchical Composition**: Support for inheritance and composition of specifications
- **Cross-Validation**: Checks references between models, prompts, and tasks
- **Multiple Formats**: Supports YAML and JSON specifications
- **CLI Interface**: Command-line tool for validation with hierarchical support
- **Comprehensive Testing**: Full test suite with Go testing
- **Static Binary**: Single executable with no dependencies

## Installation

### Using go install

```bash
go install github.com/FabioGuin/APAI/validators/go@latest
```

### Manual Installation

1. Clone the repository
2. Build the validator:
   ```bash
   go build -o apai-validator .
   ```

## Usage

### CLI Usage

```bash
# Basic validation
go run cli.go validate spec.yaml

# Hierarchical validation
go run cli.go validate spec.yaml --hierarchical

# Show hierarchy tree
go run cli.go tree spec.yaml

# Merge specifications
go run cli.go merge output.yaml spec1.yaml spec2.yaml
```

### Programmatic Usage

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    // Create validator instance
    validator := NewAPAIValidator()
    
    // Validate a file
    isValid, err := validator.ValidateFile("spec.yaml")
    if err != nil {
        log.Fatal(err)
    }
    
    // Validate an object
    spec := map[string]interface{}{
        "apai": "0.1.0",
        "info": map[string]interface{}{
            "title":       "My AI System",
            "version":     "1.0.0",
            "description": "A test AI system",
            "author":      "AI Team",
            "license":     "MIT",
        },
        "models": []interface{}{
            map[string]interface{}{
                "id":       "main_model",
                "type":     "LLM",
                "provider": "openai",
                "name":     "gpt-4",
                "purpose":  "conversation",
            },
        },
        "prompts": []interface{}{
            map[string]interface{}{
                "id":       "system_prompt",
                "role":     "system",
                "template": "You are a helpful AI assistant",
            },
        },
        "constraints": []interface{}{
            map[string]interface{}{
                "id":       "safety",
                "rule":     "output NOT contains harmful_content",
                "severity": "critical",
            },
        },
        "tasks": []interface{}{
            map[string]interface{}{
                "id":          "handle_query",
                "description": "Process user queries",
            },
        },
        "context": map[string]interface{}{
            "memory": map[string]interface{}{
                "type":      "session",
                "retention": "7d",
            },
        },
        "evaluation": map[string]interface{}{
            "metrics": []interface{}{
                map[string]interface{}{
                    "name":   "accuracy",
                    "target": 0.9,
                },
            },
        },
    }
    
    isValid = validator.ValidateSpec(spec)
    
    // Get results
    errors := validator.GetErrors()
    warnings := validator.GetWarnings()
    
    // Print results
    validator.PrintResults()
    
    // Get results as struct
    results := validator.GetResults()
    fmt.Printf("Valid: %t\n", results.Valid)
}
```

## Validation Rules

### Required Sections

The validator checks for the following required sections:

- `apai` - Specification version
- `info` - System metadata
- `models` - AI models
- `prompts` - Prompt templates
- `constraints` - Safety and ethical constraints
- `tasks` - Workflow definitions
- `context` - State management
- `evaluation` - Metrics and testing

### Model Validation

- Required fields: `id`, `type`, `provider`, `name`, `purpose`
- Valid types: `LLM`, `Vision`, `Audio`, `Multimodal`, `Classification`, `Embedding`
- Unique IDs across all models

### Prompt Validation

- Required fields: `id`, `role`, `template`
- Valid roles: `system`, `user`, `assistant`
- Unique IDs across all prompts

### Constraint Validation

- Required fields: `id`, `rule`, `severity`
- Valid severities: `low`, `medium`, `high`, `critical`
- Unique IDs across all constraints

### Task Validation

- Required fields: `id`, `description`
- Unique IDs across all tasks
- Cross-validation of model and prompt references

### Cross-Validation

The validator performs cross-validation to ensure:

- Referenced models exist in the models section
- Referenced prompts exist in the prompts section
- All references are valid and consistent

## Error Handling

### Error Types

- **Errors**: Critical issues that prevent validation
- **Warnings**: Non-critical issues that should be addressed

### Error Examples

```go
// Missing required section
"Missing required section: models"

// Missing required field
"Model 0 missing required field: type"

// Duplicate ID
"Duplicate model ID: my_model"

// Invalid reference
"Task references unknown model: nonexistent_model"

// Invalid value
"Invalid constraint severity: invalid_severity"
```

## Testing

Run the test suite:

```bash
# Run all tests
go test

# Run tests with coverage
go test -cover

# Run tests with verbose output
go test -v

# Run specific test
go test -run TestValidateSpec
```

## Development

### Requirements

- Go 1.19 or higher
- Go modules support

### Project Structure

```
validators/go/
├── validator.go          # Main validator implementation
├── cli.go               # CLI interface
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── README.md            # This file
└── validator_test.go    # Test suite
```

### Building

```bash
# Build for current platform
go build -o apai-validator .

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o apai-validator-linux .
GOOS=windows GOARCH=amd64 go build -o apai-validator.exe .
GOOS=darwin GOARCH=amd64 go build -o apai-validator-macos .
```

### Adding New Validation Rules

1. Add the validation logic to the appropriate method in `validator.go`
2. Add test cases to `validator_test.go`
3. Update documentation

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run the test suite
6. Submit a pull request

## API Reference

### APAIValidator

#### Constructor

```go
validator := NewAPAIValidator()
```

#### Methods

##### `ValidateFile(filePath string) (bool, error)`

Validates an APAI specification file.

**Parameters:**
- `filePath` (string): Path to the specification file

**Returns:** (bool, error)

##### `ValidateSpec(spec map[string]interface{}) bool`

Validates an APAI specification object.

**Parameters:**
- `spec` (map[string]interface{}): APAI specification object

**Returns:** bool

##### `GetErrors() []string`

Gets list of validation errors.

**Returns:** []string

##### `GetWarnings() []string`

Gets list of validation warnings.

**Returns:** []string

##### `PrintResults()`

Prints validation results to console.

**Returns:** void

##### `GetResults() ValidationResult`

Gets validation results as struct.

**Returns:** ValidationResult

### ValidationResult

```go
type ValidationResult struct {
    Valid    bool     `json:"valid"`
    Errors   []string `json:"errors"`
    Warnings []string `json:"warnings"`
}
```

## Performance

The Go validator is optimized for performance:

- **Fast parsing**: Uses efficient YAML and JSON parsers
- **Memory efficient**: Minimal memory allocation
- **Concurrent validation**: Can be used in concurrent environments
- **Static binary**: No runtime dependencies

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](../../LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/FabioGuin/APAI/issues)
- **Discussions**: [GitHub Discussions](https://github.com/FabioGuin/APAI/discussions)

## Changelog

### 0.1.0
- Initial release
- Basic validation for all APAI sections
- CLI interface
- Comprehensive test suite
- Cross-validation support
- High-performance implementation
