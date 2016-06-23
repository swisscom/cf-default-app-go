# CF Default App Go

The default Go app that will be pushed into the Swisscom Application cloud if no source code is provided.

## Run locally

1. Install [Go](https://golang.org/doc/install)
1. Run `go run main.go`

## Run in the cloud

1. Install the [cf CLI](https://github.com/cloudfoundry/cli#downloads)
1. Run `cf push my-go-app --random-route`
1. Visit the given URL
