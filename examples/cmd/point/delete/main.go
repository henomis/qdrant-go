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
	resp := &response.PointsDelete{}
	err := client.PointsDelete(
		context.Background(),
		&request.PointsDelete{
			CollectionName: "test",
			Wait:           &wait,
			Filter: request.Filter{
				Must: []request.M{
					{
						"key": "key",
						"match": request.M{
							"value": "value",
						},
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
