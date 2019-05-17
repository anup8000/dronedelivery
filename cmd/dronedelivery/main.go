package main

import (
	"fmt"

	"github.com/anup8000/dronedelivery/internal/app/schedule"
	"github.com/anup8000/dronedelivery/internal/pkg/errorhandler"
	"github.com/jessevdk/go-flags"
)

//Parameters define the input parameters for the program
type Parameters struct {
	InputFilePath  string `short:"i" long:"inputfilepath" description:"path of the orders file" required:"true"`
	OutputFilePath string `short:"o" long:"outputfilepath" description:"path of the output file" required:"false" default:"../../assets/testfiles/orders.txt"`
}

func main() {
	// Fetch input params
	var params Parameters
	_, err := flags.Parse(&params)
	if errorhandler.Check(err) {
		return
	}

	//Process orders and create schedule
	sch := schedule.NewSchedule(params.InputFilePath, params.OutputFilePath)
	sch.ReadOrders()
	sch.ProcessOrders()
	sch.CalculatePromoterScore()
	sch.WriteSchedules()
	fmt.Printf("\n The net promoter score is %f%% \n", sch.NetPromoterScore)
}
