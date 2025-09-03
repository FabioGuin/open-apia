# OpenAPIA 0.1 Specification

## Machine-Readable Formats

This specification is available in multiple formats for different use cases:

### Main Specification
- **YAML**: `spec/openapia-0.1.yaml` - Official specification (human-readable)

### Examples and Templates
- **JSON Example**: `spec/examples/openapia-0.1-example.json` - Complete working example in JSON format
- **Domain Examples**: `spec/examples/*.yaml` - Real-world use case examples

### Validation
- **JSON Schema**: `spec/schemas/openapia-0.1-schema.json` - Validation schema for OpenAPIA specifications

See `spec/README.md` for detailed explanation of each file type and usage.

## Table of Contents

1. [Introduction](#introduction)
2. [Specification Metadata](#specification-metadata)
3. [Hierarchical Composition](#hierarchical-composition)
4. [AI System Metadata](#ai-system-metadata)
5. [AI Models](#ai-models)
6. [Prompts](#prompts)
7. [Constraints](#constraints)
8. [Tasks](#tasks)
9. [Automations](#automations)
10. [Context](#context)
11. [Evaluation](#evaluation)
12. [Extensions](#extensions)
13. [Validation](#validation)
14. [Governance](#governance)

## Introduction

OpenAPIA (Open Artificial Intelligence Architecture) is an open standard for describing, documenting, and validating AI systems. This specification provides a comprehensive framework for defining AI architectures in a vendor-agnostic, AI-native manner.

The OpenAPIA 0.1 specification is designed to:
- Enable clear documentation of AI system architectures
- Support hierarchical composition of AI specifications
- Provide structured validation of AI system configurations
- Facilitate interoperability between different AI systems and providers
- Ensure compliance with ethical and safety standards

## Specification Metadata

### openapia

**Type:** `string`  
**Required:** Yes  
**Description:** The OpenAPIA specification version being used.

**Example:**
```yaml
openapia: "0.1.0"
```

This field indicates which version of the OpenAPIA specification the document conforms to, enabling proper validation and tooling support.

## Hierarchical Composition

### inherits

**Type:** `array[string]`  
**Required:** No  
**Description:** List of OpenAPIA specifications to inherit from, enabling hierarchical composition of AI system definitions.

**Example:**
```yaml
inherits: 
  - "../openapia-global.yaml"
  - "../../openapia-team.yaml"
```

This field allows AI systems to inherit configurations from parent specifications, supporting organizational hierarchies and shared configurations. The inheritance follows a merge strategy where child specifications can override or extend parent configurations.

## AI System Metadata

### info

**Type:** `object`  
**Required:** Yes  
**Description:** Contains metadata about the AI system being described.

#### info.title

**Type:** `string`  
**Required:** Yes  
**Description:** The name of the AI system.

#### info.version

**Type:** `string`  
**Required:** Yes  
**Description:** The version of the AI system.

#### info.description

**Type:** `string`  
**Required:** Yes  
**Description:** A detailed description of the AI system's purpose and functionality.

#### info.author

**Type:** `string`  
**Required:** Yes  
**Description:** The author or team responsible for the AI system.

#### info.license

**Type:** `string`  
**Required:** Yes  
**Description:** The license identifier for the AI system.

#### info.contact

**Type:** `object`  
**Required:** No  
**Description:** Contact information for the AI system.

##### info.contact.email

**Type:** `string`  
**Required:** No  
**Description:** Contact email address.

##### info.contact.url

**Type:** `string`  
**Required:** No  
**Description:** Contact URL or website.

#### info.ai_metadata

**Type:** `object`  
**Required:** Yes  
**Description:** AI-specific metadata for the system.

##### info.ai_metadata.domain

**Type:** `string`  
**Required:** Yes  
**Description:** The application domain of the AI system (e.g., "customer_service", "content_generation", "data_analysis").

##### info.ai_metadata.complexity

**Type:** `string`  
**Required:** Yes  
**Enum:** `["low", "medium", "high"]`  
**Description:** The complexity level of the AI system.

##### info.ai_metadata.deployment

**Type:** `string`  
**Required:** Yes  
**Enum:** `["development", "staging", "production"]`  
**Description:** The deployment environment for the AI system.

##### info.ai_metadata.last_updated

**Type:** `string`  
**Required:** Yes  
**Format:** ISO 8601 timestamp  
**Description:** The last update timestamp of the AI system specification.

##### info.ai_metadata.supported_languages

**Type:** `array[string]`  
**Required:** Yes  
**Description:** List of supported languages for the AI system.

##### info.ai_metadata.hierarchy_info

**Type:** `object`  
**Required:** Yes  
**Description:** Information about the hierarchical composition of the AI system.

###### info.ai_metadata.hierarchy_info.level

**Type:** `string`  
**Required:** Yes  
**Enum:** `["global", "regional", "department", "team", "sprint", "feature", "environment"]`  
**Description:** The hierarchy level of this specification.

###### info.ai_metadata.hierarchy_info.parent_specs

**Type:** `array[string]`  
**Required:** No  
**Description:** List of parent specifications in the hierarchy.

###### info.ai_metadata.hierarchy_info.scope

**Type:** `string`  
**Required:** Yes  
**Enum:** `["organization", "department", "team", "project", "feature", "environment"]`  
**Description:** The scope of this specification.

###### info.ai_metadata.hierarchy_info.inheritance_mode

**Type:** `string`  
**Required:** Yes  
**Enum:** `["merge", "override", "extend"]`  
**Description:** The inheritance mode for this specification.

## AI Models

### models

**Type:** `array[object]`  
**Required:** No  
**Description:** List of AI models used in the system.

#### models[].id

**Type:** `string`  
**Required:** Yes  
**Description:** Unique identifier for the model.

#### models[].type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["LLM", "Vision", "Audio", "Multimodal", "Classification", "Embedding"]`  
**Description:** The type of AI model.

#### models[].provider

**Type:** `string`  
**Required:** Yes  
**Description:** The provider of the AI model (e.g., "OpenAI", "Anthropic", "Google").

#### models[].name

**Type:** `string`  
**Required:** Yes  
**Description:** The name of the specific model.

#### models[].version

**Type:** `string`  
**Required:** Yes  
**Description:** The version of the model.

#### models[].purpose

**Type:** `string`  
**Required:** Yes  
**Description:** The intended use case or purpose of the model.

#### models[].capabilities

**Type:** `array[string]`  
**Required:** Yes  
**Description:** List of capabilities provided by the model.

#### models[].parameters

**Type:** `object`  
**Required:** No  
**Description:** Model-specific parameters and configuration.

##### models[].parameters.temperature

**Type:** `number`  
**Required:** No  
**Description:** Temperature parameter for controlling randomness in model outputs.

##### models[].parameters.max_tokens

**Type:** `number`  
**Required:** No  
**Description:** Maximum number of tokens to generate.

##### models[].parameters.top_p

**Type:** `number`  
**Required:** No  
**Description:** Top-p parameter for nucleus sampling.

##### models[].parameters.frequency_penalty

**Type:** `number`  
**Required:** No  
**Description:** Frequency penalty to reduce repetition.

##### models[].parameters.presence_penalty

**Type:** `number`  
**Required:** No  
**Description:** Presence penalty to encourage new topics.

#### models[].limits

**Type:** `object`  
**Required:** No  
**Description:** Operational limits for the model.

##### models[].limits.max_input_tokens

**Type:** `number`  
**Required:** No  
**Description:** Maximum number of input tokens supported.

##### models[].limits.max_output_tokens

**Type:** `number`  
**Required:** No  
**Description:** Maximum number of output tokens supported.

##### models[].limits.requests_per_minute

**Type:** `number`  
**Required:** No  
**Description:** Rate limit for API requests per minute.

#### models[].cost

**Type:** `object`  
**Required:** No  
**Description:** Cost information for the model.

##### models[].cost.input_per_1k_tokens

**Type:** `number`  
**Required:** No  
**Description:** Cost per 1,000 input tokens.

##### models[].cost.output_per_1k_tokens

**Type:** `number`  
**Required:** No  
**Description:** Cost per 1,000 output tokens.

##### models[].cost.currency

**Type:** `string`  
**Required:** No  
**Description:** Currency for cost calculations.

#### models[].performance

**Type:** `object`  
**Required:** No  
**Description:** Performance metrics for the model.

##### models[].performance.accuracy

**Type:** `number`  
**Required:** No  
**Description:** Model accuracy score.

##### models[].performance.latency

**Type:** `string`  
**Required:** No  
**Description:** Average response latency.

##### models[].performance.throughput

**Type:** `string`  
**Required:** No  
**Description:** Throughput capacity.

## Prompts

### prompts

**Type:** `array[object]`  
**Required:** No  
**Description:** Structured prompts used in the AI system.

#### prompts[].id

**Type:** `string`  
**Required:** Yes  
**Description:** Unique identifier for the prompt.

#### prompts[].role

**Type:** `string`  
**Required:** Yes  
**Enum:** `["system", "user", "assistant"]`  
**Description:** The role of the prompt in the conversation.

#### prompts[].style

**Type:** `string`  
**Required:** Yes  
**Description:** The style or tone of the prompt.

#### prompts[].language

**Type:** `string`  
**Required:** Yes  
**Description:** The language of the prompt.

#### prompts[].template

**Type:** `string`  
**Required:** Yes  
**Description:** The prompt template with variable placeholders.

#### prompts[].variables

**Type:** `object`  
**Required:** No  
**Description:** Definition of template variables.

##### prompts[].variables.{variable_name}

**Type:** `object`  
**Required:** No  
**Description:** Definition of a specific template variable.

###### prompts[].variables.{variable_name}.type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["string", "number", "boolean", "array", "object"]`  
**Description:** The data type of the variable.

###### prompts[].variables.{variable_name}.required

**Type:** `boolean`  
**Required:** Yes  
**Description:** Whether the variable is required.

###### prompts[].variables.{variable_name}.default

**Type:** `any`  
**Required:** No  
**Description:** Default value for the variable.

###### prompts[].variables.{variable_name}.enum

**Type:** `array[any]`  
**Required:** No  
**Description:** Allowed values for the variable.

###### prompts[].variables.{variable_name}.description

**Type:** `string`  
**Required:** Yes  
**Description:** Description of the variable.

#### prompts[].config

**Type:** `object`  
**Required:** No  
**Description:** Configuration for the prompt.

##### prompts[].config.temperature

**Type:** `number`  
**Required:** No  
**Description:** Temperature setting for this prompt.

##### prompts[].config.max_tokens

**Type:** `number`  
**Required:** No  
**Description:** Maximum tokens for this prompt.

## Constraints

### constraints

**Type:** `array[object]`  
**Required:** No  
**Description:** Safety, ethical, and operational constraints for the AI system.

#### constraints[].id

**Type:** `string`  
**Required:** Yes  
**Description:** Unique identifier for the constraint.

#### constraints[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Human-readable name of the constraint.

#### constraints[].type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["content_safety", "privacy", "performance", "budget", "fairness"]`  
**Description:** The type of constraint.

#### constraints[].rule

**Type:** `string`  
**Required:** Yes  
**Description:** The constraint rule expression or description.

#### constraints[].severity

**Type:** `string`  
**Required:** Yes  
**Enum:** `["low", "medium", "high", "critical"]`  
**Description:** The severity level of the constraint violation.

#### constraints[].enforcement

**Type:** `string`  
**Required:** Yes  
**Enum:** `["automatic", "monitoring", "manual"]`  
**Description:** How the constraint is enforced.

#### constraints[].description

**Type:** `string`  
**Required:** Yes  
**Description:** Detailed description of the constraint.

#### constraints[].actions

**Type:** `array[string]`  
**Required:** No  
**Description:** Actions to take when the constraint is violated.

## Tasks

### tasks

**Type:** `array[object]`  
**Required:** No  
**Description:** Declarative workflows that define how the AI system processes requests.

#### tasks[].id

**Type:** `string`  
**Required:** Yes  
**Description:** Unique identifier for the task.

#### tasks[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Human-readable name of the task.

#### tasks[].description

**Type:** `string`  
**Required:** Yes  
**Description:** Detailed description of what the task does.

#### tasks[].type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["conversational", "analysis", "generation", "classification"]`  
**Description:** The type of task.

#### tasks[].priority

**Type:** `string`  
**Required:** Yes  
**Enum:** `["low", "medium", "high", "critical"]`  
**Description:** The priority level of the task.

#### tasks[].input

**Type:** `object`  
**Required:** No  
**Description:** Schema definition for task input.

##### tasks[].input.{field_name}

**Type:** `object`  
**Required:** No  
**Description:** Definition of an input field.

###### tasks[].input.{field_name}.type

**Type:** `string`  
**Required:** Yes  
**Description:** The data type of the field.

###### tasks[].input.{field_name}.required

**Type:** `boolean`  
**Required:** Yes  
**Description:** Whether the field is required.

###### tasks[].input.{field_name}.description

**Type:** `string`  
**Required:** Yes  
**Description:** Description of the field.

#### tasks[].output

**Type:** `object`  
**Required:** No  
**Description:** Schema definition for task output.

##### tasks[].output.{field_name}

**Type:** `object`  
**Required:** No  
**Description:** Definition of an output field.

###### tasks[].output.{field_name}.type

**Type:** `string`  
**Required:** Yes  
**Description:** The data type of the field.

###### tasks[].output.{field_name}.description

**Type:** `string`  
**Required:** Yes  
**Description:** Description of the field.

###### tasks[].output.{field_name}.minimum

**Type:** `number`  
**Required:** No  
**Description:** Minimum value for numeric types.

###### tasks[].output.{field_name}.maximum

**Type:** `number`  
**Required:** No  
**Description:** Maximum value for numeric types.

###### tasks[].output.{field_name}.items

**Type:** `object`  
**Required:** No  
**Description:** Item definition for array types.

#### tasks[].steps

**Type:** `array[object]`  
**Required:** No  
**Description:** Execution steps for the task.

##### tasks[].steps[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Name of the execution step.

##### tasks[].steps[].action

**Type:** `string`  
**Required:** Yes  
**Enum:** `["analyze", "generate", "validate", "search", "escalate", "classify", "mcp_tool", "mcp_resource", "automation"]`  
**Description:** The action to perform in this step.

##### tasks[].steps[].model

**Type:** `string`  
**Required:** No  
**Description:** Reference to a model ID to use in this step.

##### tasks[].steps[].prompt

**Type:** `string`  
**Required:** No  
**Description:** Reference to a prompt ID to use in this step.

##### tasks[].steps[].source

**Type:** `string`  
**Required:** No  
**Description:** Data source for this step.

##### tasks[].steps[].mcp_server

**Type:** `string`  
**Required:** No  
**Description:** Reference to an MCP server ID to use in this step.

##### tasks[].steps[].mcp_tool

**Type:** `string`  
**Required:** No  
**Description:** MCP tool name to execute (if action is mcp_tool).

##### tasks[].steps[].mcp_resource

**Type:** `string`  
**Required:** No  
**Description:** MCP resource name to access (if action is mcp_resource).

##### tasks[].steps[].mcp_parameters

**Type:** `object`  
**Required:** No  
**Description:** Parameters to pass to the MCP tool or resource.

##### tasks[].steps[].automation

**Type:** `string`  
**Required:** No  
**Description:** Reference to an automation ID to trigger (if action is automation).

##### tasks[].steps[].automation_parameters

**Type:** `object`  
**Required:** No  
**Description:** Parameters to pass to the automation.

##### tasks[].steps[].check_automation

**Type:** `string`  
**Required:** No  
**Description:** Reference to an automation ID to check status (for wait actions).

##### tasks[].steps[].constraints

**Type:** `array[string]`  
**Required:** No  
**Description:** References to constraint IDs that apply to this step.

##### tasks[].steps[].conditions

**Type:** `array[object]`  
**Required:** No  
**Description:** Conditional execution logic for this step.

###### tasks[].steps[].conditions[].if

**Type:** `string`  
**Required:** Yes  
**Description:** Condition expression to evaluate.

###### tasks[].steps[].conditions[].then

**Type:** `string`  
**Required:** Yes  
**Description:** Next step or action to take if condition is true.

## Automations

### automations

**Type:** `array[object]`  
**Required:** No  
**Description:** External automation workflows that can be triggered by the AI system. This section enables declarative integration with automation platforms like n8n, Zapier, Microsoft Power Automate, and custom webhooks.

**Philosophy:** OpenAPIA follows a declarative approach to automation integration. Instead of describing the internal logic of automation workflows, it declares where and when external automations should be triggered, maintaining clear separation of concerns between AI logic and automation workflows.

**Key Benefits:**
- **Vendor Agnostic**: Works with any automation platform
- **Maintainable**: AI logic remains separate from automation details  
- **Monitored**: Built-in health checks and performance metrics
- **Secure**: Configurable authentication and data validation

#### automations[].id

**Type:** `string`  
**Required:** Yes  
**Description:** Unique identifier for the automation.

#### automations[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Human-readable name of the automation.

#### automations[].description

**Type:** `string`  
**Required:** Yes  
**Description:** Detailed description of what the automation does.

#### automations[].provider

**Type:** `string`  
**Required:** Yes  
**Description:** The automation platform provider (e.g., "n8n", "zapier", "power_automate", "custom").

#### automations[].version

**Type:** `string`  
**Required:** No  
**Description:** Version of the automation workflow.

#### automations[].trigger

**Type:** `object`  
**Required:** Yes  
**Description:** Configuration for how the automation is triggered.

##### automations[].trigger.type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["webhook", "scheduled", "event", "manual"]`  
**Description:** The type of trigger for the automation.

##### automations[].trigger.endpoint

**Type:** `string`  
**Required:** No  
**Description:** Webhook endpoint URL (if type is webhook).

##### automations[].trigger.method

**Type:** `string`  
**Required:** No  
**Description:** HTTP method for webhook triggers (e.g., "POST", "GET").

##### automations[].trigger.schedule

**Type:** `string`  
**Required:** No  
**Description:** Cron expression for scheduled triggers.

##### automations[].trigger.authentication

**Type:** `object`  
**Required:** No  
**Description:** Authentication configuration for the trigger.

###### automations[].trigger.authentication.type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["none", "api_key", "bearer", "basic"]`  
**Description:** Authentication type for the trigger.

###### automations[].trigger.authentication.header

**Type:** `string`  
**Required:** No  
**Description:** Header name for API key authentication.

###### automations[].trigger.authentication.value

**Type:** `string`  
**Required:** No  
**Description:** Authentication value (API key, token, etc.).

##### automations[].trigger.conditions

**Type:** `array[string]`  
**Required:** No  
**Description:** Conditions that must be met for the automation to trigger.

#### automations[].integration

**Type:** `object`  
**Required:** Yes  
**Description:** Integration configuration for the external automation platform.

##### automations[].integration.type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["external_workflow", "webhook", "api_call", "zap", "flow"]`  
**Description:** The type of integration with the automation platform.

##### automations[].integration.workflow_id

**Type:** `string`  
**Required:** No  
**Description:** Identifier for the external workflow (e.g., n8n workflow ID).

##### automations[].integration.environment

**Type:** `string`  
**Required:** No  
**Description:** Environment where the automation runs (e.g., "production", "staging").

##### automations[].integration.timeout

**Type:** `string`  
**Required:** No  
**Description:** Maximum time to wait for automation completion.

##### automations[].integration.retry_policy

**Type:** `object`  
**Required:** No  
**Description:** Retry configuration for failed automation executions.

###### automations[].integration.retry_policy.max_retries

**Type:** `number`  
**Required:** No  
**Description:** Maximum number of retry attempts.

###### automations[].integration.retry_policy.backoff_strategy

**Type:** `string`  
**Required:** No  
**Enum:** `["linear", "exponential"]`  
**Description:** Backoff strategy for retries.

###### automations[].integration.retry_policy.initial_delay

**Type:** `number`  
**Required:** No  
**Description:** Initial delay in milliseconds before first retry.

#### automations[].data_contract

**Type:** `object`  
**Required:** Yes  
**Description:** Contract defining the data structure for automation input and output.

##### automations[].data_contract.input

**Type:** `object`  
**Required:** No  
**Description:** Schema for input data sent to the automation.

##### automations[].data_contract.output

**Type:** `object`  
**Required:** No  
**Description:** Schema for output data returned by the automation.

#### automations[].monitoring

**Type:** `object`  
**Required:** No  
**Description:** Monitoring and health check configuration for the automation.

##### automations[].monitoring.health_check

**Type:** `object`  
**Required:** No  
**Description:** Health check configuration.

###### automations[].monitoring.health_check.endpoint

**Type:** `string`  
**Required:** No  
**Description:** Health check endpoint URL.

###### automations[].monitoring.health_check.interval

**Type:** `string`  
**Required:** No  
**Description:** Health check interval (e.g., "1m", "5m").

###### automations[].monitoring.health_check.timeout

**Type:** `string`  
**Required:** No  
**Description:** Health check timeout (e.g., "10s", "30s").

##### automations[].monitoring.metrics

**Type:** `array[object]`  
**Required:** No  
**Description:** Metrics to track for the automation.

###### automations[].monitoring.metrics[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Name of the metric.

###### automations[].monitoring.metrics[].description

**Type:** `string`  
**Required:** Yes  
**Description:** Description of what the metric measures.

###### automations[].monitoring.metrics[].target

**Type:** `string`  
**Required:** No  
**Description:** Target value for the metric.

#### automations[].metadata

**Type:** `object`  
**Required:** No  
**Description:** Additional metadata for the automation.

##### automations[].metadata.tags

**Type:** `array[string]`  
**Required:** No  
**Description:** Tags for categorizing the automation.

##### automations[].metadata.category

**Type:** `string`  
**Required:** No  
**Description:** Category of the automation (e.g., "business_process", "data_processing", "notification").

##### automations[].metadata.maintainer

**Type:** `string`  
**Required:** No  
**Description:** Team or person responsible for maintaining the automation.

##### automations[].metadata.documentation_url

**Type:** `string`  
**Required:** No  
**Description:** URL to documentation for the automation.

## Context

### context

**Type:** `object`  
**Required:** No  
**Description:** State management and memory configuration for the AI system.

#### context.memory

**Type:** `object`  
**Required:** No  
**Description:** Memory configuration for the AI system.

##### context.memory.type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["session", "persistent", "temporary"]`  
**Description:** The type of memory storage.

##### context.memory.retention

**Type:** `string`  
**Required:** Yes  
**Description:** How long to retain memory data.

##### context.memory.scope

**Type:** `string`  
**Required:** Yes  
**Enum:** `["per_user", "global", "per_session"]`  
**Description:** The scope of memory storage.

##### context.memory.storage

**Type:** `object`  
**Required:** No  
**Description:** Storage configuration for memory.

###### context.memory.storage.provider

**Type:** `string`  
**Required:** Yes  
**Description:** Storage provider for memory data.

###### context.memory.storage.ttl

**Type:** `number`  
**Required:** No  
**Description:** Time to live in seconds for stored data.

###### context.memory.storage.max_size

**Type:** `string`  
**Required:** No  
**Description:** Maximum storage size for memory data.

##### context.memory.store

**Type:** `array[string]`  
**Required:** No  
**Description:** What types of data to store in memory.

##### context.memory.exclude

**Type:** `array[string]`  
**Required:** No  
**Description:** What types of data to exclude from memory storage.

#### context.conversation

**Type:** `object`  
**Required:** No  
**Description:** Conversation management configuration.

##### context.conversation.max_turns

**Type:** `number`  
**Required:** No  
**Description:** Maximum number of conversation turns to maintain.

##### context.conversation.context_window

**Type:** `number`  
**Required:** No  
**Description:** Size of the context window for conversations.

##### context.conversation.summary_frequency

**Type:** `number`  
**Required:** No  
**Description:** How often to create conversation summaries.

##### context.conversation.summary_template

**Type:** `string`  
**Required:** No  
**Description:** Template for creating conversation summaries.

#### context.business_context

**Type:** `object`  
**Required:** No  
**Description:** Business context configuration.

##### context.business_context.company_info

**Type:** `object`  
**Required:** No  
**Description:** Company information for business context.

###### context.business_context.company_info.name

**Type:** `string`  
**Required:** No  
**Description:** Company name.

###### context.business_context.company_info.industry

**Type:** `string`  
**Required:** No  
**Description:** Industry sector.

###### context.business_context.company_info.products

**Type:** `array[string]`  
**Required:** No  
**Description:** List of company products.

###### context.business_context.company_info.policies

**Type:** `array[string]`  
**Required:** No  
**Description:** List of company policies.

##### context.business_context.knowledge_base

**Type:** `object`  
**Required:** No  
**Description:** Knowledge base configuration.

###### context.business_context.knowledge_base.type

**Type:** `string`  
**Required:** No  
**Description:** Type of knowledge base.

###### context.business_context.knowledge_base.provider

**Type:** `string`  
**Required:** No  
**Description:** Knowledge base provider.

###### context.business_context.knowledge_base.index_name

**Type:** `string`  
**Required:** No  
**Description:** Name of the knowledge base index.

###### context.business_context.knowledge_base.embedding_model

**Type:** `string`  
**Required:** No  
**Description:** Embedding model used for knowledge base.

#### context.mcp_servers

**Type:** `array[object]`  
**Required:** No  
**Description:** Model Context Protocol (MCP) servers for external data and tool access.

##### context.mcp_servers[].id

**Type:** `string`  
**Required:** Yes  
**Description:** Unique identifier for the MCP server.

##### context.mcp_servers[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Human-readable name of the MCP server.

##### context.mcp_servers[].description

**Type:** `string`  
**Required:** Yes  
**Description:** Detailed description of the MCP server's purpose and capabilities.

##### context.mcp_servers[].version

**Type:** `string`  
**Required:** Yes  
**Description:** Version of the MCP server.

##### context.mcp_servers[].transport

**Type:** `object`  
**Required:** Yes  
**Description:** Transport configuration for connecting to the MCP server.

###### context.mcp_servers[].transport.type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["stdio", "sse", "websocket"]`  
**Description:** Transport protocol type for the MCP server connection.

###### context.mcp_servers[].transport.command

**Type:** `string`  
**Required:** No  
**Description:** Command to start the MCP server (for stdio transport).

###### context.mcp_servers[].transport.args

**Type:** `array[string]`  
**Required:** No  
**Description:** Command line arguments for starting the MCP server.

###### context.mcp_servers[].transport.url

**Type:** `string`  
**Required:** No  
**Description:** Server URL for SSE or WebSocket transport.

###### context.mcp_servers[].transport.headers

**Type:** `object`  
**Required:** No  
**Description:** Custom headers for the connection.

##### context.mcp_servers[].capabilities

**Type:** `object`  
**Required:** Yes  
**Description:** Capabilities provided by the MCP server.

###### context.mcp_servers[].capabilities.tools

**Type:** `array[string]`  
**Required:** No  
**Description:** List of available tools provided by the server.

###### context.mcp_servers[].capabilities.resources

**Type:** `array[string]`  
**Required:** No  
**Description:** List of available resources provided by the server.

###### context.mcp_servers[].capabilities.prompts

**Type:** `array[string]`  
**Required:** No  
**Description:** List of available prompts provided by the server.

##### context.mcp_servers[].authentication

**Type:** `object`  
**Required:** Yes  
**Description:** Authentication configuration for the MCP server.

###### context.mcp_servers[].authentication.type

**Type:** `string`  
**Required:** Yes  
**Enum:** `["none", "api_key", "oauth", "custom"]`  
**Description:** Authentication type for the MCP server.

###### context.mcp_servers[].authentication.api_key

**Type:** `string`  
**Required:** No  
**Description:** API key for authentication (if applicable).

###### context.mcp_servers[].authentication.token

**Type:** `string`  
**Required:** No  
**Description:** Access token for authentication (if applicable).

###### context.mcp_servers[].authentication.custom_auth

**Type:** `object`  
**Required:** No  
**Description:** Custom authentication configuration.

##### context.mcp_servers[].security

**Type:** `object`  
**Required:** No  
**Description:** Security configuration for the MCP server.

###### context.mcp_servers[].security.allowed_operations

**Type:** `array[string]`  
**Required:** No  
**Description:** List of allowed operations for the server.

###### context.mcp_servers[].security.rate_limits

**Type:** `object`  
**Required:** No  
**Description:** Rate limiting configuration.

####### context.mcp_servers[].security.rate_limits.requests_per_minute

**Type:** `number`  
**Required:** No  
**Description:** Maximum requests per minute.

####### context.mcp_servers[].security.rate_limits.requests_per_hour

**Type:** `number`  
**Required:** No  
**Description:** Maximum requests per hour.

###### context.mcp_servers[].security.timeout

**Type:** `number`  
**Required:** No  
**Description:** Request timeout in seconds.

##### context.mcp_servers[].health_check

**Type:** `object`  
**Required:** No  
**Description:** Health check configuration for the MCP server.

###### context.mcp_servers[].health_check.enabled

**Type:** `boolean`  
**Required:** No  
**Description:** Whether health checks are enabled.

###### context.mcp_servers[].health_check.interval

**Type:** `number`  
**Required:** No  
**Description:** Health check interval in seconds.

###### context.mcp_servers[].health_check.timeout

**Type:** `number`  
**Required:** No  
**Description:** Health check timeout in seconds.

###### context.mcp_servers[].health_check.retry_count

**Type:** `number`  
**Required:** No  
**Description:** Number of retry attempts for failed health checks.

##### context.mcp_servers[].metadata

**Type:** `object`  
**Required:** No  
**Description:** Additional metadata for the MCP server.

###### context.mcp_servers[].metadata.tags

**Type:** `array[string]`  
**Required:** No  
**Description:** Tags for categorizing the server.

###### context.mcp_servers[].metadata.category

**Type:** `string`  
**Required:** No  
**Description:** Category of the MCP server.

###### context.mcp_servers[].metadata.maintainer

**Type:** `string`  
**Required:** No  
**Description:** Maintainer of the MCP server.

###### context.mcp_servers[].metadata.documentation_url

**Type:** `string`  
**Required:** No  
**Description:** URL to the server's documentation.

## Evaluation

### evaluation

**Type:** `object`  
**Required:** No  
**Description:** Metrics and testing configuration for the AI system.

#### evaluation.metrics

**Type:** `array[object]`  
**Required:** No  
**Description:** Performance metrics to track.

##### evaluation.metrics[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Name of the metric.

##### evaluation.metrics[].description

**Type:** `string`  
**Required:** Yes  
**Description:** Description of what the metric measures.

##### evaluation.metrics[].target

**Type:** `any`  
**Required:** No  
**Description:** Target value for the metric.

##### evaluation.metrics[].measurement

**Type:** `object`  
**Required:** No  
**Description:** Configuration for measuring the metric.

###### evaluation.metrics[].measurement.method

**Type:** `string`  
**Required:** Yes  
**Description:** Method used to measure the metric.

###### evaluation.metrics[].measurement.sample_size

**Type:** `number`  
**Required:** No  
**Description:** Sample size for measurement.

###### evaluation.metrics[].measurement.frequency

**Type:** `string`  
**Required:** No  
**Description:** How frequently to measure the metric.

###### evaluation.metrics[].measurement.criteria

**Type:** `array[string]`  
**Required:** No  
**Description:** Evaluation criteria for the metric.

###### evaluation.metrics[].measurement.percentiles

**Type:** `array[number]`  
**Required:** No  
**Description:** Percentiles to track for latency metrics.

###### evaluation.metrics[].measurement.scale

**Type:** `string`  
**Required:** No  
**Description:** Scale for rating metrics.

#### evaluation.test_cases

**Type:** `array[object]`  
**Required:** No  
**Description:** Automated test cases for the AI system.

##### evaluation.test_cases[].id

**Type:** `string`  
**Required:** Yes  
**Description:** Unique identifier for the test case.

##### evaluation.test_cases[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Name of the test case.

##### evaluation.test_cases[].input

**Type:** `string`  
**Required:** Yes  
**Description:** Test input for the test case.

##### evaluation.test_cases[].expected_behavior

**Type:** `string`  
**Required:** Yes  
**Description:** Expected behavior for the test case.

##### evaluation.test_cases[].category

**Type:** `string`  
**Required:** Yes  
**Enum:** `["functional", "safety", "privacy", "performance"]`  
**Description:** Category of the test case.

##### evaluation.test_cases[].priority

**Type:** `string`  
**Required:** Yes  
**Enum:** `["low", "medium", "high", "critical"]`  
**Description:** Priority level of the test case.

#### evaluation.performance_tests

**Type:** `array[object]`  
**Required:** No  
**Description:** Performance testing configuration.

##### evaluation.performance_tests[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Name of the performance test.

##### evaluation.performance_tests[].description

**Type:** `string`  
**Required:** Yes  
**Description:** Description of the performance test.

##### evaluation.performance_tests[].concurrent_users

**Type:** `number`  
**Required:** No  
**Description:** Number of concurrent users for the test.

##### evaluation.performance_tests[].duration

**Type:** `string`  
**Required:** No  
**Description:** Duration of the performance test.

##### evaluation.performance_tests[].target_response_time

**Type:** `string`  
**Required:** No  
**Description:** Target response time for the test.

##### evaluation.performance_tests[].target_throughput

**Type:** `string`  
**Required:** No  
**Description:** Target throughput for the test.

##### evaluation.performance_tests[].target_availability

**Type:** `string`  
**Required:** No  
**Description:** Target availability for the test.

## Extensions

### extensions

**Type:** `object`  
**Required:** No  
**Description:** Optional extensions for advanced AI system capabilities.

#### extensions.vision_support

**Type:** `boolean`  
**Required:** No  
**Description:** Whether the system supports computer vision capabilities.

#### extensions.multilingual

**Type:** `boolean`  
**Required:** No  
**Description:** Whether the system supports multiple languages.

#### extensions.audio_processing

**Type:** `boolean`  
**Required:** No  
**Description:** Whether the system supports audio processing capabilities.

#### extensions.real_time_processing

**Type:** `boolean`  
**Required:** No  
**Description:** Whether the system supports real-time processing.



#### extensions.mcp_support

**Type:** `boolean`  
**Required:** No  
**Description:** Whether the system supports Model Context Protocol (MCP) servers.

#### extensions.automation_support

**Type:** `boolean`  
**Required:** No  
**Description:** Whether the system supports external automation workflows.

#### extensions.advanced

**Type:** `object`  
**Required:** No  
**Description:** Advanced configuration options for extended capabilities.

##### extensions.advanced.computer_vision

**Type:** `object`  
**Required:** No  
**Description:** Computer vision configuration.

###### extensions.advanced.computer_vision.enabled

**Type:** `boolean`  
**Required:** No  
**Description:** Whether computer vision is enabled.

###### extensions.advanced.computer_vision.models

**Type:** `array[object]`  
**Required:** No  
**Description:** Computer vision models configuration.

##### extensions.advanced.audio_processing

**Type:** `object`  
**Required:** No  
**Description:** Audio processing configuration.

###### extensions.advanced.audio_processing.enabled

**Type:** `boolean`  
**Required:** No  
**Description:** Whether audio processing is enabled.

###### extensions.advanced.audio_processing.models

**Type:** `array[object]`  
**Required:** No  
**Description:** Audio processing models configuration.

##### extensions.advanced.multimodal

**Type:** `object`  
**Required:** No  
**Description:** Multimodal configuration.

###### extensions.advanced.multimodal.enabled

**Type:** `boolean`  
**Required:** No  
**Description:** Whether multimodal capabilities are enabled.

###### extensions.advanced.multimodal.models

**Type:** `array[object]`  
**Required:** No  
**Description:** Multimodal models configuration.

  
**Description:** Inter-agent communication configuration.

##### extensions.advanced.mcp

**Type:** `object`  
**Required:** No  
**Description:** Model Context Protocol configuration.

###### extensions.advanced.mcp.enabled

**Type:** `boolean`  
**Required:** No  
**Description:** Whether MCP support is enabled.

###### extensions.advanced.mcp.default_timeout

**Type:** `number`  
**Required:** No  
**Description:** Default timeout for MCP operations in seconds.

###### extensions.advanced.mcp.retry_policy

**Type:** `object`  
**Required:** No  
**Description:** Retry configuration for failed MCP operations.

####### extensions.advanced.mcp.retry_policy.max_retries

**Type:** `number`  
**Required:** No  
**Description:** Maximum number of retry attempts.

####### extensions.advanced.mcp.retry_policy.backoff_strategy

**Type:** `string`  
**Required:** No  
**Enum:** `["linear", "exponential"]`  
**Description:** Backoff strategy for retries.

####### extensions.advanced.mcp.retry_policy.initial_delay

**Type:** `number`  
**Required:** No  
**Description:** Initial delay in milliseconds before first retry.

###### extensions.advanced.mcp.connection_pool

**Type:** `object`  
**Required:** No  
**Description:** Connection pool configuration for MCP servers.

####### extensions.advanced.mcp.connection_pool.max_connections

**Type:** `number`  
**Required:** No  
**Description:** Maximum number of concurrent connections.

####### extensions.advanced.mcp.connection_pool.idle_timeout

**Type:** `number`  
**Required:** No  
**Description:** Idle timeout in seconds before closing connections.

###### extensions.advanced.mcp.monitoring

**Type:** `object`  
**Required:** No  
**Description:** MCP monitoring configuration.

####### extensions.advanced.mcp.monitoring.enabled

**Type:** `boolean`  
**Required:** No  
**Description:** Whether MCP monitoring is enabled.

####### extensions.advanced.mcp.monitoring.metrics_collection

**Type:** `boolean`  
**Required:** No  
**Description:** Whether to collect MCP performance metrics.

####### extensions.advanced.mcp.monitoring.health_check_interval

**Type:** `number`  
**Required:** No  
**Description:** Health check interval in seconds.

##### extensions.advanced.automation

**Type:** `object`  
**Required:** No  
**Description:** Advanced automation configuration.

###### extensions.advanced.automation.enabled

**Type:** `boolean`  
**Required:** No  
**Description:** Whether advanced automation features are enabled.

###### extensions.advanced.automation.default_timeout

**Type:** `number`  
**Required:** No  
**Description:** Default timeout for automation operations in seconds.

###### extensions.advanced.automation.retry_policy

**Type:** `object`  
**Required:** No  
**Description:** Default retry configuration for failed automation operations.

####### extensions.advanced.automation.retry_policy.max_retries

**Type:** `number`  
**Required:** No  
**Description:** Maximum number of retry attempts.

####### extensions.advanced.automation.retry_policy.backoff_strategy

**Type:** `string`  
**Required:** No  
**Enum:** `["linear", "exponential"]`  
**Description:** Backoff strategy for retries.

####### extensions.advanced.automation.retry_policy.initial_delay

**Type:** `number`  
**Required:** No  
**Description:** Initial delay in milliseconds before first retry.

###### extensions.advanced.automation.monitoring

**Type:** `object`  
**Required:** No  
**Description:** Automation monitoring configuration.

####### extensions.advanced.automation.monitoring.enabled

**Type:** `boolean`  
**Required:** No  
**Description:** Whether automation monitoring is enabled.

####### extensions.advanced.automation.monitoring.health_check_interval

**Type:** `number`  
**Required:** No  
**Description:** Health check interval in seconds.

####### extensions.advanced.automation.monitoring.metrics_collection

**Type:** `boolean`  
**Required:** No  
**Description:** Whether to collect automation performance metrics.

###### extensions.advanced.automation.providers

**Type:** `object`  
**Required:** No  
**Description:** Configuration for specific automation providers.

####### extensions.advanced.automation.providers.n8n

**Type:** `object`  
**Required:** No  
**Description:** n8n automation platform configuration.

######## extensions.advanced.automation.providers.n8n.base_url

**Type:** `string`  
**Required:** No  
**Description:** Base URL for n8n instance.

######## extensions.advanced.automation.providers.n8n.api_version

**Type:** `string`  
**Required:** No  
**Description:** API version for n8n.

######## extensions.advanced.automation.providers.n8n.timeout

**Type:** `number`  
**Required:** No  
**Description:** Default timeout for n8n operations in seconds.

####### extensions.advanced.automation.providers.zapier

**Type:** `object`  
**Required:** No  
**Description:** Zapier automation platform configuration.

######## extensions.advanced.automation.providers.zapier.base_url

**Type:** `string`  
**Required:** No  
**Description:** Base URL for Zapier webhooks.

######## extensions.advanced.automation.providers.zapier.timeout

**Type:** `number`  
**Required:** No  
**Description:** Default timeout for Zapier operations in seconds.

## Validation

### validation

**Type:** `object`  
**Required:** No  
**Description:** Validation configuration for the OpenAPIA specification.

#### validation.schema_version

**Type:** `string`  
**Required:** No  
**Description:** Version of the validation schema.

#### validation.required_sections

**Type:** `array[string]`  
**Required:** No  
**Description:** List of sections that are required in the specification.

#### validation.custom_validators

**Type:** `array[object]`  
**Required:** No  
**Description:** Custom validation rules.

##### validation.custom_validators[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Name of the custom validator.

##### validation.custom_validators[].description

**Type:** `string`  
**Required:** Yes  
**Description:** Description of what the validator checks.

## Governance

### governance

**Type:** `object`  
**Required:** No  
**Description:** Governance and maintenance information for the AI system.

#### governance.maintainers

**Type:** `array[object]`  
**Required:** No  
**Description:** List of project maintainers.

##### governance.maintainers[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Name of the maintainer.

##### governance.maintainers[].email

**Type:** `string`  
**Required:** Yes  
**Description:** Email address of the maintainer.

##### governance.maintainers[].role

**Type:** `string`  
**Required:** Yes  
**Description:** Role of the maintainer.

#### governance.contributors

**Type:** `array[object]`  
**Required:** No  
**Description:** List of project contributors.

##### governance.contributors[].name

**Type:** `string`  
**Required:** Yes  
**Description:** Name of the contributor.

##### governance.contributors[].email

**Type:** `string`  
**Required:** Yes  
**Description:** Email address of the contributor.

##### governance.contributors[].role

**Type:** `string`  
**Required:** Yes  
**Description:** Role of the contributor.

#### governance.review_process

**Type:** `string`  
**Required:** No  
**Description:** Description of the review process for changes.

#### governance.approval_required

**Type:** `number`  
**Required:** No  
**Description:** Number of approvals required for changes.

#### governance.testing_required

**Type:** `boolean`  
**Required:** No  
**Description:** Whether testing is required for changes.

#### governance.documentation_required

**Type:** `boolean`  
**Required:** No  
**Description:** Whether documentation updates are required for changes.

#### governance.update_process

**Type:** `object`  
**Required:** No  
**Description:** Process for updating the AI system.

##### governance.update_process.versioning

**Type:** `string`  
**Required:** No  
**Description:** Versioning strategy for the system.

##### governance.update_process.breaking_changes

**Type:** `string`  
**Required:** No  
**Description:** Policy for handling breaking changes.

##### governance.update_process.deprecation_notice

**Type:** `string`  
**Required:** No  
**Description:** Required notice period for deprecations.

##### governance.update_process.migration_guide

**Type:** `string`  
**Required:** No  
**Description:** Requirement for migration guides when making changes.

---

## Conclusion

This specification provides a comprehensive framework for describing AI systems using the OpenAPIA 0.1 standard. The specification is designed to be extensible, allowing for future enhancements while maintaining backward compatibility.

For more information about OpenAPIA, including examples and implementation guides, please refer to the additional documentation in the `docs/` directory.
