package fileops

import (
	"testing"
	"time"

	"github.com/anup8000/dronedelivery/internal/app/mock"
)

func TestReadOrders(t *testing.T) {
	numorders := 20
	orders := mock.GetNewData(numorders, time.Time{}, 300, 40)
	err := WriteOrders("../../../assets/testfiles/orders.txt", orders)
	if err != nil {
		t.Errorf("Writing orders failed!")
	}
	readorders, err := ReadOrders("../../../assets/testfiles/orders.txt")
	if err != nil {
		t.Errorf("Reading orders failed!")
	}
	if len(*readorders) != numorders {
		t.Errorf("Writing orders failed! Expected 1 row, got %d rows", len(*readorders))
	}
}

func TestWriteSchedules(t *testing.T) {
	orders := mock.GetNewData(20, time.Time{}, 300, 40)
	err := WriteSchedules("../../../assets/testfiles/schedules.txt", orders, 50.00000)
	if err != nil {
		t.Errorf("Writing orders failed!")
	}
}

func TestWriteOrders(t *testing.T) {
	numorders := 20
	orders := mock.GetNewData(numorders, time.Time{}, 300, 40)
	err := WriteOrders("../../../assets/testfiles/orders.txt", orders)
	if err != nil {
		t.Errorf("Writing orders failed!")
	}
	readorders, err := ReadOrders("../../../assets/testfiles/orders.txt")
	if err != nil {
		t.Errorf("Reading orders failed!")
	}
	if len(*readorders) != numorders {
		t.Errorf("Writing orders failed! Expected 1 row, got %d rows", len(*readorders))
	}
}
