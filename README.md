![](./docs/cf-purge.webp)

# CF Purge

![Go Version](https://img.shields.io/badge/Go-1.22+-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Release](https://img.shields.io/github/v/release/Dzhuneyt/cf-purge)

A powerful CLI tool for bulk-deleting AWS CloudFormation stacks using glob patterns. Perfect for cleaning up development environments, CI/CD temporary stacks, and feature branch deployments.

## ⚠️ Important Safety Notice

**CF Purge is a powerful tool that PERMANENTLY DELETES AWS CloudFormation stacks and their resources.**

- 🚨 **Always test in non-production environments first**
- 🔍 **Understand the implications of deleting your infrastructure**
- 💾 **Have proper backups of critical data**
- ✅ **Use specific glob patterns to avoid accidental deletions**
- 🎯 **Review the list of matched stacks before confirming deletion**

This tool will prompt you for confirmation before deleting any stacks, but **deletion is irreversible**.

## Repository Structure

This repository is structured as a **monorepo**:

- `core/` — Go CLI app for bulk-deleting CloudFormation stacks
- `website/` — Static website/landing page
- `docs/` — Documentation assets

## Prerequisites

Before using CF Purge, ensure you have:

- ✅ AWS CLI configured with appropriate credentials
- ✅ IAM permissions for CloudFormation operations (see [AWS Configuration](#aws-configuration))
- ✅ Go 1.22+ (only for building from source)
- ✅ Basic understanding of CloudFormation stacks
- ✅ Non-production environment for initial testing

## Installation

### Option 1: Homebrew (Recommended)

```bash
brew tap dzhuneyt/tap
brew install cf-purge
cf-purge --version  # Verify installation
```

### Option 2: Pre-built Binaries

Download the latest binary for your platform from [GitHub Releases](https://github.com/Dzhuneyt/cf-purge/releases).

Supported platforms:
- Linux (amd64, arm64, 386)
- macOS (amd64, arm64)
- Windows (amd64, 386)

### Option 3: Build from Source

```bash
git clone https://github.com/Dzhuneyt/cf-purge.git
cd cf-purge/core
go build -o cf-purge main.go
```

### Verification

```bash
cf-purge --help
```

## AWS Configuration

### Required IAM Permissions

Create an IAM policy with the following permissions:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "cloudformation:ListStacks",
                "cloudformation:DeleteStack",
                "cloudformation:DescribeStacks"
            ],
            "Resource": "*"
        }
    ]
}
```

### AWS Credentials Setup

Ensure AWS credentials are configured via one of these methods:

- **AWS CLI**: `aws configure`
- **Environment variables**: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
- **IAM roles** (for EC2/Lambda execution)
- **AWS profiles**: `export AWS_PROFILE=your-profile`

Verify your configuration:
```bash
aws sts get-caller-identity
```

## Usage

### Basic Syntax

```bash
cf-purge --glob "PATTERN"
```

### Pattern Matching Examples

```bash
# Delete all development environment stacks
cf-purge --glob "*-dev"

# Delete all stacks ending with "-api"
cf-purge --glob "*-api"

# Delete all stacks starting with "my-stack-"
cf-purge --glob "my-stack-*"

# Delete stacks matching a specific pattern
cf-purge --glob "*-feature-branch-*"
```

### Real-World Use Cases

#### 1. Clean up CI/CD temporary stacks
```bash
cf-purge --glob "ci-build-*"
cf-purge --glob "*-pr-*"
```

#### 2. Remove old feature branch stacks
```bash
cf-purge --glob "*-feature-user-auth-*"
cf-purge --glob "myapp-branch-*"
```

#### 3. Cleanup development environments
```bash
cf-purge --glob "*-dev"
cf-purge --glob "*-staging"
cf-purge --glob "test-*-tmp"
```

#### 4. Remove stacks by environment type
```bash
cf-purge --glob "*-sandbox-*"
cf-purge --glob "*-playground-*"
```

### Interactive Confirmation

CF Purge will show you exactly which stacks match your pattern and ask for confirmation:

```
Glob pattern matched the following stacks:
----------------------------------------
my-app-dev-api
my-app-dev-database
my-app-dev-frontend
----------------------------------------
This operation will delete 3 stacks
THIS OPERATION IS DESTRUCTIVE AND IRREVERSIBLE
Please, confirm that you want to delete these stacks irreversibly.
[y/N]: 
```

**Type `y` and press Enter to proceed, or any other key to cancel.**

## Troubleshooting

### Common Issues

#### "Access Denied" Error
- **Cause**: Insufficient IAM permissions
- **Solution**: Ensure your AWS credentials have the required CloudFormation permissions ([see AWS Configuration](#aws-configuration))

#### Stack Deletion Fails
- **Cause**: Stack resources have dependencies, deletion protection, or retain policies
- **Solution**: 
  - Check stack events in the AWS Console for specific resource errors
  - Manually resolve resource dependencies
  - Disable deletion protection on protected resources

#### No Stacks Match Pattern
- **Cause**: Glob pattern doesn't match any existing stack names
- **Solution**: 
  - List all stacks: `aws cloudformation list-stacks`
  - Test your pattern with a broader glob first
  - Check for typos in the pattern

### Best Practices for Pattern Testing

Before running CF Purge, test your glob pattern:

```bash
# List all stacks to see available names
aws cloudformation list-stacks --query 'StackSummaries[].StackName' --output table

# Test pattern matching with AWS CLI
aws cloudformation list-stacks --query 'StackSummaries[?contains(StackName, `your-pattern`)].StackName' --output table
```

### Getting Help

- 📖 Check [Issues](https://github.com/Dzhuneyt/cf-purge/issues) for similar problems
- 📚 Review [AWS CloudFormation documentation](https://docs.aws.amazon.com/cloudformation/) for stack deletion limitations
- 💬 Use `cf-purge --help` for command options
- 🐛 Report bugs or request features via [GitHub Issues](https://github.com/Dzhuneyt/cf-purge/issues)

## Performance & Limitations

### Performance Characteristics
- ⚡ Processes stacks sequentially to avoid AWS API throttling
- 🕐 Large numbers of stacks may take several minutes to delete
- 🔒 No parallel deletion to ensure safe operation and proper confirmation

### Current Limitations
- ❌ Stack deletion failures require manual intervention (see roadmap)
- ❌ No progress indicator for long-running operations
- ❌ Limited to glob pattern matching (regex support planned)
- ❌ No dry-run mode (planned feature)

### Best Practices
- 🎯 Use specific patterns to avoid accidental deletions
- 🧪 Test patterns in development environments first
- 📦 Delete in smaller batches for large numbers of stacks
- 📋 Keep a record of important stacks before bulk operations

## Roadmap

- [ ] **Dry-run mode** - Preview which stacks would be deleted without actually deleting them
- [ ] **Handle delete failures** due to stack dependencies with automatic retries
- [ ] **Progress indicators** for long-running operations
- [ ] **Regex pattern support** in addition to glob patterns
- [ ] **Parallel deletion** with configurable concurrency
- [ ] **Stack filtering** by tags, creation date, or status
- [x] ✅ Delete by prefix or suffix, not just by glob pattern
- [x] ✅ CI to auto-publish to HomeBrew for easier installation and usage

## Contributing

We welcome contributions! Here's how you can help:

### Development Setup

```bash
git clone https://github.com/Dzhuneyt/cf-purge.git
cd cf-purge/core
go mod download
go run main.go --help
```

### Ways to Contribute

- 🐛 Report bugs or request features via [Issues](https://github.com/Dzhuneyt/cf-purge/issues)
- 📖 Improve documentation
- 🧪 Add test cases
- ⚡ Implement features from the roadmap
- 🎨 Enhance the website

### Pull Request Process

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature-name`
3. Make your changes in the `core/` directory
4. Test thoroughly with non-production stacks
5. Commit your changes: `git commit -m "Add your feature"`
6. Push to your fork: `git push origin feature/your-feature-name`
7. Submit a pull request with a clear description

## Related Tools

- [AWS CLI CloudFormation](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/) - Official AWS CLI for CloudFormation
- [cfn-lint](https://github.com/aws-cloudformation/cfn-lint) - CloudFormation template validation
- [aws-nuke](https://github.com/rebuy-de/aws-nuke) - Nuclear option for AWS account cleanup
- [cfn-teardown](https://github.com/cloudtools/cfn-teardown) - Alternative stack deletion tool

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📧 **Issues**: [GitHub Issues](https://github.com/Dzhuneyt/cf-purge/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/Dzhuneyt/cf-purge/discussions)
- 📖 **Documentation**: This README and [GitHub Wiki](https://github.com/Dzhuneyt/cf-purge/wiki)

---

**Remember: With great power comes great responsibility. Use CF Purge wisely! 🚀**
