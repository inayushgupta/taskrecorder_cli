package todo

import (
	"time"
)

type List struct {
	name     string
	location string
}

type Task struct {
	name     string
	priority int
	date     time.Time
}

func (list *List) GetName() string {
	return list.name
}

func (list *List) SetName(newName string) {
	list.Name = newName
}

func (list *List) GetLocation() string {
	return list.location
}

func (list *List) SetLocation(newLocation string) {
	list.Location = newLocation
}

func (task *Task) GetName() string {
	return task.name
}

func (task *Task) SetName(newName string) {
	task.name = newName
}

func (task *Task) GetPriority() int {
	return task.priority
}

func (task *Task) SetPriority(newPriority int) {
	task.priority = newPriority
}

func (task *Task) GetDate() time.Time {
	return task.date
}

func (task *Task) SetDate(newDate time.Time) {
	task.date = newDate
}
