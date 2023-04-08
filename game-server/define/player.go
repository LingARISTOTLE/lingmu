package define

/*
HandlerParam
@Description: 用来作为通道获取操作的命令（也就是获取key到handler_register中找方法来调用）
*/
type HandlerParam struct {
	HandlerKey string      //处理方法
	Data       interface{} //待处理数据（比如一个用户player）
}
