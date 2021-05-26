package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rancher/norman/clientbase"
	"github.com/rancher/norman/types"
	client "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print(usage())
		return
	}

	serverURL := os.Args[1]
	// The server URL needs to have /v3 on the end for the API endpoint
	if !strings.HasSuffix(serverURL, "/v3") {
		serverURL = serverURL + "/v3"
	}

	client, err := client.NewClient(&clientbase.ClientOpts{
		URL:      serverURL,
		TokenKey: os.Args[2],
	})
	if err != nil {
		log.Fatal(err)
	}

	clusters, err := client.Cluster.List(&types.ListOpts{})
	if err != nil {
		log.Fatal(err)
	}

	for _, cluster := range clusters.Data {
		log.Println(cluster.Name)
	}

}

func usage() string {
	return `Rancher client demo
Two args are required - Rancher URL and a token to connect to rancher
`
}
