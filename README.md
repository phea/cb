# cb (beta)

## Introduction

cb is a Go client library to interact with the [CrunchBase API v2](https://developer.crunchbase.com/).

## Installation

```sh
go get github.com/phea/cb
```

### To Do
* Timeout option for requests
* More API endpoints
* Proper testing
* Handling errors

### Usage

```go
package main

import "github.com/phea/cb"

// Create a client with your CrunchBase API Key.
client := cb.NewClient("USER_KEY")

// Single resource
org, err := client.GetOrganization("name")

// List resource
// Set any params
params := url.Values{}
params.Set("order", "created_at ASC")
params.Set("page", "1")

query, err := client.GetOrganizations(&params)
for org := range query.List {
  // Do something...
}

// Next and Prev convenience functions for list resources.
query, err = query.Next()
```

### Contribution
Pull requests and feedback are welcome.

Twitter: [@phea](http://twitter.com/phea)
