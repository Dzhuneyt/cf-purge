---
title: "About"
description: "About this site"
---

<div class="prose prose-slate">

[← Back to Home]({{< relref "/_index.md" >}})


CF-Purge was created to solve a common pain point for teams using AWS CloudFormation and the AWS CDK: cleaning up multiple stacks quickly and safely.

### Why did I build this tool?

When working with the AWS CDK, it's common for a project to create many CloudFormation stacks—sometimes 10 or more per environment. If a developer leaves the project, their environment (with all its stacks) needs to be deleted. Manually finding and deleting these stacks in the AWS Console is tedious and error-prone.

CF-Purge makes this process fast and reliable. You can use glob patterns, prefixes, or suffixes to match exactly the stacks you want to delete—all from the command line, in seconds, instead of minutes of clicking in the AWS Console.

This tool is perfect for:
- Cleaning up dev/test environments
- Offboarding developers
- Automating stack management in CI/CD pipelines
- Saving time and reducing the risk of missed stacks

CF-Purge turns a slow, manual process into a single, powerful command.

</div>

