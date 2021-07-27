package coincap

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var client = NewDefaultClient()

func TestClient_GetAssets(t *testing.T) {
	assets, err := client.GetAssets(GetAssetsParams{LimitOffsetParams: LimitOffsetParams{Limit: 10, Offset: 10}})
	assert.NoError(t, err)
	assert.Equal(t, 10, len(assets.Data))
	ranks := make([]int, 10)
	for i, v := range assets.Data {
		ranks[i] = v.Rank
	}
	assert.ElementsMatch(t, ranks, []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
}

func TestClient_GetAsset(t *testing.T) {
	asset, err := client.GetAsset("polkadot")
	assert.NoError(t, err)
	assert.Equal(t, "DOT", asset.Asset.Symbol)
	assert.Equal(t, "polkadot", asset.Asset.Id)
	assert.Equal(t, "Polkadot", asset.Asset.Name)
}

func TestClient_GetAssetHistory(t *testing.T) {
	t.Run("EndExclusive", func(t *testing.T) {
		end := t1.Add(time.Hour * 24 * 14)
		history, err := client.GetAssetHistory(GetAssetHistoryParams{
			Id: "polkadot",
			HistoryParams: HistoryParams{
				Interval: M30,
				Start:    t1,
				End:      end,
			},
		})
		assert.NoError(t, err)
		l := len(history.Data)
		assert.Equal(t, 24*14*2, l)
		assert.Equal(t, t1, history.Data[0].Date)
		assert.Equal(t, end.Add(-M30.Value()), history.Data[l-1].Date)
	})

	t.Run("EndInclusive", func(t *testing.T) {
		end := t1.Add(time.Hour * 24 * 7)
		params := GetAssetHistoryParams{
			Id: "polkadot",
			HistoryParams: HistoryParams{
				Interval: M30,
				Start:    t1,
				End:      end.Add(time.Second),
			},
		}
		q, _ := params.toQuery()
		t.Log(q)
		history, err := client.GetAssetHistory(params)
		assert.NoError(t, err, params)
		l := len(history.Data)
		assert.Equal(t, 24*7*2+1, l)
		assert.Equal(t, t1, history.Data[0].Date)
		assert.Equal(t, end, history.Data[l-1].Date)
	})
}

func TestClient_GetAssetMarkets(t *testing.T) {
	ams, err := client.GetAssetMarkets(GetAssetMarketsParams{Id: "solana"})
	assert.NoError(t, err)
	assert.NotEmpty(t, ams.Data)
	for _, v := range ams.Data {
		assert.Equal(t, "solana", v.BaseId)
	}
}

func TestClient_GetRates(t *testing.T) {
	rates, err := client.GetRates()
	assert.NoError(t, err)
	assert.NotEmpty(t, rates.Data)
	symbols := make([]string, len(rates.Data))
	for i, v := range rates.Data {
		symbols[i] = v.Symbol
	}
	assert.Subset(t, symbols, []string{"USD", "GBP", "EUR", "USDC", "USDT", "AUD", "TRY", "BTC", "ETH", "CAD"})
}

func TestClient_GetRate(t *testing.T) {
	rate, err := client.GetRate("australian-dollar")
	assert.NoError(t, err)
	assert.Equal(t, "AUD", rate.Data.Symbol)
}

func TestClient_GetExchanges(t *testing.T) {
	exchanges, err := client.GetExchanges()
	assert.NoError(t, err)
	assert.NotEmpty(t, exchanges)
	assert.Equal(t, "binance", exchanges.Data[0].ExchangeId)
}

func TestClient_GetExchange(t *testing.T) {
	exchange, err := client.GetExchange("kraken")
	assert.NoError(t, err)
	assert.Equal(t, "kraken", exchange.Data.ExchangeId)
	assert.Equal(t, "Kraken", exchange.Data.Name)
}

func TestClient_GetMarkets(t *testing.T) {
	t.Run("WithAsset", func(t *testing.T) {
		markets, err := client.GetMarkets(GetMarketsParams{ExchangeId: "kraken", AssetSymbol: "ADA"})
		assert.NoError(t, err)
		assert.NotEmpty(t, markets.Data)
		for _, v := range markets.Data {
			assert.Equal(t, "cardano", v.BaseId)
		}
	})

	t.Run("WithBase/Quote", func(t *testing.T) {
		linkUsdc, err := client.GetMarkets(GetMarketsParams{ExchangeId: "binance", BaseSymbol: "link", QuoteId: "usd-coin"})
		assert.NoError(t, err)
		assert.Equal(t, "LINK", linkUsdc.Data[0].BaseSymbol)
		assert.Equal(t, "USDC", linkUsdc.Data[0].QuoteSymbol)
	})
}

func TestClient_GetCandles(t *testing.T) {
	candles, err := client.GetCandles(GetCandlesParams{
		Exchange: "binance",
		BaseId:   "bitcoin",
		QuoteId:  "tether",
		HistoryParams: HistoryParams{
			Interval: M30,
			Start:    t1,
			End:      t1.Add(time.Hour * 7 * 24),
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, 7*24*2-1, len(candles.Data))
}
