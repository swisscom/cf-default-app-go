# CF Default App Go

The default Go app that will be pushed into the Swisscom Application cloud if no source code is provided.

## Run locally

1. Install [Go](https://golang.org/doc/install)
1. Run `go run main.go`
1. Visit <http://localhost:8080>

## Run in the cloud

1. Install the [cf CLI](https://github.com/cloudfoundry/cli#downloads)
1. Run `cf push --random-route`
1. Visit the given URL

## Create ZIP for Go buildpack

1. Run `zip -r go_app.zip public templates glide.lock glide.yaml main.go`

## Create ZIP for Binary buildpack

1. Install [Go](https://golang.org/doc/install)
1. Change texts in `templates/index.html`
1. Run `GOARCH=amd64 GOOS=linux go build .`
1. Run `zip -r binary_app.zip public templates cf-default-app-go Procfile`
