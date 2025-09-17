# OpenAPIA PHP Validator

A comprehensive PHP validator for OpenAPIA specifications.

## Features

- **Full Validation**: Validates all OpenAPIA specification sections
- **Hierarchical Composition**: Support for inheritance and composition of specifications
- **Cross-Validation**: Checks references between models, prompts, and tasks
- **Multiple Formats**: Supports YAML and JSON specifications
- **CLI Interface**: Command-line tool for validation with hierarchical support
- **Comprehensive Testing**: Full test suite with PHPUnit
- **PSR-4 Autoloading**: Modern PHP standards compliance

## Installation

### Using Composer

```bash
composer require openapia/validator-php
```

### Manual Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   composer install
   ```

## Usage

### CLI Usage

```bash
# Basic validation
php cli.php validate spec.yaml

# Hierarchical validation
php cli.php validate spec.yaml --hierarchical

# Show hierarchy tree
php cli.php tree spec.yaml

# Merge specifications
php cli.php merge output.yaml spec1.yaml spec2.yaml
```

### Programmatic Usage

```php
<?php

use OpenAPIA\OpenAPIAValidator;

// Create validator instance
$validator = new OpenAPIAValidator();

// Basic validation
$isValid = $validator->validateFile('spec.yaml');

// Hierarchical validation with inheritance
$isValid = $validator->validateWithInheritance('spec.yaml');

// Show hierarchy tree
$validator->printHierarchyTree('spec.yaml');

// Merge specifications
$specs = [
    $validator->loadSpec('global.yaml'),
    $validator->loadSpec('team.yaml')
];
$validator->mergeSpecifications($specs, 'merged.yaml', 'yaml');

// Get results
$errors = $validator->getErrors();
$warnings = $validator->getWarnings();
```

## Validation Rules

### Required Sections

The validator checks for the following required sections:

- `openapia` - Specification version
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

```php
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
composer test

# Run tests with coverage
composer test-coverage

# Run static analysis
composer phpstan

# Check code style
composer cs-check

# Fix code style issues
composer cs-fix
```

## Development

### Requirements

- PHP 8.0 or higher
- Composer
- PHPUnit (for testing)
- PHPStan (for static analysis)
- PHP CodeSniffer (for code style)

### Project Structure

```
validators/php/
├── OpenAPIAValidator.php    # Main validator class
├── cli.php                  # CLI interface
├── composer.json            # Dependencies
├── README.md               # This file
└── tests/
    └── OpenAPIAValidatorTest.php  # Test suite
```

### Adding New Validation Rules

1. Add the validation logic to the appropriate method in `OpenAPIAValidator`
2. Add test cases to `OpenAPIAValidatorTest`
3. Update documentation

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run the test suite
6. Submit a pull request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](../../LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/FabioGuin/OpenAPIA/issues)
- **Discussions**: [GitHub Discussions](https://github.com/FabioGuin/OpenAPIA/discussions)

## Changelog

### 0.1.0
- Initial release
- Basic validation for all OpenAPIA sections
- CLI interface
- Comprehensive test suite
- Cross-validation support
