package coincap

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestAsset(t *testing.T) {
	var data AssetData
	err := unmarshalModel("asset_id", &data)
	asset := data.Asset
	assert.NoError(t, err)
	assert.Equal(t, int64(1627299055657), data.Timestamp)
	assert.Equal(t, "polkadot", asset.Id)
	assert.Equal(t, 9, asset.Rank)
	assert.Equal(t, "DOT", asset.Symbol)
	assert.Equal(t, "Polkadot", asset.Name)
	assert.Equal(t, 1013089106.2001500000000000, asset.Supply)
	assert.Nil(t, asset.MaxSupply)
	assert.Equal(t, 14964432257.0276593916630207, asset.MarketCapUsd)
	assert.Equal(t, 567049854.0292255884810806, asset.VolumeUsd24Hr)
	assert.Equal(t, 14.7710918668897673, asset.PriceUsd)
	assert.Equal(t, 10.0829076670949820, asset.ChangePercent24Hr)
	assert.Equal(t, 14.3158259816882364, asset.Vwap24Hr)
	assert.Equal(t, "https://polkascan.io/polkadot", *asset.Explorer)
}

func TestAssets(t *testing.T) {
	var data AssetsData
	err := unmarshalModel("assets", &data)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(data.Data))
}

func TestAssetHistory(t *testing.T) {
	var data AssetHistoriesData
	err := unmarshalModel("asset_history", &data)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(data.Data))
	assert.Equal(t, 1013089106.25717, data.Data[0].CirculatingSupply)
	date := time.Date(2021, 7, 26, 11, 0, 0, 0, time.UTC)
	for i, v := range data.Data {
		assert.Equal(t, date.Add(time.Minute*time.Duration(30*i)), v.Date)
	}
}

func TestAssetMarkets(t *testing.T) {
	var data AssetMarketsData
	err := unmarshalModel("asset_markets", &data)
	assert.NoError(t, err)
	assert.Equal(t, 7, len(data.Data))
}

func TestRates(t *testing.T) {
	var data RatesData
	err := unmarshalModel("assets", &data)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(data.Data))
}

func TestRatesID(t *testing.T) {
	var data RateData
	err := unmarshalModel("rates_id", &data)
	assert.NoError(t, err)
	assert.Equal(t, "USDC", data.Data.Symbol)
	assert.Equal(t, "usd-coin", data.Data.Id)
	assert.Equal(t, 1.0000000000000000, data.Data.RateUsd)
	var nilStringPointer *string = nil
	assert.Equal(t, nilStringPointer, data.Data.CurrencySymbol)
	assert.Equal(t, "crypto", data.Data.Type)
}

func TestExchanges(t *testing.T) {
	var data ExchangesData
	err := unmarshalModel("exchanges", &data)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(data.Data))
	assert.Equal(t, "binance", data.Data[0].ExchangeId)
}

func TestExchange(t *testing.T) {
	var data ExchangeData
	err := unmarshalModel("exchange", &data)
	assert.NoError(t, err)
	assert.Equal(t, "kraken", data.Data.ExchangeId)
	assert.Equal(t, 1.845246588138155764, data.Data.PercentTotalVolume)
	assert.Equal(t, 1239651159.14623069901081, data.Data.VolumeUsd)
	assert.Equal(t, 141, data.Data.TradingPairs)
}

func TestMarkets(t *testing.T) {
	var data MarketsData
	err := unmarshalModel("markets", &data)
	assert.NoError(t, err)
	assert.Equal(t, 7, len(data.Data))
}

func TestCandles(t *testing.T) {
	var data CandlesData
	err := unmarshalModel("candles", &data)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(data.Data))
	assert.Equal(t, 15.169, data.Data[0].Open)
	assert.Equal(t, 15.174, data.Data[0].High)
	assert.Equal(t, 15.012, data.Data[0].Low)
	assert.Equal(t, 15.13, data.Data[0].Close)
	assert.Equal(t, 183240.592, data.Data[0].Volume)
}

func unmarshalModel(filename string, ptr interface{}) error {
	bs, err := os.ReadFile(fmt.Sprintf("mock/%s.json", filename))
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}
