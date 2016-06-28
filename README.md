# CF Default App Go

The default Go app that will be pushed into the Swisscom Application cloud if no source code is provided.

## Run locally

1. Install [Go](https://golang.org/doc/install)
1. Run `go run server.go`

## Run in the cloud

1. Install the [cf CLI](https://github.com/cloudfoundry/cli#downloads)
1. Run `cf push --random-route`
1. Visit the given URL

## Create ZIP for Go

1. Run `zip -r go_app.zip Godeps public templates Procfile server.go`

## Create ZIP for Binary

1. Install [Go](https://golang.org/doc/install)
1. Change texts in `templates/index.html`
1. Change Procfile to `web: ./cf-default-app-go`
1. Run `GOARCH=amd64 GOOS=linux go build .`
1. Run `zip -r binary_app.zip public templates cf-default-app-go Procfile`
