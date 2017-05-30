package gollection

func New(slice interface{}) *Gollection{
	return &Gollection{
		slice: convertSliceToSliceOfInterfaces(slice),
	}
}

func Filter(sliceParam interface{}, whereFuncParam interface{}) []interface{} {
	slice := convertSliceToSliceOfInterfaces(sliceParam)
	return filter(slice, whereFuncParam)
}

func Each(sliceParam interface{}, eachFuncParam interface{}) []interface{} {
	slice := convertSliceToSliceOfInterfaces(sliceParam)
	return each(slice, eachFuncParam)
}

func Map(sliceParam interface{}, eachFuncParam interface{}) []interface{} {
	slice := convertSliceToSliceOfInterfaces(sliceParam)
	return mapp(slice, eachFuncParam)
}

func Reduce(sliceParam interface{}, reduceIterationFuncParam interface{}, aggregatorParam interface{}) interface{} {
	slice := convertSliceToSliceOfInterfaces(sliceParam)
	return reduce(slice, reduceIterationFuncParam, aggregatorParam)
}