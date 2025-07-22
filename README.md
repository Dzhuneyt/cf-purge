![](./docs/cf-purge.webp)

# CF Purge

This repository is now structured as a **monorepo**. The main Go CLI app is located in the `core/` folder. Other projects (such as a website) will live in sibling directories.

## Structure

- `core/` — Go CLI app for bulk-deleting CloudFormation stacks
- `website/` — (Coming soon) Static website/landing page (e.g., Hugo)
- `docs/` — Documentation assets

## Go CLI Usage

### Installation

```
brew tap dzhuneyt/tap
brew install cf-purge
```

Or grab one of the [latest releases](https://github.com/Dzhuneyt/cf-purge/releases).

**Usage**

From anywhere (after install), usage is unchanged:

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
