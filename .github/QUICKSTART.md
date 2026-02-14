# Quick Start: CI/CD Setup

This guide will help you get the GitHub Actions CI/CD pipeline running for your robohub-inventory service.

## Prerequisites

- GitHub repository for this project
- GitHub account with access to the repository

## Setup Steps

### 1. Push the Workflow Files

The workflow files are already in your repository at `.github/workflows/`. Push them to GitHub:

```bash
git add .github/
git add .dockerignore
git add README.md
git commit -m "Add GitHub Actions CI/CD for Docker builds"
git push origin main
```

### 2. Configure Repository Permissions

1. Go to your repository on GitHub
2. Click **Settings** → **Actions** → **General**
3. Scroll to **Workflow permissions**
4. Select **Read and write permissions**
5. Check **Allow GitHub Actions to create and approve pull requests**
6. Click **Save**

### 3. Verify the First Build

After pushing, GitHub Actions will automatically trigger a build:

1. Go to the **Actions** tab in your repository
2. You should see a workflow run for "Docker Build and Publish"
3. Click on it to see the build progress
4. Wait for it to complete (usually 3-5 minutes)

### 4. Find Your Published Image

Once the build succeeds:

1. Go to your repository main page
2. Look for **Packages** in the right sidebar
3. Click on `robohub-inventory`
4. You'll see your published Docker image with tags

### 5. Make the Image Public (Optional)

By default, the image is private. To make it public:

1. Click on the package name
2. Click **Package settings** (bottom right)
3. Scroll to **Danger Zone**
4. Click **Change visibility**
5. Select **Public**
6. Confirm the change

### 6. Test Pulling the Image

```bash
# Public image (no authentication needed)
docker pull ghcr.io/YOUR_GITHUB_USERNAME/robohub-inventory:latest

# Run the image
docker run -p 8080:8080 \
  -e DB_HOST=localhost \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  -e DB_NAME=robohub_inventory \
  ghcr.io/YOUR_GITHUB_USERNAME/robohub-inventory:latest
```

## Creating Releases

To create a versioned release:

```bash
# Create and push a tag
git tag v1.0.0
git push origin v1.0.0
```

This will trigger a build with the following tags:
- `v1.0.0`
- `1.0.0`
- `1.0`
- `1`

## Manual Builds

To trigger a manual build with custom settings:

1. Go to **Actions** tab
2. Select **Docker Manual Build** from the left sidebar
3. Click **Run workflow**
4. Fill in the parameters:
   - **tag**: Your custom tag (e.g., `test`, `dev`)
   - **platforms**: Target platforms (default: both amd64 and arm64)
   - **push**: Whether to push to registry
5. Click **Run workflow**

## Troubleshooting

### "Permission denied" error

**Solution**: Enable "Read and write permissions" in Settings → Actions → General

### Image not visible after build

**Solution**: Check the Actions tab for build logs. The image may be private - see step 5 above.

### Rate limit errors

**Solution**: The workflow uses GitHub Actions cache. If you still hit rate limits, consider:
- Reducing build frequency
- Authenticating to Docker Hub for higher limits

### Build fails on ARM64

**Solution**: QEMU emulation can be slow. For faster iterations during development:
1. Edit `.github/workflows/docker-publish.yml`
2. Change `platforms: linux/amd64,linux/arm64` to `platforms: linux/amd64`
3. Build ARM64 images only for releases

## Next Steps

- **Add status badge**: Copy the badge from the top of README.md and update `OWNER`
- **Set up notifications**: Configure GitHub notifications for failed builds
- **Add security scanning**: See the workflow README for Trivy integration
- **Create release workflow**: Automate GitHub releases when tags are pushed

## Support

For issues or questions:
1. Check the [workflow README](.github/workflows/README.md) for detailed documentation
2. Review build logs in the Actions tab
3. Open an issue in the repository

## Quick Reference

```bash
# Push to main → builds main, main-<sha>, latest tags
git push origin main

# Push to develop → builds develop, develop-<sha> tags
git push origin develop

# Create release → builds versioned tags
git tag v1.0.0 && git push origin v1.0.0

# Pull image
docker pull ghcr.io/OWNER/robohub-inventory:TAG
```

Replace:
- `OWNER` with your GitHub username or organization
- `TAG` with the desired tag (e.g., `latest`, `1.0.0`, `main`)
