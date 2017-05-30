package main

import (
	g "gollection/gollection"
	"fmt"
)


func main(){
	type Person struct {
		Name string
	}
	//ui.Body.AddRows(
	//	ui.NewRow(
	//		ui.NewCol(12, 0, par5)),
	//	ui.NewRow(
	//		ui.NewCol(12, 0, par3)),
	//	ui.NewRow(
	//		ui.NewCol(12, 0, par4)),
	//	ui.NewRow(
	//		ui.NewCol(12, 0, par)),
	//	ui.NewRow(
	//		ui.NewCol(12, 0, par2)))
	gollection := g.New([]string{"a", "b"}).
		Filter(func(s string) bool {return s == "b"}).
		Each(func(s string){
			fmt.Println(s)
	  })

	fmt.Println(len(gollection.Collect()))

	stringItems := []string{"a", "b"}
	results := g.Filter(stringItems, func(item string) bool { return item == "a" })
	g.Each(results, func(item string) {
		fmt.Println(item)
	})
	//prints "a"

	personItems := []Person{Person{Name: "jason"}, Person{Name: "joanne"}}
	results = g.Filter(personItems, func(p Person) bool { return p.Name == "jason" })
	g.Each(results, func(p Person) {
		fmt.Println(p)
	})
	//prints "jason"

	peopleNames := g.Map(personItems, func(p Person) string { return p.Name })
	g.Each(peopleNames, func(name string) {
		fmt.Println(name)
	})
	//prints "jason"
	//prints "joanne"

	namesString := g.Reduce(peopleNames, func(aggregator string, name string) string {
		return aggregator + name
	}, "")
	fmt.Println(namesString)
}
