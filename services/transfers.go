package services

import (
	"fmt"

	"github.com/kwanifi/numiscan-api/dao/filters"
	"github.com/kwanifi/numiscan-api/smodels"
)

func (s *ServiceFacade) GetAggTransfersVolume(filter filters.Agg) (items []smodels.AggItem, err error) {
	items, err = s.dao.GetAggTransfersVolume(filter)
	if err != nil {
		return nil, fmt.Errorf("dao.GetAggTransfersVolume: %s", err.Error())
	}
	return items, nil
}
