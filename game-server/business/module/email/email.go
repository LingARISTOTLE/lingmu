package email

/*
Mail
@Description: 邮箱接口
*/
type Mail interface {
	ToPB()
	LoadFrom()
	GetDBModel()
}
