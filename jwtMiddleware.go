package types

type JwtMiddleware func(HandlerFunc, []string) HandlerFunc
