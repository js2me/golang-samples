package main

type Saiyan struct {
	Name string
	Power int
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name: name,
		Power: power,
	}
}


func helloworld() (string,string){
	return "hello", "world"
}


func main() {
	//goku := Saiyan{
	//	Name: "Goku",
	//	Power: 9000,
	//}
	//goku := Saiyan{}

	// или

	//goku = Saiyan{Name: "Goku"}
	//goku.Power = 9000

	//или

	//goku = Saiyan{"Goku", 9000}
	goku := &Saiyan{"Sergey",46000}
	//goku := NewSaiyan("Goku", 9001)
	goku.Super()

	_, world := helloworld()
	println("it's over 9000! ", world, ", and power goku is - ", goku.Power)
}
