package stats

import (
	"github.com/a1ishm/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var total types.Money
	var len = len(payments)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			len -= 1
			continue
		}
		total += payment.Amount
	}
	avg := total / types.Money((len))
	return avg
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var total types.Money	
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}

		if payment.Status == types.StatusFail {
			continue
		}
		total += payment.Amount
	}
	return total
}

func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}