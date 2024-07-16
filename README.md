# Block Explorer Audit job

Audit block, transaction, log, internal tx base on hash and checksum

## Requirements

- Postgres 10+

# Setup Instructions
1. Export ENV or Add to .env
```
export APP_ADDR=0.0.0.0:8080
export PG_HOST=127.0.0.1
export PG_PORT=5432
export PG_USER=postgres
export PG_PASS=postgres
export PG_DB=explorer

```

4. Build & start audit job
`go build cmd/audit/main`
`./main`
