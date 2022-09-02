package server

type Middleware interface {
	PreRequest(Handler) Handler
	PostRequest(RespHandler) RespHandler
	CanSkip(string) bool
	Skip(routeName ...string)
}

type SkipMiddleware struct {
	skipRouteMap map[string]bool
}

func (m *SkipMiddleware) Skip(routeName ...string) {
	if m.skipRouteMap == nil {
		m.skipRouteMap = make(map[string]bool)
	}

	for _, name := range routeName {
		if _, ok := m.skipRouteMap[name]; ok {
			continue
		}

		m.skipRouteMap[name] = true
	}
}

func (m *SkipMiddleware) CanSkip(routeName string) bool {
	val, ok := m.skipRouteMap[routeName]

	return val && ok
}

type EmptyPostRequestMiddleware interface {
	PostRequest(*HandlerResp) *HandlerResp
}

type EmptyPostMiddleware struct {
}

func (pm *EmptyPostMiddleware) PostRequest(rh RespHandler) RespHandler {
	return rh
}
