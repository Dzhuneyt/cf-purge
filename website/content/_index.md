---
title: '🚀 CF-Purge'
description: 'Quickly delete AWS CloudFormation stacks that match a pattern. Install instantly with Homebrew.'
---

<div class="prose prose-slate">
  <h2>Quickly delete CloudFormation stacks that match a pattern.</h2>
  <p>
    CF-Purge is a blazing-fast CLI tool to bulk-delete AWS CloudFormation stacks using simple glob patterns. Perfect for cleaning up dev/test environments or automating stack management.
  </p>
  <a href="https://github.com/Dzhuneyt/cf-purge" target="_blank" class="bg-blue-700 text-white px-6 py-3 rounded-md shadow hover:bg-blue-800 transition">⭐ Star on GitHub</a>
</div>

<div class="grid grid-cols-1 md:grid-cols-2 gap-10 mb-12">
  <div>
    <h3>🍺 Install with Homebrew</h3>
    <pre class="bg-gray-900 text-green-300 rounded px-4 py-3 text-sm overflow-x-auto"><code>brew tap dzhuneyt/tap
brew install cf-purge</code></pre>
    <p>Or grab one of the <a href="https://github.com/Dzhuneyt/cf-purge/releases" class="text-blue-700 underline">latest releases</a>.</p>
  </div>
  <div>
    <h3>⚡ Usage Examples</h3>
    <p>Delete all stacks that match a glob pattern:</p>
    <pre class="bg-gray-900 text-green-300 rounded px-4 py-3 text-sm overflow-x-auto"><code>cf-purge --glob "*-some-glob-pattern-*"</code></pre>
    <p>Delete all stacks that end with <code>-api</code>:</p>
    <pre class="bg-gray-900 text-green-300 rounded px-4 py-3 text-sm overflow-x-auto"><code>cf-purge --glob "*-api"</code></pre>
    <p>Delete all stacks that start with <code>my-stack-</code>:</p>
    <pre class="bg-gray-900 text-green-300 rounded px-4 py-3 text-sm overflow-x-auto"><code>cf-purge --glob "my-stack-*"</code></pre>
  </div>
</div>

<div class="mb-12">
  <h3>Why CF-Purge?</h3>
  <ul>
    <li>🏎️ <strong>Fast</strong>: Instantly finds and deletes stacks in bulk</li>
    <li>🔍 <strong>Flexible</strong>: Use glob patterns to match exactly what you want</li>
    <li>🛡️ <strong>Safe</strong>: Only deletes stacks you specify</li>
    <li>🍺 <strong>Easy Install</strong>: Homebrew support for one-command setup</li>
    <li>🧑‍💻 <strong>Open Source</strong>: <a href="https://github.com/Dzhuneyt/cf-purge" class="text-blue-700 underline">View on GitHub</a></li>
  </ul>
</div>

<div class="flex gap-4">
  <a href="https://github.com/Dzhuneyt/cf-purge" target="_blank" class="bg-blue-700 text-white px-6 py-3 rounded-md font-bold text-base shadow hover:bg-blue-800 transition">GitHub</a>
  {{< raw >}}<a href="{{ "/about/" | relURL }}" class="bg-gray-100 text-gray-800 px-6 py-3 rounded-md font-bold text-base shadow hover:bg-gray-200 transition">Learn More</a>{{< /raw >}}
</div>
