package main

import "fmt"

func mainx(){

	var maps map[int]string
	maps = map[int]string {
		1: "Hello World",
		2: "Programming in Go Lang"}

	fmt.Println(maps[1])
	maps[3] = "Mathematical Programming"
	maps[16] = "Sixteenth"

	for _, value := range maps {
		fmt.Println(value)
	}

	delete(maps, 16)
	for _, value := range maps {
		fmt.Println(value)
	}

	/*
	var data []float64
	var people []Person

	fmt.Println(data)

	data = append(data, 89.2)
	//wdata := [3]float64{32, 45, 78}
	data = append(data, []float64{32, 45, 78}...)

	fmt.Println(data)
	people = append(people, Person{5,"Dessi", "M", 35})

	for _, value := range people{
		value.Display()
		//people[index].Display()
	}
	*/

	/*
	var p Person
	p.New(1,"Jesse",'M', 32)
	q := _Person(2,"Jane",'F', 23)
	var r Person = Person{3,"Amina",'F', 33}
	p.Display()
	q.Display()
	r.Display()

	fmt.Println("Enter new Person details")
	fmt.Print("Id: ")
	fmt.Scan(&p.id)

	fmt.Print("Name: ")
	fmt.Scan(&p.name)

	fmt.Print("Gender: ")
	fmt.Scan(&p.gender)

	fmt.Print("Age: ")
	fmt.Scan(&p.age)

	p.Display()*/

	/*var pp Person
	var (
		id int
		nm string
		gn string
		gg int
	)
	fmt.Println("Enter new Person details")
	fmt.Print("Id: ")
	fmt.Scan(&id)

	fmt.Print("Name: ")
	fmt.Scan(&nm)

	fmt.Print("Gender: ")
	fmt.Scan(&gn)

	fmt.Print("Age: ")
	fmt.Scan(&gg)

	pp.New(id, nm, gn, gg)
	pp.Display()
	pp.Store("data.dat")*/
}
