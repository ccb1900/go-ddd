package http

import "goddd/internal/ports/http/v0/handler"

func ProvideV0Routers(
    bh *handler.BookHandler,
) []IRegisterRoute {
    return []IRegisterRoute{bh}
}