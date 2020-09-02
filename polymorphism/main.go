/*DEMONSTRATION OF HOW WE CAN ACHIEVE POLYMORPHISM IN GOLANG
 */

package main

import "fmt"

// Income ...
/*so this interface will gonna have two methods, those struct define
these methods the will implement this interface
so using this interface we will calculate monthly income for different struct*/
type Income interface {
	SourceName() string
	MonthlyIncome() int
}

/*let's say a programmer have two sets of income*/

// Freelance ...
type Freelance struct {
	Name               string
	HourlyRate         int
	TotalHoursPerMonth int
}

// SourceName ...
func (f Freelance) SourceName() string {
	return f.Name
}

// MonthlyIncome ...
func (f Freelance) MonthlyIncome() int {
	return f.HourlyRate * f.TotalHoursPerMonth
}

// Blogs ...
// this project has fixed income
type Blogs struct {
	Name           string
	IncomePerMonth int // rupees
}

// SourceName ...
func (b Blogs) SourceName() string {
	return b.Name
}

// MonthlyIncome ...
func (b Blogs) MonthlyIncome() int {
	return b.IncomePerMonth
}

// calculateNetIncome ...
func calculateNetIncome(inc []Income) {
	netIncome := 0
	for i := range inc {
		netIncome += inc[i].MonthlyIncome()
	}
	fmt.Println(netIncome)
}

// we want to calculate the net income of this dev.
func main() {
	project1 := Freelance{Name: "Project 1", HourlyRate: 500, TotalHoursPerMonth: 40}
	project2 := Blogs{Name: "Project 2", IncomePerMonth: 10000}
	incomeStreams := []Income{project1, project2}
	calculateNetIncome(incomeStreams)
}
