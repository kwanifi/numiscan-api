package smodels

import "github.com/kwanifi/numiscan-api/dmodels"

type HistoricalState struct {
	Current        dmodels.HistoricalState `json:"current"`
	PriceAgg       []AggItem               `json:"price_agg"`
	MarketCapAgg   []AggItem               `json:"market_cap_agg"`
	StakedRatioAgg []AggItem               `json:"staked_ratio"`
}
