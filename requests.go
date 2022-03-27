package coincap

import "fmt"

func (c *Client) GetAssets(params GetAssetsParams) (AssetsData, error) {
	var data AssetsData
	err := c.Do(fmt.Sprintf("%s/assets", url), params, &data)
	return data, err
}

func (c *Client) GetAsset(id string) (AssetData, error) {
	var data AssetData
	if len(id) == 0 {
		return data, MissingParameterError
	}
	err := c.Do(fmt.Sprintf("%s/assets/%s", url, id), nil, &data)
	return data, err
}

func (c *Client) GetAssetHistory(params GetAssetHistoryParams) (AssetHistoriesData, error) {
	var data AssetHistoriesData
	if len(params.Id) == 0 {
		return data, MissingParameterError
	}
	err := c.Do(fmt.Sprintf("%s/assets/%s/history", url, params.Id), params, &data)
	return data, err
}

func (c *Client) GetAssetMarkets(params GetAssetMarketsParams) (AssetMarketsData, error) {
	var data AssetMarketsData
	if len(params.Id) == 0 {
		return data, MissingParameterError
	}
	err := c.Do(fmt.Sprintf("%s/assets/%s/markets", url, params.Id), params, &data)
	return data, err
}

func (c *Client) GetRates() (RatesData, error) {
	var data RatesData
	err := c.Do(fmt.Sprintf("%s/rates", url), nil, &data)
	return data, err
}

func (c *Client) GetRate(id string) (RateData, error) {
	var data RateData
	if len(id) == 0 {
		return data, MissingParameterError
	}
	err := c.Do(fmt.Sprintf("%s/rates/%s", url, id), nil, &data)
	return data, err
}

func (c *Client) GetExchanges() (ExchangesData, error) {
	var data ExchangesData
	err := c.Do(fmt.Sprintf("%s/exchanges", url), nil, &data)
	return data, err
}

func (c *Client) GetExchange(id string) (ExchangeData, error) {
	var data ExchangeData
	if len(id) == 0 {
		return data, MissingParameterError
	}
	err := c.Do(fmt.Sprintf("%s/exchanges/%s", url, id), nil, &data)
	return data, err
}

func (c *Client) GetMarkets(params GetMarketsParams) (MarketsData, error) {
	var data MarketsData
	err := c.Do(fmt.Sprintf("%s/markets", url), params, &data)
	return data, err
}

func (c *Client) GetCandles(params GetCandlesParams) (CandlesData, error) {
	var data CandlesData
	err := c.Do(fmt.Sprintf("%s/candles", url), params, &data)
	return data, err
}
