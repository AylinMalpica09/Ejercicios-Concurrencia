package models

import (
	"fmt"
)
	


type Letter struct{
	a chan bool
	b chan bool
	c chan bool
	f chan bool
}

func NewLetter (a chan bool, b chan bool, c chan bool, f chan bool) * Letter {
	return &Letter {
		a:a,
		b:b,
		c:c,
		f:f,
	}
}

func (l *Letter) PrintA(){
	for i:=1; i<10; i++{
		fmt.Print("A")
		l.b<- true
		<-l.a
	}
}

func (l *Letter) PrintB(){
	for i:=1; i<10; i++{
		<-l.b
		fmt.Print("B")
		l.c<- true
	}
}

func (l *Letter) PrintC(){
	for i:=1; i<10; i++{
		<-l.c
		fmt.Print("C\n")
		l.a<- true
	}
	l.f<- true
}