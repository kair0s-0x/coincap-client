package coincap

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var client = NewClient()

func TestClient_GetAssets(t *testing.T) {
	assets, err := client.GetAssets(GetAssetsParams{LimitOffsetParams: LimitOffsetParams{Limit: 10, Offset: 10}})
	require.NoError(t, err)
	require.Equal(t, 10, len(assets.Data))
	ranks := make([]int, 10)
	for i, v := range assets.Data {
		ranks[i] = v.Rank
	}
	require.ElementsMatch(t, ranks, []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
}

func TestClient_GetAsset(t *testing.T) {
	asset, err := client.GetAsset("polkadot")
	require.NoError(t, err)
	require.Equal(t, "DOT", asset.Asset.Symbol)
	require.Equal(t, "polkadot", asset.Asset.Id)
	require.Equal(t, "Polkadot", asset.Asset.Name)
}

func TestClient_GetAssetHistory(t *testing.T) {
	interval := time.Hour * 24 * 7
	start := time.Now().Truncate(time.Hour * 24).UTC().Add(-interval)
	end := start.Add(interval)
	t.Run("EndExclusive", func(t *testing.T) {
		history, err := client.GetAssetHistory(GetAssetHistoryParams{
			Id:            "polkadot",
			HistoryParams: HistoryParams{Interval: H1, Start: start, End: end},
		})
		require.NoError(t, err)
		l := len(history.Data)
		require.Equal(t, 24*7, l)
		require.Equal(t, start, history.Data[0].Date)
		require.Equal(t, end.Add(-H1.Value()), history.Data[l-1].Date)
	})

	t.Run("EndInclusive", func(t *testing.T) {
		params := GetAssetHistoryParams{
			Id:            "polkadot",
			HistoryParams: HistoryParams{Interval: H1, Start: start, End: end.Add(time.Second)},
		}
		history, err := client.GetAssetHistory(params)
		require.NoError(t, err, params)
		l := len(history.Data)
		require.Equal(t, 24*7+1, l)
		require.Equal(t, start, history.Data[0].Date)
		require.Equal(t, end, history.Data[l-1].Date)
	})
}

func TestClient_GetAssetMarkets(t *testing.T) {
	ams, err := client.GetAssetMarkets(GetAssetMarketsParams{Id: "solana"})
	require.NoError(t, err)
	require.NotEmpty(t, ams.Data)
	for _, v := range ams.Data {
		require.Equal(t, "solana", v.BaseId)
	}
}

func TestClient_GetRates(t *testing.T) {
	rates, err := client.GetRates()
	require.NoError(t, err)
	require.NotEmpty(t, rates.Data)
	symbols := make([]string, len(rates.Data))
	for i, v := range rates.Data {
		symbols[i] = v.Symbol
	}
	require.Subset(t, symbols, []string{"USD", "GBP", "EUR", "USDT", "AUD", "TRY", "BTC", "ETH", "CAD"})
}

func TestClient_GetRate(t *testing.T) {
	rate, err := client.GetRate("australian-dollar")
	require.NoError(t, err)
	require.Equal(t, "AUD", rate.Data.Symbol)
}

func TestClient_GetExchanges(t *testing.T) {
	exchanges, err := client.GetExchanges()
	require.NoError(t, err)
	require.NotEmpty(t, exchanges)
	require.Equal(t, "binance", exchanges.Data[0].ExchangeId)
}

func TestClient_GetExchange(t *testing.T) {
	exchange, err := client.GetExchange("kraken")
	require.NoError(t, err)
	require.Equal(t, "kraken", exchange.Data.ExchangeId)
	require.Equal(t, "Kraken", exchange.Data.Name)
}

func TestClient_GetMarkets(t *testing.T) {
	t.Run("WithAsset", func(t *testing.T) {
		markets, err := client.GetMarkets(GetMarketsParams{ExchangeId: "kraken", AssetSymbol: "ADA"})
		require.NoError(t, err)
		require.NotEmpty(t, markets.Data)
		for _, v := range markets.Data {
			require.Equal(t, "cardano", v.BaseId)
		}
	})

	t.Run("WithBase/Quote", func(t *testing.T) {
		linkUsdc, err := client.GetMarkets(GetMarketsParams{ExchangeId: "binance", BaseSymbol: "link", QuoteId: "usd-coin"})
		require.NoError(t, err)
		require.Equal(t, "LINK", linkUsdc.Data[0].BaseSymbol)
		require.Equal(t, "USDC", linkUsdc.Data[0].QuoteSymbol)
	})
}

func TestClient_GetCandles(t *testing.T) {
	candles, err := client.GetCandles(GetCandlesParams{
		Exchange: "binance",
		BaseId:   "bitcoin",
		QuoteId:  "tether",
		HistoryParams: HistoryParams{
			Interval: D1,
			Start:    t1,
			End:      t1.Add(time.Hour * 30 * 24),
		},
	})
	require.NoError(t, err)
	require.Equal(t, 30, len(candles.Data))
}
