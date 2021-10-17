# plausible-proxy

[![CircleCI](https://circleci.com/gh/mtlynch/plausible-proxy.svg?style=svg)](https://circleci.com/gh/mtlynch/plausible-proxy)

A simple HTTP proxy meant for proxying plausible feeds for podcasts. Host this function on your own domain and set your podcast host's plausible feed as the URL. That way, you use a custom domain and can change podcast hosts without it ever affecting your listeners.

## Testing

```bash
PORT=8080 \
  go run dev-scripts/cli.go
```

## Deployment

### With CircleCI

To deploy to Google Cloud Platform, replace the ENV variables in the CircleCI config (`.circleci/config.yml`) with your own project's details.

### With gcloud CLI tool

```bash
GO_RUNTIME="go113"

FUNCTION_NAME="ProxyPlausibleEvent"
gcloud functions deploy \
  "${FUNCTION_NAME}" \
  --runtime "${GO_RUNTIME}" \
  --source proxy \
  --trigger-http \
  --allow-unauthenticated

FUNCTION_NAME="ProxyPlausibleScript"
gcloud functions deploy \
  "${FUNCTION_NAME}" \
  --runtime "${GO_RUNTIME}" \
  --source proxy \
  --trigger-http \
  --allow-unauthenticated
```
