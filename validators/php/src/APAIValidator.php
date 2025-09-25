<?php

/**
 * OpenAPIA Validator - PHP Implementation
 * 
 * A comprehensive validator for OpenAPIA specifications.
 * 
 * @package OpenAPIA
 * @version 0.1.0
 * @license Apache-2.0
 */

namespace OpenAPIA;

use Symfony\Component\Yaml\Yaml;
use Symfony\Component\Yaml\Exception\ParseException;

class APAIValidator
{
    private array $errors = [];
    private array $warnings = [];
    private string $schemaVersion = '0.1.0';
    
    // Hierarchical composition properties
    private array $inheritedSpecs = [];
    private array $mergeCache = [];
    
    /**
     * Validate an OpenAPIA specification file
     */
    public function validateFile(string $filePath): bool
    {
        try {
            if (!file_exists($filePath)) {
                $this->errors[] = "File not found: {$filePath}";
                return false;
            }
            
            $content = file_get_contents($filePath);
            if ($content === false) {
                $this->errors[] = "Cannot read file: {$filePath}";
                return false;
            }
            
            $spec = $this->parseFile($filePath, $content);
            if ($spec === null) {
                return false;
            }
            
            return $this->validateSpec($spec);
            
        } catch (\Exception $e) {
            $this->errors[] = "Unexpected error: " . $e->getMessage();
            return false;
        }
    }
    
    /**
     * Parse file content based on file extension
     */
    private function parseFile(string $filePath, string $content): ?array
    {
        $extension = strtolower(pathinfo($filePath, PATHINFO_EXTENSION));
        
        switch ($extension) {
            case 'yaml':
            case 'yml':
                try {
                    return Yaml::parse($content);
                } catch (ParseException $e) {
                    $this->errors[] = "YAML parsing error: " . $e->getMessage();
                    return null;
                }
                
            case 'json':
                $decoded = json_decode($content, true);
                if (json_last_error() !== JSON_ERROR_NONE) {
                    $this->errors[] = "JSON parsing error: " . json_last_error_msg();
                    return null;
                }
                return $decoded;
                
            default:
                $this->errors[] = "Unsupported file format: {$extension}";
                return null;
        }
    }
    
    /**
     * Validate an OpenAPIA specification array
     */
    public function validateSpec(array $spec): bool
    {
        $this->errors = [];
        $this->warnings = [];
        
        // Validate required sections
        $this->validateRequiredSections($spec);
        
        // Validate each section
        if (isset($spec['openapia'])) {
            $this->validateOpenAPIAVersion($spec['openapia']);
        }
        
        if (isset($spec['info'])) {
            $this->validateInfo($spec['info']);
        }
        
        if (isset($spec['models'])) {
            $this->validateModels($spec['models']);
        }
        
        if (isset($spec['prompts'])) {
            $this->validatePrompts($spec['prompts']);
        }
        
        if (isset($spec['constraints'])) {
            $this->validateConstraints($spec['constraints']);
        }
        
        if (isset($spec['tasks'])) {
            $this->validateTasks($spec['tasks']);
        }
        
        if (isset($spec['context'])) {
            $this->validateContext($spec['context']);
        }
        
        if (isset($spec['evaluation'])) {
            $this->validateEvaluation($spec['evaluation']);
        }
        
        // Cross-validation
        $this->crossValidate($spec);
        
        return empty($this->errors);
    }
    
    /**
     * Validate that all required sections are present
     */
    private function validateRequiredSections(array $spec): void
    {
        $requiredSections = [
            'openapia', 'info', 'models', 'prompts',
            'constraints', 'tasks', 'context', 'evaluation'
        ];
        
        foreach ($requiredSections as $section) {
            if (!array_key_exists($section, $spec)) {
                $this->errors[] = "Missing required section: {$section}";
            }
        }
    }
    
    /**
     * Validate the OpenAPIA version
     */
    private function validateOpenAPIAVersion($version): void
    {
        if (!is_string($version)) {
            $this->errors[] = "openapia version must be a string";
            return;
        }
        
        if (!str_starts_with($version, '0.1')) {
            $this->warnings[] = "Version {$version} may not be supported";
        }
    }
    
    /**
     * Validate the info section
     */
    private function validateInfo(array $info): void
    {
        $requiredFields = ['title', 'version', 'description'];
        
        foreach ($requiredFields as $field) {
            if (!array_key_exists($field, $info)) {
                $this->errors[] = "Missing required field in info: {$field}";
            }
        }
        
        if (isset($info['ai_metadata'])) {
            $this->validateAIMetadata($info['ai_metadata']);
        }
    }
    
    /**
     * Validate AI-specific metadata
     */
    private function validateAIMetadata(array $metadata): void
    {
        if (!array_key_exists('domain', $metadata)) {
            $this->warnings[] = "ai_metadata.domain is recommended";
        }
        
        if (isset($metadata['complexity'])) {
            $validComplexities = ['low', 'medium', 'high'];
            if (!in_array($metadata['complexity'], $validComplexities)) {
                $this->errors[] = "Invalid complexity: {$metadata['complexity']}";
            }
        }
    }
    
    /**
     * Validate the models section
     */
    private function validateModels(array $models): void
    {
        if (!is_array($models)) {
            $this->errors[] = "models must be an array";
            return;
        }
        
        if (empty($models)) {
            $this->errors[] = "At least one model is required";
            return;
        }
        
        $modelIds = [];
        foreach ($models as $index => $model) {
            if (!is_array($model)) {
                $this->errors[] = "Model {$index} must be an array";
                continue;
            }
            
            // Validate required fields
            $requiredFields = ['id', 'type', 'provider', 'name', 'purpose'];
            foreach ($requiredFields as $field) {
                if (!array_key_exists($field, $model)) {
                    $this->errors[] = "Model {$index} missing required field: {$field}";
                }
            }
            
            // Check for duplicate IDs
            if (isset($model['id'])) {
                if (in_array($model['id'], $modelIds)) {
                    $this->errors[] = "Duplicate model ID: {$model['id']}";
                }
                $modelIds[] = $model['id'];
            }
            
            // Validate model type
            if (isset($model['type'])) {
                $validTypes = ['LLM', 'Vision', 'Audio', 'Multimodal', 'Classification', 'Embedding'];
                if (!in_array($model['type'], $validTypes)) {
                    $this->warnings[] = "Unknown model type: {$model['type']}";
                }
            }
        }
    }
    
    /**
     * Validate the prompts section
     */
    private function validatePrompts(array $prompts): void
    {
        if (!is_array($prompts)) {
            $this->errors[] = "prompts must be an array";
            return;
        }
        
        $promptIds = [];
        foreach ($prompts as $index => $prompt) {
            if (!is_array($prompt)) {
                $this->errors[] = "Prompt {$index} must be an array";
                continue;
            }
            
            // Validate required fields
            $requiredFields = ['id', 'role', 'template'];
            foreach ($requiredFields as $field) {
                if (!array_key_exists($field, $prompt)) {
                    $this->errors[] = "Prompt {$index} missing required field: {$field}";
                }
            }
            
            // Check for duplicate IDs
            if (isset($prompt['id'])) {
                if (in_array($prompt['id'], $promptIds)) {
                    $this->errors[] = "Duplicate prompt ID: {$prompt['id']}";
                }
                $promptIds[] = $prompt['id'];
            }
            
            // Validate role
            if (isset($prompt['role'])) {
                $validRoles = ['system', 'user', 'assistant'];
                if (!in_array($prompt['role'], $validRoles)) {
                    $this->errors[] = "Invalid prompt role: {$prompt['role']}";
                }
            }
        }
    }
    
    /**
     * Validate the constraints section
     */
    private function validateConstraints(array $constraints): void
    {
        if (!is_array($constraints)) {
            $this->errors[] = "constraints must be an array";
            return;
        }
        
        $constraintIds = [];
        foreach ($constraints as $index => $constraint) {
            if (!is_array($constraint)) {
                $this->errors[] = "Constraint {$index} must be an array";
                continue;
            }
            
            // Validate required fields
            $requiredFields = ['id', 'rule', 'severity'];
            foreach ($requiredFields as $field) {
                if (!array_key_exists($field, $constraint)) {
                    $this->errors[] = "Constraint {$index} missing required field: {$field}";
                }
            }
            
            // Check for duplicate IDs
            if (isset($constraint['id'])) {
                if (in_array($constraint['id'], $constraintIds)) {
                    $this->errors[] = "Duplicate constraint ID: {$constraint['id']}";
                }
                $constraintIds[] = $constraint['id'];
            }
            
            // Validate severity
            if (isset($constraint['severity'])) {
                $validSeverities = ['low', 'medium', 'high', 'critical'];
                if (!in_array($constraint['severity'], $validSeverities)) {
                    $this->errors[] = "Invalid constraint severity: {$constraint['severity']}";
                }
            }
        }
    }
    
    /**
     * Validate the tasks section
     */
    private function validateTasks(array $tasks): void
    {
        if (!is_array($tasks)) {
            $this->errors[] = "tasks must be an array";
            return;
        }
        
        $taskIds = [];
        foreach ($tasks as $index => $task) {
            if (!is_array($task)) {
                $this->errors[] = "Task {$index} must be an array";
                continue;
            }
            
            // Validate required fields
            $requiredFields = ['id', 'description'];
            foreach ($requiredFields as $field) {
                if (!array_key_exists($field, $task)) {
                    $this->errors[] = "Task {$index} missing required field: {$field}";
                }
            }
            
            // Check for duplicate IDs
            if (isset($task['id'])) {
                if (in_array($task['id'], $taskIds)) {
                    $this->errors[] = "Duplicate task ID: {$task['id']}";
                }
                $taskIds[] = $task['id'];
            }
            
            // Validate task steps if present
            if (isset($task['steps']) && is_array($task['steps'])) {
                $this->validateTaskSteps($task['steps'], $index);
            }
        }
    }
    
    /**
     * Validate task steps
     */
    private function validateTaskSteps(array $steps, int $taskIndex): void
    {
        foreach ($steps as $stepIndex => $step) {
            if (!is_array($step)) {
                $this->errors[] = "Task {$taskIndex} step {$stepIndex} must be an array";
                continue;
            }
            
            // Validate required fields
            $requiredFields = ['name', 'action'];
            foreach ($requiredFields as $field) {
                if (!array_key_exists($field, $step)) {
                    $this->errors[] = "Task {$taskIndex} step {$stepIndex} missing required field: {$field}";
                }
            }
            
            // Validate action type
            if (isset($step['action'])) {
                $validActions = ['analyze', 'generate', 'validate', 'search', 'escalate', 'classify', 'mcp_tool', 'mcp_resource'];
                if (!in_array($step['action'], $validActions)) {
                    $this->warnings[] = "Task {$taskIndex} step {$stepIndex} unknown action: {$step['action']}";
                }
            }
            
            // Validate MCP-specific fields
            if (isset($step['action']) && in_array($step['action'], ['mcp_tool', 'mcp_resource'])) {
                if (!array_key_exists('mcp_server', $step)) {
                    $this->errors[] = "Task {$taskIndex} step {$stepIndex} MCP action missing mcp_server field";
                }
                
                if ($step['action'] === 'mcp_tool' && !array_key_exists('mcp_tool', $step)) {
                    $this->errors[] = "Task {$taskIndex} step {$stepIndex} mcp_tool action missing mcp_tool field";
                }
                
                if ($step['action'] === 'mcp_resource' && !array_key_exists('mcp_resource', $step)) {
                    $this->errors[] = "Task {$taskIndex} step {$stepIndex} mcp_resource action missing mcp_resource field";
                }
            }
        }
    }
    
    /**
     * Validate the context section
     */
    private function validateContext(array $context): void
    {
        if (!is_array($context)) {
            $this->errors[] = "context must be an array";
            return;
        }
        
        if (!array_key_exists('memory', $context)) {
            $this->warnings[] = "context.memory is recommended";
        }
        
        // Validate MCP servers if present
        if (array_key_exists('mcp_servers', $context)) {
            $this->validateMcpServers($context['mcp_servers']);
        }
    }
    
    /**
     * Validate MCP servers section
     */
    private function validateMcpServers(array $mcpServers): void
    {
        if (!is_array($mcpServers)) {
            $this->errors[] = "mcp_servers must be an array";
            return;
        }
        
        $serverIds = [];
        foreach ($mcpServers as $index => $server) {
            if (!is_array($server)) {
                $this->errors[] = "MCP server {$index} must be an array";
                continue;
            }
            
            // Validate required fields
            $requiredFields = ['id', 'name', 'description', 'version', 'transport', 'capabilities', 'authentication'];
            foreach ($requiredFields as $field) {
                if (!array_key_exists($field, $server)) {
                    $this->errors[] = "MCP server {$index} missing required field: {$field}";
                }
            }
            
            // Check for duplicate IDs
            if (isset($server['id'])) {
                if (in_array($server['id'], $serverIds)) {
                    $this->errors[] = "Duplicate MCP server ID: {$server['id']}";
                }
                $serverIds[] = $server['id'];
            }
            
            // Validate transport configuration
            if (isset($server['transport']) && is_array($server['transport'])) {
                $this->validateMcpTransport($server['transport'], $index);
            }
            
            // Validate authentication configuration
            if (isset($server['authentication']) && is_array($server['authentication'])) {
                $this->validateMcpAuthentication($server['authentication'], $index);
            }
        }
    }
    
    /**
     * Validate MCP transport configuration
     */
    private function validateMcpTransport(array $transport, int $serverIndex): void
    {
        if (!array_key_exists('type', $transport)) {
            $this->errors[] = "MCP server {$serverIndex} transport missing required field: type";
            return;
        }
        
        $validTypes = ['stdio', 'sse', 'websocket'];
        if (!in_array($transport['type'], $validTypes)) {
            $this->errors[] = "MCP server {$serverIndex} invalid transport type: {$transport['type']}";
        }
        
        // Validate transport-specific fields
        if ($transport['type'] === 'stdio') {
            if (!array_key_exists('command', $transport)) {
                $this->errors[] = "MCP server {$serverIndex} stdio transport missing command";
            }
        } elseif (in_array($transport['type'], ['sse', 'websocket'])) {
            if (!array_key_exists('url', $transport)) {
                $this->errors[] = "MCP server {$serverIndex} {$transport['type']} transport missing url";
            }
        }
    }
    
    /**
     * Validate MCP authentication configuration
     */
    private function validateMcpAuthentication(array $auth, int $serverIndex): void
    {
        if (!array_key_exists('type', $auth)) {
            $this->errors[] = "MCP server {$serverIndex} authentication missing required field: type";
            return;
        }
        
        $validTypes = ['none', 'api_key', 'oauth', 'custom'];
        if (!in_array($auth['type'], $validTypes)) {
            $this->errors[] = "MCP server {$serverIndex} invalid authentication type: {$auth['type']}";
        }
        
        // Validate authentication-specific fields
        if ($auth['type'] === 'api_key' && !array_key_exists('api_key', $auth)) {
            $this->warnings[] = "MCP server {$serverIndex} api_key authentication missing api_key field";
        }
        
        if ($auth['type'] === 'oauth' && !array_key_exists('token', $auth)) {
            $this->warnings[] = "MCP server {$serverIndex} oauth authentication missing token field";
        }
    }
    
    /**
     * Validate the evaluation section
     */
    private function validateEvaluation(array $evaluation): void
    {
        if (!is_array($evaluation)) {
            $this->errors[] = "evaluation must be an array";
            return;
        }
        
        if (!array_key_exists('metrics', $evaluation)) {
            $this->warnings[] = "evaluation.metrics is recommended";
        }
    }
    
    /**
     * Perform cross-validation between sections
     */
    private function crossValidate(array $spec): void
    {
        // Validate that referenced models exist
        if (isset($spec['tasks']) && isset($spec['models'])) {
            $modelIds = [];
            foreach ($spec['models'] as $model) {
                if (isset($model['id'])) {
                    $modelIds[] = $model['id'];
                }
            }
            
            foreach ($spec['tasks'] as $task) {
                if (isset($task['steps'])) {
                    foreach ($task['steps'] as $step) {
                        if (isset($step['model']) && !in_array($step['model'], $modelIds)) {
                            $this->errors[] = "Task references unknown model: {$step['model']}";
                        }
                    }
                }
            }
        }
        
        // Validate that referenced prompts exist
        if (isset($spec['tasks']) && isset($spec['prompts'])) {
            $promptIds = [];
            foreach ($spec['prompts'] as $prompt) {
                if (isset($prompt['id'])) {
                    $promptIds[] = $prompt['id'];
                }
            }
            
            foreach ($spec['tasks'] as $task) {
                if (isset($task['steps'])) {
                    foreach ($task['steps'] as $step) {
                        if (isset($step['prompt']) && !in_array($step['prompt'], $promptIds)) {
                            $this->errors[] = "Task references unknown prompt: {$step['prompt']}";
                        }
                    }
                }
            }
        }
        
        // Validate that referenced MCP servers exist
        if (isset($spec['tasks']) && isset($spec['context']['mcp_servers'])) {
            $mcpServerIds = [];
            foreach ($spec['context']['mcp_servers'] as $server) {
                if (isset($server['id'])) {
                    $mcpServerIds[] = $server['id'];
                }
            }
            
            foreach ($spec['tasks'] as $task) {
                if (isset($task['steps'])) {
                    foreach ($task['steps'] as $step) {
                        if (isset($step['mcp_server']) && !in_array($step['mcp_server'], $mcpServerIds)) {
                            $this->errors[] = "Task references unknown MCP server: {$step['mcp_server']}";
                        }
                    }
                }
            }
        }
    }
    
    /**
     * Get list of validation errors
     */
    public function getErrors(): array
    {
        return $this->errors;
    }
    
    /**
     * Get list of validation warnings
     */
    public function getWarnings(): array
    {
        return $this->warnings;
    }
    
    /**
     * Print validation results
     */
    public function printResults(): void
    {
        if (!empty($this->errors)) {
            echo "❌ Validation Errors:\n";
            foreach ($this->errors as $error) {
                echo "  - {$error}\n";
            }
        }
        
        if (!empty($this->warnings)) {
            echo "⚠️  Validation Warnings:\n";
            foreach ($this->warnings as $warning) {
                echo "  - {$warning}\n";
            }
        }
        
        if (empty($this->errors) && empty($this->warnings)) {
            echo "✅ Validation passed with no issues\n";
        } elseif (empty($this->errors)) {
            echo "✅ Validation passed with warnings\n";
        }
    }
    
    /**
     * Get validation results as array
     */
    public function getResults(): array
    {
        return [
            'valid' => empty($this->errors),
            'errors' => $this->errors,
            'warnings' => $this->warnings
        ];
    }
    
    // ============================================================================
    // HIERARCHICAL COMPOSITION METHODS
    // ============================================================================
    
    /**
     * Validate specification with inheritance support
     */
    public function validateWithInheritance(string $filePath): bool
    {
        try {
            if (!file_exists($filePath)) {
                $this->errors[] = "File not found: {$filePath}";
                return false;
            }
            
            $content = file_get_contents($filePath);
            if ($content === false) {
                $this->errors[] = "Cannot read file: {$filePath}";
                return false;
            }
            
            $spec = $this->parseFile($filePath, $content);
            if ($spec === null) {
                return false;
            }
            
            // Load and merge inherited specifications
            $mergedSpec = $this->mergeInheritedSpecifications($spec, $filePath);
            
            // Validate merged specification
            return $this->validateSpec($mergedSpec);
            
        } catch (\Exception $e) {
            $this->errors[] = "Unexpected error during hierarchical validation: " . $e->getMessage();
            return false;
        }
    }
    
    /**
     * Load specification from file (for hierarchical use)
     */
    public function loadSpec(string $filePath): ?array
    {
        try {
            if (!file_exists($filePath)) {
                return null;
            }
            
            $content = file_get_contents($filePath);
            if ($content === false) {
                return null;
            }
            
            return $this->parseFile($filePath, $content);
            
        } catch (\Exception $e) {
            $this->errors[] = "Error loading specification {$filePath}: " . $e->getMessage();
            return null;
        }
    }
    
    /**
     * Resolve inheritance path to absolute path
     */
    private function resolveInheritancePath(string $inheritPath, string $currentSpecPath): string
    {
        $currentDir = dirname($currentSpecPath);
        return realpath($currentDir . DIRECTORY_SEPARATOR . $inheritPath) ?: 
               $currentDir . DIRECTORY_SEPARATOR . $inheritPath;
    }
    
    /**
     * Load all inherited specifications
     */
    private function loadInheritedSpecs(array $spec, string $specPath): void
    {
        if (!isset($spec['inherits']) || !is_array($spec['inherits'])) {
            return;
        }
        
        foreach ($spec['inherits'] as $inheritPath) {
            $resolvedPath = $this->resolveInheritancePath($inheritPath, $specPath);
            
            if (isset($this->inheritedSpecs[$resolvedPath])) {
                continue; // Already loaded
            }
            
            $inheritedSpec = $this->loadSpec($resolvedPath);
            if ($inheritedSpec !== null) {
                $this->inheritedSpecs[$resolvedPath] = $inheritedSpec;
                
                // Recursively load inherited specs
                $this->loadInheritedSpecs($inheritedSpec, $resolvedPath);
            } else {
                $this->errors[] = "Inherited specification not found: {$inheritPath}";
            }
        }
    }
    
    /**
     * Deep merge two arrays
     */
    private function deepMerge(array $base, array $override): array
    {
        $result = $base;
        
        foreach ($override as $key => $value) {
            if (isset($result[$key]) && is_array($result[$key]) && is_array($value)) {
                $result[$key] = $this->deepMerge($result[$key], $value);
            } else {
                $result[$key] = $value;
            }
        }
        
        return $result;
    }
    
    /**
     * Merge specifications based on inheritance
     */
    private function mergeInheritedSpecifications(array $spec, string $specPath): array
    {
        if (isset($this->mergeCache[$specPath])) {
            return $this->mergeCache[$specPath];
        }
        
        // Load inherited specifications
        $this->loadInheritedSpecs($spec, $specPath);
        
        // Start with base specification
        $merged = $spec;
        
        // Apply inheritance in reverse order (so later specs override earlier ones)
        if (isset($spec['inherits']) && is_array($spec['inherits'])) {
            foreach (array_reverse($spec['inherits']) as $inheritPath) {
                $resolvedPath = $this->resolveInheritancePath($inheritPath, $specPath);
                if (isset($this->inheritedSpecs[$resolvedPath])) {
                    $inheritedSpec = $this->inheritedSpecs[$resolvedPath];
                    // Recursively merge inherited spec
                    $inheritedMerged = $this->mergeInheritedSpecifications($inheritedSpec, $resolvedPath);
                    $merged = $this->deepMerge($inheritedMerged, $merged);
                }
            }
        }
        
        // Cache the result
        $this->mergeCache[$specPath] = $merged;
        return $merged;
    }
    
    /**
     * Get hierarchy information from specification
     */
    public function getHierarchyInfo(array $spec): array
    {
        $hierarchyInfo = [];
        
        if (isset($spec['info']['ai_metadata']['hierarchy_info'])) {
            $hierarchyInfo = $spec['info']['ai_metadata']['hierarchy_info'];
        }
        
        return $hierarchyInfo;
    }
    
    /**
     * Print hierarchy tree for a specification
     */
    public function printHierarchyTree(string $specPath, int $level = 0): void
    {
        $indent = str_repeat("  ", $level);
        
        try {
            $spec = $this->loadSpec($specPath);
            if ($spec === null) {
                echo "{$indent}❌ Error loading {$specPath}\n";
                return;
            }
            
            $title = $spec['info']['title'] ?? 'Unknown';
            $hierarchyInfo = $this->getHierarchyInfo($spec);
            $levelName = $hierarchyInfo['level'] ?? 'unknown';
            $scope = $hierarchyInfo['scope'] ?? 'unknown';
            
            echo "{$indent}📄 {$title} ({$levelName}/{$scope})\n";
            echo "{$indent}   Path: {$specPath}\n";
            
            if (isset($spec['inherits']) && is_array($spec['inherits'])) {
                foreach ($spec['inherits'] as $inheritPath) {
                    $resolvedPath = $this->resolveInheritancePath($inheritPath, $specPath);
                    $this->printHierarchyTree($resolvedPath, $level + 1);
                }
            }
        } catch (\Exception $e) {
            echo "{$indent}❌ Error loading {$specPath}: " . $e->getMessage() . "\n";
        }
    }
    
    /**
     * Merge multiple specifications
     */
    public function mergeSpecifications(array $specs, string $outputPath, string $format = 'yaml'): bool
    {
        try {
            if (empty($specs)) {
                $this->errors[] = "No specifications to merge";
                return false;
            }
            
            // Start with first specification
            $merged = $specs[0];
            
            // Merge with remaining specifications
            for ($i = 1; $i < count($specs); $i++) {
                $merged = $this->deepMerge($merged, $specs[$i]);
            }
            
            // Save merged specification
            if ($format === 'yaml') {
                $content = Yaml::dump($merged, 10, 2);
            } else {
                $content = json_encode($merged, JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES);
            }
            
            if (file_put_contents($outputPath, $content) === false) {
                $this->errors[] = "Cannot write to output file: {$outputPath}";
                return false;
            }
            
            return true;
            
        } catch (\Exception $e) {
            $this->errors[] = "Error merging specifications: " . $e->getMessage();
            return false;
        }
    }
}
