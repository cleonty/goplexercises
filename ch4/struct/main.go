// main.go
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Salary        int
	Position      string
	ManagerID     int
}

func EmployeeByID(id int) *Employee {
	return &Employee{
		ID:       id,
		Position: "босс",
	}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary *= 105 / 100
}

type Point struct{ X, Y int }

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func main() {
	var dilbert Employee = Employee{
		Position:  "программист",
		ManagerID: 5,
	}
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior " + *position
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (активный участник команды)"
	(*employeeOfTheMonth).Position += " (активный участник команды)"
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)
	fmt.Println(dilbert)
	EmployeeByID(dilbert.ManagerID).Salary = 0

	seen := make(map[string]struct{})
	s := "hello"
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
	}
	p := Point{1, 2}
	fmt.Println(p)
	{
		pp := &Point{1, 2}
		fmt.Println(pp)
	}
	{
		pp := new(Point)
		*pp = Point{1, 2}
		fmt.Println(pp)
	}
	a := Point{1, 2}
	b := Point{1, 2}
	fmt.Println(a == b)

	type address struct {
		host string
		port int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	hits[address{"gooogle.com", 80}]++
	fmt.Println(hits)
	{
		type Point struct {
			X, Y int
		}
		type Circle struct {
			Point
			Radius int
		}
		type Wheel struct {
			Circle
			Spikes int
		}

		var w Wheel
		w.X = 20
		w.Y = 20
		w.Radius = 50
		w.Spikes = 8
		w = Wheel{Circle{Point{1, 2}, 5}, 4}
		w = Wheel{
			Circle: Circle{
				Point: Point{
					X: 5,
					Y: 5,
				},
				Radius: 10,
			},
			Spikes: 15,
		}
		fmt.Printf("%#v\n", w)
	}

}
