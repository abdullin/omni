package views

import "github.com/abdullin/omni/lang"

type store struct {
	all map[lang.TaskId]*taskItem
}

type taskItem struct {
	TaskId lang.TaskId
	Name   string
	Inbox  bool
}

func newStore() *store {
	return &store{make(map[lang.TaskId]*taskItem)}
}

func (s *store) addTaskToInbox(id lang.TaskId, name string, inbox bool) {
	s.all[id] = &taskItem{id, name, inbox}

}
