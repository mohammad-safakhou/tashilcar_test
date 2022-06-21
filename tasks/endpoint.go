package tasks

import "tashil/tasks/push"

type scheduler struct {
	tempRepository
}

func (s *scheduler) SchedulerTaskHandler(everySecond int, request func(interface{}) error) (err error, scheduleId int) {
	// schedule{
	request()
}
