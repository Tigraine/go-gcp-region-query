# go-gcp-region-query

Retrieves the Google GCP Region name where your VM is currently running in.

Installation:

```
go get github.com/tigraine/go-gcp-region-query
```

Usage:

```
import "github.com/tigraine/go-gcp-region-query"

region, err := gcpregionquery.GetLocalRegion()
if err != nil {
  // You are either not running in a GCP Datacenter or there was a problem calling the API
}
fmt.Println(region) // Example: europe-west1-c
```
