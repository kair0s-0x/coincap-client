# coincap-go

[CoinCap.io](https://coincap.io/) REST API wrapper written with Go. For docs checkout: https://docs.coincap.io/

## Installation

```sh
go get -u github.com/kair0s-0x/coincap-go
```

## Usage

```go
// Instance
client := coincap.DefaultClient() // uses http.DefaultClient
// Customization
retry := retryablehttp.NewClient() // or you can use a customized http.Client
retry.RetryMax = 3
retry.RetryWaitMax = time.Second * 5
client := coincap.CustomClient(retry.StandardClient())

// API interface
GetAssets(GetAssetsParams) (AssetsData, error)
GetAsset(id string) (AssetData, error)
GetAssetHistory(GetAssetHistoryParams) (AssetHistoriesData, error)
GetAssetMarkets(GetAssetMarketsParams) (AssetMarketsData, error)
GetRates() (RatesData, error)
GetRate(id string) (RateData, error)
GetExchanges() (ExchangesData, error)
GetExchange(id string) (ExchangeData, error)
GetMarkets(GetMarketsParams) (MarketsData, error)
GetCandles(GetCandlesParams) (CandlesData, error)

// Examples
assets, err := client.GetAssets(GetAssetsParams{Ids: []string{"bitcoin", "ethereum"}})
polkadot, err := client.GetAsset("polkadot")
linkUsdc, err := client.GetMarkets(GetMarketsParams{ExchangeId: "binance", BaseSymbol: "link", QuoteId: "usd-coin"})
```

## Notes

Each `response` and `parameter` declared as `struct`.

Some parameter logics implemented (required parameters, api limits or start/end timestamp relations etc.).

`gzip` encoding enabled by default.

## ToDo

- WebSocket support
- Extensive error handling
