# deploytag

Zero-config deploy previews from git tags.

## Install

```bash
git clone https://github.com/openkickstartai/deploytag.git
cd deploytag && go build -o deploytag .
```

## Usage

```bash
# Deploy a preview from current branch
deploytag preview --ttl 24h

# Deploy from a specific tag
deploytag deploy v1.2.3-preview

# List active previews
deploytag list

# Clean up expired previews
deploytag cleanup
```

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

## Config (.deploytag.yml)

```yaml
provider: docker  # docker | static | serverless
port: 8080
build: docker build -t app .
run: docker run -p 8080:8080 app
ttl: 24h
```

## Testing

```bash
go test -v ./...
```
