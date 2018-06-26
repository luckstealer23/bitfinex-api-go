package rest

import (
	"net/url"

	"github.com/luckstealer23/bitfinex-api-go/v2"
)

type TickerService struct {
	Synchronous
}

func (t *TickerService) All(symbols ...string) (*bitfinex.TickerSnapshot, error) {
	req := NewRequestWithMethod("tickers", "GET")
	req.Params = make(url.Values)

	for _, symbol := range symbols {
		req.Params.Add("symbols", symbol)
	}
	raw, err := t.Request(req)

	if err != nil {
		return nil, err
	}

	data := make([][]interface{}, 0, len(raw))
	for _, rawTicker := range raw {
		if rawTicker, ok := rawTicker.([]interface{}); ok {
			data = append(data, rawTicker)
		}
	}

	tickers, err := bitfinex.NewTickerSnapshotFromRawRest(data)
	if err != nil {
		return nil, err
	}

	return tickers, nil
}
