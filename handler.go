package hbcp

type Handler struct {
	OnMsg   func(context *Context, msg Msg)
	OnJoin  func(context *Context)
	OnClose func(context *Context)
}

func (handler *Handler) EmitMsg(context *Context, msg Msg) {
	if handler.OnMsg != nil {
		go handler.OnMsg(context, msg)
	}
}

func (handler *Handler) EmitJoin(context *Context) {
	if handler.OnJoin != nil {
		go handler.OnJoin(context)
	}
}

func (handler *Handler) EmitClose(context *Context) {
	if handler.OnClose != nil {
		go handler.OnClose(context)
	}
}
