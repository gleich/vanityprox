# vanityprox

[![build](https://github.com/gleich/vanityprox/actions/workflows/build.yml/badge.svg)](https://github.com/gleich/vanityprox/actions/workflows/build.yml)
[![deploy](https://github.com/gleich/vanityprox/actions/workflows/deploy.yml/badge.svg)](https://github.com/gleich/vanityprox/actions/workflows/deploy.yml)
[![lint](https://github.com/gleich/vanityprox/actions/workflows/lint.yml/badge.svg)](https://github.com/gleich/vanityprox/actions/workflows/lint.yml)
[![go report card](https://goreportcard.com/badge/pkg.mattglei.ch/vanityprox)](https://goreportcard.com/report/pkg.mattglei.ch/vanityprox)

A simple webserver that allows for custom vanity URLs for go modules. Allows for go modules to be installed like `go get pkg.mattglei.ch/timber` instead of `go get github.com/gleich/timber`.

## How to use

Want to use vanityprox for your own domain and GitHub account/organization? Simply run the following docker command to get started:

```bash
docker run -p 8000:8000 \
  -e VANITYPROX_HOST="https://pkg.mattglei.ch" \
  -e VANITYPROX_SOURCE_PREFIX="https://github.com/gleich" \
  ghcr.io/gleich/vanityprox
```

This will then start the server on port `8000`. Its that easy!

As you can see vanityprox is configured with environment variables. Below is details on what all of the environment variables mean.

### Configuration

- `VANITYPROX_HOST`
  - Host vanity URL that the server will get requests from.
  - Example: `https://pkg.matglei.ch`
  - **REQUIRED**
- `VANITYPROX_SOURCE_PREFIX`
  - Prefix for where all of the code for the go modules is stored. If you're using GitHub it should simply just be the name of the GitHub account/organization.
  - Example: `https://github.com/gleich`
  - **REQUIRED**
- `VANITYPROX_ROOT_REDIRECT`
  - If a request is made to GET / where should the server send the user? Set this environment variable to have them be redirected (via a 301 permanent redirect) to the given URL. If no URL is supplied then it simply will just return a 404 error when something hits GET /.
  - Example: `https://github.com/gleich/vanityprox`
  - **OPTIONAL**
- `VANITYPROX_LOG_TIMEZONE`
  - Timezone that logs should be outputted in. Default is UTC. Select from the [Wikipedia list of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
  - Example: `America/New_York`
  - **OPTIONAL**
- `VANITYPROX_LOG_TIME_FORMAT`
  - Format that should be used for outputting the logs. Used [golang's built in time formatting system](https://go.dev/src/time/format.go).
  - Example: `01/02 03:04:05 PM MST`
  - **OPTIONAL**
- `VANITYPROX_FAVICON`
  - URL for a favicon of your choosing. This will be displayed in the HTML that gets returned from the server.
  - **OPTIONAL**
