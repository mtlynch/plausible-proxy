# plausible-proxy

[![CircleCI](https://circleci.com/gh/mtlynch/plausible-proxy.svg?style=svg)](https://circleci.com/gh/mtlynch/plausible-proxy)

A reverse proxy for Plausible Analytics, implemented in Google Cloud Functions.

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
FUNCTION_NAME="ProxyPlausible"

gcloud functions deploy \
  "${FUNCTION_NAME}" \
  --runtime "${GO_RUNTIME}" \
  --memory 128MB \
  --source proxy \
  --trigger-http \
  --allow-unauthenticated
```

## Usage

After you deploy the proxy, check the Cloud Functions output or dashbord for the serving URL. It will look something like this:

* <https://YOUR-PROJECT.cloudfunctions.net/ProxyPlausible>

### Without redirects

Once you deploy your proxy, you can refer to it in your Plausible JS snippet as follows:

```html
<script
  defer
  data-api="https://YOUR-PROJECT.cloudfunctions.net/ProxyPlausible/api/event"
  data-domain="YOUR-DOMAIN"
  src="https://YOUR-PROJECT.cloudfunctions.net/ProxyPlausible/js/script.js">
</script>
```

### With Firebase rewrites

If you're using Firebase, the cleaner way to access your function is to use a Firebase rewrite to refer to the proxy so that :

```json
"rewrites": [
  {
    "source": "/plaus-proxy/**",
    "function": "ProxyPlausible"
  }
]
```

Then your `<script>` include simplifies to the following:

```html
<script
  defer
  data-api="/plaus-proxy/api/event"
  data-domain="YOUR-DOMAIN"
  src="/plaus-proxy/js/script.js">
</script>
```
