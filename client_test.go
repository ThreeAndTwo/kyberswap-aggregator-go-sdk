package kyberswap_aggregator_go_sdk

import (
	"testing"
)

var agg *Aggregator

func init() {
	agg = NewAggregator("")
}

func TestAggregator_GetSwapInfo(t *testing.T) {
	type args struct {
		params *QuerySwapTokenParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test kyberswap aggregator",
			args: args{
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
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := agg.GetSwapInfo(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSwapInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%v", got)
		})
	}
}
