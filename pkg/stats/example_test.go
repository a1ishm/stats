package stats

import (
	"fmt"
	"github.com/a1ishm/bank/v2/pkg/types")


func ExampleAvg() {
	payments := []types.Payment{
		{Amount: 10_000_00, Status: types.StatusOk},
		{Amount: 12_000_00, Status: types.StatusOk},
		{Amount: 16_000_00, Status: types.StatusOk},
		{Amount: 18_000_00, Status: types.StatusOk},
		{Amount: 30_000_00, Status: types.StatusFail},
	}
	fmt.Println(Avg(payments))

	//Output: 1400000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{Amount: 10_000_00, Category: "A"},
		{Amount: 12_000_00, Category: "B"},
		{Amount: 16_000_00, Category: "C"},
		{Amount: 18_000_00, Category: "A"},
		{Amount: 30_000_00, Category: "A", Status: types.StatusFail},
	}
	fmt.Println(TotalInCategory(payments, "A"))

	//Output: 2800000
}