package types

type ProtectedHandlerFunc struct {
	handler HandlerFunc
	roles   []string
}

func NewProtectedHandler(hnd HandlerFunc, roles []string) ProtectedHandlerFunc {
	return ProtectedHandlerFunc{
		handler: hnd,
		roles:   roles,
	}
}

func (ph *ProtectedHandlerFunc) GetHandler() HandlerFunc {
	return ph.handler
}

func (ph *ProtectedHandlerFunc) GetRoles() []string {
	return ph.roles
}
