# Getting Started with deploytag

## Prerequisites

- Go 1.22+
- Docker (for container deployments)
- Git

## Quick Start

### 1. Install

```bash
git clone https://github.com/openkickstartai/deploytag.git
cd deploytag && go build -o deploytag .
```

### 2. Create a config

```bash
cat > .deploytag.yml << EOF
provider: docker
port: 8080
build: docker build -t myapp .
run: docker run -p 8080:8080 myapp
ttl: 24h
EOF
```

### 3. Deploy a preview

```bash
./deploytag preview --ttl 48h
```

You'll get a URL like `https://a1b2c3d4.preview.deploytag.dev`

### 4. Share with your team

Send the URL to QA, PM, or stakeholders. The preview auto-expires after the TTL.

## Architecture

```
git tag -> deploytag detect -> build image -> provision URL -> serve traffic -> TTL cleanup
```

1. **Detection**: Watches for git tags matching `*-preview` pattern
2. **Build**: Runs your configured build command (Docker, static, etc.)
3. **Provision**: Allocates a unique subdomain with HTTPS
4. **Serve**: Routes traffic to your container/files
5. **Cleanup**: Automatically removes expired previews
