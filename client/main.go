package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1/emoine_rv1connect"
)

func main() {
	client := emoine_rv1connect.NewGeneralAPIServiceClient(
		http.DefaultClient,
		"http://localhost:8090",
	)

	res, err := client.GetEvent(context.Background(), connect.NewRequest(&emoine_rv1.GetEventRequest{
		Id: "hoge",
	}))
	log.Println(res, err)
}
