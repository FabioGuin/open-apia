/**
 * OpenAPIA Validator - JavaScript Implementation
 * 
 * A comprehensive validator for OpenAPIA specifications.
 * 
 * @package OpenAPIA
 * @version 0.1.0
 * @license Apache-2.0
 */

const yaml = require('js-yaml');
const Ajv = require('ajv');
const addFormats = require('ajv-formats');

class OpenAPIAValidator {
    constructor() {
        this.errors = [];
        this.warnings = [];
        this.schemaVersion = '0.1.0';

        // Initialize JSON Schema validator
        this.ajv = new Ajv({ allErrors: true });
        addFormats(this.ajv);

        // Load OpenAPIA schema
        this.schema = this.loadSchema();
    }

    /**
     * Load the OpenAPIA JSON schema
     */
    loadSchema() {
        // In a real implementation, this would load from the schema file
        // For now, we'll define the basic structure
        return {
            type: 'object',
            required: ['openapia', 'info', 'models', 'prompts', 'constraints', 'tasks', 'context', 'evaluation'],
            properties: {
                openapia: {
                    type: 'string',
                    pattern: '^0\\.1\\.\\d+$'
                },
                info: {
                    type: 'object',
                    required: ['title', 'version', 'description', 'author', 'license'],
                    properties: {
                        title: { type: 'string' },
                        version: { type: 'string' },
                        description: { type: 'string' },
                        author: { type: 'string' },
                        license: { type: 'string' }
                    }
                },
                models: {
                    type: 'array',
                    minItems: 1,
                    items: {
                        type: 'object',
                        required: ['id', 'type', 'provider', 'name', 'purpose'],
                        properties: {
                            id: { type: 'string' },
                            type: {
                                type: 'string',
                                enum: ['LLM', 'Vision', 'Audio', 'Multimodal', 'Classification', 'Embedding']
                            },
                            provider: { type: 'string' },
                            name: { type: 'string' },
                            purpose: { type: 'string' }
                        }
                    }
                },
                prompts: {
                    type: 'array',
                    items: {
                        type: 'object',
                        required: ['id', 'role', 'template'],
                        properties: {
                            id: { type: 'string' },
                            role: {
                                type: 'string',
                                enum: ['system', 'user', 'assistant']
                            },
                            template: { type: 'string' }
                        }
                    }
                },
                constraints: {
                    type: 'array',
                    items: {
                        type: 'object',
                        required: ['id', 'rule', 'severity'],
                        properties: {
                            id: { type: 'string' },
                            rule: { type: 'string' },
                            severity: {
                                type: 'string',
                                enum: ['low', 'medium', 'high', 'critical']
                            }
                        }
                    }
                },
                tasks: {
                    type: 'array',
                    items: {
                        type: 'object',
                        required: ['id', 'description'],
                        properties: {
                            id: { type: 'string' },
                            description: { type: 'string' }
                        }
                    }
                },
                context: {
                    type: 'object'
                },
                evaluation: {
                    type: 'object'
                }
            }
        };
    }

    /**
     * Validate an OpenAPIA specification file
     */
    async validateFile(filePath) {
        try {
            const fs = require('fs').promises;
            const content = await fs.readFile(filePath, 'utf8');

            let spec;
            if (filePath.endsWith('.yaml') || filePath.endsWith('.yml')) {
                spec = yaml.load(content);
            } else if (filePath.endsWith('.json')) {
                spec = JSON.parse(content);
            } else {
                this.errors.push(`Unsupported file format: ${filePath}`);
                return false;
            }

            return this.validateSpec(spec);
        } catch (error) {
            if (error.code === 'ENOENT') {
                this.errors.push(`File not found: ${filePath}`);
            } else if (error.name === 'YAMLException') {
                this.errors.push(`YAML parsing error: ${error.message}`);
            } else if (error.name === 'SyntaxError') {
                this.errors.push(`JSON parsing error: ${error.message}`);
            } else {
                this.errors.push(`Unexpected error: ${error.message}`);
            }
            return false;
        }
    }

    /**
     * Validate an OpenAPIA specification object
     */
    validateSpec(spec) {
        this.errors = [];
        this.warnings = [];

        // Validate against JSON schema
        const validate = this.ajv.compile(this.schema);
        const valid = validate(spec);

        if (!valid) {
            validate.errors.forEach(error => {
                this.errors.push(`${error.instancePath}: ${error.message}`);
            });
        }

        // Custom validation
        this.validateRequiredSections(spec);
        this.validateOpenAPIAVersion(spec.openapia);
        this.validateInfo(spec.info);
        this.validateModels(spec.models);
        this.validatePrompts(spec.prompts);
        this.validateConstraints(spec.constraints);
        this.validateTasks(spec.tasks);
        this.validateContext(spec.context);
        this.validateEvaluation(spec.evaluation);

        // Cross-validation
        this.crossValidate(spec);

        return this.errors.length === 0;
    }

    /**
     * Validate that all required sections are present
     */
    validateRequiredSections(spec) {
        const requiredSections = [
            'openapia', 'info', 'models', 'prompts',
            'constraints', 'tasks', 'context', 'evaluation'
        ];

        requiredSections.forEach(section => {
            if (!(section in spec)) {
                this.errors.push(`Missing required section: ${section}`);
            }
        });
    }

    /**
     * Validate the OpenAPIA version
     */
    validateOpenAPIAVersion(version) {
        if (typeof version !== 'string') {
            this.errors.push('openapia version must be a string');
            return;
        }

        if (!version.startsWith('0.1')) {
            this.warnings.push(`Version ${version} may not be supported`);
        }
    }

    /**
     * Validate the info section
     */
    validateInfo(info) {
        if (!info || typeof info !== 'object') {
            this.errors.push('info must be an object');
            return;
        }

        const requiredFields = ['title', 'version', 'description', 'author', 'license'];
        requiredFields.forEach(field => {
            if (!(field in info)) {
                this.errors.push(`Missing required field in info: ${field}`);
            }
        });

        if (info.ai_metadata) {
            this.validateAIMetadata(info.ai_metadata);
        }
    }

    /**
     * Validate AI-specific metadata
     */
    validateAIMetadata(metadata) {
        if (!metadata.domain) {
            this.warnings.push('ai_metadata.domain is recommended');
        }

        if (metadata.complexity) {
            const validComplexities = ['low', 'medium', 'high'];
            if (!validComplexities.includes(metadata.complexity)) {
                this.errors.push(`Invalid complexity: ${metadata.complexity}`);
            }
        }
    }

    /**
     * Validate the models section
     */
    validateModels(models) {
        if (!Array.isArray(models)) {
            this.errors.push('models must be an array');
            return;
        }

        if (models.length === 0) {
            this.errors.push('At least one model is required');
            return;
        }

        const modelIds = new Set();
        models.forEach((model, index) => {
            if (typeof model !== 'object') {
                this.errors.push(`Model ${index} must be an object`);
                return;
            }

            // Validate required fields
            const requiredFields = ['id', 'type', 'provider', 'name', 'purpose'];
            requiredFields.forEach(field => {
                if (!(field in model)) {
                    this.errors.push(`Model ${index} missing required field: ${field}`);
                }
            });

            // Check for duplicate IDs
            if (model.id) {
                if (modelIds.has(model.id)) {
                    this.errors.push(`Duplicate model ID: ${model.id}`);
                }
                modelIds.add(model.id);
            }

            // Validate model type
            if (model.type) {
                const validTypes = ['LLM', 'Vision', 'Audio', 'Multimodal', 'Classification', 'Embedding'];
                if (!validTypes.includes(model.type)) {
                    this.warnings.push(`Unknown model type: ${model.type}`);
                }
            }
        });
    }

    /**
     * Validate the prompts section
     */
    validatePrompts(prompts) {
        if (!Array.isArray(prompts)) {
            this.errors.push('prompts must be an array');
            return;
        }

        const promptIds = new Set();
        prompts.forEach((prompt, index) => {
            if (typeof prompt !== 'object') {
                this.errors.push(`Prompt ${index} must be an object`);
                return;
            }

            // Validate required fields
            const requiredFields = ['id', 'role', 'template'];
            requiredFields.forEach(field => {
                if (!(field in prompt)) {
                    this.errors.push(`Prompt ${index} missing required field: ${field}`);
                }
            });

            // Check for duplicate IDs
            if (prompt.id) {
                if (promptIds.has(prompt.id)) {
                    this.errors.push(`Duplicate prompt ID: ${prompt.id}`);
                }
                promptIds.add(prompt.id);
            }

            // Validate role
            if (prompt.role) {
                const validRoles = ['system', 'user', 'assistant'];
                if (!validRoles.includes(prompt.role)) {
                    this.errors.push(`Invalid prompt role: ${prompt.role}`);
                }
            }
        });
    }

    /**
     * Validate the constraints section
     */
    validateConstraints(constraints) {
        if (!Array.isArray(constraints)) {
            this.errors.push('constraints must be an array');
            return;
        }

        const constraintIds = new Set();
        constraints.forEach((constraint, index) => {
            if (typeof constraint !== 'object') {
                this.errors.push(`Constraint ${index} must be an object`);
                return;
            }

            // Validate required fields
            const requiredFields = ['id', 'rule', 'severity'];
            requiredFields.forEach(field => {
                if (!(field in constraint)) {
                    this.errors.push(`Constraint ${index} missing required field: ${field}`);
                }
            });

            // Check for duplicate IDs
            if (constraint.id) {
                if (constraintIds.has(constraint.id)) {
                    this.errors.push(`Duplicate constraint ID: ${constraint.id}`);
                }
                constraintIds.add(constraint.id);
            }

            // Validate severity
            if (constraint.severity) {
                const validSeverities = ['low', 'medium', 'high', 'critical'];
                if (!validSeverities.includes(constraint.severity)) {
                    this.errors.push(`Invalid constraint severity: ${constraint.severity}`);
                }
            }
        });
    }

    /**
     * Validate the tasks section
     */
    validateTasks(tasks) {
        if (!Array.isArray(tasks)) {
            this.errors.push('tasks must be an array');
            return;
        }

        const taskIds = new Set();
        tasks.forEach((task, index) => {
            if (typeof task !== 'object') {
                this.errors.push(`Task ${index} must be an object`);
                return;
            }

            // Validate required fields
            const requiredFields = ['id', 'description'];
            requiredFields.forEach(field => {
                if (!(field in task)) {
                    this.errors.push(`Task ${index} missing required field: ${field}`);
                }
            });

            // Check for duplicate IDs
            if (task.id) {
                if (taskIds.has(task.id)) {
                    this.errors.push(`Duplicate task ID: ${task.id}`);
                }
                taskIds.add(task.id);
            }
        });
    }

    /**
     * Validate the context section
     */
    validateContext(context) {
        if (!context || typeof context !== 'object') {
            this.errors.push('context must be an object');
            return;
        }

        if (!context.memory) {
            this.warnings.push('context.memory is recommended');
        }
    }

    /**
     * Validate the evaluation section
     */
    validateEvaluation(evaluation) {
        if (!evaluation || typeof evaluation !== 'object') {
            this.errors.push('evaluation must be an object');
            return;
        }

        if (!evaluation.metrics) {
            this.warnings.push('evaluation.metrics is recommended');
        }
    }

    /**
     * Perform cross-validation between sections
     */
    crossValidate(spec) {
        // Validate that referenced models exist
        if (spec.tasks && spec.models) {
            const modelIds = new Set(spec.models.map(model => model.id).filter(Boolean));

            spec.tasks.forEach(task => {
                if (task.steps) {
                    task.steps.forEach(step => {
                        if (step.model && !modelIds.has(step.model)) {
                            this.errors.push(`Task references unknown model: ${step.model}`);
                        }
                    });
                }
            });
        }

        // Validate that referenced prompts exist
        if (spec.tasks && spec.prompts) {
            const promptIds = new Set(spec.prompts.map(prompt => prompt.id).filter(Boolean));

            spec.tasks.forEach(task => {
                if (task.steps) {
                    task.steps.forEach(step => {
                        if (step.prompt && !promptIds.has(step.prompt)) {
                            this.errors.push(`Task references unknown prompt: ${step.prompt}`);
                        }
                    });
                }
            });
        }
    }

    /**
     * Get list of validation errors
     */
    getErrors() {
        return this.errors;
    }

    /**
     * Get list of validation warnings
     */
    getWarnings() {
        return this.warnings;
    }

    /**
     * Print validation results
     */
    printResults() {
        if (this.errors.length > 0) {
            console.log('❌ Validation Errors:');
            this.errors.forEach(error => {
                console.log(`  - ${error}`);
            });
        }

        if (this.warnings.length > 0) {
            console.log('⚠️  Validation Warnings:');
            this.warnings.forEach(warning => {
                console.log(`  - ${warning}`);
            });
        }

        if (this.errors.length === 0 && this.warnings.length === 0) {
            console.log('✅ Validation passed with no issues');
        } else if (this.errors.length === 0) {
            console.log('✅ Validation passed with warnings');
        }
    }

    /**
     * Get validation results as object
     */
    getResults() {
        return {
            valid: this.errors.length === 0,
            errors: this.errors,
            warnings: this.warnings
        };
    }
}

module.exports = OpenAPIAValidator;
