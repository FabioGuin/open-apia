# OpenAPIA Multi-Agent Systems

## Overview

OpenAPIA supports multi-agent systems through its existing features: **Hierarchical Composition**, **MCP Integration**, and **Automation Workflows**. Instead of adding dedicated multi-agent syntax, OpenAPIA leverages these proven patterns to create sophisticated agent coordination.

## Philosophy

Multi-agent systems in OpenAPIA follow these principles:

- **Composition over Configuration**: Build agents by composing existing OpenAPIA specifications
- **Declarative Coordination**: Define agent interactions through clear contracts
- **Separation of Concerns**: Each agent maintains its own capabilities and constraints
- **Scalable Architecture**: Easily add, remove, or modify agents without system-wide changes
- **No Special Syntax**: Multi-agent capabilities are achieved through existing OpenAPIA features

## Multi-Agent Patterns

### Pattern 1: Hierarchical Agent Composition

Create specialized agents as separate OpenAPIA specifications and compose them hierarchically.

#### Example: Customer Support Multi-Agent System

```yaml
# Main orchestrator specification
openapia: "0.1.0"
info:
  title: "Customer Support Orchestrator"
  description: "Coordinates multiple specialized agents for customer support"
  ai_metadata:
    domain: "customer_support"
    complexity: "high"
    hierarchy_info:
      level: "orchestrator"
      scope: "multi_agent"

# Inherit from specialized agent specifications
inherits:
  - "./agents/sentiment-analyzer.yaml"
  - "./agents/order-processor.yaml"
  - "./agents/faq-responder.yaml"
  - "./agents/escalation-manager.yaml"

models:
  - id: "orchestrator_llm"
    type: "LLM"
    provider: "OpenAI"
    name: "GPT-4"
    purpose: "Agent coordination and routing"

tasks:
  - id: "route_customer_inquiry"
    name: "Route Customer Inquiry"
    description: "Analyze and route customer inquiries to appropriate agents"
    type: "analysis"
    priority: "high"
    
    steps:
      - name: "analyze_inquiry"
        action: "analyze"
        model: "orchestrator_llm"
        prompt: "routing_prompt"
      
      - name: "route_to_agent"
        action: "escalate"
        conditions:
          - if: "inquiry_type == 'order'"
            then: "order_processor"
          - if: "inquiry_type == 'faq'"
            then: "faq_responder"
          - if: "sentiment == 'negative'"
            then: "escalation_manager"
```

#### Specialized Agent: Sentiment Analyzer

```yaml
# agents/sentiment-analyzer.yaml
openapia: "0.1.0"
info:
  title: "Sentiment Analysis Agent"
  description: "Specialized agent for analyzing customer sentiment"
  ai_metadata:
    domain: "sentiment_analysis"
    complexity: "medium"
    hierarchy_info:
      level: "agent"
      scope: "specialized"

models:
  - id: "sentiment_model"
    type: "Classification"
    provider: "HuggingFace"
    name: "sentiment-analyzer"
    purpose: "Customer sentiment analysis"

tasks:
  - id: "analyze_sentiment"
    name: "Analyze Customer Sentiment"
    description: "Analyze customer message sentiment and emotional state"
    type: "classification"
    
    input:
      message:
        type: "string"
        required: true
        description: "Customer message to analyze"
    
    output:
      sentiment:
        type: "string"
        enum: ["positive", "negative", "neutral"]
        description: "Overall sentiment"
      confidence:
        type: "number"
        minimum: 0
        maximum: 1
        description: "Confidence score"
      emotions:
        type: "array"
        description: "Detected emotions"
```

#### Specialized Agent: Order Processor

```yaml
# agents/order-processor.yaml
openapia: "0.1.0"
info:
  title: "Order Processing Agent"
  description: "Specialized agent for handling order-related inquiries"
  ai_metadata:
    domain: "order_management"
    complexity: "medium"
    hierarchy_info:
      level: "agent"
      scope: "specialized"

models:
  - id: "order_llm"
    type: "LLM"
    provider: "OpenAI"
    name: "GPT-4"
    purpose: "Order processing and customer assistance"

context:
  mcp_servers:
    - id: "order-database"
      name: "Order Database MCP Server"
      capabilities:
        tools: ["get_order", "update_order", "search_orders"]
        resources: ["db://orders/*"]

tasks:
  - id: "process_order_inquiry"
    name: "Process Order Inquiry"
    description: "Handle customer order-related questions and requests"
    type: "conversational"
    
    steps:
      - name: "get_order_info"
        action: "mcp_tool"
        mcp_server: "order-database"
        mcp_tool: "get_order"
        mcp_parameters:
          order_id: "${input.order_id}"
      
      - name: "generate_response"
        action: "generate"
        model: "order_llm"
        prompt: "order_assistance_prompt"
```

### Pattern 2: MCP-Based Agent Coordination

Use MCP servers to create agent-like capabilities that can be invoked by other agents.

#### Example: Agent Network via MCP

```yaml
openapia: "0.1.0"
info:
  title: "MCP-Based Agent Network"
  description: "Multi-agent system using MCP servers for agent coordination"

context:
  mcp_servers:
    # Sentiment Analysis Agent as MCP Server
    - id: "sentiment-agent"
      name: "Sentiment Analysis Agent"
      description: "Provides sentiment analysis capabilities"
      transport:
        type: "stdio"
        command: "python"
        args: ["-m", "sentiment_agent_mcp"]
      capabilities:
        tools: ["analyze_sentiment", "classify_emotion", "detect_urgency"]
        resources: ["sentiment://models/*"]
    
    # Order Processing Agent as MCP Server
    - id: "order-agent"
      name: "Order Processing Agent"
      description: "Handles order-related operations"
      transport:
        type: "stdio"
        command: "python"
        args: ["-m", "order_agent_mcp"]
      capabilities:
        tools: ["process_order", "check_status", "update_order"]
        resources: ["orders://database/*"]
    
    # FAQ Agent as MCP Server
    - id: "faq-agent"
      name: "FAQ Response Agent"
      description: "Provides FAQ and knowledge base responses"
      transport:
        type: "stdio"
        command: "python"
        args: ["-m", "faq_agent_mcp"]
      capabilities:
        tools: ["search_faq", "generate_response", "escalate_query"]
        resources: ["faq://knowledge-base/*"]

tasks:
  - id: "multi_agent_customer_support"
    name: "Multi-Agent Customer Support"
    description: "Coordinate multiple agents for comprehensive customer support"
    type: "conversational"
    
    steps:
      # Step 1: Analyze sentiment
      - name: "analyze_sentiment"
        action: "mcp_tool"
        mcp_server: "sentiment-agent"
        mcp_tool: "analyze_sentiment"
        mcp_parameters:
          message: "${input.customer_message}"
      
      # Step 2: Route based on sentiment and content
      - name: "route_inquiry"
        action: "mcp_tool"
        mcp_server: "faq-agent"
        mcp_tool: "search_faq"
        conditions:
          - if: "sentiment.urgency == 'high'"
            then: "escalate_immediately"
        mcp_parameters:
          query: "${input.customer_message}"
          sentiment: "${sentiment.sentiment}"
      
      # Step 3: Handle order-specific inquiries
      - name: "process_order_inquiry"
        action: "mcp_tool"
        mcp_server: "order-agent"
        mcp_tool: "process_order"
        conditions:
          - if: "inquiry_type == 'order'"
            then: "generate_final_response"
        mcp_parameters:
          inquiry: "${input.customer_message}"
          customer_id: "${input.customer_id}"
      
      # Step 4: Generate coordinated response
      - name: "generate_final_response"
        action: "generate"
        model: "coordinator_llm"
        prompt: "multi_agent_response_prompt"
```

### Pattern 3: Automation-Based Agent Orchestration

Use automation workflows to coordinate agent interactions and manage complex multi-agent processes.

#### Example: E-commerce Multi-Agent System

```yaml
openapia: "0.1.0"
info:
  title: "E-commerce Multi-Agent System"
  description: "Multi-agent system for e-commerce using automation workflows"

automations:
  # Agent Coordination Workflow
  - id: "agent_coordination_workflow"
    name: "Multi-Agent Coordination"
    description: "Coordinates multiple AI agents for customer service"
    provider: "n8n"
    
    trigger:
      type: "webhook"
      endpoint: "/webhooks/customer-inquiry"
      method: "POST"
    
    integration:
      type: "external_workflow"
      workflow_id: "n8n://workflows/multi-agent-coordination"
      timeout: "10m"
    
    data_contract:
      input:
        customer_message:
          type: "string"
          required: true
        customer_id:
          type: "string"
          required: false
        session_id:
          type: "string"
          required: true
      
      output:
        agent_responses:
          type: "array"
          description: "Responses from multiple agents"
        final_response:
          type: "string"
          description: "Coordinated final response"
        escalation_required:
          type: "boolean"
          description: "Whether human escalation is needed"
    
    monitoring:
      health_check:
        endpoint: "https://n8n.company.com/health/agent-coordination"
        interval: "2m"
      metrics:
        - name: "agent_coordination_time"
          description: "Time to coordinate all agents"
          target: "< 30s"
        - name: "agent_success_rate"
          description: "Success rate of agent coordination"
          target: "> 95%"

  # Individual Agent Workflows
  - id: "sentiment_analysis_workflow"
    name: "Sentiment Analysis Agent Workflow"
    description: "Workflow for sentiment analysis agent"
    provider: "n8n"
    
    trigger:
      type: "webhook"
      endpoint: "/webhooks/sentiment-analysis"
    
    integration:
      type: "external_workflow"
      workflow_id: "n8n://workflows/sentiment-analysis"
      timeout: "2m"
    
    data_contract:
      input:
        message:
          type: "string"
          required: true
      output:
        sentiment:
          type: "string"
          enum: ["positive", "negative", "neutral"]
        confidence:
          type: "number"
        urgency:
          type: "string"
          enum: ["low", "medium", "high"]

tasks:
  - id: "orchestrate_multi_agent_support"
    name: "Orchestrate Multi-Agent Support"
    description: "Orchestrate multiple agents for comprehensive customer support"
    type: "conversational"
    
    steps:
      # Step 1: Trigger agent coordination workflow
      - name: "coordinate_agents"
        action: "automation"
        automation: "agent_coordination_workflow"
        automation_parameters:
          customer_message: "${input.customer_message}"
          customer_id: "${input.customer_id}"
          session_id: "${input.session_id}"
      
      # Step 2: Wait for coordination to complete
      - name: "wait_for_coordination"
        action: "wait"
        timeout: "10m"
        check_automation: "agent_coordination_workflow"
        conditions:
          - if: "coordination_status == 'completed'"
            then: "process_responses"
          - if: "coordination_status == 'failed'"
            then: "escalate_to_human"
      
      # Step 3: Process coordinated responses
      - name: "process_responses"
        action: "generate"
        model: "coordinator_llm"
        prompt: "multi_agent_response_synthesis"
        constraints: ["response_time", "data_privacy"]
```

## Best Practices for Multi-Agent Systems

### 1. Agent Design Principles

- **Single Responsibility**: Each agent should have a clear, focused purpose
- **Loose Coupling**: Agents should interact through well-defined interfaces
- **High Cohesion**: Related functionality should be grouped within the same agent
- **Stateless Design**: Agents should be stateless when possible for better scalability

### 2. Communication Patterns

- **Request-Response**: Simple query-response pattern for most interactions
- **Event-Driven**: Use automation workflows for complex event coordination
- **Pipeline**: Chain agents in sequence for data processing workflows
- **Broadcast**: Send information to multiple agents simultaneously

### 3. Error Handling and Resilience

```yaml
# Example: Resilient multi-agent task
tasks:
  - id: "resilient_multi_agent_task"
    steps:
      - name: "try_primary_agent"
        action: "mcp_tool"
        mcp_server: "primary-agent"
        mcp_tool: "process_request"
        conditions:
          - if: "agent_available == true"
            then: "complete"
          - if: "agent_available == false"
            then: "try_fallback_agent"
      
      - name: "try_fallback_agent"
        action: "mcp_tool"
        mcp_server: "fallback-agent"
        mcp_tool: "process_request"
        conditions:
          - if: "fallback_success == true"
            then: "complete"
          - if: "fallback_success == false"
            then: "escalate_to_human"
      
      - name: "escalate_to_human"
        action: "escalate"
        reason: "All agents unavailable"
```

### 4. Monitoring and Observability

```yaml
evaluation:
  metrics:
    - name: "agent_response_time"
      description: "Average response time across all agents"
      target: "< 5s"
      measurement:
        method: "automated"
        frequency: "real_time"
        percentiles: [50, 95, 99]
    
    - name: "agent_coordination_success_rate"
      description: "Success rate of multi-agent coordination"
      target: 0.95
      measurement:
        method: "automated"
        frequency: "hourly"
    
    - name: "agent_escalation_rate"
      description: "Rate of escalations to human agents"
      target: "< 10%"
      measurement:
        method: "automated"
        frequency: "daily"

  test_cases:
    - id: "multi_agent_coordination_test"
      name: "Multi-Agent Coordination Test"
      input: "Test multi-agent coordination with various scenarios"
      expected_behavior: "All agents should coordinate successfully"
      category: "functional"
      priority: "high"
    
    - id: "agent_failure_resilience_test"
      name: "Agent Failure Resilience Test"
      input: "Simulate agent failures"
      expected_behavior: "System should handle agent failures gracefully"
      category: "resilience"
      priority: "high"
```

## Implementation Examples

### Example 1: Customer Support Multi-Agent System

See `../examples/multi-agent/multi-agent-customer-support.yaml` for a complete implementation.

### Example 2: E-commerce Multi-Agent System

See `../examples/automation/ecommerce-automation.yaml` for a complete implementation with automation integration.

### Example 3: Content Moderation Multi-Agent System

See `../examples/core/content-moderator.yaml` for a complete implementation.

## Migration from Single-Agent to Multi-Agent

### Step 1: Identify Agent Boundaries
- Analyze existing tasks and identify natural boundaries
- Group related functionality into potential agents
- Define clear interfaces between agents

### Step 2: Create Agent Specifications
- Create separate OpenAPIA specifications for each agent
- Define agent-specific models, prompts, and constraints
- Implement agent-specific MCP servers if needed

### Step 3: Implement Coordination
- Choose coordination pattern (hierarchical, MCP-based, or automation-based)
- Implement orchestrator or coordination logic
- Define communication protocols between agents

### Step 4: Add Monitoring and Testing
- Implement multi-agent specific metrics
- Add coordination testing
- Monitor agent performance and interactions

### Step 5: Gradual Rollout
- Start with simple multi-agent scenarios
- Gradually increase complexity
- Monitor and optimize based on real-world usage

## Conclusion

OpenAPIA's multi-agent capabilities leverage existing, proven features to create sophisticated agent coordination without adding complexity to the core specification. This approach provides:

- **Flexibility**: Multiple patterns for different use cases
- **Scalability**: Easy to add, remove, or modify agents
- **Maintainability**: Reuses existing OpenAPIA patterns and tools
- **Consistency**: Maintains OpenAPIA's declarative philosophy

By using hierarchical composition, MCP integration, and automation workflows, you can build complex multi-agent systems that are both powerful and maintainable.
