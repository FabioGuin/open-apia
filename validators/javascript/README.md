# OpenAPIA JavaScript Validator

A comprehensive JavaScript/Node.js validator for OpenAPIA specifications.

## Features

- **Full Validation**: Validates all OpenAPIA specification sections
- **Hierarchical Composition**: Support for inheritance and composition of specifications
- **Cross-Validation**: Checks references between models, prompts, and tasks
- **Multiple Formats**: Supports YAML and JSON specifications
- **CLI Interface**: Command-line tool for validation with hierarchical support
- **Browser Support**: Works in both Node.js and browser environments
- **TypeScript Support**: Full TypeScript definitions included
- **Comprehensive Testing**: Full test suite with Jest

## Installation

### Using npm

```bash
npm install openapia-validator-js
```

### Using yarn

```bash
yarn add openapia-validator-js
```

### Manual Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```

## Usage

### CLI Usage

```bash
# Basic validation
node cli.js validate spec.yaml

# Hierarchical validation
node cli.js validate spec.yaml --hierarchical

# Show hierarchy tree
node cli.js tree spec.yaml

# Merge specifications
node cli.js merge output.yaml spec1.yaml spec2.yaml
```

### Programmatic Usage

```javascript
const OpenAPIAValidator = require('openapia-validator-js');

// Create validator instance
const validator = new OpenAPIAValidator();

// Validate a file
const isValid = await validator.validateFile('spec.yaml');

// Validate an object
const spec = {
  openapia: '0.1.0',
  info: {
    title: 'My AI System',
    version: '1.0.0',
    description: 'A test AI system',
    author: 'AI Team',
    license: 'MIT'
  },
  models: [
    {
      id: 'main_model',
      type: 'LLM',
      provider: 'openai',
      name: 'gpt-4',
      purpose: 'conversation'
    }
  ],
  prompts: [
    {
      id: 'system_prompt',
      role: 'system',
      template: 'You are a helpful AI assistant'
    }
  ],
  constraints: [
    {
      id: 'safety',
      rule: 'output NOT contains harmful_content',
      severity: 'critical'
    }
  ],
  tasks: [
    {
      id: 'handle_query',
      description: 'Process user queries'
    }
  ],
  context: {
    memory: {
      type: 'session',
      retention: '7d'
    }
  },
  evaluation: {
    metrics: [
      {
        name: 'accuracy',
        target: 0.9
      }
    ]
  }
};

const isValid = validator.validateSpec(spec);

// Get results
const errors = validator.getErrors();
const warnings = validator.getWarnings();

// Print results
validator.printResults();

// Get results as object
const results = validator.getResults();
```

### Browser Usage

```html
<!DOCTYPE html>
<html>
<head>
    <title>OpenAPIA Validator</title>
</head>
<body>
    <script src="node_modules/openapia-validator-js/dist/index.js"></script>
    <script>
        const validator = new OpenAPIAValidator();
        const spec = { /* your specification */ };
        const isValid = validator.validateSpec(spec);
        validator.printResults();
    </script>
</body>
</html>
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

```javascript
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
npm test

# Run tests with coverage
npm run test:coverage

# Run tests in watch mode
npm run test:watch

# Lint code
npm run lint

# Fix linting issues
npm run lint:fix
```

## Development

### Requirements

- Node.js 16.0 or higher
- npm or yarn
- Jest (for testing)
- ESLint (for code style)

### Project Structure

```
validators/javascript/
├── src/
│   └── OpenAPIAValidator.js    # Main validator class
├── cli.js                      # CLI interface
├── package.json                # Dependencies and scripts
├── README.md                   # This file
└── tests/
    └── OpenAPIAValidator.test.js  # Test suite
```

### Building

```bash
# Build for production
npm run build

# Build in development mode
npm run dev
```

### Adding New Validation Rules

1. Add the validation logic to the appropriate method in `OpenAPIAValidator`
2. Add test cases to `OpenAPIAValidator.test.js`
3. Update documentation

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run the test suite
6. Submit a pull request

## API Reference

### OpenAPIAValidator

#### Constructor

```javascript
const validator = new OpenAPIAValidator();
```

#### Methods

##### `validateFile(filePath)`

Validates an OpenAPIA specification file.

**Parameters:**
- `filePath` (string): Path to the specification file

**Returns:** Promise<boolean>

##### `validateSpec(spec)`

Validates an OpenAPIA specification object.

**Parameters:**
- `spec` (object): OpenAPIA specification object

**Returns:** boolean

##### `getErrors()`

Gets list of validation errors.

**Returns:** string[]

##### `getWarnings()`

Gets list of validation warnings.

**Returns:** string[]

##### `printResults()`

Prints validation results to console.

**Returns:** void

##### `getResults()`

Gets validation results as object.

**Returns:** object

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](../../LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/FabioGuin/OpenAPIA/issues)
- **Discussions**: [GitHub Discussions](https://github.com/FabioGuin/OpenAPIA/discussions)
- **Email**: random@starzero.it

## Changelog

### 0.1.0
- Initial release
- Basic validation for all OpenAPIA sections
- CLI interface
- Comprehensive test suite
- Cross-validation support
- Browser support
- TypeScript definitions
