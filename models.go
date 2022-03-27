package coincap

import (
	"time"
)

type Asset struct {
	Id                string   `json:"id"`
	Rank              int      `json:"rank,string"`
	Symbol            string   `json:"symbol"`
	Name              string   `json:"name"`
	Supply            float64  `json:"supply,string"`
	MaxSupply         *float64 `json:"maxSupply,string"`
	MarketCapUsd      float64  `json:"marketCapUsd,string"`
	VolumeUsd24Hr     float64  `json:"volumeUsd24Hr,string"`
	PriceUsd          float64  `json:"priceUsd,string"`
	ChangePercent24Hr float64  `json:"changePercent24Hr,string"`
	Vwap24Hr          float64  `json:"vwap24Hr,string"`
	Explorer          *string  `json:"explorer"`
}

type AssetData struct {
	Asset     Asset `json:"data,omitempty"`
	Timestamp int64 `json:"timestamp"`
}

type AssetsData struct {
	Data      []Asset `json:"data,omitempty"`
	Timestamp int64   `json:"timestamp"`
}

type AssetHistory struct {
	PriceUsd          float64   `json:"priceUsd,string"`
	Time              int64     `json:"time"`
	CirculatingSupply float64   `json:"circulatingSupply,string"`
	Date              time.Time `json:"date"`
}
type AssetHistoriesData struct {
	Data      []AssetHistory `json:"data,omitempty"`
	Timestamp int64          `json:"timestamp"`
}

type AssetMarket struct {
	ExchangeId    string  `json:"exchangeId"`
	BaseId        string  `json:"baseId"`
	QuoteId       string  `json:"quoteId"`
	BaseSymbol    string  `json:"baseSymbol"`
	QuoteSymbol   string  `json:"quoteSymbol"`
	VolumeUsd24Hr float64 `json:"volumeUsd24Hr,string"`
	PriceUsd      float64 `json:"priceUsd,string"`
	VolumePercent float64 `json:"volumePercent,string"`
}
type AssetMarketsData struct {
	Data      []AssetMarket `json:"data,omitempty"`
	Timestamp int64         `json:"timestamp"`
}

type Rate struct {
	Id             string  `json:"Id"`
	Symbol         string  `json:"symbol"`
	CurrencySymbol *string `json:"currencySymbol"`
	Type           string  `json:"type"`
	RateUsd        float64 `json:"rateUsd,string"`
}
type RateData struct {
	Data      Rate  `json:"data,omitempty"`
	Timestamp int64 `json:"timestamp"`
}
type RatesData struct {
	Data      []Rate `json:"data,omitempty"`
	Timestamp int64  `json:"timestamp"`
}

type Exchange struct {
	ExchangeId         string  `json:"exchangeId"`
	Name               string  `json:"name"`
	Rank               int     `json:"rank,string"`
	PercentTotalVolume float64 `json:"percentTotalVolume,string"`
	VolumeUsd          float64 `json:"volumeUsd,string"`
	TradingPairs       int     `json:"tradingPairs,string"`
	Socket             bool    `json:"socket"`
	ExchangeUrl        string  `json:"exchangeUrl"`
	Updated            int64   `json:"updated"`
}
type ExchangeData struct {
	Data      Exchange `json:"data,omitempty"`
	Timestamp int64    `json:"timestamp"`
}
type ExchangesData struct {
	Data      []Exchange `json:"data,omitempty"`
	Timestamp int64      `json:"timestamp"`
}

type Market struct {
	ExchangeId            string  `json:"exchangeId"`
	Rank                  int     `json:"rank,string"`
	BaseSymbol            string  `json:"baseSymbol"`
	BaseId                string  `json:"baseId"`
	QuoteSymbol           string  `json:"quoteSymbol"`
	QuoteId               string  `json:"quoteId"`
	PriceQuote            float64 `json:"priceQuote,string"`
	PriceUsd              float64 `json:"priceUsd,string"`
	VolumeUsd24Hr         float64 `json:"volumeUsd24Hr,string"`
	PercentExchangeVolume float64 `json:"percentExchangeVolume,string"`
	TradesCount24Hr       int64   `json:"tradesCount24Hr,string"`
	Updated               int64   `json:"updated"`
}
type MarketsData struct {
	Data      []Market `json:"data,omitempty"`
	Timestamp int64    `json:"timestamp"`
}

type Candle struct {
	Open   float64 `json:"open,string"`
	High   float64 `json:"high,string"`
	Low    float64 `json:"low,string"`
	Close  float64 `json:"close,string"`
	Volume float64 `json:"volume,string"`
	Period int64   `json:"period"`
}
type CandlesData struct {
	Data      []Candle `json:"data,omitempty"`
	Timestamp int64    `json:"timestamp"`
}
