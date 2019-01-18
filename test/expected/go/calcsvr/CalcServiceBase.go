// Code generated by protoapi:go; DO NOT EDIT.

package calcsvr

import (
	"github.com/labstack/echo"
	"github.com/yoozoo/protoapi/protoapigo"
)

// CalcService is the interface contains all the controllers
type CalcService interface {
	Add(c echo.Context, req *AddReq) (resp *AddResp, bizError *AddError, err error)
}

func _add_Handler(srv CalcService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(AddReq)

		if err = c.Bind(req); err != nil {

			return c.JSON(500, err)

		}
		/*

		 */
		resp, bizError, err := srv.Add(c, req)
		if err != nil {

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}

// RegisterCalcService is used to bind routers
func RegisterCalcService(e *echo.Echo, srv CalcService) {
	RegisterCalcServiceWithPrefix(e, srv, "")
}

// RegisterCalcServiceWithPrefix is used to bind routers with custom prefix
func RegisterCalcServiceWithPrefix(e *echo.Echo, srv CalcService, prefix string) {
	// switch to strict JSONAPIBinder, if using echo's DefaultBinder
	if _, ok := e.Binder.(*echo.DefaultBinder); ok {
		e.Binder = new(protoapigo.JSONAPIBinder)
	}
	e.POST(prefix+"/CalcService.add", _add_Handler(srv))
}
