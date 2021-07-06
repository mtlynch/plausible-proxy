# rss-proxy

[![CircleCI](https://circleci.com/gh/mtlynch/rss-proxy.svg?style=svg)](https://circleci.com/gh/mtlynch/rss-proxy)

## Testing

```bash
TESTER="$(mktemp -d)/cli"
go build -o "${TESTER}" ./dev-scripts/cli.go
"${TESTER}"
```
