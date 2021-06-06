package filters

import "github.com/kwanifi/numiscan-api/dmodels"

type Stats struct {
	Titles []string     `schema:"-"`
	To     dmodels.Time `schema:"to"`
	From   dmodels.Time `schema:"-"`
}
