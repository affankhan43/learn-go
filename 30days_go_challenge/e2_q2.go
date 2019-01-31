// MODIFY THIS AREA
// [filename]: e2_q2.go
// [question]: 2
// [instagram username]: affan.ak43

package main

import (
	"fmt"
	"github.com/leekchan/accounting"
)

type Project struct{
	Name string
	Monthly_Cost  int
}
type Employe struct{
	Name  string
	Salary  int
}
type Expenses interface{
	CalculateMonthly(int) int
}

func main() {
	P1:=Project{Name:"Marketing",Monthly_Cost:70000,}
	P2:=Project{Name:"CRM",Monthly_Cost:15000,}
	E1,E2,E3:=Employe{Name:"Affan",Salary:10000,},Employe{Name:"Hasan",Salary:12000,},Employe{Name:"John",Salary:5000,}
	expenses:=[]Expenses{P1,P2,E1,E2,E3}
	months:=5
	total:=CalculateTotal(expenses,months)
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println("Total expenses for 2 Projects and 3 employees in",months,"months is",ac.FormatMoney(total))
}

func(p Project) CalculateMonthly(months int) int{
	cost:= p.Monthly_Cost*months
	if cost > 50000{
		cost = cost+5000
	}
	return cost
}

func(e Employe) CalculateMonthly(months int) int{
	salary:=e.Salary*months
	if salary > 80000{
		salary = salary+8000
	}
	return salary
}

func CalculateTotal(e []Expenses,months int) int{
	total:=0
	for _,expense := range e{
		total+=expense.CalculateMonthly(months)
	}
	return total
}