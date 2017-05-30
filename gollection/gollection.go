package gollection

type(
	Gollection struct {
		slice []interface{}
	}
)


func (g *Gollection) Filter(whereFunc interface{}) *Gollection{
	g.slice = filter(g.slice, whereFunc)
	return g
}

func (g *Gollection) Each(eachFunc interface{}) *Gollection{
	each(g.slice, eachFunc)
	return g
}

func (g *Gollection) Map(mapFunc interface{}) *Gollection{
	g.slice = mapp(g.slice, mapFunc)
	return g
}

//Collect returns the underlying slice, which has likely been modified by Filter, Reduce, etc
func (g *Gollection) Collect() []interface{}{
	return g.slice
}



