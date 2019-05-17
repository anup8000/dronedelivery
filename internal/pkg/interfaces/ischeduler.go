package interfaces

//IScheduler is the interface for scheduling
type IScheduler interface {
	ReadOrders()
	WriteSchedules()
	ProcessOrders()
	CalculatePromoterScore()
}
