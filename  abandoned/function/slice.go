package function

/*
CheckInNumberSlice
@Description:判断s中有没有a
@param a
@param s
@return bool
*/
func CheckInNumberSlice[T uint64 | int32](a T, s []T) bool {
	for _, val := range s {
		if a == val {
			return true
		}
	}
	return false
}

/*
DelEleInSlice
@Description:移除已有元素 只支持元素不重复场景
@param a
@param old
@return new
*/
func DelEleInSlice[T uint64 | int32](a T, old []T) (new []T) {
	for i, val := range old {
		if a == val {
			new = append(old[:i], old[i+1:]...)
			return
		}
	}
	return old
}
