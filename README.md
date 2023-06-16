# Unofficial Qdrant Go SDK


[![GoDoc](https://godoc.org/github.com/henomis/qdrant-go?status.svg)](https://godoc.org/github.com/henomis/qdrant-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/qdrant-go)](https://goreportcard.com/report/github.com/henomis/qdrant-go) [![GitHub release](https://img.shields.io/github/release/henomis/qdrant-go.svg)](https://github.com/henomis/qdrant-go/releases)

This is [Qdrant](https://qdrant.tech/)'s **unofficial** Go client, designed to enable you to use Qdrant's services easily from your own applications.

## Qdrant

[Qdrant](https://qdrant.tech/) is a vector database that allows you to build high-performance vector search applications.

## API support

### collections
- ✅ list
- ✅ create
- ✅ collect info
- ❌ update
- ❌ delete
- ❌ update aliases
- ❌ create index
- ❌ delete index
- ❌ cluster info
- ❌ update cluster setup
- ❌ list aliases
- ❌ recover from uploaded snapshot
- ❌ recover from snapshot
- ❌ create snapshot
- ❌ delete snapshot
- ❌ download snapshot


### points 
- ❌ get point
- ❌ get points
- ✅ upsert points
- ❌ delete points
- ❌ update vectors
- ❌ delete vectors
- ❌ set payload
- ❌ overwrite payload
- ❌ delete payload
- ❌ clear payload
- ❌ scroll payload
- ✅ search points
- ❌ search batch points
- ❌ search point groups
- ❌ recommend points
- ❌ recommend batch points
- ❌ recommend point groups
- ❌ count points

### cluster
- ❌ cluster status info
- ❌ tries to recover current peer Raft state
- ❌ remove peer
- ❌ collection cluster info
- ❌ update collection cluster setup

### snapshots
- ❌ recover from uploaded snapshot
- ❌ recover from snapshot
- ❌ list collection snapshots
- ❌ create collection snapshot
- ❌ delete collection snapshot
- ❌ download collection snapshot
- ❌ list storage snapshots
- ❌ create storage snapshot
- ❌ delete storage snapshot
- ❌ download storage snapshot

### service
- ❌ collect telemetry data
- ❌ collect Prometheus metrics data
- ❌ set lock options
- ❌ get lock options


## Getting started

### Installation

You can load qdrant-go into your project by using:
```
go get github.com/henomis/qdrant-go
```


### Configuration

The only thing you need to start using Qdrant's APIs is the API key. Copy and paste it in the corresponding place in the code, select the API and the parameters you want to use, and that's it.


### Usage

Please refer to the [examples folder](examples/cmd/) to see how to use the SDK.

Here below a simple usage example:

```go
package main

import (
	"context"
	"fmt"

	qdrantgo "github.com/henomis/qdrant-go"
	"github.com/henomis/qdrant-go/request"
	"github.com/henomis/qdrant-go/response"
)

func main() {

	client := qdrantgo.New("http://localhost:6333", "secret-api-key")

	onDisk := true
	resp := &response.CollectionCreate{}
	err := client.CollectionCreate(
		context.Background(),
		&request.CollectionCreate{
			CollectionName: "test",
			Vectors: request.VectorsParams{
				Size:     1536,
				Distance: request.DistanceCosine,
				OnDisk:   &onDisk,
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
```

## Who uses qdrant-go?

* [LinGoose](https://github.com/henomis/lingoose) Go framework for building awesome LLM apps