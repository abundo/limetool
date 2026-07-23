# Limetool

A Go library and CLI for communicating with the [Lime CRM](https://www.lime.se/) REST API.

It fetches companies and their deliveries (with embedded product, service,
agreement and delivery point data) and prints them as JSON.

## Usage

Configure the API URL and key in `/etc/limetool/limetool.yaml` (or pass
`-f <path>` to use another location):

```yaml
lime:
  apiurl: https://your-instance.lime-crm.com/api/v1
  apikey: your-api-key
```

Then run:

    limetool get                    # fetch all companies
    limetool get -c "Acme AB"       # fetch one or more companies (repeatable -c)
    limetool get -r                 # bypass the local cache and refresh from Lime
    limetool show-config            # print the resolved configuration

Global flags: `-d` for debug logging, `--loglevel` (error|warn|info|debug).

Responses are cached under `/tmp/lime-<hash>` to avoid refetching unchanged
data; use `-r`/`--refresh` to force a live request.

## Development

See [DEV.md](DEV.md).
