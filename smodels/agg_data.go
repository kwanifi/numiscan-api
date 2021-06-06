package smodels

import (
	"github.com/kwanifi/numiscan-api/dmodels"
	"github.com/shopspring/decimal"
)

type AggItem struct {
	Time  dmodels.Time    `db:"time" json:"time"`
	Value decimal.Decimal `db:"value" json:"value"`
}
