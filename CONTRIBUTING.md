# Contributing to APAI

Thank you for your interest in contributing to APAI! This document provides guidelines and information for contributors.

## How to Contribute

### 1. Open an Issue

Before making changes, please open an issue to:
- Discuss your idea or problem
- Get feedback from the community
- Ensure your contribution aligns with project goals

### 2. Fork and Branch

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature-name`
3. Make your changes
4. Test your changes thoroughly

### 3. Submit a Pull Request

1. Push your branch to your fork
2. Open a pull request with a clear description
3. Include tests and documentation
4. Reference any related issues

## Development Setup

### Prerequisites

- Python 3.8+
- Node.js 16+ (for JavaScript tools)
- Go 1.19+ (for Go tools)

### Local Development

```bash
# Clone your fork
git clone https://github.com/FabioGuin/APAI.git
cd APAI

# Install Python dependencies
cd validators/python
pip install -r requirements.txt

# Install JavaScript dependencies
cd ../../validators/javascript
npm install

# Install Go dependencies
cd ../../validators/go
go mod tidy
```

## Contribution Guidelines

### Code Style

- **Python**: Follow PEP 8, use Black for formatting
- **JavaScript**: Follow ESLint configuration
- **Go**: Follow `gofmt` and `golint`
- **YAML**: Use consistent indentation (2 spaces)

### Documentation

- Update documentation for any changes
- Include examples for new features
- Follow the existing documentation style
- Use clear, concise language

### Testing

- Write tests for new features
- Ensure all tests pass
- Maintain or improve test coverage
- Include integration tests for validators

### Commit Messages

Use clear, descriptive commit messages:

```
feat: add support for multimodal models
fix: resolve validation error in prompt variables
docs: update README with new examples
test: add tests for constraint validation
```

## Types of Contributions

### Specification Changes

- **Minor Changes**: Bug fixes, clarifications
- **Major Changes**: New features, breaking changes
- **Process**: RFC-style discussion in issues

### Validators

- **Python**: Full-featured validation library
- **JavaScript**: Node.js and browser support
- **Go**: High-performance validation

### Tools

- **CLI**: Command-line validation and generation
- **Generators**: Documentation and code generation
- **Integrations**: IDE plugins, CI/CD tools

### Documentation

- **Specification**: Core specification documentation
- **Examples**: Real-world usage examples
- **Tutorials**: Getting started guides
- **API Docs**: Tool and library documentation

## Review Process

### Pull Request Review

1. **Automated Checks**: CI/CD runs tests and validation
2. **Maintainer Review**: Code review by maintainer
3. **Community Feedback**: Open for community input
4. **Testing**: Manual testing of changes
5. **Approval**: Maintainer approves and merges

### Review Criteria

- **Functionality**: Does it work as intended?
- **Quality**: Is the code well-written and tested?
- **Documentation**: Is it properly documented?
- **Compatibility**: Does it maintain backward compatibility?
- **Security**: Are there any security concerns?

## Release Process

### Versioning

We use [Semantic Versioning](https://semver.org/):

- **MAJOR**: Breaking changes to the specification
- **MINOR**: New features, backward compatible
- **PATCH**: Bug fixes, backward compatible

### Release Schedule

- **Patch Releases**: As needed for bug fixes
- **Minor Releases**: Every 3-6 months
- **Major Releases**: Periodic releases

## Community Guidelines

### Be Respectful

- Use welcoming and inclusive language
- Be respectful of differing viewpoints
- Focus on what is best for the community
- Show empathy towards other community members

### Be Constructive

- Provide constructive feedback
- Help others learn and improve
- Share knowledge and experience
- Be patient with newcomers

### Be Professional

- Stay on topic in discussions
- Use appropriate channels for different types of communication
- Follow the project's code of conduct
- Respect maintainers' time and decisions

## Getting Help

### Resources

- **Documentation**: Check the [docs](docs/) directory
- **Examples**: See [examples](examples/)
- **Issues**: Search existing issues for similar problems
- **Discussions**: Use GitHub Discussions for questions

### Contact

- **GitHub Issues**: [GitHub Issues](https://github.com/FabioGuin/APAI/issues)
- **Discussions**: [GitHub Discussions](https://github.com/FabioGuin/APAI/discussions)

## Recognition

Contributors will be recognized in:
- **CONTRIBUTORS.md**: List of all contributors
- **Release Notes**: Credit for significant contributions
- **Documentation**: Attribution for documentation contributions

## License

By contributing to APAI, you agree that your contributions will be licensed under the Apache License 2.0. See [LICENSE](LICENSE) for details.

---

Thank you for contributing to APAI! Your contributions help make AI systems more transparent, safe, and interoperable.
