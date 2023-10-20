package models

import (
	"sync"
	"fmt"
)

type LetraC struct {
	letra string
	wg *sync.WaitGroup
}

func NewLetraC(lC string, wg *sync.WaitGroup) *LetraC {
	return &LetraC{
		letra: lC,
		wg: wg,
	}
}

func (lC *LetraC) ImprimirLC() {
	defer lC.wg.Done()
	fmt.Println(lC.letra)
}
