package coincap

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c *MockClient) GetAssets(limit int) ([]Asset, error) {
	mockData := []Asset{
		{Name: "Bitcoin", PriceUsd: "57435.2862859343831301"},
		{Name: "Ethereum", PriceUsd: "3686.1846963999934439"},
		{Name: "Tether", PriceUsd: "1.000384928392"},
		{Name: "Cardano", PriceUsd: "1.23"},
		{Name: "Solana", PriceUsd: "160.45"},
	}

	if limit > len(mockData) {
		limit = len(mockData)
	}

	return mockData[:limit], nil
}

// const baseURL = "https://api.coincap.io/v2"

// type Client struct {
// 	apiKey string
// 	http   *http.Client
// }

// func NewClient(apiKey string) *Client {
// 	return &Client{
// 		apiKey: apiKey,
// 		http: &http.Client{
// 			Timeout: 10 * time.Second,
// 		},
// 	}
// }

// func (c *Client) GetAssets(limit int) ([]Asset, error) {
// 	req, err := http.NewRequest(
// 		http.MethodGet,
// 		fmt.Sprintf("%s/assets?limit=%d", baseURL, limit),
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+c.apiKey)
// 	req.Header.Set("Accept", "application/json")

// 	resp, err := c.http.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
// 	}

// 	var result AssetResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return nil, err
// 	}

// 	return result.Data, nil
// }
