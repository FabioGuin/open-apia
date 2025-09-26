<?php

/**
 * APAI Validator Tests
 * 
 * @package APAI
 * @version 0.1.0
 * @license Apache-2.0
 */

namespace APAI\Tests;

use PHPUnit\Framework\TestCase;
use APAI\APAIValidator;

class APAIValidatorTest extends TestCase
{
    private APAIValidator $validator;
    
    protected function setUp(): void
    {
        $this->validator = new APAIValidator();
    }
    
    /**
     * Test valid specification
     */
    public function testValidSpecification(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system'
            ],
            'models' => [
                [
                    'id' => 'test_model',
                    'type' => 'LLM',
                    'provider' => 'openai',
                    'name' => 'gpt-4',
                    'purpose' => 'conversation'
                ]
            ],
            'prompts' => [
                [
                    'id' => 'test_prompt',
                    'role' => 'system',
                    'template' => 'You are a helpful assistant'
                ]
            ],
            'constraints' => [
                [
                    'id' => 'safety',
                    'rule' => 'output NOT contains harmful_content',
                    'severity' => 'critical'
                ]
            ],
            'tasks' => [
                [
                    'id' => 'test_task',
                    'description' => 'Test task'
                ]
            ],
            'context' => [
                'memory' => [
                    'type' => 'session',
                    'retention' => '7d'
                ]
            ],
            'evaluation' => [
                'metrics' => [
                    [
                        'name' => 'accuracy',
                        'target' => 0.9
                    ]
                ]
            ]
        ];
        
        $this->assertTrue($this->validator->validateSpec($spec));
        $this->assertEmpty($this->validator->getErrors());
    }
    
    /**
     * Test missing required sections
     */
    public function testMissingRequiredSections(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system'
            ]
            // Missing other required sections
        ];
        
        $this->assertFalse($this->validator->validateSpec($spec));
        $errors = $this->validator->getErrors();
        $this->assertNotEmpty($errors);
        $this->assertStringContainsString('Missing required section', $errors[0]);
    }
    
    /**
     * Test missing required fields in info
     */
    public function testMissingRequiredFieldsInInfo(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System'
                // Missing version and description
            ],
            'models' => [],
            'prompts' => [],
            'constraints' => [],
            'tasks' => [],
            'context' => [],
            'evaluation' => []
        ];
        
        $this->assertFalse($this->validator->validateSpec($spec));
        $errors = $this->validator->getErrors();
        $this->assertNotEmpty($errors);
        $this->assertStringContainsString('Missing required field in info', $errors[0]);
    }
    
    /**
     * Test missing required fields in models
     */
    public function testMissingRequiredFieldsInModels(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system'
            ],
            'models' => [
                [
                    'id' => 'test_model'
                    // Missing type, provider, name, purpose
                ]
            ],
            'prompts' => [],
            'constraints' => [],
            'tasks' => [],
            'context' => [],
            'evaluation' => []
        ];
        
        $this->assertFalse($this->validator->validateSpec($spec));
        $errors = $this->validator->getErrors();
        $this->assertNotEmpty($errors);
        $this->assertStringContainsString('Model 0 missing required field', $errors[0]);
    }
    
    /**
     * Test duplicate model IDs
     */
    public function testDuplicateModelIds(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system'
            ],
            'models' => [
                [
                    'id' => 'duplicate_id',
                    'type' => 'LLM',
                    'provider' => 'openai',
                    'name' => 'gpt-4',
                    'purpose' => 'conversation'
                ],
                [
                    'id' => 'duplicate_id', // Duplicate ID
                    'type' => 'LLM',
                    'provider' => 'openai',
                    'name' => 'gpt-3.5',
                    'purpose' => 'conversation'
                ]
            ],
            'prompts' => [],
            'constraints' => [],
            'tasks' => [],
            'context' => [],
            'evaluation' => []
        ];
        
        $this->assertFalse($this->validator->validateSpec($spec));
        $errors = $this->validator->getErrors();
        $this->assertNotEmpty($errors);
        $this->assertStringContainsString('Duplicate model ID', $errors[0]);
    }
    
    /**
     * Test invalid model type
     */
    public function testInvalidModelType(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system'
            ],
            'models' => [
                [
                    'id' => 'test_model',
                    'type' => 'INVALID_TYPE',
                    'provider' => 'openai',
                    'name' => 'gpt-4',
                    'purpose' => 'conversation'
                ]
            ],
            'prompts' => [],
            'constraints' => [],
            'tasks' => [],
            'context' => [],
            'evaluation' => []
        ];
        
        $this->assertTrue($this->validator->validateSpec($spec)); // Should pass with warning
        $warnings = $this->validator->getWarnings();
        $this->assertNotEmpty($warnings);
        $this->assertStringContainsString('Unknown model type', $warnings[0]);
    }
    
    /**
     * Test invalid constraint severity
     */
    public function testInvalidConstraintSeverity(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system'
            ],
            'models' => [],
            'prompts' => [],
            'constraints' => [
                [
                    'id' => 'test_constraint',
                    'rule' => 'test rule',
                    'severity' => 'INVALID_SEVERITY'
                ]
            ],
            'tasks' => [],
            'context' => [],
            'evaluation' => []
        ];
        
        $this->assertFalse($this->validator->validateSpec($spec));
        $errors = $this->validator->getErrors();
        $this->assertNotEmpty($errors);
        $this->assertStringContainsString('Invalid constraint severity', $errors[0]);
    }
    
    /**
     * Test cross-validation of model references
     */
    public function testCrossValidationModelReferences(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system'
            ],
            'models' => [
                [
                    'id' => 'existing_model',
                    'type' => 'LLM',
                    'provider' => 'openai',
                    'name' => 'gpt-4',
                    'purpose' => 'conversation'
                ]
            ],
            'prompts' => [],
            'constraints' => [],
            'tasks' => [
                [
                    'id' => 'test_task',
                    'description' => 'Test task',
                    'steps' => [
                        [
                            'action' => 'generate',
                            'model' => 'nonexistent_model' // References non-existent model
                        ]
                    ]
                ]
            ],
            'context' => [],
            'evaluation' => []
        ];
        
        $this->assertFalse($this->validator->validateSpec($spec));
        $errors = $this->validator->getErrors();
        $this->assertNotEmpty($errors);
        $this->assertStringContainsString('Task references unknown model', $errors[0]);
    }
    
    /**
     * Test cross-validation of prompt references
     */
    public function testCrossValidationPromptReferences(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system'
            ],
            'models' => [],
            'prompts' => [
                [
                    'id' => 'existing_prompt',
                    'role' => 'system',
                    'template' => 'You are a helpful assistant'
                ]
            ],
            'constraints' => [],
            'tasks' => [
                [
                    'id' => 'test_task',
                    'description' => 'Test task',
                    'steps' => [
                        [
                            'action' => 'generate',
                            'prompt' => 'nonexistent_prompt' // References non-existent prompt
                        ]
                    ]
                ]
            ],
            'context' => [],
            'evaluation' => []
        ];
        
        $this->assertFalse($this->validator->validateSpec($spec));
        $errors = $this->validator->getErrors();
        $this->assertNotEmpty($errors);
        $this->assertStringContainsString('Task references unknown prompt', $errors[0]);
    }
    
    /**
     * Test AI metadata validation
     */
    public function testAIMetadataValidation(): void
    {
        $spec = [
            'apai' => '0.1.0',
            'info' => [
                'title' => 'Test AI System',
                'version' => '1.0.0',
                'description' => 'A test AI system',
                'ai_metadata' => [
                    'complexity' => 'INVALID_COMPLEXITY'
                ]
            ],
            'models' => [],
            'prompts' => [],
            'constraints' => [],
            'tasks' => [],
            'context' => [],
            'evaluation' => []
        ];
        
        $this->assertFalse($this->validator->validateSpec($spec));
        $errors = $this->validator->getErrors();
        $this->assertNotEmpty($errors);
        $this->assertStringContainsString('Invalid complexity', $errors[0]);
    }
}
