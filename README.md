# Golang service boilerplate

## Running

```bash
go run cmd/main.go config/config.yaml
```

Default environment is `development`. To run in `production`:

```bash
ENVIRONMENT=production go run cmd/main.go config/config.yaml
```