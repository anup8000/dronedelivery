package fileops

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anup8000/dronedelivery/internal/pkg/errorhandler"
	"github.com/anup8000/dronedelivery/internal/pkg/models"
)

const basedir = "../../assets/files"
const tab = "\t"
const space = " "
const newline = "\n"

//ReadOrders is catered sepcifically to the order file
func ReadOrders(fpath string) (*[]models.Order, error) {
	if fpath == "" {
		fpath = basedir + "/orders.txt"
	}
	filefullpath := fpath
	fileHandle, _ := os.Open(filefullpath)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var orders []models.Order
	for fileScanner.Scan() {
		str := fileScanner.Text()
		arr := strings.Split(str, space)
		if len(arr) < 3 {

		} else {
			//ts, err := time.Parse("15:04:05", arr[2])
			ts, err := gettimefromstring(arr[2])
			if errorhandler.Checkandlog(err, "ReadOrders") {
				return nil, err
			}
			order := models.Order{OrderID: arr[0], Coordinates: arr[1], Timestamp: ts}
			orders = append(orders, order)
		}
	}
	return &orders, nil
}

func gettimefromstring(timestr string) (time.Time, error) {
	arr := strings.Split(timestr, ":")
	if arr == nil || len(arr) < 3 {
		return time.Time{}, errors.New("not a valid time string")
	}
	hours, err := strconv.Atoi(arr[0])
	mins, err1 := strconv.Atoi(arr[1])
	secs, err2 := strconv.Atoi(arr[2])
	if err != nil || err1 != nil || err2 != nil {
		return time.Time{}, errors.New("not a valid time string")
	}
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hours, mins, secs, 0, time.Local), nil
}

//WriteSchedules writes schedules to file
func WriteSchedules(fpath string, orders *[]models.Order, nps float64) error {
	if fpath == "" {
		fpath = basedir + "/schedule.txt"
	}
	filefullpath := fpath
	f, err := os.Create(filefullpath)
	errorhandler.Checkandpanic(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, order := range *orders {
		var sb strings.Builder
		sb.WriteString(order.OrderID)
		sb.WriteString(space)
		sb.WriteString(order.ProcessedTime.Format("15:04:05"))
		sb.WriteString(newline)
		_, err := w.WriteString(sb.String())
		if errorhandler.Checkandlog(err, "WriteSchedules") {
			return err
		}
	}
	_, err = w.WriteString(fmt.Sprintf("%.2f", nps))
	if errorhandler.Checkandlog(err, "WriteSchedules") {
		return err
	}
	w.Flush()
	return nil
}

//WriteOrders writes orders to a file
func WriteOrders(fpath string, orders *[]models.Order) error {
	if fpath == "" {
		fpath = basedir + "/orders.txt"
	}
	filefullpath := fpath
	f, err := os.Create(filefullpath)
	if errorhandler.Checkandlog(err, "WriteOrders") {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, order := range *orders {
		var sb strings.Builder
		sb.WriteString(order.OrderID)
		sb.WriteString(space)
		sb.WriteString(order.Coordinates)
		sb.WriteString(space)
		sb.WriteString(order.Timestamp.Format("15:04:05"))
		sb.WriteString(newline)
		_, err := w.WriteString(sb.String())
		if errorhandler.Checkandlog(err, "WriteOrders") {
			return err
		}
	}

	w.Flush()
	return nil
}
