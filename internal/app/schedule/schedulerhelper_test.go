package schedule

import (
	"testing"
	"time"

	"github.com/anup8000/dronedelivery/internal/app/mock"
	"github.com/anup8000/dronedelivery/internal/pkg/models"
)

func getfakeorder() *models.Order {
	initTimestamp := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 6, 0, 0, 0, time.Local)
	orders := *mock.GetNewData(1, initTimestamp, 300, 40)
	order := orders[0]
	order.Coordinates = "N4E3"
	return &order
}

func getfakeorders() *[]models.Order {
	return mock.GetNewData(2, time.Time{}, 300, 40)
}

func TestCalculateDistanceForCoordinates(t *testing.T) {
	type args struct {
		coord string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "TestCalculateDistanceForCoordinates", args: args{coord: "S4E3"}, want: 10.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDistanceForCoordinates(tt.args.coord); got != tt.want {
				t.Errorf("CalculateDistanceForCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateWaitTime(t *testing.T) {
	initTimestamp := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 6, 30, 0, 0, time.Local)
	order := getfakeorder()
	waittime := CalculateWaitTime(initTimestamp, order)
	if waittime > 100 {
		t.Errorf("Test for wait time calculations failed. Expected less than 100, got %f", waittime)
	}
}

func TestPickOrderToSubmit(t *testing.T) {
	initTimestamp := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 6, 0, 0, 0, time.Local)
	orders := getfakeorders()
	_, index := PickOrderToSubmit(initTimestamp, orders)
	if index != 1 {
		t.Errorf("Test for picking order to submit failed. Expected 1st order to pick, got %d", index)
	}
}

func TestCalculatePromoterScore(t *testing.T) {
	order := getfakeorder()
	CalculatePromoterScore(order)
	if !order.IsPromoter {
		t.Errorf("Test for checking promoter. Expected order to be a promoter, but got %v", order.IsPromoter)
	}
}
