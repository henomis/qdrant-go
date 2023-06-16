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

	withPayload := true
	resp := &response.PointSearch{}
	err := client.PointSearch(
		context.Background(),
		&request.PointSearch{
			CollectionName: "test",
			Vector:         []float64{1.1, 2.2, 3.3, 4.4},
			Limit:          10,
			WithPayload:    &withPayload,
			WithVector:     &withPayload,
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
