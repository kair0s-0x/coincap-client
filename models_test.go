package coincap

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestAsset(t *testing.T) {
	var data AssetData
	err := unmarshalModel("asset_id", &data)
	asset := data.Asset
	require.NoError(t, err)
	require.Equal(t, int64(1627299055657), data.Timestamp)
	require.Equal(t, "polkadot", asset.Id)
	require.Equal(t, 9, asset.Rank)
	require.Equal(t, "DOT", asset.Symbol)
	require.Equal(t, "Polkadot", asset.Name)
	require.Equal(t, 1013089106.2001500000000000, asset.Supply)
	require.Nil(t, asset.MaxSupply)
	require.Equal(t, 14964432257.0276593916630207, asset.MarketCapUsd)
	require.Equal(t, 567049854.0292255884810806, asset.VolumeUsd24Hr)
	require.Equal(t, 14.7710918668897673, asset.PriceUsd)
	require.Equal(t, 10.0829076670949820, asset.ChangePercent24Hr)
	require.Equal(t, 14.3158259816882364, asset.Vwap24Hr)
	require.Equal(t, "https://polkascan.io/polkadot", *asset.Explorer)
}

func TestAssets(t *testing.T) {
	var data AssetsData
	err := unmarshalModel("assets", &data)
	require.NoError(t, err)
	require.Equal(t, 5, len(data.Data))
}

func TestAssetHistory(t *testing.T) {
	var data AssetHistoriesData
	err := unmarshalModel("asset_history", &data)
	require.NoError(t, err)
	require.Equal(t, 5, len(data.Data))
	require.Equal(t, 1013089106.25717, data.Data[0].CirculatingSupply)
	date := time.Date(2021, 7, 26, 11, 0, 0, 0, time.UTC)
	for i, v := range data.Data {
		require.Equal(t, date.Add(time.Minute*time.Duration(30*i)), v.Date)
	}
}

func TestAssetMarkets(t *testing.T) {
	var data AssetMarketsData
	err := unmarshalModel("asset_markets", &data)
	require.NoError(t, err)
	require.Equal(t, 7, len(data.Data))
}

func TestRates(t *testing.T) {
	var data RatesData
	err := unmarshalModel("assets", &data)
	require.NoError(t, err)
	require.Equal(t, 5, len(data.Data))
}

func TestRatesID(t *testing.T) {
	var data RateData
	err := unmarshalModel("rates_id", &data)
	require.NoError(t, err)
	require.Equal(t, "USDC", data.Data.Symbol)
	require.Equal(t, "usd-coin", data.Data.Id)
	require.Equal(t, 1.0000000000000000, data.Data.RateUsd)
	var nilStringPointer *string = nil
	require.Equal(t, nilStringPointer, data.Data.CurrencySymbol)
	require.Equal(t, "crypto", data.Data.Type)
}

func TestExchanges(t *testing.T) {
	var data ExchangesData
	err := unmarshalModel("exchanges", &data)
	require.NoError(t, err)
	require.Equal(t, 5, len(data.Data))
	require.Equal(t, "binance", data.Data[0].ExchangeId)
}

func TestExchange(t *testing.T) {
	var data ExchangeData
	err := unmarshalModel("exchange", &data)
	require.NoError(t, err)
	require.Equal(t, "kraken", data.Data.ExchangeId)
	require.Equal(t, 1.845246588138155764, data.Data.PercentTotalVolume)
	require.Equal(t, 1239651159.14623069901081, data.Data.VolumeUsd)
	require.Equal(t, 141, data.Data.TradingPairs)
}

func TestMarkets(t *testing.T) {
	var data MarketsData
	err := unmarshalModel("markets", &data)
	require.NoError(t, err)
	require.Equal(t, 7, len(data.Data))
}

func TestCandles(t *testing.T) {
	var data CandlesData
	err := unmarshalModel("candles", &data)
	require.NoError(t, err)
	require.Equal(t, 10, len(data.Data))
	require.Equal(t, 15.169, data.Data[0].Open)
	require.Equal(t, 15.174, data.Data[0].High)
	require.Equal(t, 15.012, data.Data[0].Low)
	require.Equal(t, 15.13, data.Data[0].Close)
	require.Equal(t, 183240.592, data.Data[0].Volume)
}

func unmarshalModel(filename string, ptr interface{}) error {
	bs, err := os.ReadFile(fmt.Sprintf("mock/%s.json", filename))
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}
