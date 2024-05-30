package kyberswap_aggregator_go_sdk

type QuerySwapTokenParams struct {
	Network           string
	TokenIn           string
	TokenOut          string
	AmountIn          string
	To                string
	SlippageTolerance int
	UseMeta           bool
	SaveGas           int
	GasInclude        int
	ClientData        string
	Deadline          int64
	ExcludedSources   string
}

type QuerySwapResult struct {
	InputAmount     string               `json:"inputAmount"`
	OutputAmount    string               `json:"outputAmount"`
	TotalGas        int                  `json:"totalGas"`
	GasPriceGwei    string               `json:"gasPriceGwei"`
	GasUsd          float64              `json:"gasUsd"`
	AmountInUsd     float64              `json:"amountInUsd"`
	AmountOutUsd    float64              `json:"amountOutUsd"`
	ReceivedUsd     float64              `json:"receivedUsd"`
	Swaps           [][]QuerySwapData    `json:"swaps"`
	Tokens          map[string]TokenInfo `json:"tokens"`
	EncodedSwapData string               `json:"encodedSwapData"`
	RouterAddress   string               `json:"routerAddress"`
}

type QuerySwapData struct {
	Pool              string        `json:"pool"`
	TokenIn           string        `json:"tokenIn"`
	TokenOut          string        `json:"tokenOut"`
	LimitReturnAmount string        `json:"limitReturnAmount"`
	SwapAmount        string        `json:"swapAmount"`
	AmountOut         string        `json:"amountOut"`
	Exchange          string        `json:"exchange"`
	PoolLength        int           `json:"poolLength"`
	PoolType          string        `json:"poolType"`
	PoolExtra         SwapPoolExtra `json:"poolExtra"`
	Extra             interface{}   `json:"extra"`
	MaxPrice          string        `json:"maxPrice"`
}

type SwapPoolExtra struct {
	TokenInIndex     int  `json:"tokenInIndex"`
	TokenOutIndex    int  `json:"tokenOutIndex"`
	Underlying       bool `json:"underlying"`
	TokenInIsNative  bool `json:"TokenInIsNative"`
	TokenOutIsNative bool `json:"TokenOutIsNative"`
}

type TokenInfo struct {
	Address  string  `json:"address"`
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Decimals int     `json:"decimals"`
}
