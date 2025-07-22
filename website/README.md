# Website Subfolder

This directory contains the source code and configuration for the documentation website of the `cf-purge` project.

## What is this folder?
This `website` subfolder is dedicated to the static site, powered by Hugo. It contains all the files necessary to build and serve the documentation or marketing site for the project.

## How to Start the Project

### Prerequisites
- [Hugo](https://gohugo.io/getting-started/installing/) installed on your system.

### Tailwind CSS Workflow

This project uses [Tailwind CSS](https://tailwindcss.com/) via the Tailwind CLI. **You must pre-compile your CSS before serving or building the site.**

#### Local Development
1. In one terminal, start the Tailwind CLI in watch mode:
   ```sh
   pnpm run tailwind:dev
   ```
   This will rebuild `static/css/style.css` whenever you change content, layouts, or Tailwind config.
2. In another terminal, start the Hugo development server:
   ```sh
   hugo server -D
   ```
3. Open your browser and go to the address shown in the terminal (usually http://localhost:1313).

#### Building for Production
Before running the Hugo build, generate the production (minified) CSS:
```sh
pnpm run tailwind:build
hugo
```
- The output will be in the `public/` directory, ready for deployment.
- Make sure to commit the latest `static/css/style.css` if your deployment pipeline does not run the Tailwind build step automatically.

#### CI / Automated Production Builds
If building in CI, add a pre-build step:
```sh
pnpm install --frozen-lockfile
pnpm run tailwind:build
hugo
```
- Ensure Node.js, PNPM, and Hugo are available in your CI environment.
- You can cache `node_modules` for faster builds.

## More Information
- See the main project README for overall project details.
- For Hugo-specific configuration, check `hugo.toml` in this folder.
