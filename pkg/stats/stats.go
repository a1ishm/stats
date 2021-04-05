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

func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	sum := map[types.Category]int{}
	categories := CategoriesTotal(payments)

	for _, payment := range payments {
		sum[payment.Category] += 1
	}

	for key, value := range categories {
		categories[key] = value / types.Money(sum[key])
	}

	return categories
}