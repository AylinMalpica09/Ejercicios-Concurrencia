package models

import (
	"sync"
)

type Suma struct {
	 total int 
	 wg *sync.WaitGroup
	 mutex *sync.Mutex
}

func NewSuma(t int, w *sync.WaitGroup, m *sync.Mutex) *Suma {
	return &Suma{
		total: t,
		wg: w,
		mutex: m,
	}
	
}

func (s *Suma) Increment(){
	defer s.wg.Done() 
	for i := 1; i <= 215000; i++ {
		s.mutex.Lock()
		s.total++ //seccion critica 
		s.mutex.Unlock()
	}
}

func (s *Suma) GetTotal() int{
	return s.total
}