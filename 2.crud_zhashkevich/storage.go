package main

import (
	"errors"
	"sync"
)

type Employee struct {
	ID     int    `json:"id"`   //	Оказывается это не одинарная кавычка
	Name   string `json:"name"` //	а английская ё, то есть там, где тильда через шифт (` или ~)
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
	GetAll() []Employee
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()

	e.ID = s.counter
	s.data[e.ID] = *e

	s.counter++

	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	employee, ok := s.data[id]
	if !ok {
		return employee, errors.New("employee not found")
	}
	return employee, nil
}

func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()
	s.data[id] = e
	s.Unlock()
}

func (s *MemoryStorage) GetAll() []Employee {
	s.Lock()
	emplSlice := make([]Employee, 0)
	for _, value := range s.data {
		emplSlice = append(emplSlice, value)
	}
	s.Unlock()
	return emplSlice
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	delete(s.data, id)
	s.Unlock()
}
