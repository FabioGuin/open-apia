# APAI Hierarchical Composition

APAI supports hierarchical composition of specifications, enabling the creation of complex AI systems through inheritance and specialization of configurations.

## Table of Contents

1. [Key Concepts](#key-concepts)
   - [1. Inheritance](#1-inheritance)
   - [2. Composition](#2-composition)
   - [3. Specialization](#3-specialization)

2. [Hierarchical Structure](#hierarchical-structure)
   - [Organizational Levels](#organizational-levels)
   - [File Organization](#file-organization)

3. [Inheritance Configuration](#inheritance-configuration)
   - [Basic Inheritance](#basic-inheritance)
   - [Multiple Inheritance](#multiple-inheritance)
   - [Inheritance Modes](#inheritance-modes)

4. [Composition Algorithms](#composition-algorithms)
   - [Merge Mode](#merge-mode)
   - [Override Mode](#override-mode)
   - [Extend Mode](#extend-mode)

5. [Hierarchy Information](#hierarchy-information)
   - [Level Definition](#level-definition)
   - [Scope Definition](#scope-definition)
   - [Parent Specifications](#parent-specifications)

6. [Best Practices](#best-practices)
   - [Design Principles](#design-principles)
   - [Naming Conventions](#naming-conventions)
   - [File Organization](#file-organization)

7. [Advanced Patterns](#advanced-patterns)
   - [Cross-Department Inheritance](#cross-department-inheritance)
   - [Environment-Specific Configurations](#environment-specific-configurations)
   - [Feature Flags and Conditional Inheritance](#feature-flags-and-conditional-inheritance)

8. [Tooling and Automation](#tooling-and-automation)
   - [Validation Tools](#validation-tools)
   - [Composition Tools](#composition-tools)
   - [Documentation Generation](#documentation-generation)

9. [Examples](#examples)
   - [Simple Inheritance](#simple-inheritance)
   - [Complex Multi-Level Hierarchy](#complex-multi-level-hierarchy)
   - [Cross-Organizational Inheritance](#cross-organizational-inheritance)

10. [Troubleshooting](#troubleshooting)
    - [Common Issues](#common-issues)
    - [Debugging Tips](#debugging-tips)

---

## Key Concepts

### 1. **Inheritance**
APAI specifications can inherit from other specifications, creating a hierarchy of configurations.

### 2. **Composition**
Inherited specifications are "composed" (merged) with the local specification, with priority given to the local specification.

### 3. **Specialization**
Each level of the hierarchy can specialize or extend inherited configurations.

## Hierarchical Structure

```
organization/
├── apai-global.yaml          # Level: global
├── apai-regional/
│   ├── apai-eu.yaml          # Level: regional
│   └── apai-us.yaml          # Level: regional
├── apai-department/
│   ├── apai-cs.yaml          # Level: department
│   └── apai-sales.yaml       # Level: department
├── apai-team/
│   ├── apai-dev-team.yaml    # Level: team
│   └── apai-qa-team.yaml     # Level: team
├── apai-sprint/
│   ├── apai-sprint-24.yaml   # Level: sprint
│   └── apai-sprint-25.yaml   # Level: sprint
├── features/
│   ├── feature-chatbot/
│   │   └── apai.yaml         # Level: feature
│   └── feature-sentiment/
│       └── apai.yaml         # Level: feature
└── environments/
    ├── dev/
    │   └── apai.yaml         # Level: environment
    ├── staging/
    │   └── apai.yaml         # Level: environment
    └── prod/
        └── apai.yaml         # Level: environment
```

## Hierarchy Levels

### **Global** (Organization)
- **Scope**: Entire organization
- **Content**: Global models, constraints, and policies
- **Example**: `apai-global.yaml`

### **Regional** (Region)
- **Scope**: Geographic region
- **Content**: Regional compliance, localized models
- **Example**: `apai-eu.yaml`, `apai-us.yaml`

### **Department** (Department)
- **Scope**: Business department
- **Content**: Department-specific models and constraints
- **Example**: `apai-customer-support.yaml`

### **Team** (Team)
- **Scope**: Development team
- **Content**: Team standards, development models
- **Example**: `apai-dev-team.yaml`

### **Sprint** (Sprint)
- **Scope**: Development sprint
- **Content**: Sprint-specific configuration
- **Example**: `apai-sprint-24.yaml`

### **Feature** (Feature)
- **Scope**: Specific feature
- **Content**: Feature-specific models and constraints
- **Example**: `features/chatbot/apai.yaml`

### **Environment** (Environment)
- **Scope**: Deployment environment
- **Content**: Environment-specific configuration
- **Example**: `environments/prod/apai.yaml`

## Inheritance Syntax

### **`inherits` Field**
```yaml
# Local specification
apai: "0.1.0"

# Inheritance from parent specifications
inherits:
  - "../apai-global.yaml"      # Inherit from global
  - "../../apai-team.yaml"     # Inherit from team
  - "../apai-sprint-24.yaml"   # Inherit from sprint

info:
  title: "Feature Chatbot"
  # ... rest of configuration
```

### **Hierarchical Metadata**
```yaml
info:
  ai_metadata:
    hierarchy_info:
      level: "feature"                    # Level in hierarchy
      parent_specs:                       # Parent specifications
        - "../apai-global.yaml"
        - "../../apai-team.yaml"
      scope: "project"                    # Specification scope
      inheritance_mode: "merge"           # Inheritance mode
```

## Inheritance Modes

### **1. Merge (Default)**
Inherited specifications are "merged" with the local specification.

```yaml
# Parent: apai-global.yaml
models:
  - id: "global_llm"
    type: "LLM"
    provider: "openai"
    name: "gpt-4"

# Child: apai-feature.yaml
inherits:
  - "../apai-global.yaml"

models:
  - id: "feature_specific_model"  # Adds new model
    type: "Classification"
    provider: "huggingface"

# Result: Both models are available
```

### **2. Override**
The local specification completely overrides inherited specifications.

```yaml
# Child: apai-feature.yaml
inherits:
  - "../apai-global.yaml"

info:
  ai_metadata:
    hierarchy_info:
      inheritance_mode: "override"

models:
  - id: "feature_llm"  # Completely replaces parent models
    type: "LLM"
    provider: "anthropic"
    name: "claude-3"
```

### **3. Extend**
The local specification extends inherited specifications without overriding them.

```yaml
# Child: apai-feature.yaml
inherits:
  - "../apai-global.yaml"

info:
  ai_metadata:
    hierarchy_info:
      inheritance_mode: "extend"

models:
  - id: "feature_extension"  # Adds only new models
    type: "Classification"
    provider: "custom"
```

## Composition Algorithm

### **Application Order**
1. **Global** (base)
2. **Regional** (override global)
3. **Department** (override regional)
4. **Team** (override department)
5. **Sprint** (override team)
6. **Feature** (override sprint)
7. **Environment** (override feature)

### **Merge Rules**
- **Models**: Merge by ID, local has priority
- **Constraints**: Merge by ID, local has priority
- **Tasks**: Merge by ID, local has priority
- **Context**: Deep merge, local has priority
- **Evaluation**: Merge by name, local has priority

## Practical Examples

### **Example 1: Feature inheriting from Team**
```yaml
# apai-team.yaml
apai: "0.1.0"
info:
  title: "Development Team Standards"
  ai_metadata:
    hierarchy_info:
      level: "team"
      scope: "team"

models:
  - id: "team_llm"
    type: "LLM"
    provider: "openai"
    name: "gpt-4"

constraints:
  - id: "team_safety"
    type: "content_safety"
    rule: "output NOT contains harmful_content"
    severity: "medium"

# features/chatbot/apai.yaml
apai: "0.1.0"
inherits:
  - "../../apai-team.yaml"

info:
  title: "Chatbot Feature"
  ai_metadata:
    hierarchy_info:
      level: "feature"
      scope: "project"
      parent_specs: ["../../apai-team.yaml"]
      inheritance_mode: "merge"

models:
  - id: "chatbot_specific_model"  # Adds new model
    type: "Classification"
    provider: "huggingface"

constraints:
  - id: "chatbot_response_time"   # Adds new constraint
    type: "performance"
    rule: "response_time < 3s"
    severity: "high"
```

### **Example 2: Environment inheriting from Feature**
```yaml
# environments/prod/apai.yaml
apai: "0.1.0"
inherits:
  - "../apai-feature.yaml"

info:
  title: "Production Environment"
  ai_metadata:
    hierarchy_info:
      level: "environment"
      scope: "environment"
      parent_specs: ["../apai-feature.yaml"]
      inheritance_mode: "merge"

constraints:
  - id: "prod_performance"        # Override existing constraint
    type: "performance"
    rule: "response_time < 2s"    # More strict for prod
    severity: "critical"
  
  - id: "prod_security"           # Adds prod-specific constraint
    type: "content_safety"
    rule: "output NOT contains security_vulnerabilities"
    severity: "critical"
```

## Hierarchical Validation

### **Validation with Inheritance**
```bash
# Simple validation
python validators/python/apai_validator.py validate apai.yaml

# Validation with inheritance
python validators/python/apai_validator.py validate apai.yaml --hierarchical

# Validation of entire hierarchy
python validators/python/apai_validator.py validate . --recursive
```

### **Automatic Validation**
The APAI validator automatically supports inheritance:

```python
from apai_validator import APAIValidator

validator = APAIValidator()

# Validation with inheritance
result = validator.validate_with_inheritance("features/chatbot/apai.yaml")

if result.is_valid:
    print("✅ Hierarchical validation completed")
else:
    print("❌ Validation errors:")
    for error in result.errors:
        print(f"  - {error}")
```

## Best Practices

### **1. Hierarchical Structure**
- Maintain a logical and consistent hierarchy
- Use descriptive names for levels
- Document dependencies between specifications

### **2. Inheritance**
- Use `merge` for most cases
- Use `override` only when necessary
- Avoid inheritance chains that are too long

### **3. Naming Convention**
- `apai-global.yaml` for global level
- `apai-{level}.yaml` for specific levels
- `apai.yaml` for local specifications

### **4. Versioning**
- Maintain compatible versions between levels
- Document breaking changes
- Use semantic versioning

### **5. Testing**
- Test each level of the hierarchy
- Validate the final composition
- Test fallback scenarios

## Tooling and Automation

### **CLI Commands**
```bash
# Validate with inheritance
python validators/python/apai_validator.py validate spec.yaml --hierarchical
node validators/javascript/cli.js validate spec.yaml --hierarchical
php validators/php/cli.php validate spec.yaml --hierarchical
go run validators/go/cli.go validate spec.yaml --hierarchical

# Show hierarchy tree
python validators/python/apai_validator.py tree spec.yaml
node validators/javascript/cli.js tree spec.yaml
php validators/php/cli.php tree spec.yaml
go run validators/go/cli.go tree spec.yaml

# Merge specifications
python validators/python/apai_validator.py merge output.yaml spec1.yaml spec2.yaml
node validators/javascript/cli.js merge output.yaml spec1.yaml spec2.yaml
php validators/php/cli.php merge output.yaml spec1.yaml spec2.yaml
go run validators/go/cli.go merge output.yaml spec1.yaml spec2.yaml
```

### **CI/CD Integration**
```yaml
# .github/workflows/apai-hierarchy.yml
name: APAI Hierarchical Validation

on:
  pull_request:
    paths:
      - '**/apai*.yaml'

jobs:
  validate-hierarchy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Python
        uses: actions/setup-python@v4
        with:
          python-version: '3.9'
      - name: Install dependencies
        run: pip install pyyaml jsonschema
      - name: Validate Global Specs
        run: python validators/python/apai_validator.py validate apai-global.yaml --hierarchical
      
      - name: Validate Team Specs
        run: python validators/python/apai_validator.py validate apai-team.yaml --hierarchical
      
      - name: Validate Feature Specs
        run: |
          for spec in features/*/apai.yaml; do
            python validators/python/apai_validator.py validate "$spec" --hierarchical
          done
      
      - name: Validate Environment Specs
        run: |
          for spec in environments/*/apai.yaml; do
            python validators/python/apai_validator.py validate "$spec" --hierarchical
          done
```

## Conclusions

APAI's hierarchical composition enables:

1. **Reuse** common configurations
2. **Specialize** for specific contexts
3. **Maintain** organizational consistency
4. **Scale** complex AI systems
5. **Govern** specification evolution

This approach makes APAI suitable for both small teams and large enterprise organizations.
