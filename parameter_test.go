package coincap

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var t1 = time.Date(2021, 6, 26, 0, 0, 0, 0, time.UTC)
var t2 = t1.Add(time.Hour * 24)

func TestLimitOffsetParams(t *testing.T) {
	p := LimitOffsetParams{Limit: 10, Offset: 10}
	q, err := p.toQuery()
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"limit": "10", "offset": "10"}, q)
}

func TestHistoryParams(t *testing.T) {
	p := HistoryParams{Interval: M30, Start: t1, End: t2}
	q, err := p.toQuery()
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"interval": "m30", "start": "1624665600000", "end": "1624752000000"}, q)
}

func TestGetAssetsParams(t *testing.T) {
	p := GetAssetsParams{Search: "bt", LimitOffsetParams: LimitOffsetParams{Limit: 10, Offset: 10}}
	q, err := p.toQuery()
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"search": "bt", "limit": "10", "offset": "10"}, q)
	p = GetAssetsParams{Ids: []string{"polkadot", "solana"}}
	q, err = p.toQuery()
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"ids": "polkadot,solana"}, q)
}
