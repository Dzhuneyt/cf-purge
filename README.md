![](./docs/cf-purge.webp)

# CF Purge

Quickly delete multiple CloudFormation stacks as a bulk operation.

Delete stacks that start with a prefix, end with a
suffix, or contain a string in their name.

# Usage

## Installation

```
brew tap dzhuneyt/tap
brew install cf-purge
```

Or grab one of the [latest releases](https://github.com/Dzhuneyt/cf-purge/releases).

**Usage**

cf-purge is a command-line tool. It takes a single argument - a glob pattern string that matches the CloudFormation
stack names you want to delete.

Some examples:

```bash
cf-purge --glob "*-some-glob-pattern-*" # Deletes all stacks that match the glob pattern
cf-purge --glob "*-api" # Deletes all stacks that end with "-api"
cf-purge --glob "my-stack-*" # Deletes all stacks that start with "my-stack-"
```

# Roadmap

- [ ] Handle delete failures due to stack dependencies
- [X] Delete by prefix or suffix, not just by glob pattern
- [X] CI to autopublish to HomeBrew for easier installation and usage

# Contributions

All contributions are welcome. Please open an issue or a PR if you have any ideas or improvements.
