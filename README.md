# vanityprox

[![build](https://github.com/gleich/vanityprox/actions/workflows/build.yml/badge.svg)](https://github.com/gleich/vanityprox/actions/workflows/build.yml)
[![deploy](https://github.com/gleich/vanityprox/actions/workflows/deploy.yml/badge.svg)](https://github.com/gleich/vanityprox/actions/workflows/deploy.yml)
[![lint](https://github.com/gleich/vanityprox/actions/workflows/lint.yml/badge.svg)](https://github.com/gleich/vanityprox/actions/workflows/lint.yml)
[![go report card](https://goreportcard.com/badge/go.mattglei.ch/vanityprox)](https://goreportcard.com/report/go.mattglei.ch/vanityprox)

A simple webserver that allows for custom vanity URLs for go modules. Allows for go modules to be installed like `go get go.mattglei.ch/timber` instead of `go get github.com/gleich/timber`.

## How to use

Want to use vanityprox for your own domain and GitHub account/organization? Simply run the following docker command to get started:

```bash
docker run -p 8000:8000 ghcr.io/gleich/vanityprox
```

This will then start the server on port `8000`. It's that easy!

Vanityprox is configured with a TOML file called [vanityprox.toml](./vanityprox.toml). Below is details on what all of the configuration options mean.

### Configuration

```toml
host = "https://go.mattglei.ch"
source_prefix = "https://github.com/gleich"
favicon = "https://mattglei.ch/favicon.ico"

packages = [
    "timber",
    "lcp",
    "vanityprox",
    "ritcs",
]

[logs]
timezone = "America/New_York"
time_format = "01/02 03:04:05 PM MST"
```

- `host`
  - Host vanity URL that the server will get requests from.
  - Example: `https://go.mattglei.ch`
  - **REQUIRED**
- `source_prefix`
  - Prefix for where all of the code for the go modules is stored. If you're using GitHub it should simply just be the URL of the GitHub account/organization.
  - Example: `https://github.com/gleich`
  - **REQUIRED**
- `packages`
  - All of the packages that vanityprox should handle
  - Example: `["timber", "lcp"]`
  - **REQUIRED**
- `root_redirect`
  - If a request is made to GET / where should the server send the user? Set this environment variable to have them be redirected (via a 308 permanent redirect) to the given URL. If no URL is supplied then it simply will just return a 404 error when something hits GET /.
  - Example: `https://github.com/gleich/vanityprox`
  - **OPTIONAL**
- `favicon`
  - URL for a favicon. This will be displayed in the HTML that gets returned from the server.
  - **OPTIONAL**
- `logs.timezone`
  - Timezone that logs should be outputted in. Default is UTC. Select from the [Wikipedia list of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
  - Example: `America/New_York`
  - **OPTIONAL**
- `logs.time_format`
  - Format that should be used for outputting the logs. Uses [golang's built in time formatting system](https://go.dev/src/time/format.go).
  - Example: `01/02 03:04:05 PM MST`
  - **OPTIONAL**
