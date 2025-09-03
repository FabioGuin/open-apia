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
        
        # Hierarchical composition properties
        self.inherited_specs = {}
        self.merge_cache = {}
    
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
            
            # Validate task steps if present
            if 'steps' in task and isinstance(task['steps'], list):
                self._validate_task_steps(task['steps'], i)
    
    def _validate_task_steps(self, steps: List[Dict[str, Any]], task_index: int) -> None:
        """Validate task steps."""
        for step_index, step in enumerate(steps):
            if not isinstance(step, dict):
                self.errors.append(f"Task {task_index} step {step_index} must be a dictionary")
                continue
            
            # Validate required fields
            required_fields = ['name', 'action']
            for field in required_fields:
                if field not in step:
                    self.errors.append(f"Task {task_index} step {step_index} missing required field: {field}")
            
            # Validate action type
            if 'action' in step:
                valid_actions = ['analyze', 'generate', 'validate', 'search', 'escalate', 'classify', 'mcp_tool', 'mcp_resource']
                if step['action'] not in valid_actions:
                    self.warnings.append(f"Task {task_index} step {step_index} unknown action: {step['action']}")
            
            # Validate MCP-specific fields
            if 'action' in step and step['action'] in ['mcp_tool', 'mcp_resource']:
                if 'mcp_server' not in step:
                    self.errors.append(f"Task {task_index} step {step_index} MCP action missing mcp_server field")
                
                if step['action'] == 'mcp_tool' and 'mcp_tool' not in step:
                    self.errors.append(f"Task {task_index} step {step_index} mcp_tool action missing mcp_tool field")
                
                if step['action'] == 'mcp_resource' and 'mcp_resource' not in step:
                    self.errors.append(f"Task {task_index} step {step_index} mcp_resource action missing mcp_resource field")
    
    def _validate_context(self, context: Dict[str, Any]) -> None:
        """Validate the context section."""
        if not isinstance(context, dict):
            self.errors.append("context must be a dictionary")
            return
        
        if 'memory' not in context:
            self.warnings.append("context.memory is recommended")
        
        # Validate MCP servers if present
        if 'mcp_servers' in context:
            self._validate_mcp_servers(context['mcp_servers'])
    
    def _validate_mcp_servers(self, mcp_servers: List[Dict[str, Any]]) -> None:
        """Validate MCP servers section."""
        if not isinstance(mcp_servers, list):
            self.errors.append("mcp_servers must be a list")
            return
        
        server_ids = set()
        for index, server in enumerate(mcp_servers):
            if not isinstance(server, dict):
                self.errors.append(f"MCP server {index} must be a dictionary")
                continue
            
            # Validate required fields
            required_fields = ['id', 'name', 'description', 'version', 'transport', 'capabilities', 'authentication']
            for field in required_fields:
                if field not in server:
                    self.errors.append(f"MCP server {index} missing required field: {field}")
            
            # Check for duplicate IDs
            if 'id' in server:
                if server['id'] in server_ids:
                    self.errors.append(f"Duplicate MCP server ID: {server['id']}")
                server_ids.add(server['id'])
            
            # Validate transport configuration
            if 'transport' in server and isinstance(server['transport'], dict):
                self._validate_mcp_transport(server['transport'], index)
            
            # Validate authentication configuration
            if 'authentication' in server and isinstance(server['authentication'], dict):
                self._validate_mcp_authentication(server['authentication'], index)
    
    def _validate_mcp_transport(self, transport: Dict[str, Any], server_index: int) -> None:
        """Validate MCP transport configuration."""
        if 'type' not in transport:
            self.errors.append(f"MCP server {server_index} transport missing required field: type")
            return
        
        valid_types = ['stdio', 'sse', 'websocket']
        if transport['type'] not in valid_types:
            self.errors.append(f"MCP server {server_index} invalid transport type: {transport['type']}")
        
        # Validate transport-specific fields
        if transport['type'] == 'stdio':
            if 'command' not in transport:
                self.errors.append(f"MCP server {server_index} stdio transport missing command")
        elif transport['type'] in ['sse', 'websocket']:
            if 'url' not in transport:
                self.errors.append(f"MCP server {server_index} {transport['type']} transport missing url")
    
    def _validate_mcp_authentication(self, auth: Dict[str, Any], server_index: int) -> None:
        """Validate MCP authentication configuration."""
        if 'type' not in auth:
            self.errors.append(f"MCP server {server_index} authentication missing required field: type")
            return
        
        valid_types = ['none', 'api_key', 'oauth', 'custom']
        if auth['type'] not in valid_types:
            self.errors.append(f"MCP server {server_index} invalid authentication type: {auth['type']}")
        
        # Validate authentication-specific fields
        if auth['type'] == 'api_key' and 'api_key' not in auth:
            self.warnings.append(f"MCP server {server_index} api_key authentication missing api_key field")
        
        if auth['type'] == 'oauth' and 'token' not in auth:
            self.warnings.append(f"MCP server {server_index} oauth authentication missing token field")
    
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
        
        # Validate that referenced MCP servers exist
        if 'tasks' in spec and 'context' in spec and 'mcp_servers' in spec['context']:
            mcp_server_ids = {server.get('id') for server in spec['context']['mcp_servers'] if 'id' in server}
            
            for task in spec['tasks']:
                if 'steps' in task:
                    for step in task['steps']:
                        if 'mcp_server' in step and step['mcp_server'] not in mcp_server_ids:
                            self.errors.append(f"Task references unknown MCP server: {step['mcp_server']}")
    
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
    
    def get_results(self) -> Dict[str, Any]:
        """Get validation results as dictionary."""
        return {
            'valid': len(self.errors) == 0,
            'errors': self.errors,
            'warnings': self.warnings
        }
    
    # ============================================================================
    # HIERARCHICAL COMPOSITION METHODS
    # ============================================================================
    
    def validate_with_inheritance(self, file_path: str) -> bool:
        """Validate specification with inheritance support."""
        try:
            spec = self.load_spec(file_path)
            if spec is None:
                return False
            
            # Load and merge inherited specifications
            merged_spec = self.merge_inherited_specifications(spec, file_path)
            
            # Validate merged specification
            return self.validate_spec(merged_spec)
            
        except Exception as e:
            self.errors.append(f"Unexpected error during hierarchical validation: {e}")
            return False
    
    def load_spec(self, spec_path: str) -> Optional[Dict[str, Any]]:
        """Load OpenAPIA specification from file."""
        try:
            with open(spec_path, 'r', encoding='utf-8') as f:
                if spec_path.endswith('.yaml') or spec_path.endswith('.yml'):
                    return yaml.safe_load(f)
                elif spec_path.endswith('.json'):
                    return json.load(f)
                else:
                    self.errors.append(f"Unsupported file format: {spec_path}")
                    return None
        except FileNotFoundError:
            self.errors.append(f"File not found: {spec_path}")
            return None
        except yaml.YAMLError as e:
            self.errors.append(f"YAML parsing error in {spec_path}: {e}")
            return None
        except json.JSONDecodeError as e:
            self.errors.append(f"JSON parsing error in {spec_path}: {e}")
            return None
        except Exception as e:
            self.errors.append(f"Error loading specification {spec_path}: {e}")
            return None
    
    def resolve_inheritance_path(self, inherit_path: str, current_spec_path: str) -> str:
        """Resolve inheritance path to absolute path."""
        current_dir = os.path.dirname(current_spec_path)
        return os.path.join(current_dir, inherit_path)
    
    def load_inherited_specs(self, spec: Dict[str, Any], spec_path: str) -> None:
        """Load all inherited specifications."""
        if 'inherits' not in spec or not isinstance(spec['inherits'], list):
            return
        
        for inherit_path in spec['inherits']:
            resolved_path = self.resolve_inheritance_path(inherit_path, spec_path)
            
            if resolved_path in self.inherited_specs:
                continue  # Already loaded
            
            inherited_spec = self.load_spec(resolved_path)
            if inherited_spec is not None:
                self.inherited_specs[resolved_path] = inherited_spec
                
                # Recursively load inherited specs
                self.load_inherited_specs(inherited_spec, resolved_path)
            else:
                self.errors.append(f"Inherited specification not found: {inherit_path}")
    
    def deep_merge(self, base: Dict[str, Any], override: Dict[str, Any]) -> Dict[str, Any]:
        """Deep merge two dictionaries."""
        result = base.copy()
        
        for key, value in override.items():
            if key in result and isinstance(result[key], dict) and isinstance(value, dict):
                result[key] = self.deep_merge(result[key], value)
            else:
                result[key] = value
        
        return result
    
    def merge_inherited_specifications(self, spec: Dict[str, Any], spec_path: str) -> Dict[str, Any]:
        """Merge specifications based on inheritance."""
        if spec_path in self.merge_cache:
            return self.merge_cache[spec_path]
        
        # Load inherited specifications
        self.load_inherited_specs(spec, spec_path)
        
        # Start with base specification
        merged = spec.copy()
        
        # Apply inheritance in reverse order (so later specs override earlier ones)
        if 'inherits' in spec and isinstance(spec['inherits'], list):
            for inherit_path in reversed(spec['inherits']):
                resolved_path = self.resolve_inheritance_path(inherit_path, spec_path)
                if resolved_path in self.inherited_specs:
                    inherited_spec = self.inherited_specs[resolved_path]
                    # Recursively merge inherited spec
                    inherited_merged = self.merge_inherited_specifications(inherited_spec, resolved_path)
                    merged = self.deep_merge(inherited_merged, merged)
        
        # Cache the result
        self.merge_cache[spec_path] = merged
        return merged
    
    def get_hierarchy_info(self, spec: Dict[str, Any]) -> Dict[str, Any]:
        """Get hierarchy information from specification."""
        hierarchy_info = {}
        
        if 'info' in spec and 'ai_metadata' in spec['info'] and 'hierarchy_info' in spec['info']['ai_metadata']:
            hierarchy_info = spec['info']['ai_metadata']['hierarchy_info']
        
        return hierarchy_info
    
    def print_hierarchy_tree(self, spec_path: str, level: int = 0) -> None:
        """Print hierarchy tree for a specification."""
        indent = "  " * level
        
        try:
            spec = self.load_spec(spec_path)
            if spec is None:
                print(f"{indent}❌ Error loading {spec_path}")
                return
            
            title = spec.get('info', {}).get('title', 'Unknown')
            hierarchy_info = self.get_hierarchy_info(spec)
            level_name = hierarchy_info.get('level', 'unknown')
            scope = hierarchy_info.get('scope', 'unknown')
            
            print(f"{indent}📄 {title} ({level_name}/{scope})")
            print(f"{indent}   Path: {spec_path}")
            
            if 'inherits' in spec and isinstance(spec['inherits'], list):
                for inherit_path in spec['inherits']:
                    resolved_path = self.resolve_inheritance_path(inherit_path, spec_path)
                    self.print_hierarchy_tree(resolved_path, level + 1)
        except Exception as e:
            print(f"{indent}❌ Error loading {spec_path}: {e}")
    
    def merge_specifications(self, specs: List[Dict[str, Any]], output_path: str, format: str = 'yaml') -> bool:
        """Merge multiple specifications."""
        try:
            if not specs:
                self.errors.append("No specifications to merge")
                return False
            
            # Start with first specification
            merged = specs[0].copy()
            
            # Merge with remaining specifications
            for i in range(1, len(specs)):
                merged = self.deep_merge(merged, specs[i])
            
            # Save merged specification
            with open(output_path, 'w', encoding='utf-8') as f:
                if format == 'yaml':
                    yaml.dump(merged, f, default_flow_style=False, indent=2)
                else:
                    json.dump(merged, f, indent=2, ensure_ascii=False)
            
            return True
            
        except Exception as e:
            self.errors.append(f"Error merging specifications: {e}")
            return False


def main():
    """Main CLI entry point."""
    args = sys.argv[1:]
    
    if len(args) == 0 or args[0] in ['-h', '--help', 'help']:
        show_help()
        sys.exit(0)
    
    command = args[0]
    options = args[1:]
    
    if command == 'validate':
        handle_validate(options)
    elif command == 'tree':
        handle_tree(options)
    elif command == 'merge':
        handle_merge(options)
    else:
        print(f"Unknown command: {command}")
        show_help()
        sys.exit(1)

def handle_validate(options):
    """Handle validate command."""
    if len(options) == 0:
        print("Error: No file specified")
        print("Usage: python openapia_validator.py validate <file> [--hierarchical]")
        sys.exit(1)
    
    file_path = options[0]
    hierarchical = '--hierarchical' in options
    
    print(f"Validating OpenAPIA specification{' with inheritance' if hierarchical else ''}: {file_path}")
    print('-' * 60)
    
    validator = OpenAPIAValidator()
    
    if hierarchical:
        is_valid = validator.validate_with_inheritance(file_path)
    else:
        is_valid = validator.validate_file(file_path)
    
    if is_valid:
        print("✅ Validation successful!")
    else:
        print("❌ Validation failed!")
        print("\nErrors:")
        for error in validator.get_errors():
            print(f"  • {error}")
    
    warnings = validator.get_warnings()
    if warnings:
        print("\nWarnings:")
        for warning in warnings:
            print(f"  ⚠️  {warning}")
    
    sys.exit(0 if is_valid else 1)

def handle_tree(options):
    """Handle tree command."""
    if len(options) == 0:
        print("Error: No file specified")
        print("Usage: python openapia_validator.py tree <file>")
        sys.exit(1)
    
    file_path = options[0]
    
    print("OpenAPIA Specification Hierarchy Tree")
    print("=" * 50)
    
    validator = OpenAPIAValidator()
    validator.print_hierarchy_tree(file_path)

def handle_merge(options):
    """Handle merge command."""
    if len(options) < 2:
        print("Error: Missing required arguments")
        print("Usage: python openapia_validator.py merge <output> <file1> [file2] ...")
        sys.exit(1)
    
    output_path = options[0]
    input_files = options[1:]
    
    print("Merging OpenAPIA specifications...")
    print(f"Output: {output_path}")
    print(f"Input files: {', '.join(input_files)}")
    print('-' * 60)
    
    validator = OpenAPIAValidator()
    specs = []
    
    for file in input_files:
        if not os.path.exists(file):
            print(f"Error: File not found: {file}")
            sys.exit(1)
        
        spec = validator.load_spec(file)
        if not spec:
            print(f"Error: Cannot load specification: {file}")
            sys.exit(1)
        
        specs.append(spec)
        print(f"✅ Loaded: {file}")
    
    format_type = 'yaml'
    if output_path.endswith('.json'):
        format_type = 'json'
    
    success = validator.merge_specifications(specs, output_path, format_type)
    
    if success:
        print("\n✅ Merge completed successfully!")
        print(f"Merged specification saved to: {output_path}")
    else:
        print("\n❌ Merge failed!")
        for error in validator.get_errors():
            print(f"  • {error}")
        sys.exit(1)

def show_help():
    """Show help information."""
    print("OpenAPIA Validator CLI - Python Implementation")
    print("=" * 50)
    print("")
    
    print("USAGE:")
    print("  python openapia_validator.py <command> [options]")
    print("")
    
    print("COMMANDS:")
    print("  validate <file> [--hierarchical]  Validate OpenAPIA specification")
    print("  tree <file>                       Show hierarchy tree for specification")
    print("  merge <output> <files...>         Merge multiple specifications")
    print("")
    
    print("OPTIONS:")
    print("  --hierarchical                   Use hierarchical validation with inheritance")
    print("  -h, --help                       Show this help message")
    print("")
    
    print("EXAMPLES:")
    print("  python openapia_validator.py validate spec.yaml")
    print("  python openapia_validator.py validate spec.yaml --hierarchical")
    print("  python openapia_validator.py tree spec.yaml")
    print("  python openapia_validator.py merge output.yaml spec1.yaml spec2.yaml")
    print("")
    
    print("For more information, visit: https://github.com/openapia/openapia")


if __name__ == '__main__':
    main()
