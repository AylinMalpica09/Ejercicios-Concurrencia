package models

import (
	"sync"
	"fmt"
)

type LetraA struct {
	letra string
	wg *sync.WaitGroup
}

func NewLetraA(lA string, wg *sync.WaitGroup) *LetraA {
	return &LetraA{
		letra: lA,
		wg: wg,
	}
}

func (lA *LetraA) ImprimirLA() {
	defer lA.wg.Done()
	fmt.Println(lA.letra)
}
