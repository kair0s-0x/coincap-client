# coincap-go

[CoinCap](https://coincap.io/) REST API wrapper written with `Go`. For `docs` checkout: https://docs.coincap.io/

## Installation

```shell
go get -u github.com/softronaut/coincap-go
```

## Usage

```go
client := coincap.DefaultClient()
// or with retryablehttp
retry := retryablehttp.NewClient()
retry.RetryMax = 3
retry.RetryWaitMax = time.Second * 5
coincap.CustomClient(retry.StandardClient())

// api interface
GetAssets(GetAssetsParams) (AssetsData, error)
GetAsset(string) (AssetData, error)
GetAssetHistory(GetAssetHistoryParams) (AssetHistoriesData, error)
GetAssetMarkets(GetAssetMarketsParams) (AssetMarketsData, error)
GetRates() (RatesData, error)
GetRate(string) (RateData, error)
GetExchanges() (ExchangesData, error)
GetExchange(string) (ExchangeData, error)
GetMarkets(GetMarketsParams) (MarketsData, error)
GetCandles(GetCandlesParams) (CandlesData, error)

// request examples
assets, err := client.GetAssets(GetAssetsParams{Ids: []string{"bitcoin", "ethereum"}})
polkadot, err := client.GetAsset("polkadot")
linkUsdc, err := client.GetMarkets(GetMarketsParams{ExchangeId: "binance", BaseSymbol: "link", QuoteId: "usd-coin"})
```

## Notes

All return and parameter types created.

Required parameters and limits handled before the requests, so you will know some errors during the development.

`gzip` encoding enabled by default.

## ToDo

- WebSocket support
- Extensive error handling
