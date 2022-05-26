package mapping

import (
	"time"

	"github.com/tgidk/go-api-365-add-in/entities"
	"github.com/tgidk/go-api-365-add-in/resources"
)

type Mapper interface {
	ToAsset(s entities.Spark) resources.Asset
}

type mapper struct{}

func GetNewMapper() Mapper {
	return &mapper{}
}

func (*mapper) ToAsset(s entities.Spark) resources.Asset {
	asset := resources.Asset{Symbol: s.Symbol}
	if len(s.Close) == len(s.Timestamp) {
		for i := range s.Close {
			quote := resources.Quote{Price: s.Close[i], Date: time.Unix(s.Timestamp[i], 0)}
			asset.Quotes = append(asset.Quotes, quote)
		}
	}
	return asset
}
