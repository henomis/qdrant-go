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

	wait := true
	resp := &response.PointUpsert{}
	err := client.PointUpsert(
		context.Background(),
		&request.PointUpsert{
			CollectionName: "test",
			Wait:           &wait,
			Points: []request.Point{
				{
					ID:     "45b07125-f592-414f-a9d0-160c8ecc283a",
					Vector: []float64{1.1, 2.2, 3.3, 4.4},
					Payload: map[string]interface{}{
						"key": "value",
					},
				},
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
