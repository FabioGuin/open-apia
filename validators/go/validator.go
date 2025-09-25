package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// APAIValidator represents the main validator struct
type APAIValidator struct {
	Errors      []string
	Warnings    []string
	SchemaVersion string
	
	// Hierarchical composition properties
	inheritedSpecs map[string]map[string]interface{}
	mergeCache     map[string]map[string]interface{}
}

// ValidationResult represents the result of validation
type ValidationResult struct {
	Valid    bool     `json:"valid"`
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
}

// NewAPAIValidator creates a new validator instance
func NewAPAIValidator() *APAIValidator {
	return &APAIValidator{
		Errors:        make([]string, 0),
		Warnings:      make([]string, 0),
		SchemaVersion: "0.1.0",
		inheritedSpecs: make(map[string]map[string]interface{}),
		mergeCache:     make(map[string]map[string]interface{}),
	}
}

// ValidateFile validates an OpenAPIA specification file
func (v *APAIValidator) ValidateFile(filePath string) (bool, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("file not found: %s", filePath)
	}

	var spec map[string]interface{}
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(content, &spec)
		if err != nil {
			return false, fmt.Errorf("YAML parsing error: %v", err)
		}
	case ".json":
		err = json.Unmarshal(content, &spec)
		if err != nil {
			return false, fmt.Errorf("JSON parsing error: %v", err)
		}
	default:
		return false, fmt.Errorf("unsupported file format: %s", ext)
	}

	return v.ValidateSpec(spec), nil
}

// ValidateSpec validates an OpenAPIA specification map
func (v *APAIValidator) ValidateSpec(spec map[string]interface{}) bool {
	v.Errors = make([]string, 0)
	v.Warnings = make([]string, 0)

	// Validate required sections
	v.validateRequiredSections(spec)

	// Validate each section
	if openapia, exists := spec["openapia"]; exists {
		v.validateOpenAPIAVersion(openapia)
	}

	if info, exists := spec["info"]; exists {
		v.validateInfo(info)
	}

	if models, exists := spec["models"]; exists {
		v.validateModels(models)
	}

	if prompts, exists := spec["prompts"]; exists {
		v.validatePrompts(prompts)
	}

	if constraints, exists := spec["constraints"]; exists {
		v.validateConstraints(constraints)
	}

	if tasks, exists := spec["tasks"]; exists {
		v.validateTasks(tasks)
	}

	if context, exists := spec["context"]; exists {
		v.validateContext(context)
	}

	if evaluation, exists := spec["evaluation"]; exists {
		v.validateEvaluation(evaluation)
	}

	// Cross-validation
	v.crossValidate(spec)

	return len(v.Errors) == 0
}

// validateRequiredSections validates that all required sections are present
func (v *APAIValidator) validateRequiredSections(spec map[string]interface{}) {
	requiredSections := []string{
		"openapia", "info", "models", "prompts",
		"constraints", "tasks", "context", "evaluation",
	}

	for _, section := range requiredSections {
		if _, exists := spec[section]; !exists {
			v.Errors = append(v.Errors, fmt.Sprintf("Missing required section: %s", section))
		}
	}
}

// validateOpenAPIAVersion validates the OpenAPIA version
func (v *APAIValidator) validateOpenAPIAVersion(version interface{}) {
	versionStr, ok := version.(string)
	if !ok {
		v.Errors = append(v.Errors, "openapia version must be a string")
		return
	}

	matched, _ := regexp.MatchString(`^0\.1\.\d+$`, versionStr)
	if !matched {
		v.Warnings = append(v.Warnings, fmt.Sprintf("Version %s may not be supported", versionStr))
	}
}

// validateInfo validates the info section
func (v *APAIValidator) validateInfo(info interface{}) {
	infoMap, ok := info.(map[string]interface{})
	if !ok {
		v.Errors = append(v.Errors, "info must be an object")
		return
	}

	requiredFields := []string{"title", "version", "description", "author", "license"}
	for _, field := range requiredFields {
		if _, exists := infoMap[field]; !exists {
			v.Errors = append(v.Errors, fmt.Sprintf("Missing required field in info: %s", field))
		}
	}

	if aiMetadata, exists := infoMap["ai_metadata"]; exists {
		v.validateAIMetadata(aiMetadata)
	}
}

// validateAIMetadata validates AI-specific metadata
func (v *APAIValidator) validateAIMetadata(metadata interface{}) {
	metadataMap, ok := metadata.(map[string]interface{})
	if !ok {
		return
	}

	if _, exists := metadataMap["domain"]; !exists {
		v.Warnings = append(v.Warnings, "ai_metadata.domain is recommended")
	}

	if complexity, exists := metadataMap["complexity"]; exists {
		complexityStr, ok := complexity.(string)
		if ok {
			validComplexities := []string{"low", "medium", "high"}
			valid := false
			for _, validComplexity := range validComplexities {
				if complexityStr == validComplexity {
					valid = true
					break
				}
			}
			if !valid {
				v.Errors = append(v.Errors, fmt.Sprintf("Invalid complexity: %s", complexityStr))
			}
		}
	}
}

// validateModels validates the models section
func (v *APAIValidator) validateModels(models interface{}) {
	modelsSlice, ok := models.([]interface{})
	if !ok {
		v.Errors = append(v.Errors, "models must be an array")
		return
	}

	if len(modelsSlice) == 0 {
		v.Errors = append(v.Errors, "At least one model is required")
		return
	}

	modelIds := make(map[string]bool)
	for i, model := range modelsSlice {
		modelMap, ok := model.(map[string]interface{})
		if !ok {
			v.Errors = append(v.Errors, fmt.Sprintf("Model %d must be an object", i))
			continue
		}

		// Validate required fields
		requiredFields := []string{"id", "type", "provider", "name", "purpose"}
		for _, field := range requiredFields {
			if _, exists := modelMap[field]; !exists {
				v.Errors = append(v.Errors, fmt.Sprintf("Model %d missing required field: %s", i, field))
			}
		}

		// Check for duplicate IDs
		if id, exists := modelMap["id"]; exists {
			idStr, ok := id.(string)
			if ok {
				if modelIds[idStr] {
					v.Errors = append(v.Errors, fmt.Sprintf("Duplicate model ID: %s", idStr))
				}
				modelIds[idStr] = true
			}
		}

		// Validate model type
		if modelType, exists := modelMap["type"]; exists {
			typeStr, ok := modelType.(string)
			if ok {
				validTypes := []string{"LLM", "Vision", "Audio", "Multimodal", "Classification", "Embedding"}
				valid := false
				for _, validType := range validTypes {
					if typeStr == validType {
						valid = true
						break
					}
				}
				if !valid {
					v.Warnings = append(v.Warnings, fmt.Sprintf("Unknown model type: %s", typeStr))
				}
			}
		}
	}
}

// validatePrompts validates the prompts section
func (v *APAIValidator) validatePrompts(prompts interface{}) {
	promptsSlice, ok := prompts.([]interface{})
	if !ok {
		v.Errors = append(v.Errors, "prompts must be an array")
		return
	}

	promptIds := make(map[string]bool)
	for i, prompt := range promptsSlice {
		promptMap, ok := prompt.(map[string]interface{})
		if !ok {
			v.Errors = append(v.Errors, fmt.Sprintf("Prompt %d must be an object", i))
			continue
		}

		// Validate required fields
		requiredFields := []string{"id", "role", "template"}
		for _, field := range requiredFields {
			if _, exists := promptMap[field]; !exists {
				v.Errors = append(v.Errors, fmt.Sprintf("Prompt %d missing required field: %s", i, field))
			}
		}

		// Check for duplicate IDs
		if id, exists := promptMap["id"]; exists {
			idStr, ok := id.(string)
			if ok {
				if promptIds[idStr] {
					v.Errors = append(v.Errors, fmt.Sprintf("Duplicate prompt ID: %s", idStr))
				}
				promptIds[idStr] = true
			}
		}

		// Validate role
		if role, exists := promptMap["role"]; exists {
			roleStr, ok := role.(string)
			if ok {
				validRoles := []string{"system", "user", "assistant"}
				valid := false
				for _, validRole := range validRoles {
					if roleStr == validRole {
						valid = true
						break
					}
				}
				if !valid {
					v.Errors = append(v.Errors, fmt.Sprintf("Invalid prompt role: %s", roleStr))
				}
			}
		}
	}
}

// validateConstraints validates the constraints section
func (v *APAIValidator) validateConstraints(constraints interface{}) {
	constraintsSlice, ok := constraints.([]interface{})
	if !ok {
		v.Errors = append(v.Errors, "constraints must be an array")
		return
	}

	constraintIds := make(map[string]bool)
	for i, constraint := range constraintsSlice {
		constraintMap, ok := constraint.(map[string]interface{})
		if !ok {
			v.Errors = append(v.Errors, fmt.Sprintf("Constraint %d must be an object", i))
			continue
		}

		// Validate required fields
		requiredFields := []string{"id", "rule", "severity"}
		for _, field := range requiredFields {
			if _, exists := constraintMap[field]; !exists {
				v.Errors = append(v.Errors, fmt.Sprintf("Constraint %d missing required field: %s", i, field))
			}
		}

		// Check for duplicate IDs
		if id, exists := constraintMap["id"]; exists {
			idStr, ok := id.(string)
			if ok {
				if constraintIds[idStr] {
					v.Errors = append(v.Errors, fmt.Sprintf("Duplicate constraint ID: %s", idStr))
				}
				constraintIds[idStr] = true
			}
		}

		// Validate severity
		if severity, exists := constraintMap["severity"]; exists {
			severityStr, ok := severity.(string)
			if ok {
				validSeverities := []string{"low", "medium", "high", "critical"}
				valid := false
				for _, validSeverity := range validSeverities {
					if severityStr == validSeverity {
						valid = true
						break
					}
				}
				if !valid {
					v.Errors = append(v.Errors, fmt.Sprintf("Invalid constraint severity: %s", severityStr))
				}
			}
		}
	}
}

// validateTasks validates the tasks section
func (v *APAIValidator) validateTasks(tasks interface{}) {
	tasksSlice, ok := tasks.([]interface{})
	if !ok {
		v.Errors = append(v.Errors, "tasks must be an array")
		return
	}

	taskIds := make(map[string]bool)
	for i, task := range tasksSlice {
		taskMap, ok := task.(map[string]interface{})
		if !ok {
			v.Errors = append(v.Errors, fmt.Sprintf("Task %d must be an object", i))
			continue
		}

		// Validate required fields
		requiredFields := []string{"id", "description"}
		for _, field := range requiredFields {
			if _, exists := taskMap[field]; !exists {
				v.Errors = append(v.Errors, fmt.Sprintf("Task %d missing required field: %s", i, field))
			}
		}

		// Check for duplicate IDs
		if id, exists := taskMap["id"]; exists {
			idStr, ok := id.(string)
			if ok {
				if taskIds[idStr] {
					v.Errors = append(v.Errors, fmt.Sprintf("Duplicate task ID: %s", idStr))
				}
				taskIds[idStr] = true
			}
		}

		// Validate task steps if present
		if steps, exists := taskMap["steps"]; exists {
			v.validateTaskSteps(steps, i)
		}
	}
}

// validateTaskSteps validates task steps
func (v *APAIValidator) validateTaskSteps(steps interface{}, taskIndex int) {
	stepsSlice, ok := steps.([]interface{})
	if !ok {
		v.Errors = append(v.Errors, fmt.Sprintf("Task %d steps must be an array", taskIndex))
		return
	}

	for stepIndex, step := range stepsSlice {
		stepMap, ok := step.(map[string]interface{})
		if !ok {
			v.Errors = append(v.Errors, fmt.Sprintf("Task %d step %d must be an object", taskIndex, stepIndex))
			continue
		}

		// Validate required fields
		requiredFields := []string{"name", "action"}
		for _, field := range requiredFields {
			if _, exists := stepMap[field]; !exists {
				v.Errors = append(v.Errors, fmt.Sprintf("Task %d step %d missing required field: %s", taskIndex, stepIndex, field))
			}
		}

		// Validate action type
		if action, exists := stepMap["action"]; exists {
			if actionStr, ok := action.(string); ok {
				validActions := []string{"analyze", "generate", "validate", "search", "escalate", "classify", "mcp_tool", "mcp_resource"}
				isValid := false
				for _, validAction := range validActions {
					if actionStr == validAction {
						isValid = true
						break
					}
				}
				if !isValid {
					v.Warnings = append(v.Warnings, fmt.Sprintf("Task %d step %d unknown action: %s", taskIndex, stepIndex, actionStr))
				}
			}
		}

		// Validate MCP-specific fields
		if action, exists := stepMap["action"]; exists {
			if actionStr, ok := action.(string); ok {
				if actionStr == "mcp_tool" || actionStr == "mcp_resource" {
					if _, exists := stepMap["mcp_server"]; !exists {
						v.Errors = append(v.Errors, fmt.Sprintf("Task %d step %d MCP action missing mcp_server field", taskIndex, stepIndex))
					}

					if actionStr == "mcp_tool" {
						if _, exists := stepMap["mcp_tool"]; !exists {
							v.Errors = append(v.Errors, fmt.Sprintf("Task %d step %d mcp_tool action missing mcp_tool field", taskIndex, stepIndex))
						}
					}

					if actionStr == "mcp_resource" {
						if _, exists := stepMap["mcp_resource"]; !exists {
							v.Errors = append(v.Errors, fmt.Sprintf("Task %d step %d mcp_resource action missing mcp_resource field", taskIndex, stepIndex))
						}
					}
				}
			}
		}
	}
}

// validateContext validates the context section
func (v *APAIValidator) validateContext(context interface{}) {
	contextMap, ok := context.(map[string]interface{})
	if !ok {
		v.Errors = append(v.Errors, "context must be an object")
		return
	}

	if _, exists := contextMap["memory"]; !exists {
		v.Warnings = append(v.Warnings, "context.memory is recommended")
	}

	// Validate MCP servers if present
	if mcpServers, exists := contextMap["mcp_servers"]; exists {
		v.validateMcpServers(mcpServers)
	}
}

// validateMcpServers validates MCP servers section
func (v *APAIValidator) validateMcpServers(mcpServers interface{}) {
	mcpServersSlice, ok := mcpServers.([]interface{})
	if !ok {
		v.Errors = append(v.Errors, "mcp_servers must be an array")
		return
	}

	serverIds := make(map[string]bool)
	for index, server := range mcpServersSlice {
		serverMap, ok := server.(map[string]interface{})
		if !ok {
			v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d must be an object", index))
			continue
		}

		// Validate required fields
		requiredFields := []string{"id", "name", "description", "version", "transport", "capabilities", "authentication"}
		for _, field := range requiredFields {
			if _, exists := serverMap[field]; !exists {
				v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d missing required field: %s", index, field))
			}
		}

		// Check for duplicate IDs
		if id, exists := serverMap["id"]; exists {
			if idStr, ok := id.(string); ok {
				if serverIds[idStr] {
					v.Errors = append(v.Errors, fmt.Sprintf("Duplicate MCP server ID: %s", idStr))
				}
				serverIds[idStr] = true
			}
		}

		// Validate transport configuration
		if transport, exists := serverMap["transport"]; exists {
			v.validateMcpTransport(transport, index)
		}

		// Validate authentication configuration
		if auth, exists := serverMap["authentication"]; exists {
			v.validateMcpAuthentication(auth, index)
		}
	}
}

// validateMcpTransport validates MCP transport configuration
func (v *APAIValidator) validateMcpTransport(transport interface{}, serverIndex int) {
	transportMap, ok := transport.(map[string]interface{})
	if !ok {
		v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d transport must be an object", serverIndex))
		return
	}

	if transportType, exists := transportMap["type"]; exists {
		if typeStr, ok := transportType.(string); ok {
			validTypes := []string{"stdio", "sse", "websocket"}
			isValid := false
			for _, validType := range validTypes {
				if typeStr == validType {
					isValid = true
					break
				}
			}
			if !isValid {
				v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d invalid transport type: %s", serverIndex, typeStr))
			}

			// Validate transport-specific fields
			if typeStr == "stdio" {
				if _, exists := transportMap["command"]; !exists {
					v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d stdio transport missing command", serverIndex))
				}
			} else if typeStr == "sse" || typeStr == "websocket" {
				if _, exists := transportMap["url"]; !exists {
					v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d %s transport missing url", serverIndex, typeStr))
				}
			}
		}
	} else {
		v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d transport missing required field: type", serverIndex))
	}
}

// validateMcpAuthentication validates MCP authentication configuration
func (v *APAIValidator) validateMcpAuthentication(auth interface{}, serverIndex int) {
	authMap, ok := auth.(map[string]interface{})
	if !ok {
		v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d authentication must be an object", serverIndex))
		return
	}

	if authType, exists := authMap["type"]; exists {
		if typeStr, ok := authType.(string); ok {
			validTypes := []string{"none", "api_key", "oauth", "custom"}
			isValid := false
			for _, validType := range validTypes {
				if typeStr == validType {
					isValid = true
					break
				}
			}
			if !isValid {
				v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d invalid authentication type: %s", serverIndex, typeStr))
			}

			// Validate authentication-specific fields
			if typeStr == "api_key" {
				if _, exists := authMap["api_key"]; !exists {
					v.Warnings = append(v.Warnings, fmt.Sprintf("MCP server %d api_key authentication missing api_key field", serverIndex))
				}
			}
			if typeStr == "oauth" {
				if _, exists := authMap["token"]; !exists {
					v.Warnings = append(v.Warnings, fmt.Sprintf("MCP server %d oauth authentication missing token field", serverIndex))
				}
			}
		}
	} else {
		v.Errors = append(v.Errors, fmt.Sprintf("MCP server %d authentication missing required field: type", serverIndex))
	}
}

// validateEvaluation validates the evaluation section
func (v *APAIValidator) validateEvaluation(evaluation interface{}) {
	evaluationMap, ok := evaluation.(map[string]interface{})
	if !ok {
		v.Errors = append(v.Errors, "evaluation must be an object")
		return
	}

	if _, exists := evaluationMap["metrics"]; !exists {
		v.Warnings = append(v.Warnings, "evaluation.metrics is recommended")
	}
}

// crossValidate performs cross-validation between sections
func (v *APAIValidator) crossValidate(spec map[string]interface{}) {
	// Validate that referenced models exist
	if tasks, tasksExists := spec["tasks"]; tasksExists {
		if models, modelsExists := spec["models"]; modelsExists {
			modelIds := make(map[string]bool)
			if modelsSlice, ok := models.([]interface{}); ok {
				for _, model := range modelsSlice {
					if modelMap, ok := model.(map[string]interface{}); ok {
						if id, exists := modelMap["id"]; exists {
							if idStr, ok := id.(string); ok {
								modelIds[idStr] = true
							}
						}
					}
				}
			}

			if tasksSlice, ok := tasks.([]interface{}); ok {
				for _, task := range tasksSlice {
					if taskMap, ok := task.(map[string]interface{}); ok {
						if steps, exists := taskMap["steps"]; exists {
							if stepsSlice, ok := steps.([]interface{}); ok {
								for _, step := range stepsSlice {
									if stepMap, ok := step.(map[string]interface{}); ok {
										if model, exists := stepMap["model"]; exists {
											if modelStr, ok := model.(string); ok {
												if !modelIds[modelStr] {
													v.Errors = append(v.Errors, fmt.Sprintf("Task references unknown model: %s", modelStr))
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// Validate that referenced prompts exist
	if tasks, tasksExists := spec["tasks"]; tasksExists {
		if prompts, promptsExists := spec["prompts"]; promptsExists {
			promptIds := make(map[string]bool)
			if promptsSlice, ok := prompts.([]interface{}); ok {
				for _, prompt := range promptsSlice {
					if promptMap, ok := prompt.(map[string]interface{}); ok {
						if id, exists := promptMap["id"]; exists {
							if idStr, ok := id.(string); ok {
								promptIds[idStr] = true
							}
						}
					}
				}
			}

			if tasksSlice, ok := tasks.([]interface{}); ok {
				for _, task := range tasksSlice {
					if taskMap, ok := task.(map[string]interface{}); ok {
						if steps, exists := taskMap["steps"]; exists {
							if stepsSlice, ok := steps.([]interface{}); ok {
								for _, step := range stepsSlice {
									if stepMap, ok := step.(map[string]interface{}); ok {
										if prompt, exists := stepMap["prompt"]; exists {
											if promptStr, ok := prompt.(string); ok {
												if !promptIds[promptStr] {
													v.Errors = append(v.Errors, fmt.Sprintf("Task references unknown prompt: %s", promptStr))
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// Validate that referenced MCP servers exist
	if tasks, tasksExists := spec["tasks"]; tasksExists {
		if context, contextExists := spec["context"]; contextExists {
			if contextMap, ok := context.(map[string]interface{}); ok {
				if mcpServers, mcpServersExists := contextMap["mcp_servers"]; mcpServersExists {
					mcpServerIds := make(map[string]bool)
					if mcpServersSlice, ok := mcpServers.([]interface{}); ok {
						for _, server := range mcpServersSlice {
							if serverMap, ok := server.(map[string]interface{}); ok {
								if id, exists := serverMap["id"]; exists {
									if idStr, ok := id.(string); ok {
										mcpServerIds[idStr] = true
									}
								}
							}
						}
					}

					if tasksSlice, ok := tasks.([]interface{}); ok {
						for _, task := range tasksSlice {
							if taskMap, ok := task.(map[string]interface{}); ok {
								if steps, exists := taskMap["steps"]; exists {
									if stepsSlice, ok := steps.([]interface{}); ok {
										for _, step := range stepsSlice {
											if stepMap, ok := step.(map[string]interface{}); ok {
												if mcpServer, exists := stepMap["mcp_server"]; exists {
													if mcpServerStr, ok := mcpServer.(string); ok {
														if !mcpServerIds[mcpServerStr] {
															v.Errors = append(v.Errors, fmt.Sprintf("Task references unknown MCP server: %s", mcpServerStr))
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// GetErrors returns the list of validation errors
func (v *APAIValidator) GetErrors() []string {
	return v.Errors
}

// GetWarnings returns the list of validation warnings
func (v *APAIValidator) GetWarnings() []string {
	return v.Warnings
}

// PrintResults prints validation results
func (v *APAIValidator) PrintResults() {
	if len(v.Errors) > 0 {
		fmt.Println("❌ Validation Errors:")
		for _, error := range v.Errors {
			fmt.Printf("  - %s\n", error)
		}
	}

	if len(v.Warnings) > 0 {
		fmt.Println("⚠️  Validation Warnings:")
		for _, warning := range v.Warnings {
			fmt.Printf("  - %s\n", warning)
		}
	}

	if len(v.Errors) == 0 && len(v.Warnings) == 0 {
		fmt.Println("✅ Validation passed with no issues")
	} else if len(v.Errors) == 0 {
		fmt.Println("✅ Validation passed with warnings")
	}
}

// GetResults returns validation results as a struct
func (v *APAIValidator) GetResults() ValidationResult {
	return ValidationResult{
		Valid:    len(v.Errors) == 0,
		Errors:   v.Errors,
		Warnings: v.Warnings,
	}
}

// ============================================================================
// HIERARCHICAL COMPOSITION METHODS
// ============================================================================

// ValidateWithInheritance validates specification with inheritance support
func (v *APAIValidator) ValidateWithInheritance(filePath string) (bool, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("file not found: %s", filePath)
	}

	var spec map[string]interface{}
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(content, &spec)
		if err != nil {
			return false, fmt.Errorf("invalid YAML: %v", err)
		}
	case ".json":
		err = json.Unmarshal(content, &spec)
		if err != nil {
			return false, fmt.Errorf("invalid JSON: %v", err)
		}
	default:
		return false, fmt.Errorf("unsupported file format: %s", ext)
	}

	// Load and merge inherited specifications
	mergedSpec := v.mergeInheritedSpecifications(spec, filePath)

	// Validate merged specification
	isValid := v.ValidateSpec(mergedSpec)
	return isValid, nil
}

// loadSpec loads specification from file (for hierarchical use)
func (v *APAIValidator) loadSpec(filePath string) (map[string]interface{}, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("file not found: %s", filePath)
	}

	var spec map[string]interface{}
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(content, &spec)
		if err != nil {
			return nil, fmt.Errorf("invalid YAML: %v", err)
		}
	case ".json":
		err = json.Unmarshal(content, &spec)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	return spec, nil
}

// resolveInheritancePath resolves inheritance path to absolute path
func (v *APAIValidator) resolveInheritancePath(inheritPath, currentSpecPath string) string {
	currentDir := filepath.Dir(currentSpecPath)
	return filepath.Join(currentDir, inheritPath)
}

// loadInheritedSpecs loads all inherited specifications
func (v *APAIValidator) loadInheritedSpecs(spec map[string]interface{}, specPath string) {
	inherits, exists := spec["inherits"]
	if !exists {
		return
	}

	inheritsSlice, ok := inherits.([]interface{})
	if !ok {
		return
	}

	for _, inheritPath := range inheritsSlice {
		inheritPathStr, ok := inheritPath.(string)
		if !ok {
			continue
		}

		resolvedPath := v.resolveInheritancePath(inheritPathStr, specPath)

		if _, exists := v.inheritedSpecs[resolvedPath]; exists {
			continue // Already loaded
		}

		inheritedSpec, err := v.loadSpec(resolvedPath)
		if err != nil {
			v.Errors = append(v.Errors, fmt.Sprintf("Inherited specification not found: %s", inheritPathStr))
			continue
		}

		v.inheritedSpecs[resolvedPath] = inheritedSpec

		// Recursively load inherited specs
		v.loadInheritedSpecs(inheritedSpec, resolvedPath)
	}
}

// deepMerge performs deep merge of two maps
func (v *APAIValidator) deepMerge(base, override map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Copy base values
	for key, value := range base {
		result[key] = value
	}

	// Override with override values
	for key, value := range override {
		if baseValue, exists := result[key]; exists {
			if baseMap, baseIsMap := baseValue.(map[string]interface{}); baseIsMap {
				if overrideMap, overrideIsMap := value.(map[string]interface{}); overrideIsMap {
					result[key] = v.deepMerge(baseMap, overrideMap)
					continue
				}
			}
		}
		result[key] = value
	}

	return result
}

// mergeInheritedSpecifications merges specifications based on inheritance
func (v *APAIValidator) mergeInheritedSpecifications(spec map[string]interface{}, specPath string) map[string]interface{} {
	if cached, exists := v.mergeCache[specPath]; exists {
		return cached
	}

	// Load inherited specifications
	v.loadInheritedSpecs(spec, specPath)

	// Start with base specification
	merged := make(map[string]interface{})
	for key, value := range spec {
		merged[key] = value
	}

	// Apply inheritance in reverse order (so later specs override earlier ones)
	if inherits, exists := spec["inherits"]; exists {
		if inheritsSlice, ok := inherits.([]interface{}); ok {
			// Reverse the slice
			for i := len(inheritsSlice) - 1; i >= 0; i-- {
				inheritPath := inheritsSlice[i].(string)
				resolvedPath := v.resolveInheritancePath(inheritPath, specPath)
				if inheritedSpec, exists := v.inheritedSpecs[resolvedPath]; exists {
					// Recursively merge inherited spec
					inheritedMerged := v.mergeInheritedSpecifications(inheritedSpec, resolvedPath)
					merged = v.deepMerge(inheritedMerged, merged)
				}
			}
		}
	}

	// Cache the result
	v.mergeCache[specPath] = merged
	return merged
}

// getHierarchyInfo extracts hierarchy information from specification
func (v *APAIValidator) getHierarchyInfo(spec map[string]interface{}) map[string]interface{} {
	info, exists := spec["info"]
	if !exists {
		return make(map[string]interface{})
	}

	infoMap, ok := info.(map[string]interface{})
	if !ok {
		return make(map[string]interface{})
	}

	aiMetadata, exists := infoMap["ai_metadata"]
	if !exists {
		return make(map[string]interface{})
	}

	aiMetadataMap, ok := aiMetadata.(map[string]interface{})
	if !ok {
		return make(map[string]interface{})
	}

	hierarchyInfo, exists := aiMetadataMap["hierarchy_info"]
	if !exists {
		return make(map[string]interface{})
	}

	hierarchyInfoMap, ok := hierarchyInfo.(map[string]interface{})
	if !ok {
		return make(map[string]interface{})
	}

	return hierarchyInfoMap
}

// PrintHierarchyTree prints hierarchy tree for a specification
func (v *APAIValidator) PrintHierarchyTree(specPath string, level int) {
	indent := strings.Repeat("  ", level)

	spec, err := v.loadSpec(specPath)
	if err != nil {
		fmt.Printf("%s❌ Error loading %s: %v\n", indent, specPath, err)
		return
	}

	title := "Unknown"
	if info, exists := spec["info"]; exists {
		if infoMap, ok := info.(map[string]interface{}); ok {
			if titleValue, exists := infoMap["title"]; exists {
				if titleStr, ok := titleValue.(string); ok {
					title = titleStr
				}
			}
		}
	}

	hierarchyInfo := v.getHierarchyInfo(spec)
	levelName := "unknown"
	scope := "unknown"

	if levelValue, exists := hierarchyInfo["level"]; exists {
		if levelStr, ok := levelValue.(string); ok {
			levelName = levelStr
		}
	}

	if scopeValue, exists := hierarchyInfo["scope"]; exists {
		if scopeStr, ok := scopeValue.(string); ok {
			scope = scopeStr
		}
	}

	fmt.Printf("%s📄 %s (%s/%s)\n", indent, title, levelName, scope)
	fmt.Printf("%s   Path: %s\n", indent, specPath)

	if inherits, exists := spec["inherits"]; exists {
		if inheritsSlice, ok := inherits.([]interface{}); ok {
			for _, inheritPath := range inheritsSlice {
				if inheritPathStr, ok := inheritPath.(string); ok {
					resolvedPath := v.resolveInheritancePath(inheritPathStr, specPath)
					v.PrintHierarchyTree(resolvedPath, level+1)
				}
			}
		}
	}
}

// MergeSpecifications merges multiple specifications
func (v *APAIValidator) MergeSpecifications(specs []map[string]interface{}, outputPath, format string) error {
	if len(specs) == 0 {
		return fmt.Errorf("no specifications to merge")
	}

	// Start with first specification
	merged := make(map[string]interface{})
	for key, value := range specs[0] {
		merged[key] = value
	}

	// Merge with remaining specifications
	for i := 1; i < len(specs); i++ {
		merged = v.deepMerge(merged, specs[i])
	}

	// Save merged specification
	var content []byte
	var err error

	if format == "yaml" {
		content, err = yaml.Marshal(merged)
	} else {
		content, err = json.MarshalIndent(merged, "", "  ")
	}

	if err != nil {
		return fmt.Errorf("error marshaling merged specification: %v", err)
	}

	err = ioutil.WriteFile(outputPath, content, 0644)
	if err != nil {
		return fmt.Errorf("error writing output file: %v", err)
	}

	return nil
}
