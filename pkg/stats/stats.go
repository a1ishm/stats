package stats

import (
	"github.com/a1ishm/bank/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var total types.Money
	for _, payment := range payments {
		total += payment.Amount
	}
	avg := total / types.Money((len(payments)))
	return avg
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		total += payment.Amount
	}
	return total
}