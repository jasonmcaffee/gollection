# gollection
gollection provides commonly needed high order functions for iterating, filtering, reducing, etc slices in go.

## Gollection Usage
The gollection object provides a fluent interface for chaining multiple functions together.

```go
type Person struct { Name string }
personNames := []string{"jason", "sarah", "john"}
var persons []Person
g.New(personNames).
  Filter(func(name string) bool{ return name != "john" }).    //get rid of anyone named john
  Map(func(name string) Person{ return Person{Name:name} }).  //convert to Person
  Filter(func(p Person) bool{ return p.Name == "jason" }).    //get rid of anyone not named jason
  CollectAs(func(people []Person){ persons = people })        //get back the slice as a slice of Person
```

### Filter
The filter function allows you to filter a given slice utilizing a bool function which indicates the item in the slice should be included in the result.

```go
g.New([]string{"a", "b"}).
    Filter(func(s string) bool {return s == "b"}).
    Each(func(s string){ fmt.Println(s) })
//prints b
```
### Map
The map function provides a mechanism for mapping a slice X to a slice of type Y.

e.g. if we had a list of names, and wanted to create a new list of Person, where each Person had a name assigned to it, we could:
```go
type Person struct { Name string }
personNames := []string{"jason", "sarah", "john"}
var persons []Person
g.New(personNames).
	Map(func(name string) Person{ return Person{Name:name} }).  //convert to Person
	CollectAs(func(people []Person){ persons = people })			  //get back the slice as a slice of Person
```
### Reduce
The reduce function allows us to combine values from a given slice into a single value
```go
peopleAges := []int{20, 30, 40}
totalAge := g.New(peopleAges).Reduce(func(aggregator int, age int) int{
	return aggregator + age
}, 0)

totalAgeInt := totalAge.(int)
//would have value of 90
```

### CollectAs
Converts a slice of type X to slice of type Y.

Since Gollection Collect can only return a []interface{}, rather than the concrete type, we need to provide a mechanism for easy conversion

```go
var result []string
g.New([]interface{}{"a", "b"}).
	CollectAs(func(strings []string){ result = strings })
//result is now a slice of strings with values "a" and "b"
```

## Public Function Usage
The gollection package also provides public functions which accept slices and perform map, reduce, etc on them

### CollectAs
Converts a slice of type X to slice of type Y.

Since Gollection Collect can only return a []interface{}, rather than the concrete type, we need to provide a mechanism for easy conversion

```go
input := []interface{}{"a", "b"}
var output []string
g.CollectAs(input, func(strings []string){output = strings})

//output is now a slice of strings with values "a" and "b"
```

