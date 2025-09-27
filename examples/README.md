# APAI Examples

This directory contains comprehensive examples demonstrating the Architecture Protocol for Artificial Intelligence capabilities across different use cases and complexity levels.

## Directory Structure

```
examples/
├── README.md                           # This file
├── apai-0.1-example.json          # Complete JSON example for tools
├── agents/                            # Multi-agent system components
│   └── sentiment-analyzer.yaml        # Specialized sentiment analysis agent
├── core/                              # Core AI system examples
│   ├── customer-support.yaml          # E-commerce customer support
│   ├── content-moderator.yaml         # AI-powered content filtering
│   └── multilingual-chatbot.yaml      # Multi-language chatbot
├── multi-agent/                       # Multi-agent system examples
│   └── multi-agent-customer-support.yaml  # Complete multi-agent system
├── automation/                        # Automation integration examples
│   ├── ecommerce-automation.yaml      # n8n integration example
│   ├── zapier-automation.yaml         # Zapier webhook integration
│   └── mcp-integration.yaml           # Model Context Protocol integration
└── templates/                         # Template examples
    ├── basic-template.yaml            # Minimal starting template
    └── security-template.yaml         # Security-focused production template
```

## Security and Environment Variables

**IMPORTANT**: All examples use environment variables for sensitive data like API keys, passwords, and tokens. 

### Setup Required
1. Copy the environment template: `cp .env.example .env`
2. Fill in your actual credentials in the `.env` file
3. Never commit the `.env` file to version control

### Environment Variable Usage
Examples use the `${VARIABLE_NAME}` syntax for secure credential management:
```yaml
authentication:
  type: "api_key"
  api_key: "${OPENAI_API_KEY}"  # Secure API key from environment
```

See [Security Best Practices](../docs/security-best-practices.md) for complete guidelines.

## Templates

### [Basic Template](templates/basic-template.yaml)
**Complexity**: Low  
**Use Case**: Getting started, simple AI systems  
**Features**: Minimal configuration, basic security

### [Security Template](templates/security-template.yaml)
**Complexity**: High  
**Use Case**: Production systems, enterprise deployments  
**Features**: Comprehensive security controls, compliance, monitoring

## Examples by Category

### Core AI Systems
Basic AI system implementations for common use cases.

#### [Customer Support AI](core/customer-support.yaml)
**Complexity**: Medium  
**Domain**: Customer Service  
**Features**: Order management, returns, product assistance, multilingual support

Complete e-commerce customer support assistant with:
- Order tracking and management
- Return and refund processing
- Product recommendations
- Multilingual conversation support
- Sentiment analysis integration
- Escalation to human agents

#### [Content Moderator](core/content-moderator.yaml)
**Complexity**: Medium  
**Domain**: Content Generation  
**Features**: Content filtering, safety checks, automated moderation

AI-powered content moderation system with:
- Text content analysis
- Image content filtering
- Safety constraint enforcement
- Automated moderation actions
- Human review escalation
- Performance monitoring

#### [Multilingual Chatbot](core/multilingual-chatbot.yaml)
**Complexity**: Low  
**Domain**: Customer Service  
**Features**: Multi-language support, conversation management

Multi-language chatbot implementation with:
- Automatic language detection
- Language-specific responses
- Conversation context management
- Fallback language support
- Cultural adaptation

### Multi-Agent Systems
Complex systems using multiple specialized AI agents.

#### [Multi-Agent Customer Support](multi-agent/multi-agent-customer-support.yaml)
**Complexity**: High  
**Domain**: Customer Support  
**Features**: Agent orchestration, specialized agents, hierarchical composition

Complete multi-agent customer support system with:
- **Sentiment Analysis Agent**: Analyzes customer emotional state
- **Order Processing Agent**: Handles order-related inquiries
- **FAQ Response Agent**: Provides knowledge base responses
- **Escalation Manager**: Manages human agent handoffs
- **Orchestrator**: Coordinates all agents

#### [Sentiment Analyzer Agent](agents/sentiment-analyzer.yaml)
**Complexity**: Low  
**Domain**: Sentiment Analysis  
**Features**: Specialized agent, reusable component

Standalone sentiment analysis agent that can be:
- Used independently
- Integrated into larger systems
- Inherited by other specifications
- Extended with additional capabilities

### Automation Integration
Examples showing integration with external automation platforms.

#### [E-commerce with n8n](automation/ecommerce-automation.yaml)
**Complexity**: High  
**Domain**: E-commerce  
**Features**: n8n integration, complex workflows, business process automation

Complete e-commerce system with n8n automation:
- **Order Processing Workflow**: Automated order fulfillment
- **Inventory Management**: Real-time inventory updates
- **Customer Onboarding**: Automated welcome sequences
- **Payment Processing**: Secure payment handling
- **Shipping Integration**: Carrier API integration

#### [Customer Support with Zapier](automation/zapier-automation.yaml)
**Complexity**: Medium  
**Domain**: Customer Support  
**Features**: Zapier integration, simple webhooks, notification automation

Customer support system with Zapier automation:
- **Ticket Creation**: Automatic ticket generation
- **Email Notifications**: Customer and agent notifications
- **CRM Integration**: Customer data synchronization
- **Slack Alerts**: Team notifications
- **Follow-up Automation**: Automated follow-up sequences

#### [MCP Integration](automation/mcp-integration.yaml)
**Complexity**: Medium  
**Domain**: Data Integration  
**Features**: Model Context Protocol, external data access, tool integration

AI system with MCP server integration:
- **Database Access**: Direct database queries
- **File System Access**: Document and file operations
- **API Integration**: External service connections
- **Tool Execution**: Custom tool usage
- **Resource Management**: Efficient resource handling

### Templates and References

#### [Complete JSON Example](apai-0.1-example.json)
**Purpose**: Machine-readable reference  
**Use Case**: Tool development, SDK generation

Complete APAI specification in JSON format with:
- All specification sections populated
- Realistic example data
- Proper validation structure
- Tool-friendly format

## Getting Started with Examples

### 1. Choose Your Starting Point

**For Beginners:**
- Start with [Multilingual Chatbot](core/multilingual-chatbot.yaml)
- Then try [Customer Support AI](core/customer-support.yaml)

**For Intermediate Users:**
- Explore [Multi-Agent Customer Support](multi-agent/multi-agent-customer-support.yaml)
- Try [Zapier Automation](automation/zapier-automation.yaml)

**For Advanced Users:**
- Study [E-commerce with n8n](automation/ecommerce-automation.yaml)
- Build custom multi-agent systems

### 2. Validate Examples

All examples can be validated using APAI validators:

```bash
# Validate a specific example
python validators/python/apai_validator.py validate examples/core/customer-support.yaml

# Validate all examples
find examples -name "*.yaml" -exec python validators/python/apai_validator.py validate {} \;

# Validate with hierarchical composition
python validators/python/apai_validator.py validate examples/multi-agent/multi-agent-customer-support.yaml --hierarchical
```

### 3. Customize for Your Use Case

1. **Copy** an example that matches your needs
2. **Modify** the configuration for your domain
3. **Add** your specific models and constraints
4. **Validate** your customized specification
5. **Test** with your AI systems

## Example Comparison

| Example | Complexity | Domain | Key Features |
|---------|------------|--------|--------------|
| Multilingual Chatbot | Low | Customer Service | Multi-language, conversation |
| Customer Support AI | Medium | Customer Service | Orders, returns, recommendations |
| Content Moderator | Medium | Content Generation | Safety, filtering, moderation |
| Multi-Agent Support | High | Customer Service | Agent orchestration, specialization |
| E-commerce + n8n | High | Customer Service | Complex workflows, automation |
| Zapier Integration | Medium | Customer Service | Simple automation, notifications |
| MCP Integration | Medium | Customer Service | External data, tools, resources |

## Customization Guide

### Modifying Examples

1. **Change Models**: Update provider, model name, or parameters
2. **Adjust Constraints**: Modify safety, performance, or business rules
3. **Customize Prompts**: Adapt templates for your use case
4. **Add Tasks**: Include additional workflows or processes
5. **Configure Context**: Set up memory, business context, or MCP servers

### Best Practices

- **Start Simple**: Begin with basic examples and add complexity gradually
- **Validate Often**: Use validators to check your modifications
- **Document Changes**: Keep track of customizations for your team
- **Test Thoroughly**: Validate with real data and scenarios
- **Follow Patterns**: Use established patterns from examples

## Contributing Examples

We welcome new examples that demonstrate APAI capabilities:

1. **Identify** a use case not covered by existing examples
2. **Create** a complete, working specification
3. **Validate** using APAI validators
4. **Document** the example with clear descriptions
5. **Submit** a pull request with your example

### Example Requirements

- **Complete**: All required sections populated
- **Valid**: Passes APAI validation
- **Documented**: Clear descriptions and comments
- **Realistic**: Uses realistic data and scenarios
- **Educational**: Demonstrates best practices

## Support

- **Issues**: [Report problems with examples](https://github.com/FabioGuin/APAI/issues)
- **Discussions**: [Ask questions about examples](https://github.com/FabioGuin/APAI/discussions)
- **Documentation**: [Complete APAI documentation](../docs/)

---

**Note**: All examples are compatible with APAI 0.1.0. Check the [CHANGELOG](../CHANGELOG.md) for updates.
