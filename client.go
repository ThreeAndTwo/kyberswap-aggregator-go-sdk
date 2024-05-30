package kyberswap_aggregator_go_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const defaultKyberswapHost = "https://aggregator-api.kyberswap.com"

type Aggregator struct {
	host   string
	client *resty.Client
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05.000"})
}

func NewAggregator(host string) *Aggregator {
	if host == "" {
		host = defaultKyberswapHost
	}
	client := resty.New()
	return &Aggregator{host: host, client: client}
}

// GetSwapInfo
// https://docs.kyberswap.com/kyberswap-solutions/kyberswap-aggregator/aggregator-api-specification/evm-swaps#legacy
func (agg *Aggregator) GetSwapInfo(params *QuerySwapTokenParams) (*QuerySwapResult, error) {
	url := fmt.Sprintf("%s/%s/route/encode?tokenIn=%s&tokenOut=%s&amountIn=%s&to=%s&slippageTolerance=%d&useMeta=%t&saveGas=%d&gasInclude=%d&clientData=%s&deadline=%d&excludedSources=%s",
		agg.host, params.Network, params.TokenIn, params.TokenOut, params.AmountIn, params.To, params.SlippageTolerance,
		params.UseMeta, params.SaveGas, params.GasInclude, params.ClientData, params.Deadline, params.ExcludedSources)

	resp, err := agg.client.R().Get(url)
	if err != nil {
		log.Error().Msgf("get swap data error: %s", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		msg := fmt.Sprintf("status code %d, request kyberswap server error: %v", resp.StatusCode(), resp.Error())
		log.Error().Msgf(msg)
		return nil, fmt.Errorf(msg)
	}

	querySwapRes := new(QuerySwapResult)
	err = json.Unmarshal(resp.Body(), querySwapRes)
	if err != nil {
		log.Error().Msgf("parse unmarshal error: %s", err)
		return nil, err
	}
	return querySwapRes, nil
}
