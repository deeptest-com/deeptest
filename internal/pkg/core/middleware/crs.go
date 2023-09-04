package middleware

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"
)

// CrsAuth 跨域中间件
func CrsAuth(app string) context.Handler {
	origins := make([]string, 0)
	if app == "server" {
	} else if app == "agent" {
	}

	return cors.New(cors.Options{
		AllowedOrigins: origins,
		AllowedMethods: []string{
			consts.GET.String(), consts.POST.String(), consts.PUT.String(), consts.DELETE.String(),
			consts.PATCH.String(), consts.HEAD.String(), consts.CONNECT.String(), consts.OPTIONS.String(),
			consts.TRACE.String(),
		},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
}
