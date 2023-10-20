package main


import (
	"concurrente/models"
	
)

func main() {
	//Función para suma
	/*wg := sync.WaitGroup{}

	wg.Add(2)
	m := sync.Mutex{}
	s := models.NewSuma(0, &wg, &m)

	go s.Increment()
	go s.Increment()

	wg.Wait()
	fmt.Print("Total: ", s.GetTotal() )*/

	//Funcion de orden con indeterminismo
	/*wg := sync.WaitGroup{}
	wg.Add(3)
	LA := models.NewLetraA("a", &wg)
	LB := models.NewLetraA("b", &wg)
	LC := models.NewLetraA("c", &wg)
	go LA.ImprimirLA()
	go LB.ImprimirLA()
	go LC.ImprimirLA()
	wg.Wait()*/
	
	//
	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)
	f := make(chan bool)

	letter := models.NewLetter(a,b,c,f )
	go letter.PrintA()
	go letter.PrintB()
	go letter.PrintC()
	<-f
}