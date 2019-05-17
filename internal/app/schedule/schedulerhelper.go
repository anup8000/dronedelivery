package schedule

import (
	"math"
	"strconv"
	"time"

	"github.com/anup8000/dronedelivery/internal/pkg/models"
)

//PickOrderToSubmit picks an order to submit from a slice of orders
func PickOrderToSubmit(scheduletime time.Time, ordersptr *[]models.Order) (*models.Order, int) {
	minwaittime := 9999.0
	selectedorderindex := 0
	orders := *ordersptr
	if orders == nil || len(orders) == 0 {
		return nil, 0
	}
	for index, order := range orders {
		waittime := CalculateWaitTime(scheduletime, &order)
		if waittime < minwaittime {
			minwaittime = waittime
			selectedorderindex = index
			orders[index].Waitinminutes = waittime
		}
	}
	return &orders[selectedorderindex], selectedorderindex
}

//CalculateDistanceForCoordinates calculates distance given coordinates
func CalculateDistanceForCoordinates(coord string) float64 {
	numcount := 0
	num := 0
	var arr [2]int

	for i, chr := range coord {
		//escape the first character as we know it is a letter
		if i == 0 {
			continue
		}

		if x, err := strconv.Atoi(string(chr)); err == nil {
			num = num*10 + x
		} else {
			arr[numcount] = num
			num = 0
			numcount = numcount + 1
		}
	}
	arr[numcount] = num
	return math.Sqrt(math.Pow(float64(arr[0]), 2)+math.Pow(float64(arr[1]), 2)) * 2
}

//CalculateWaitTime calculates wait time
//(scheduledtime + hypotenuse dist * 2) - order timestamp
func CalculateWaitTime(scheduletime time.Time, order *models.Order) float64 {
	return scheduletime.Add(time.Second * time.Duration(order.Distance*60)).Sub(order.Timestamp).Minutes()
}

//CalculatePromoterScore calculates promoter score based on the given rules
func CalculatePromoterScore(order *models.Order) {
	if order.Waitinminutes < 120 {
		order.IsPromoter = true
	}
	if order.Waitinminutes >= 240 {
		order.IsDetractor = true
	}
	if order.Waitinminutes >= 120 && order.Waitinminutes < 240 {
		order.IsNeutral = true
	}
}
