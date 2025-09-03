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

// OpenAPIAValidator represents the main validator struct
type OpenAPIAValidator struct {
	Errors      []string
	Warnings    []string
	SchemaVersion string
}

// ValidationResult represents the result of validation
type ValidationResult struct {
	Valid    bool     `json:"valid"`
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
}

// NewOpenAPIAValidator creates a new validator instance
func NewOpenAPIAValidator() *OpenAPIAValidator {
	return &OpenAPIAValidator{
		Errors:        make([]string, 0),
		Warnings:      make([]string, 0),
		SchemaVersion: "0.1.0",
	}
}

// ValidateFile validates an OpenAPIA specification file
func (v *OpenAPIAValidator) ValidateFile(filePath string) (bool, error) {
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
func (v *OpenAPIAValidator) ValidateSpec(spec map[string]interface{}) bool {
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
func (v *OpenAPIAValidator) validateRequiredSections(spec map[string]interface{}) {
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
func (v *OpenAPIAValidator) validateOpenAPIAVersion(version interface{}) {
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
func (v *OpenAPIAValidator) validateInfo(info interface{}) {
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
func (v *OpenAPIAValidator) validateAIMetadata(metadata interface{}) {
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
func (v *OpenAPIAValidator) validateModels(models interface{}) {
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
func (v *OpenAPIAValidator) validatePrompts(prompts interface{}) {
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
func (v *OpenAPIAValidator) validateConstraints(constraints interface{}) {
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
func (v *OpenAPIAValidator) validateTasks(tasks interface{}) {
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
	}
}

// validateContext validates the context section
func (v *OpenAPIAValidator) validateContext(context interface{}) {
	contextMap, ok := context.(map[string]interface{})
	if !ok {
		v.Errors = append(v.Errors, "context must be an object")
		return
	}

	if _, exists := contextMap["memory"]; !exists {
		v.Warnings = append(v.Warnings, "context.memory is recommended")
	}
}

// validateEvaluation validates the evaluation section
func (v *OpenAPIAValidator) validateEvaluation(evaluation interface{}) {
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
func (v *OpenAPIAValidator) crossValidate(spec map[string]interface{}) {
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
}

// GetErrors returns the list of validation errors
func (v *OpenAPIAValidator) GetErrors() []string {
	return v.Errors
}

// GetWarnings returns the list of validation warnings
func (v *OpenAPIAValidator) GetWarnings() []string {
	return v.Warnings
}

// PrintResults prints validation results
func (v *OpenAPIAValidator) PrintResults() {
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
func (v *OpenAPIAValidator) GetResults() ValidationResult {
	return ValidationResult{
		Valid:    len(v.Errors) == 0,
		Errors:   v.Errors,
		Warnings: v.Warnings,
	}
}
