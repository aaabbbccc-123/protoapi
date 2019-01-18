// Code generated by protoapi:go; DO NOT EDIT.

package apisvr

import (
	"github.com/labstack/echo"
	"github.com/yoozoo/protoapi/protoapigo"
)

// AppService is the interface contains all the controllers
type AppService interface {
	GetEnv(c echo.Context, req *EnvListRequest) (resp *EnvListResponse, bizError *Error, err error)

	RegisterService(c echo.Context, req *RegisterServiceRequest) (resp *RegisterServiceResponse, bizError *Error, err error)

	UpdateService(c echo.Context, req *UpdateServiceRequest) (resp *UpdateServiceResponse, bizError *Error, err error)

	UploadProtoFile(c echo.Context, req *UploadProtoFileRequest) (resp *UploadProtoFileResponse, bizError *Error, err error)

	GetTags(c echo.Context, req *TagListRequest) (resp *TagListResponse, bizError *Error, err error)

	GetProducts(c echo.Context, req *ProductListRequest) (resp *ProductListResponse, bizError *Error, err error)

	GetServices(c echo.Context, req *ServiceListRequest) (resp *ServiceListResponse, bizError *Error, err error)

	SearchServices(c echo.Context, req *ServiceSearchRequest) (resp *ServiceListResponse, bizError *Error, err error)

	GetKeyList(c echo.Context, req *KeyListRequest) (resp *KeyListResponse, bizError *Error, err error)

	GetKeyValueList(c echo.Context, req *KeyValueListRequest) (resp *KeyValueListResponse, bizError *Error, err error)

	SearchKeyValueList(c echo.Context, req *SearchKeyValueListRequest) (resp *KeyValueListResponse, bizError *Error, err error)

	UpdateKeyValue(c echo.Context, req *KeyValueRequest) (resp *KeyValueResponse, bizError *Error, err error)

	FetchKeyHistory(c echo.Context, req *KVHistoryRequest) (resp *KVHistoryResponse, bizError *Error, err error)
}

func _getEnv_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(EnvListRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.GetEnv(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _registerService_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(RegisterServiceRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.RegisterService(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _updateService_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(UpdateServiceRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.UpdateService(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _uploadProtoFile_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(UploadProtoFileRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.UploadProtoFile(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _getTags_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(TagListRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.GetTags(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _getProducts_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(ProductListRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.GetProducts(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _getServices_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(ServiceListRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.GetServices(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _searchServices_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(ServiceSearchRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.SearchServices(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _getKeyList_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(KeyListRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.GetKeyList(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _getKeyValueList_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(KeyValueListRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.GetKeyValueList(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _searchKeyValueList_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(SearchKeyValueListRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.SearchKeyValueList(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _updateKeyValue_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(KeyValueRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.UpdateKeyValue(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _fetchKeyHistory_Handler(srv AppService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(KVHistoryRequest)

		if err = c.Bind(req); err != nil {

			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)

		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.FetchKeyHistory(c, req)
		if err != nil {

			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}

			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}

// RegisterAppService is used to bind routers
func RegisterAppService(e *echo.Echo, srv AppService) {
	RegisterAppServiceWithPrefix(e, srv, "")
}

// RegisterAppServiceWithPrefix is used to bind routers with custom prefix
func RegisterAppServiceWithPrefix(e *echo.Echo, srv AppService, prefix string) {
	// switch to strict JSONAPIBinder, if using echo's DefaultBinder
	if _, ok := e.Binder.(*echo.DefaultBinder); ok {
		e.Binder = new(protoapigo.JSONAPIBinder)
	}
	e.POST(prefix+"/AppService.getEnv", _getEnv_Handler(srv))
	e.POST(prefix+"/AppService.registerService", _registerService_Handler(srv))
	e.POST(prefix+"/AppService.updateService", _updateService_Handler(srv))
	e.POST(prefix+"/AppService.uploadProtoFile", _uploadProtoFile_Handler(srv))
	e.POST(prefix+"/AppService.getTags", _getTags_Handler(srv))
	e.POST(prefix+"/AppService.getProducts", _getProducts_Handler(srv))
	e.POST(prefix+"/AppService.getServices", _getServices_Handler(srv))
	e.POST(prefix+"/AppService.searchServices", _searchServices_Handler(srv))
	e.POST(prefix+"/AppService.getKeyList", _getKeyList_Handler(srv))
	e.POST(prefix+"/AppService.getKeyValueList", _getKeyValueList_Handler(srv))
	e.POST(prefix+"/AppService.searchKeyValueList", _searchKeyValueList_Handler(srv))
	e.POST(prefix+"/AppService.updateKeyValue", _updateKeyValue_Handler(srv))
	e.POST(prefix+"/AppService.fetchKeyHistory", _fetchKeyHistory_Handler(srv))
}
