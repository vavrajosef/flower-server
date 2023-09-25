package flowerfilter

import (
	"time"

	api "github.com/vavrajosef/flower-server/api"
)

func FilterUnwatered(flowers []api.FlowerDetail) []api.FlowerDetail {
	var result []api.FlowerDetail
	today := time.Now()
	for _, flower := range flowers {
		lastWatered := flower.LastWatered
		daysSince := int(today.Sub(lastWatered).Hours() / 24)
		if daysSince >= flower.Period {
			result = append(result, flower)
		}
	}
	return result
}
