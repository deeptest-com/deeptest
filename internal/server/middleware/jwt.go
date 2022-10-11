package middleware

import (
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/multi"
)

func JwtHandler() iris.Handler {
	verifier := multi.NewVerifier()
	verifier.Extractors = []multi.TokenExtractor{multi.FromHeader} // extract token only from Authorization: Bearer $token
	verifier.ErrorHandler = func(ctx *context.Context, err error) {
		//ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(_domain.Response{
			Code: _domain.AuthErr.Code,
		})
		// ctx.StopWithError(http.StatusUnauthorized, err)
	} // extract token only from Authorization: Bearer $token
	return verifier.Verify()
}
