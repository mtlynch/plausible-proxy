# rss-proxy

[![CircleCI](https://circleci.com/gh/mtlynch/rss-proxy.svg?style=svg)](https://circleci.com/gh/mtlynch/rss-proxy)

A simple HTTP proxy meant for proxying RSS feeds for podcasts. Host this function on your own domain and set your podcast host's RSS feed as the URL. That way, you use a custom domain and can change podcast hosts without it ever affecting your listeners.

## Testing

```bash
TARGET_URL=https://anchor.fm/s/someID/podcast/rss PORT=8080 \
  go run dev-scripts/cli.go
```

## Deployment

To deploy to Google Cloud Platform, configure a service account and use the CircleCI config (`.circleci/config.yml`) or just manually save your target URL to a file called `.env.yaml` like this:

```yaml
# .env.yaml
TARGET_URL: 'https://anchor.fm/s/123abc/podcast/rss' # Replace with your podcast's RSS feed
```

And deploy with:

```bash
FUNCTION_NAME="ProxyRequest"
GO_RUNTIME="go113"

gcloud functions deploy \
  "${FUNCTION_NAME}" \
  --runtime "${GO_RUNTIME}" \
  --env-vars-file .env.yaml \
  --source proxy \
  --trigger-http \
  --allow-unauthenticated
```
