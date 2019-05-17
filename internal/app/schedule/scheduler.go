package schedule

import (
	"fmt"
	"time"

	"github.com/anup8000/dronedelivery/internal/pkg/errorhandler"
	"github.com/anup8000/dronedelivery/internal/pkg/fileops"
	"github.com/anup8000/dronedelivery/internal/pkg/models"
)

//Schedule defines alias for Schedule in models
type Schedule models.Schedule

//NewSchedule returns a schedule with a default start time
func NewSchedule(inputfilepath string, outputfilepath string) *Schedule {
	return &Schedule{ScheduleStartTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 6, 0, 0, 0, time.Local),
		ScheduleEndTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 0, 0, 0, time.Local),
		InputFilePath:   inputfilepath, OutputFilePath: outputfilepath,
	}
}

/*************************************************CLASS METHODS******************************************
*********************************************************************************************************/

//ReadOrders will read from file and generate orders
func (schedule *Schedule) ReadOrders() {
	var err error
	schedule.Orders, err = fileops.ReadOrders(schedule.InputFilePath)
	//Calculate distance
	for i, order := range *schedule.Orders {
		(*schedule.Orders)[i].Distance = CalculateDistanceForCoordinates(order.Coordinates)
	}

	errorhandler.Checkandpanic(err)
}

//WriteSchedules will schedule the orders and write to a file
func (schedule *Schedule) WriteSchedules() {
	if len(*schedule.Schedules) > 0 {
		err := fileops.WriteSchedules(schedule.OutputFilePath, schedule.Schedules, schedule.NetPromoterScore)
		errorhandler.Checkandpanic(err)
	}
}

//ProcessOrders will process the orders and schedule the processing time as the orders come
func (schedule *Schedule) ProcessOrders() {
	//While schedule is not less than end-schedule time
	//Loop through available orders, schedule the one yeilding best score
	//score is based on waittime calculated as (scheduledtime + hypotenuse dist * 2) - order timestamp
	//increment schedule time by hypotenuse distance * 2
	scheduletime := schedule.ScheduleStartTime
	scheduleendtime := schedule.ScheduleEndTime

	orders := *schedule.Orders

	//orders now becomes our priority queue, on every run, we prioritize (sort on score/waittime)
	//and pop the first element
	for scheduletime.Unix() < scheduleendtime.Unix() && len(orders) > 0 {
		//pick all available orders to prioritize
		//can start from first order as incoming orders are sorted by time
		tmporderpointer := 0
		for tmporderpointer < len(orders) && scheduletime.Unix() >= orders[tmporderpointer].Timestamp.Unix() {
			tmporderpointer++
		}

		//if no order is ready (ordertime is in future), jump to the next ready order in the file
		if tmporderpointer == 0 {
			scheduletime = orders[tmporderpointer].Timestamp
			tmporderpointer++
		}

		competingorders := orders[0:tmporderpointer]
		ordertosubmit, index := PickOrderToSubmit(scheduletime, &competingorders)
		if ordertosubmit == nil {
			break
		}
		submittedorder := *ordertosubmit
		submittedorder.ProcessedTime = scheduletime

		CalculatePromoterScore(&submittedorder)

		if schedule.Schedules != nil {
			*schedule.Schedules = append(*schedule.Schedules, submittedorder)
		} else {
			schedule.Schedules = &[]models.Order{submittedorder}
		}

		//delete the order which was selected from the orders queue
		//this is Go's variadic syntax (akin to spread operators in es6)
		orders = append(orders[:index], orders[index+1:]...)
		//increment schedule time

		scheduletime = scheduletime.Add(time.Second * time.Duration(submittedorder.Distance*60))
	}

}

//CalculatePromoterScore calculates the global promoter score
func (schedule *Schedule) CalculatePromoterScore() {
	if schedule.Schedules != nil && len(*schedule.Schedules) > 0 {
		numorders := len(*schedule.Orders)
		numschedules := len(*schedule.Schedules)
		numpromoters := 0
		numdetractors := 0
		for _, sch := range *schedule.Schedules {
			if sch.IsPromoter {
				numpromoters++
			}
			if sch.IsDetractor {
				numdetractors++
			}
		}
		//add any unprocessed orders to detractors
		numdetractors = numdetractors + numorders - numschedules
		fmt.Println("***************************************************************************************")
		fmt.Printf(" Promoters: %d, Detractors (Including Incomplete Orders): %d", numpromoters, numdetractors)
		promoterpercentage := float64(float64(numpromoters)/float64(numorders)) * 100.0
		detractorpercentage := float64(float64(numdetractors)/float64(numorders)) * 100.0
		schedule.NetPromoterScore = promoterpercentage - detractorpercentage
		fmt.Printf("\n The net promoter score is %f%%", schedule.NetPromoterScore)
		fmt.Printf("\n Schedules written to: %s\n", schedule.OutputFilePath)
		fmt.Println("***************************************************************************************")
	} else {
		fmt.Println("No schedules to calculate scores on...")
	}
}
