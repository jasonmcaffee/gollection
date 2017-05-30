package gollection

import "reflect"

func filter(slice []interface{}, whereFuncParam interface{}) []interface{} {
	whereFunc := reflect.ValueOf(whereFuncParam)

	results := []interface{}{} //reflect.MakeSlice(reflect.TypeOf(sliceParam), 0, 90000)
	for _, itemObj := range slice {
		paramValues := []reflect.Value{reflect.ValueOf(itemObj)}
		funcResults := whereFunc.Call(paramValues)
		shouldIncludeItem := funcResults[0].Interface().(bool) //we only expect one result
		if shouldIncludeItem {
			results = append(results, itemObj)
		}
	}
	return results
}

func each(slice []interface{}, eachFuncParam interface{}) []interface{} {
	eachFunc := reflect.ValueOf(eachFuncParam)
	results := []interface{}{}
	for _, item := range slice {
		paramValues := []reflect.Value{reflect.ValueOf(item)}
		funcResults := eachFunc.Call(paramValues)
		if len(funcResults) > 0 {
			results = append(results, funcResults[0].Interface())
		}
	}
	return results
}

func mapp(slice []interface{}, eachFuncParam interface{}) []interface{} {
	return each(slice, eachFuncParam)
}

func collectAs(slice []interface{}, collectFuncParam interface{}){
	funcValue := reflect.ValueOf(collectFuncParam)
	funcType := funcValue.Type()
	//get the type of first param
	param0SliceType := funcType.In(0)
	//create a new slice of the type the func is expecting
	convertedSlice := reflect.MakeSlice(param0SliceType, 0, len(slice))
	//iterate over each value and at it to the convertedSlice
	for i:=0; i < len(slice); i++{
		iValue := reflect.ValueOf(slice[i])
		convertedSlice = reflect.Append(convertedSlice, iValue)
	}
	//execute the collectFunc and pass it the slice of type it was expecting.
	funcValue.Call([]reflect.Value{convertedSlice})
	return
}

func reduce(slice []interface{}, reduceIterationFuncParam interface{}, aggregatorParam interface{}) interface{} {
	reduceIterationFunc := reflect.ValueOf(reduceIterationFuncParam)
	aggregator := aggregatorParam
	for _, item := range slice {
		//reduceIterationFunc should accept 2 params: (aggregator, currentItem)
		paramValues := []reflect.Value{reflect.ValueOf(aggregator), reflect.ValueOf(item)}
		funcResults := reduceIterationFunc.Call(paramValues)
		if len(funcResults) > 0 {
			aggregator = funcResults[0].Interface()
		}
	}
	return aggregator
}

//since we cant cast slice of typed entries (e.g. []int) to []interface, we must
func convertSliceToSliceOfInterfaces(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	result := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}
	return result
}
