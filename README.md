# CF Purge - Bulk CloudFormation Stacks Deletion

Delete one or more CloudFormation stacks (that match a glob-pattern string) from an AWS account

### Usage

1. Clone the repo
2. `go get`
3. `go run . --glob "*-some-glob-pattern-*"`

### Roadmap

- [ ] Handle delete failures due to stack dependencies
- [ ] Delete by prefix or suffix, not just by glob pattern
- [ ] CI (GH Actions) to auto publish binaries + update documentation to simplify usage of the tool

### Contributions

All contributions are welcome. This is still an alpha project, so expect bugs.
