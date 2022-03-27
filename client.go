package coincap

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.coincap.io/v2"

//const wsUrl = "wss://ws.coincap.io/"

type CompressionType string

const (
	Gzip    CompressionType = "gzip"
	Deflate CompressionType = "deflate"
)

type Client struct {
	httpClient  *http.Client
	bearerToken string
	compression CompressionType
}

func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient:  http.DefaultClient,
		compression: Gzip,
	}
	for _, option := range options {
		option(client)
	}
	return client
}

type Option func(*Client)

func WithHttpClient(hc *http.Client) Option { return func(c *Client) { c.httpClient = hc } }
func WithBearerToken(bt string) Option      { return func(c *Client) { c.bearerToken = bt } }
func WithGzipCompression() Option           { return func(c *Client) { c.compression = Gzip } }
func WithDeflateCompression() Option        { return func(c *Client) { c.compression = Deflate } }

func (c *Client) Do(url string, params queryParams, ptr interface{}) error {
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
	req.Header.Set("Accept-Encoding", *(*string)(&c.compression))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
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

type Api interface {
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
	var _ Api = (*Client)(nil)
}
