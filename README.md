# deploytag

Zero-config deploy previews from git tags.

Push a tag → get a preview URL. Supports Docker, static sites, and serverless providers.

## Install

```bash
git clone https://github.com/openkickstartai/deploytag.git
cd deploytag && go build -o deploytag .
```

## Quick Start

```bash
# 1. Create a .deploytag.yml (optional — sensible defaults apply)
# 2. Deploy a preview from the current branch
deploytag preview --ttl 24h

# 3. Or deploy a specific tag
deploytag deploy v1.2.3-preview
```

## CLI Reference

| Command | Description | Example |
|---------|-------------|---------|
| `preview` | Create a preview from the current branch | `deploytag preview --ttl 12h` |
| `deploy <tag>` | Deploy a specific git tag | `deploytag deploy v1.2.3-preview` |
| `list` | List all active preview deployments | `deploytag list` |
| `cleanup` | Remove expired previews | `deploytag cleanup` |
| `version` | Print version | `deploytag version` |
| `help` | Show usage information | `deploytag --help` |

### Global Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--ttl <duration>` | `24h` | How long the preview stays alive (e.g. `30m`, `12h`, `7d`) |
| `--help`, `-h` | — | Show help message |

### Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | Error (invalid args, build failure, unknown command, etc.) |

## Configuration (`.deploytag.yml`)

Place a `.deploytag.yml` in your repository root. Every field is optional.

```yaml
# Provider determines the deploy strategy
# Options: docker | static | serverless
provider: docker

# Port exposed by the preview (used for health-check & URL generation)
port: 8080

# Shell command executed during the build phase
build: docker build -t app .

# Shell command executed to start the preview
run: docker run -p 8080:8080 app

# Time-to-live — preview is automatically cleaned up after this duration
ttl: 24h
```

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `provider` | `string` | `docker` | Deploy backend: `docker`, `static`, or `serverless` |
| `port` | `int` | `8080` | Port the application listens on |
| `build` | `string` | `""` | Build command (skipped if empty) |
| `run` | `string` | `""` | Run/start command |
| `ttl` | `string` | `24h` | Auto-expiry duration (`30m`, `12h`, `7d`, …) |

## GitHub Action

```yaml
on:
  push:
    tags: ['*-preview']
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: openkickstartai/deploytag-action@v1
        with:
          ttl: 48h
```

## Testing

```bash
go test -v ./...
```

## License

MIT
