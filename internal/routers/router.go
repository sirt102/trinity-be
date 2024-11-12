package routers

import (
	"trinity-be/internal/routers/admin"
	"trinity-be/internal/routers/public"
)

type RouterGroup struct {
		Admin admin.AdminRouterGroup
    Public public.PublicRouterGroup
}

var RouterGroupApp = new(RouterGroup)
