![](./docs/cf-purge.webp)

# CF Purge

Quickly delete multiple CloudFormation stacks as a bulk operation.

Delete stacks that start with a prefix, end with a
suffix, or contain a string in their name.

# Usage

**Download**

_MacOS:_

```bash
curl -Los cf-purge "https://github.com/Dzhuneyt/cf-purge/releases/latest/download/cf-purge-darwin-$(uname -m)"
chmod +x cf-purge
sudo mv cf-purge /usr/local/bin/
```

_Linux:_

```bash
curl -Los cf-purge "https://github.com/Dzhuneyt/cf-purge/releases/latest/download/cf-purge-linux-$(uname -m)"
chmod +x cf-purge
sudo mv cf-purge /usr/local/bin/
```

_Windows:_

Just grab the latest .exe from the [Releases](https://github.com/Dzhuneyt/cf-purge/releases) page, based on your CPU
architecture.

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
- [ ] Delete by prefix or suffix, not just by glob pattern
- [ ] CI (GH Actions) to auto publish binaries + update documentation to simplify usage of the tool

# Contributions

All contributions are welcome. Please open an issue or a PR if you have any ideas or improvements.
