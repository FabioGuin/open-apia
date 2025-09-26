# Security Best Practices for APAI

This document outlines security best practices for managing sensitive data in APAI specifications and implementations.

## Table of Contents

1. [Environment Variables and Secrets Management](#environment-variables-and-secrets-management)
   - [Overview](#overview)
   - [File Structure](#file-structure)
   - [Environment Variables Template](#environment-variables-template)
   - [YAML Configuration](#yaml-configuration)
   - [Common Patterns](#common-patterns)

2. [Security Guidelines](#security-guidelines)
   - [1. Never Commit Secrets](#1-never-commit-secrets)
   - [2. Use Strong, Unique Secrets](#2-use-strong-unique-secrets)
   - [3. Principle of Least Privilege](#3-principle-of-least-privilege)
   - [4. Secure Storage](#4-secure-storage)
   - [5. Monitoring and Logging](#5-monitoring-and-logging)

3. [Implementation Examples](#implementation-examples)
   - [Basic Setup](#basic-setup)
   - [Validation](#validation)

4. [Common Security Anti-Patterns](#common-security-anti-patterns)

5. [Environment-Specific Configurations](#environment-specific-configurations)
   - [Development](#development)
   - [Production](#production)

6. [Compliance and Standards](#compliance-and-standards)

7. [Additional Resources](#additional-resources)

8. [Support](#support)

---

## Environment Variables and Secrets Management

### Overview

APAI specifications may contain sensitive information such as API keys, database credentials, tokens, and other secrets. To maintain security and follow industry standards, all sensitive data should be managed through environment variables.

### File Structure

```
project/
├── .env.example          # Template with placeholder values (committed to git)
├── .env                  # Actual values (NEVER committed to git)
├── .gitignore           # Excludes .env files
└── apai-spec.yaml   # Uses ${VARIABLE_NAME} syntax
```

### Environment Variables Template

Use the provided `.env.example` file as a template. This file contains only the environment variables actually used in the APAI examples and documentation:

- **AI Model Providers**: OpenAI API key
- **Database Connections**: Database API keys, tokens, passwords, and connection URLs
- **Customer & Order Databases**: Specific database credentials
- **Knowledge Base MCP Server**: Authentication tokens for MCP servers
- **Redis & Caching**: Cache server connection details
- **Automation Integration**: n8n and Zapier webhook configurations (including multiple webhook URLs)
- **External Services**: Slack bot tokens and other service integrations
- **Generic API Keys**: For documentation examples and templates

### YAML Configuration

In your APAI YAML files, use the `${VARIABLE_NAME}` syntax to reference environment variables:

```yaml
# ❌ NEVER do this - hardcoded secrets
authentication:
  type: "api_key"
  api_key: "sk-1234567890abcdef"

# ✅ DO this - use environment variables
authentication:
  type: "api_key"
  api_key: "${OPENAI_API_KEY}"
```

### Common Patterns

#### Database Connections
```yaml
context:
  memory:
    storage:
      provider: "redis"
      url: "${REDIS_URL}"
      password: "${REDIS_PASSWORD}"
```

#### MCP Server Authentication
```yaml
mcp_servers:
  - id: "database-server"
    authentication:
      type: "api_key"
      api_key: "${DB_API_KEY}"
    security:
      headers:
        "Authorization": "Bearer ${DB_TOKEN}"
```

#### External Service Integration
```yaml
automation:
  triggers:
    - type: "webhook"
      url: "${ZAPIER_WEBHOOK_URL}"
      headers:
        "X-API-Key": "${ZAPIER_API_KEY}"
```

## Security Guidelines

### 1. Never Commit Secrets

- **ALWAYS** add `.env` to `.gitignore`
- **NEVER** commit files containing real API keys or passwords
- Use `.env.example` as a template for other developers

### 2. Use Strong, Unique Secrets

- Generate strong, random API keys and passwords
- Use different secrets for different environments (development, staging, production)
- Rotate secrets regularly

### 3. Principle of Least Privilege

- Only grant necessary permissions to API keys
- Use environment-specific credentials
- Regularly audit and revoke unused access

### 4. Secure Storage

- Use secure secret management systems in production (AWS Secrets Manager, Azure Key Vault, etc.)
- Encrypt secrets at rest
- Use secure communication channels (HTTPS/TLS)

### 5. Monitoring and Logging

- Monitor for unauthorized access attempts
- Log security events (without logging the actual secrets)
- Set up alerts for suspicious activity

## Implementation Examples

### Basic Setup

1. **Copy the template**:
   ```bash
   cp .env.example .env
   ```

2. **Fill in your values**:
   ```bash
   # Edit .env with your actual credentials
   OPENAI_API_KEY=sk-your-actual-key-here
   DB_PASSWORD=your-secure-password
   ```

3. **Load in your application**:
   ```python
   import os
   from dotenv import load_dotenv
   
   load_dotenv()
   
   openai_key = os.getenv('OPENAI_API_KEY')
   ```

### Validation

Always validate that required environment variables are present:

```python
import os
import sys

required_vars = ['OPENAI_API_KEY', 'DB_PASSWORD', 'REDIS_URL']
missing_vars = [var for var in required_vars if not os.getenv(var)]

if missing_vars:
    print(f"Missing required environment variables: {missing_vars}")
    sys.exit(1)
```

## Common Security Anti-Patterns

### ❌ Don't Do This

```yaml
# Hardcoded secrets
api_key: "sk-1234567890abcdef"
password: "mypassword123"

# Inline credentials
database_url: "postgresql://user:password@localhost/db"

# Plain text tokens
slack_token: "xoxb-1234567890-abcdef"
```

### ✅ Do This Instead

```yaml
# Environment variables
api_key: "${OPENAI_API_KEY}"
password: "${DB_PASSWORD}"

# Environment-based URLs
database_url: "${DATABASE_URL}"

# Secure token references
slack_token: "${SLACK_BOT_TOKEN}"
```

## Environment-Specific Configurations

### Development
```bash
# .env.development
APP_ENV=development
APP_DEBUG=true
DATABASE_URL=postgresql://localhost:5432/dev_db
```

### Production
```bash
# .env.production
APP_ENV=production
APP_DEBUG=false
DATABASE_URL=postgresql://prod-server:5432/prod_db
```

## Compliance and Standards

This security approach aligns with:

- **ISO/IEC 27001**: Information security management
- **OWASP Top 10**: Web application security risks
- **NIST Cybersecurity Framework**: Security best practices
- **GDPR**: Data protection and privacy requirements

## Additional Resources

- [OWASP Secrets Management Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Secrets_Management_Cheat_Sheet.html)
- [12-Factor App Methodology](https://12factor.net/config)
- [Environment Variables Best Practices](https://blog.doppler.com/environment-variables-in-python)

## Support

For security-related questions or to report vulnerabilities, please contact the APAI security team or create an issue in the project repository.
