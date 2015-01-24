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

func (s *store) reset() {
	s.all = make(map[lang.TaskId]*taskItem)
}

func (s *store) addTask(id lang.TaskId, name string) {
	s.all[id] = &taskItem{id, name, false}
}

func (s *store) removeTask(id lang.TaskId) {
	delete(s.all, id)
}
func (s *store) moveTaskToInbox(id lang.TaskId) {
	s.all[id].Inbox = true
}
