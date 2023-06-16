package main

import (
	"context"
	"fmt"

	qdrantgo "github.com/henomis/qdrant-go"
	"github.com/henomis/qdrant-go/request"
	"github.com/henomis/qdrant-go/response"
)

func main() {

	client := qdrantgo.New("http://localhost:6333", "")

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
