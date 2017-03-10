package hbcp

type Handler struct {
	HandleMsgFunc   func(context *Context, msg Msg)
	HandleJoinFunc  func(context *Context)
	HandleCloseFunc func(context *Context)
}

func (handler *Handler) CallMsgFunc(context *Context, msg Msg) {
	if handler.HandleMsgFunc != nil {
		go handler.HandleMsgFunc(context, msg)
	}
}

func (handler *Handler) CallJoinFunc(context *Context) {
	if handler.HandleJoinFunc != nil {
		go handler.HandleJoinFunc(context)
	}
}

func (handler *Handler) CallCloseFunc(context *Context) {
	if handler.HandleCloseFunc != nil {
		go handler.HandleCloseFunc(context)
	}
}
