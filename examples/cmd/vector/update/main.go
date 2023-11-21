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
	resp := &response.VectorsUpdate{}
	err := client.VectorsUpdate(
		context.Background(),
		&request.VectorsUpdate{
			CollectionName: "test",
			Wait:           &wait,
			Points: []request.PointVectors{
				{
					ID:     "45b07125-f592-414f-a9d0-160c8ecc283a",
					Vector: []float64{5, 6, 7, 8},
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
