# coincap-go

[CoinCap](https://coincap.io/) REST API wrapper written with `Go`. For `docs` checkout: https://docs.coincap.io/

## Installation

```shell
go get -u github.com/softronaut/coincap-go
```

## Usage

```go
// Instance
client := coincap.DefaultClient() // uses http.DefaultClient
// Customization
retry := retryablehttp.NewClient() // or you can use retryablehttp
retry.RetryMax = 3
retry.RetryWaitMax = time.Second * 5
client := coincap.CustomClient(retry.StandardClient())

// API interface
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

// Examples
assets, err := client.GetAssets(GetAssetsParams{Ids: []string{"bitcoin", "ethereum"}})
polkadot, err := client.GetAsset("polkadot")
linkUsdc, err := client.GetMarkets(GetMarketsParams{ExchangeId: "binance", BaseSymbol: "link", QuoteId: "usd-coin"})
```

## Notes

Each response and parameter declared as `struct`.

Some parameter logics implemented (required parameters, api limits or start/end timestamp relations etc.).

`gzip` encoding enabled by default.

## ToDo

- WebSocket support
- Extensive error handling
