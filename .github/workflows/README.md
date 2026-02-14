# Docker CI/CD with GitHub Actions

This directory contains GitHub Actions workflows for building and publishing Docker images for the robohub-inventory service.

## Workflows

### 1. Docker Build and Publish (`docker-publish.yml`)

**Automatic workflow** that builds and publishes Docker images to GitHub Container Registry (ghcr.io).

#### Triggers

- **Push to main/develop branches**: Builds and publishes images tagged with the branch name
- **Git tags** (e.g., `v1.0.0`): Builds and publishes release images with semantic versioning tags
- **Pull requests**: Builds images (but doesn't push) to validate the Dockerfile

#### Image Tags

The workflow automatically generates multiple tags for each build:

| Event | Generated Tags | Example |
|-------|---------------|---------|
| Push to `main` | `main`, `main-<sha>`, `latest` | `main`, `main-abc1234`, `latest` |
| Push to `develop` | `develop`, `develop-<sha>` | `develop`, `develop-xyz5678` |
| Tag `v1.2.3` | `1.2.3`, `1.2`, `1`, `v1.2.3` | Semantic version tags |
| Pull Request | `pr-<number>` | `pr-42` |

#### Features

- **Multi-platform builds**: Builds for both `linux/amd64` and `linux/arm64`
- **Build cache**: Uses GitHub Actions cache for faster builds
- **Artifact attestation**: Generates provenance attestation for security
- **Metadata**: Includes labels with build information

### 2. Docker Manual Build (`docker-manual.yml`)

**Manual workflow** that can be triggered via GitHub UI for custom builds.

#### Usage

1. Go to **Actions** tab in your GitHub repository
2. Select **Docker Manual Build** workflow
3. Click **Run workflow**
4. Configure options:
   - **tag**: Custom tag for the image (e.g., `test`, `v1.0.0-beta`)
   - **platforms**: Target platforms (default: `linux/amd64,linux/arm64`)
   - **push**: Whether to push to registry (default: `true`)

#### Use Cases

- Testing Dockerfile changes with custom tags
- Building single-platform images for faster iteration
- Creating special release candidates or beta versions

## Setup Instructions

### 1. Enable GitHub Container Registry

GitHub Container Registry (ghcr.io) is automatically enabled for your repository. No additional setup required!

### 2. Configure Repository Permissions

The workflows use `GITHUB_TOKEN` which is automatically provided by GitHub Actions. Ensure your repository has the following permissions:

1. Go to repository **Settings** → **Actions** → **General**
2. Under "Workflow permissions":
   - Select **Read and write permissions**
   - Check **Allow GitHub Actions to create and approve pull requests**

### 3. First Build

Push a commit to `main` branch or create a tag to trigger the first build:

```bash
# Option 1: Push to main
git add .
git commit -m "Add GitHub Actions workflow"
git push origin main

# Option 2: Create a release tag
git tag v1.0.0
git push origin v1.0.0
```

## Using the Published Images

### Pull the Latest Image

```bash
docker pull ghcr.io/<OWNER>/robohub-inventory:latest
```

Replace `<OWNER>` with your GitHub username or organization name.

### Pull a Specific Version

```bash
docker pull ghcr.io/<OWNER>/robohub-inventory:1.2.3
```

### Run the Image

```bash
docker run -p 8080:8080 \
  -e DB_HOST=your-db-host \
  -e DB_USER=postgres \
  -e DB_PASSWORD=your-password \
  -e DB_NAME=robohub_inventory \
  ghcr.io/<OWNER>/robohub-inventory:latest
```

### Authentication

Public images can be pulled without authentication. For private repositories:

```bash
# Login to GitHub Container Registry
echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin

# Pull the image
docker pull ghcr.io/<OWNER>/robohub-inventory:latest
```

## Image Information

Published images include the following metadata:

- **Source**: Link to the GitHub repository
- **Version**: Git tag or branch name
- **Commit**: Git commit SHA
- **Build Date**: Timestamp of the build
- **License**: Repository license

View image metadata:

```bash
docker inspect ghcr.io/<OWNER>/robohub-inventory:latest
```

## Advanced Configuration

### Building for Single Platform

For faster development builds, you can modify the workflow to build for a single platform:

```yaml
platforms: linux/amd64  # Only build for amd64
```

### Using Docker Hub Instead of GHCR

To publish to Docker Hub instead:

1. Add Docker Hub credentials to repository secrets:
   - `DOCKER_USERNAME`
   - `DOCKER_PASSWORD`

2. Update the workflow:

```yaml
env:
  REGISTRY: docker.io
  IMAGE_NAME: <your-dockerhub-username>/robohub-inventory

# In login step:
- name: Log into Docker Hub
  uses: docker/login-action@v3
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
```

### Custom Build Arguments

Add build arguments to your Dockerfile and pass them in the workflow:

```yaml
build-args: |
  VERSION=${{ steps.meta.outputs.version }}
  COMMIT=${{ github.sha }}
  BUILD_DATE=${{ steps.meta.outputs.created }}
  GO_VERSION=1.21
```

## Monitoring Builds

### View Build Status

- **Badge**: Add a status badge to your README:
  ```markdown
  ![Docker Build](https://github.com/<OWNER>/robohub-inventory/actions/workflows/docker-publish.yml/badge.svg)
  ```

- **Actions Tab**: View detailed logs in the Actions tab of your repository

### Email Notifications

GitHub automatically sends email notifications for failed workflows. Configure in your GitHub notification settings.

## Troubleshooting

### Build Fails: Permission Denied

Ensure workflow permissions are set to "Read and write" in repository settings.

### Build Fails: Rate Limited

Multi-platform builds may hit rate limits on Docker Hub for base images. Solution:
- Use GitHub Actions cache (already configured)
- Authenticate to Docker Hub to increase rate limits

### Image Not Visible in GHCR

Make images public:
1. Go to your package at `https://github.com/users/<OWNER>/packages/container/robohub-inventory`
2. Click **Package settings**
3. Under **Danger Zone**, click **Change visibility**
4. Select **Public**

## Security Best Practices

1. **Don't commit secrets**: Use GitHub Secrets for sensitive data
2. **Scan images**: Consider adding Trivy or Snyk scanning to the workflow
3. **Pin action versions**: Workflows use pinned versions (e.g., `@v4`)
4. **Minimal base images**: Using Alpine Linux for smaller attack surface
5. **Attestation**: Build provenance is automatically generated

## Example: Adding Image Scanning

Add this step before pushing:

```yaml
- name: Run Trivy vulnerability scanner
  uses: aquasecurity/trivy-action@master
  with:
    image-ref: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:${{ steps.meta.outputs.tags }}
    format: 'sarif'
    output: 'trivy-results.sarif'

- name: Upload Trivy results to GitHub Security
  uses: github/codeql-action/upload-sarif@v2
  with:
    sarif_file: 'trivy-results.sarif'
```

## Support

For issues with the workflows, please open an issue in the repository.
