package coincap

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var MissingParameterError = errors.New("missing parameter")
var InvalidParameterError = errors.New("invalid parameter")

type queryParams interface {
	toQuery() (map[string]string, error)
}

type LimitOffsetParams struct {
	Limit  int // optional, max limit of 2000
	Offset int // optional
}

func (r LimitOffsetParams) toQuery() (map[string]string, error) {
	if r.Limit > 2000 {
		return nil, InvalidParameterError
	}
	q := make(map[string]string, 2)
	if r.Limit > 0 {
		q["limit"] = strconv.Itoa(r.Limit)
	}
	if r.Offset > 0 {
		q["offset"] = strconv.Itoa(r.Offset)
	}
	return q, nil
}

func (r LimitOffsetParams) include(m map[string]string) (map[string]string, error) {
	q, err := r.toQuery()
	if err != nil {
		return nil, err
	}
	for k, v := range q {
		m[k] = v
	}
	return m, nil
}

type HistoryParams struct {
	Interval Interval  // required
	Start    time.Time // optional, UNIX milliseconds. omitting returns the most recent.
	End      time.Time // optional, end timestamp excluded from results
}

func (r HistoryParams) toQuery() (map[string]string, error) {
	hasStart := !r.Start.IsZero()
	hasEnd := !r.End.IsZero()
	if r.Interval == 0 || (hasStart && !hasEnd) || (!hasStart && hasEnd) {
		return nil, MissingParameterError
	}
	if hasStart && hasEnd {
		if r.Start.Add(r.Interval.Value()).After(r.End) {
			return nil, InvalidParameterError // todo specific error message
		}
	}
	q := make(map[string]string, 3)
	q["interval"] = fmt.Sprint(r.Interval)
	if hasStart {
		q["start"] = strconv.FormatInt(r.Start.UTC().Unix()*1e3, 10)
	}
	if hasEnd {
		q["end"] = strconv.FormatInt(r.End.UTC().Unix()*1e3, 10)
	}
	return q, nil
}

func (r HistoryParams) include(m map[string]string) (map[string]string, error) {
	q, err := r.toQuery()
	if err != nil {
		return nil, err
	}
	for k, v := range q {
		m[k] = v
	}
	return m, nil
}

type GetAssetsParams struct {
	Search string   // optional, Search by asset Id (bitcoin) or symbol (BTC)
	Ids    []string // optional, query with multiple Ids=bitcoin,ethereum,...
	LimitOffsetParams
}

func (r GetAssetsParams) toQuery() (map[string]string, error) {
	q := make(map[string]string, 4)
	if r.Limit > 2000 || len(r.Ids) > 2000 {
		return nil, InvalidParameterError
	}
	if len(r.Search) > 0 {
		q["search"] = r.Search
	}
	if len(r.Ids) > 0 {
		q["ids"] = strings.Join(r.Ids, ",")
	}
	return r.LimitOffsetParams.include(q)
}

type GetAssetHistoryParams struct {
	Id string // required, asset Id
	HistoryParams
}

func (r GetAssetHistoryParams) toQuery() (map[string]string, error) {
	if len(r.Id) == 0 {
		return nil, MissingParameterError
	}
	return r.HistoryParams.toQuery()
}

type GetAssetMarketsParams struct {
	Id string // required
	LimitOffsetParams
}

func (r GetAssetMarketsParams) toQuery() (map[string]string, error) {
	if len(r.Id) == 0 {
		return nil, MissingParameterError
	}
	return r.LimitOffsetParams.include(make(map[string]string, 2))
}

type GetMarketsParams struct {
	ExchangeId  string // optional,	search by exchange id
	BaseSymbol  string // optional, returns all containing the base symbol
	QuoteSymbol string // optional, returns all containing the quote symbol
	BaseId      string // optional, returns all containing the base id
	QuoteId     string // optional, returns all containing the quote id
	AssetSymbol string // optional, returns all assets containing symbol (base and quote)
	AssetId     string // optional, returns all assets containing id (base and quote)
	LimitOffsetParams
}

func (r GetMarketsParams) toQuery() (map[string]string, error) {
	q := map[string]string{}
	if len(r.ExchangeId) > 0 {
		q["exchangeId"] = r.ExchangeId
	}
	if len(r.BaseSymbol) > 0 {
		q["baseSymbol"] = r.BaseSymbol
	}
	if len(r.QuoteSymbol) > 0 {
		q["quoteSymbol"] = r.QuoteSymbol
	}
	if len(r.BaseId) > 0 {
		q["baseId"] = r.BaseId
	}
	if len(r.QuoteId) > 0 {
		q["quoteId"] = r.QuoteId
	}
	if len(r.AssetSymbol) > 0 {
		q["assetSymbol"] = r.AssetSymbol
	}
	if len(r.AssetId) > 0 {
		q["assetId"] = r.AssetId
	}
	return r.LimitOffsetParams.include(q)
}

type GetCandlesParams struct {
	Exchange string
	BaseId   string
	QuoteId  string
	HistoryParams
}

func (r GetCandlesParams) toQuery() (map[string]string, error) {
	if len(r.Exchange) == 0 || len(r.BaseId) == 0 || len(r.QuoteId) == 0 {
		return nil, MissingParameterError
	}
	return r.HistoryParams.include(map[string]string{
		"exchange": r.Exchange,
		"baseId":   r.BaseId,
		"quoteId":  r.QuoteId,
	})
}
