package main

import (
	"fmt"
	"log"

	"coincapclient/internal/coincap"
)

func main() {
	client := coincap.NewMockClient()

	assets, err := client.GetAssets(5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[")
	for _, asset := range assets {
		fmt.Printf("  {\n    name: %s,\n    priceUsd: %s,\n  },\n",
			asset.Name,
			asset.PriceUsd,
		)
	}
	fmt.Println("]")
}

// func main() {
// 	apiKey := os.Getenv("COINCAP_API_KEY")
// 	if apiKey == "" {
// 		log.Fatal("COINCAP_API_KEY is not set")
// 	}

// 	client := coincap.NewClient(apiKey)

// 	assets, err := client.GetAssets(10)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("[")
// 	for _, asset := range assets {
// 		fmt.Printf("  {\n    name: %s,\n    priceUsd: %s,\n  },\n",
// 			asset.Name,
// 			asset.PriceUsd,
// 		)
// 	}
// 	fmt.Println("]")
// }
