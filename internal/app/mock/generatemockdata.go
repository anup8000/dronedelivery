package mock

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/anup8000/dronedelivery/internal/pkg/models"
)

// func GetMockSchedule(numschedules int, schedulefreq int, orderstarthour int) *schedule.Schedule {
// 	schedule := schedule.NewSchedule()
// 	initTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), orderstarthour, 00, 0, 0, time.Local)
// 	schedule.Orders = mock.GetNewData(numschedules, initTime, schedulefreq)
// 	return schedule
// }

//GetNewData function generates random rows of orders
func GetNewData(numOrders int, initTimestamp time.Time, randfrequency int, gridsize int) *[]models.Order {
	if reflect.DeepEqual(initTimestamp, time.Time{}) {
		initTimestamp = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 5, 0, 0, 0, time.Local)
	}
	var orders []models.Order
	prevTimestamp := initTimestamp
	for i := 0; i < numOrders; i++ {
		if prevTimestamp.Hour() > 20 {
			break
		}
		order := getRandomOrder(i, prevTimestamp, gridsize)
		orders = append(orders, *order)
		//timestamp
		prevTimestamp = prevTimestamp.Add(time.Second * time.Duration(rand.Intn(randfrequency)))
	}
	return &orders
}

func getRandomOrder(prevorderid int, timeStamp time.Time, gridsize int) *models.Order {
	var sb strings.Builder
	//order id
	sb.WriteString("W")
	sb.WriteString(fmt.Sprintf("%03d", prevorderid+1)) // strconv.Itoa(prevorderid + 1))
	//coordinates
	coord := getRandomDirection(1, gridsize) //passing a predictable seed
	distance := getdistancefromcoords(coord)
	return &models.Order{OrderID: sb.String(), Coordinates: coord, Timestamp: timeStamp, Distance: distance}
}

var directionNS = [2]string{"N", "S"}
var directionEW = [2]string{"E", "W"}

func getRandomDirection(seed int, gridsize int) string {
	var sb strings.Builder
	//rand.Seed(time.Now().Unix())
	rand.Seed(int64(seed))
	sb.WriteString(directionNS[rand.Intn(len(directionNS))])
	sb.WriteString(strconv.Itoa(rand.Intn(gridsize)))
	sb.WriteString(directionEW[rand.Intn(len(directionEW))])
	sb.WriteString(strconv.Itoa(rand.Intn(gridsize)))

	return sb.String()
}

func getdistancefromcoords(coords string) float64 {
	numcount := 0
	num := 0
	var arr [2]int

	for i, chr := range coords {
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
