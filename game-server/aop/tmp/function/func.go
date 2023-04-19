package function

import "reflect"

/*
IsNil1
@Description: 判空
@param i
@return bool
*/
func IsNil1(i interface{}) bool {
	ret := i == nil
	if !ret {
		//异常捕获
		defer func() {
			recover()
		}()
		ret = reflect.ValueOf(i).IsNil()
	}
	return ret
}

/*
IsNil2
@Description: 判空
@param i
@return bool
*/
func IsNil2(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
