package main

import (
	"github.com/girishramnani-crest/ndc-elasticsearch/types"
	"github.com/hasura/ndc-sdk-go/connector"
)

// Start the connector server at http://localhost:8080
//
//	go run . serve
//
// See [NDC Go SDK] for more information.
//
// [NDC Go SDK]: https://github.com/hasura/ndc-sdk-go
func main() {
	if err := connector.Start[types.Configuration, types.State](
		&Connector{},
		connector.WithMetricsPrefix("ndc-elasticsearch"),
		connector.WithDefaultServiceName("ndc-elasticsearch"),
	); err != nil {
		panic(err)
	}
}
