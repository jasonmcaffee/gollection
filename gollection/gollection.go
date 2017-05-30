package gollection

type(
	//Gollection is a wrapper for a given slice, providing convenience functions for filtering, reducing, etc
	Gollection struct {
		slice []interface{}
	}
)

//Filter allows us to include only the items that the whereFunc returns true for.
//e.g. g.New([]string{"a", "b"}).Filter(func(s string){ return s == "a"}).Collect()
// would result in a slice having only 1 element, with a value of "a"
func (g *Gollection) Filter(whereFunc interface{}) *Gollection{
	g.slice = filter(g.slice, whereFunc)
	return g
}

//Each iterates over each item in the slice, passing it to the eachFunc param.
//e.g. g.New([]string{"a", "b"}).Each(func(s string){ fmt.Println(s) })
//would print "a" and "b"
func (g *Gollection) Each(eachFunc interface{}) *Gollection{
	each(g.slice, eachFunc)
	return g
}

//Map iterates over each item in the slice, passing it to the mapFunc, which is expected to return a converted value.
//e.g. g.New([]string{"john", "sally"}).Map(func(s string) Person{ return Person{Name:s} }).Each(func(p Person){ fmt.Println(p.Name)})
//would print "john" and "sally"
func (g *Gollection) Map(mapFunc interface{}) *Gollection{
	g.slice = mapp(g.slice, mapFunc)
	return g
}

//Reduce combines items in a slice into a single value (aggregator)
func (g *Gollection) Reduce(reduceFunc interface{}, aggregator interface{}) interface{} {
	return reduce(g.slice, reduceFunc, aggregator)
}

//Collect returns the underlying slice, which has likely been modified by Filter, Reduce, etc
//for typed slices, use CollectAs
func (g *Gollection) Collect() []interface{}{
	return g.slice
}

//CollectAs passes the underlying slice to the collectFunc as the type the collectFunc is expecting as its first param
//e.g. g.New([]interface{"a"}).CollectAs(func(strings []string){ var result []strings = strings })
func (g *Gollection) CollectAs(collectFunc interface{}) *Gollection{
	collectAs(g.slice, collectFunc)
	return g
}



