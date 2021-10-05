package main

import (
	"fmt"
	"os"
)

type Person struct{
	id int
	name string
	gender string
	age int
}

func (p *Person) New(id int, name string, gender string, age int) {
	if id < 0 {
		p.id = (-1) * id
	}else{
		p.id = id
	}

	p.name = name			/* simplified */
	p.gender = gender
	if age < 0 {
		p.age = (-1) * age
	}else{
		p.age = age
	}
}

func _Person(id int, name string, gender string, age int) Person{
	var p Person
	p.id = id
	p.name = name
	p.gender = gender
	p.age = age
	return p
}

func (p Person) Display() {
	fmt.Println("ID: ", p.id)
	fmt.Println("Name", p.name)
	fmt.Println("Gender: ", string(p.gender))
	fmt.Println("Age: ", p.age)
	fmt.Println("+======================")
}

func (p Person) Store(filename string) {
	content := ""
	content = content + fmt.Sprintln("ID: ", p.id)
	content = content + fmt.Sprintln("Name", p.name)
	content = content + fmt.Sprintln("Gender: ", string(p.gender))
	content = content + fmt.Sprintln("Age: ", p.age)
	content = content + fmt.Sprintln("+======================")

	file, _ := os.Create(filename)
	file.WriteString(content)
	file.Close()
}
