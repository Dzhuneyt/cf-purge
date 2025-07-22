# Website Subfolder

This directory contains the source code and configuration for the documentation website of the `cf-purge` project.

## What is this folder?
This `website` subfolder is dedicated to the static site, powered by Hugo. It contains all the files necessary to build and serve the documentation or marketing site for the project.

## How to Start the Project

### Prerequisites
- [Hugo](https://gohugo.io/getting-started/installing/) installed on your system.

### Running Locally

1. Open a terminal and navigate to this directory:
   ```sh
   cd website
   ```
2. Start the Hugo development server:
   ```sh
   hugo server -D
   ```
3. Open your browser and go to the address shown in the terminal (usually http://localhost:1313).

### Building for Production
To generate the static files for deployment:
```sh
hugo
```
The output will be in the `public/` directory.

## More Information
- See the main project README for overall project details.
- For Hugo-specific configuration, check `hugo.toml` in this folder.
