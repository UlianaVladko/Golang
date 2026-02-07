package coincap

// type AssetResponse struct {
// 	Data []Asset `json:"data"`
// }

type Asset struct {
	Name     string `json:"name"`
	PriceUsd string `json:"priceUsd"`
}
