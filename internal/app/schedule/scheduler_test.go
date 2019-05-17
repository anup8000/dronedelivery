package schedule

import (
	"testing"
	"time"

	"github.com/anup8000/dronedelivery/internal/app/mock"
)

func getMockSchedule(numschedules int, schedulefreq int, orderstarthour int, gridsize int) *Schedule {
	schedule := NewSchedule("../../../assets/testfiles/orders.txt", "../../../assets/testfiles/schedules.txt")
	initTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), orderstarthour, 0, 0, 0, time.Local)
	schedule.Orders = mock.GetNewData(numschedules, initTime, schedulefreq, gridsize)
	return schedule
}

func TestSchedule_ReadOrders(t *testing.T) {
	testschedule := &Schedule{InputFilePath: "../../../assets/testfiles/orders.txt"}
	testschedule.ReadOrders()
	orders := *testschedule.Orders
	if len(orders) != 20 {
		t.Errorf("Reading orders failed. Expected 1 order, found %d", len(orders))
	}
}

func TestSchedule_ReadBlankFile(t *testing.T) {
	testschedule := &Schedule{InputFilePath: "../../../assets/testfiles/orders1.txt"}
	testschedule.ReadOrders()
	orders := *testschedule.Orders
	if len(orders) != 0 {
		t.Errorf("Reading orders failed. Expected 0 order, found %d", len(orders))
	}
}

func TestSchedule_ReadBadFile(t *testing.T) {
	testschedule := &Schedule{InputFilePath: "../../../assets/testfiles/badorders.txt"}
	testschedule.ReadOrders()
	orders := *testschedule.Orders
	if len(orders) != 0 {
		t.Errorf("Reading orders failed. Expected 0 order, found %d", len(orders))
	}
}

//This test throws exception as we have set error mode to panic mode within schedule processor
// func TestSchedule_ReadBadFileWithSomeValidRows(t *testing.T) {
// 	testschedule := &Schedule{InputFilePath: "../../../assets/testfiles/badorders1.txt"}
// 	testschedule.ReadOrders()
// 	orders := *testschedule.Orders
// 	if len(orders) != 0 {
// 		t.Errorf("Reading orders failed. Expected 0 order, found %d", len(orders))
// 	}
// }

func TestSchedule_ProcessOrders(t *testing.T) {
	numschedules := 20
	testschedule := getMockSchedule(numschedules, 300, 5, 40)
	testschedule.ProcessOrders()
	if len(*testschedule.Schedules) != numschedules {
		t.Errorf("Failed test for processing schedules. Expected %d schedules, processed %d", numschedules, len(*testschedule.Schedules))
	}
}

func TestSchedule_CalculatePromoterScore(t *testing.T) {
	numschedules := 20
	testschedule := getMockSchedule(numschedules, 3000, 5, 5)
	testschedule.ProcessOrders()
	testschedule.CalculatePromoterScore()
	if testschedule.NetPromoterScore < 70 {
		t.Errorf("Tests for promoter score failed. The promoter score for test data is %f", testschedule.NetPromoterScore)
	} else {
		t.Logf("The promoter score is %f", testschedule.NetPromoterScore)
	}
}

func TestSchedule_WriteOrders(t *testing.T) {
	numschedules := 20
	testschedule := getMockSchedule(numschedules, 3000, 5, 5)
	testschedule.ProcessOrders()
	testschedule.CalculatePromoterScore()
	testschedule.WriteSchedules()
}
