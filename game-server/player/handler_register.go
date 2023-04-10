package player

/*
HandlerRegister
@Description: 用来注册p的处理方法，到时候只需要输入字符串就能调用其方法
@receiver p
*/
func (p *Player) HandlerRegister() {
	p.handlers[111] = p.AddFriend
	p.handlers[222] = p.DelFriend
	p.handlers[333] = p.ResolveChatMsg
}
