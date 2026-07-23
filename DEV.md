# Limetool development

Assumes the code is stored in `~/code/limetool`.

## Checkout the source code

    cd code
    git clone https://github.com/abundo/limetool

## Install dependencies

    go mod tidy

## Build binaries

    make

Builds `build/limetool` from `cmd/limetool_cli.go`.

## Install binary

Default install location is `/usr/bin`:

    sudo make install

## Project layout

- `limetool.go` — library: `Lime` client, `GetCompanies`/`GetDeliveries`
  against the Lime CRM REST API.
- `models/` — API response and domain types (`LimeCompany`, `LimeDelivery`,
  `LimeAgreement`, `LimeProduct`, `LimeService`, `LimeDeliverypoint`, ...).
- `cmd/limetool_cli.go` — CLI built with [boa](https://github.com/GiGurra/boa)
  and [cobra](https://github.com/spf13/cobra); subcommands `get` and
  `show-config`.

## Configuration

The CLI reads a YAML config file, default `/etc/limetool/limetool.yaml`,
overridable with `-f`:

```yaml
lime:
  apiurl: https://your-instance.lime-crm.com/api/v1
  apikey: your-api-key
```

## Caching

`limeAPI()` in `limetool.go` caches raw API responses on disk at
`/tmp/lime-<sha256(url)>` so repeated runs don't re-hit the Lime API. Pass
`-r`/`--refresh` on the CLI (or `refresh=true` when calling the lib) to
bypass the cache and force a live fetch.

## Notes

- `go.mod` targets Go 1.25.
- No test suite yet.
