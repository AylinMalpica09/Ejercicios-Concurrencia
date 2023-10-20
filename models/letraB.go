package models

import (
	"sync"
	"fmt"
)

type LetraB struct {
	letra string
	wg *sync.WaitGroup
}

func NewLetraB(lB string, wg *sync.WaitGroup) *LetraB {
	return &LetraB{
		letra: lB,
		wg: wg,
	}
}

func (lB *LetraB) ImprimirLB() {
	defer lB.wg.Done()
	fmt.Println(lB.letra)
}
