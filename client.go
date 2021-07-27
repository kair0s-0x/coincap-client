package coincap

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const httpUrl = "https://api.coincap.io/v2"
const wsUrl = "wss://ws.coincap.io/"

type Client struct {
	httpClient *http.Client
}

func NewDefaultClient() *Client {
	return &Client{httpClient: http.DefaultClient}
}

func NewCustomClient(client *http.Client) *Client {
	return &Client{httpClient: client}
}

func (c *Client) do(url string, params queryParams, ptr interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if params != nil {
		qp, err := params.toQuery()
		if err != nil {
			return err
		}
		q := req.URL.Query()
		for k, v := range qp {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("Accept-Encoding", "gzip")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	var reader io.ReadCloser
	switch enc := res.Header.Get("Content-Encoding"); enc {
	case "deflate":
		fallthrough
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return err
		}
		defer reader.Close()
	default:
		reader = res.Body
	}
	return json.NewDecoder(reader).Decode(ptr)
}

var ClientError = errors.New("client error")
var ServerError = errors.New("server error")

type api interface {
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
}

func assertApiInterface() {
	var _ api = (*Client)(nil)
}
