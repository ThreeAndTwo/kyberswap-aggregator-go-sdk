## kyberswap-aggregator-go-sdk


[kyberswap aggregator](https://docs.kyberswap.com/kyberswap-solutions/kyberswap-aggregator/aggregator-api-specification/evm-swaps)

### Usage

```go
agg := NewAggregator("https://aggregator-api.kyberswap.com")

params: &QuerySwapTokenParams{
    Network:           "ethereum",
    TokenIn:           "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
    TokenOut:          "0xcd5fe23c85820f7b72d0926fc9b05b43e359b7ee",
    AmountIn:          "1000000000000000",
    To:                "0x888888888889758f76e7103c6cbf23abbf58f946",
    SlippageTolerance: 10,
    UseMeta:           false,
    SaveGas:           1,
    GasInclude:        1,
    ClientData:        "%7B%22source%22:%22Pendle%22%7D",
    Deadline:          2147483647,
    ExcludedSources:   "",
}

// get swap info call data
got, _ := agg.GetSwapInfo(tt.args.params)
fmt.Println(got)
```


