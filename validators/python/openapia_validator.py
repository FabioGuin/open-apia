#!/usr/bin/env python3
"""
OpenAPIA Validator - Python Implementation

A comprehensive validator for OpenAPIA specifications.
"""

import yaml
import json
import sys
from typing import Dict, List, Any, Optional
from pathlib import Path
import argparse


class OpenAPIAValidator:
    """Main validator class for OpenAPIA specifications."""
    
    def __init__(self):
        self.errors = []
        self.warnings = []
        self.schema_version = "0.1.0"
    
    def validate_file(self, file_path: str) -> bool:
        """Validate an OpenAPIA specification file."""
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                if file_path.endswith('.yaml') or file_path.endswith('.yml'):
                    spec = yaml.safe_load(f)
                elif file_path.endswith('.json'):
                    spec = json.load(f)
                else:
                    self.errors.append(f"Unsupported file format: {file_path}")
                    return False
            
            return self.validate_spec(spec)
        
        except FileNotFoundError:
            self.errors.append(f"File not found: {file_path}")
            return False
        except yaml.YAMLError as e:
            self.errors.append(f"YAML parsing error: {e}")
            return False
        except json.JSONDecodeError as e:
            self.errors.append(f"JSON parsing error: {e}")
            return False
        except Exception as e:
            self.errors.append(f"Unexpected error: {e}")
            return False
    
    def validate_spec(self, spec: Dict[str, Any]) -> bool:
        """Validate an OpenAPIA specification dictionary."""
        self.errors = []
        self.warnings = []
        
        # Validate required sections
        self._validate_required_sections(spec)
        
        # Validate each section
        if 'openapia' in spec:
            self._validate_openapia_version(spec['openapia'])
        
        if 'info' in spec:
            self._validate_info(spec['info'])
        
        if 'models' in spec:
            self._validate_models(spec['models'])
        
        if 'prompts' in spec:
            self._validate_prompts(spec['prompts'])
        
        if 'constraints' in spec:
            self._validate_constraints(spec['constraints'])
        
        if 'tasks' in spec:
            self._validate_tasks(spec['tasks'])
        
        if 'context' in spec:
            self._validate_context(spec['context'])
        
        if 'evaluation' in spec:
            self._validate_evaluation(spec['evaluation'])
        
        # Cross-validation
        self._cross_validate(spec)
        
        return len(self.errors) == 0
    
    def _validate_required_sections(self, spec: Dict[str, Any]) -> None:
        """Validate that all required sections are present."""
        required_sections = [
            'openapia', 'info', 'models', 'prompts', 
            'constraints', 'tasks', 'context', 'evaluation'
        ]
        
        for section in required_sections:
            if section not in spec:
                self.errors.append(f"Missing required section: {section}")
    
    def _validate_openapia_version(self, version: str) -> None:
        """Validate the OpenAPIA version."""
        if not isinstance(version, str):
            self.errors.append("openapia version must be a string")
            return
        
        if not version.startswith('0.1'):
            self.warnings.append(f"Version {version} may not be supported")
    
    def _validate_info(self, info: Dict[str, Any]) -> None:
        """Validate the info section."""
        required_fields = ['title', 'version', 'description']
        
        for field in required_fields:
            if field not in info:
                self.errors.append(f"Missing required field in info: {field}")
        
        if 'ai_metadata' in info:
            self._validate_ai_metadata(info['ai_metadata'])
    
    def _validate_ai_metadata(self, metadata: Dict[str, Any]) -> None:
        """Validate AI-specific metadata."""
        if 'domain' not in metadata:
            self.warnings.append("ai_metadata.domain is recommended")
        
        if 'complexity' in metadata:
            valid_complexities = ['low', 'medium', 'high']
            if metadata['complexity'] not in valid_complexities:
                self.errors.append(f"Invalid complexity: {metadata['complexity']}")
    
    def _validate_models(self, models: List[Dict[str, Any]]) -> None:
        """Validate the models section."""
        if not isinstance(models, list):
            self.errors.append("models must be a list")
            return
        
        if len(models) == 0:
            self.errors.append("At least one model is required")
            return
        
        model_ids = set()
        for i, model in enumerate(models):
            if not isinstance(model, dict):
                self.errors.append(f"Model {i} must be a dictionary")
                continue
            
            # Validate required fields
            required_fields = ['id', 'type', 'provider', 'name', 'purpose']
            for field in required_fields:
                if field not in model:
                    self.errors.append(f"Model {i} missing required field: {field}")
            
            # Check for duplicate IDs
            if 'id' in model:
                if model['id'] in model_ids:
                    self.errors.append(f"Duplicate model ID: {model['id']}")
                model_ids.add(model['id'])
            
            # Validate model type
            if 'type' in model:
                valid_types = ['LLM', 'Vision', 'Audio', 'Multimodal', 'Classification', 'Embedding']
                if model['type'] not in valid_types:
                    self.warnings.append(f"Unknown model type: {model['type']}")
    
    def _validate_prompts(self, prompts: List[Dict[str, Any]]) -> None:
        """Validate the prompts section."""
        if not isinstance(prompts, list):
            self.errors.append("prompts must be a list")
            return
        
        prompt_ids = set()
        for i, prompt in enumerate(prompts):
            if not isinstance(prompt, dict):
                self.errors.append(f"Prompt {i} must be a dictionary")
                continue
            
            # Validate required fields
            required_fields = ['id', 'role', 'template']
            for field in required_fields:
                if field not in prompt:
                    self.errors.append(f"Prompt {i} missing required field: {field}")
            
            # Check for duplicate IDs
            if 'id' in prompt:
                if prompt['id'] in prompt_ids:
                    self.errors.append(f"Duplicate prompt ID: {prompt['id']}")
                prompt_ids.add(prompt['id'])
            
            # Validate role
            if 'role' in prompt:
                valid_roles = ['system', 'user', 'assistant']
                if prompt['role'] not in valid_roles:
                    self.errors.append(f"Invalid prompt role: {prompt['role']}")
    
    def _validate_constraints(self, constraints: List[Dict[str, Any]]) -> None:
        """Validate the constraints section."""
        if not isinstance(constraints, list):
            self.errors.append("constraints must be a list")
            return
        
        constraint_ids = set()
        for i, constraint in enumerate(constraints):
            if not isinstance(constraint, dict):
                self.errors.append(f"Constraint {i} must be a dictionary")
                continue
            
            # Validate required fields
            required_fields = ['id', 'rule', 'severity']
            for field in required_fields:
                if field not in constraint:
                    self.errors.append(f"Constraint {i} missing required field: {field}")
            
            # Check for duplicate IDs
            if 'id' in constraint:
                if constraint['id'] in constraint_ids:
                    self.errors.append(f"Duplicate constraint ID: {constraint['id']}")
                constraint_ids.add(constraint['id'])
            
            # Validate severity
            if 'severity' in constraint:
                valid_severities = ['low', 'medium', 'high', 'critical']
                if constraint['severity'] not in valid_severities:
                    self.errors.append(f"Invalid constraint severity: {constraint['severity']}")
    
    def _validate_tasks(self, tasks: List[Dict[str, Any]]) -> None:
        """Validate the tasks section."""
        if not isinstance(tasks, list):
            self.errors.append("tasks must be a list")
            return
        
        task_ids = set()
        for i, task in enumerate(tasks):
            if not isinstance(task, dict):
                self.errors.append(f"Task {i} must be a dictionary")
                continue
            
            # Validate required fields
            required_fields = ['id', 'description']
            for field in required_fields:
                if field not in task:
                    self.errors.append(f"Task {i} missing required field: {field}")
            
            # Check for duplicate IDs
            if 'id' in task:
                if task['id'] in task_ids:
                    self.errors.append(f"Duplicate task ID: {task['id']}")
                task_ids.add(task['id'])
    
    def _validate_context(self, context: Dict[str, Any]) -> None:
        """Validate the context section."""
        if not isinstance(context, dict):
            self.errors.append("context must be a dictionary")
            return
        
        if 'memory' not in context:
            self.warnings.append("context.memory is recommended")
    
    def _validate_evaluation(self, evaluation: Dict[str, Any]) -> None:
        """Validate the evaluation section."""
        if not isinstance(evaluation, dict):
            self.errors.append("evaluation must be a dictionary")
            return
        
        if 'metrics' not in evaluation:
            self.warnings.append("evaluation.metrics is recommended")
    
    def _cross_validate(self, spec: Dict[str, Any]) -> None:
        """Perform cross-validation between sections."""
        # Validate that referenced models exist
        if 'tasks' in spec and 'models' in spec:
            model_ids = {model.get('id') for model in spec['models'] if 'id' in model}
            
            for task in spec['tasks']:
                if 'steps' in task:
                    for step in task['steps']:
                        if 'model' in step and step['model'] not in model_ids:
                            self.errors.append(f"Task references unknown model: {step['model']}")
        
        # Validate that referenced prompts exist
        if 'tasks' in spec and 'prompts' in spec:
            prompt_ids = {prompt.get('id') for prompt in spec['prompts'] if 'id' in prompt}
            
            for task in spec['tasks']:
                if 'steps' in task:
                    for step in task['steps']:
                        if 'prompt' in step and step['prompt'] not in prompt_ids:
                            self.errors.append(f"Task references unknown prompt: {step['prompt']}")
    
    def get_errors(self) -> List[str]:
        """Get list of validation errors."""
        return self.errors
    
    def get_warnings(self) -> List[str]:
        """Get list of validation warnings."""
        return self.warnings
    
    def print_results(self) -> None:
        """Print validation results."""
        if self.errors:
            print("❌ Validation Errors:")
            for error in self.errors:
                print(f"  - {error}")
        
        if self.warnings:
            print("⚠️  Validation Warnings:")
            for warning in self.warnings:
                print(f"  - {warning}")
        
        if not self.errors and not self.warnings:
            print("✅ Validation passed with no issues")
        elif not self.errors:
            print("✅ Validation passed with warnings")


def main():
    """Main CLI entry point."""
    parser = argparse.ArgumentParser(description='Validate OpenAPIA specifications')
    parser.add_argument('file', help='OpenAPIA specification file to validate')
    parser.add_argument('--json', action='store_true', help='Output results as JSON')
    parser.add_argument('--quiet', action='store_true', help='Only output errors')
    
    args = parser.parse_args()
    
    validator = OpenAPIAValidator()
    is_valid = validator.validate_file(args.file)
    
    if args.json:
        result = {
            'valid': is_valid,
            'errors': validator.get_errors(),
            'warnings': validator.get_warnings()
        }
        print(json.dumps(result, indent=2))
    elif not args.quiet:
        validator.print_results()
    
    sys.exit(0 if is_valid else 1)


if __name__ == '__main__':
    main()
