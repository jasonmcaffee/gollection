package gollection_test

import (
	g "gollection/gollection"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gollection", func() {

	It("should provide a mechanism for retrieving the underlying slice as a slice of concrete types", func(){
		var result []string
		g.New([]interface{}{"a", "b"}).
			CollectAs(func(strings []string){ result = strings })

		Expect(len(result)).To(Equal(2))
		Expect(result[0]).To(Equal("a"))
		Expect(result[1]).To(Equal("b"))
	})

	It("should provide a mechanism for retrieving the underlying slice", func(){
		var result2 []interface{} = g.New([]string{"a"}).
			Collect()
		Expect(len(result2)).To(Equal(1))
		Expect(result2[0].(string)).To(Equal("a"))
	})

	It("should provide slice filtering functionality", func(){
		result := g.New([]string{"a", "b"}).
			Filter(func(s string) bool {return s == "b"}).
		  Collect()

		Expect(len(result)).To(Equal(1))
		Expect(result[0].(string)).To(Equal("b"))
	})

	It("should provide a means to mapping a slice of type X to a slice of type Y", func(){
		type Person struct { Name string }
		personNames := []string{"jason", "sarah", "john"}
		var persons []Person
		g.New(personNames).
			Map(func(name string) Person{ return Person{Name:name} }).  //convert to Person
			CollectAs(func(people []Person){ persons = people })			  //get back the slice as a slice of Person

		Expect(len(persons)).To(Equal(len(personNames)))
		Expect(persons[0].Name).To(Equal(personNames[0]))
		Expect(persons[1].Name).To(Equal(personNames[1]))
		Expect(persons[2].Name).To(Equal(personNames[2]))
	})

	It("should provide a mechanism for reducing a slice to a single value", func(){
		peopleAges := []int{20, 30, 40}
		totalAge := g.New(peopleAges).Reduce(func(aggregator int, age int) int{
			return aggregator + age
		}, 0)
		totalAgeInt := totalAge.(int)
		Expect(totalAgeInt).To(Equal(90))

		peopleNames := []string{"jason", "sarah", "john"}
		namesString := g.New(peopleNames).Reduce(func(aggregator string, name string) string {
			return aggregator + name
		}, "")

		Expect(namesString.(string)).To(Equal("jasonsarahjohn"))
	})

})

var _ = Describe("Public Functions", func(){

	It("should provide a mechanism for converting slices of type X to slices of type Y", func(){
		//convert []interface{} to []string
		input := []interface{}{"a", "b"}
		var output []string
		g.CollectAs(input, func(strings []string){output = strings})

		Expect(len(output)).To(Equal(len(input)))
		Expect(output[0]).To(Equal(input[0]))
		Expect(output[1]).To(Equal(input[1]))

		input2 := []string{"a", "b"}
		var output2 []interface{}
		g.CollectAs(input2, func(interfaces []interface{}){output2=interfaces})

		Expect(len(output2)).To(Equal(len(input2)))
		Expect(output2[0].(string)).To(Equal(input2[0]))
		Expect(output2[1].(string)).To(Equal(input2[1]))

	})

  //todo: test rest of public funcs. lower priority
})