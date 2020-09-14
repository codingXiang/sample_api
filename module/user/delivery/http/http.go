package http

import (
	"github.com/astaxie/beego/validation"
	cx "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/cxgateway/pkg/i18n"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"sample_api/model"
	"sample_api/module/user"
	"sample_api/module/user/delivery"
)

const (
	MODULE          = "user"
)

type UserHttpHandler struct {
	i18nMsg i18n.I18nMessageHandlerInterface
	gateway cx.HttpHandler
	svc     user.Service
}

func NewUserHttpHandler(gateway cx.HttpHandler, svc user.Service) delivery.HttpHandler {
	var handler = &UserHttpHandler{
		i18nMsg: i18n.NewI18nMessageHandler(MODULE),
		gateway: gateway,
		svc:     svc,
	}
	/*
		v1 版本的 User API
	*/
	v1 := gateway.GetApiRoute().Group("/v1/user")
	v1.GET("", e.Wrapper(handler.GetUserList))
	v1.GET("/:id", e.Wrapper(handler.GetUser))
	v1.POST("", e.Wrapper(handler.CreateUser))
	v1.PUT("/:id", e.Wrapper(handler.UpdateUser))
	v1.PATCH("/:id", e.Wrapper(handler.ModifyUser))
	v1.DELETE("/:id", e.Wrapper(handler.DeleteUser))

	return handler
}

func (g *UserHttpHandler) GetUserList(c *gin.Context) error {
	var (
		data = map[string]interface{}{}
	)
	g.i18nMsg.SetModule(MODULE)
	g.i18nMsg.SetCore(util.GetI18nData(c))

	//抓取 query string

	if in, isExist := c.GetQuery("id"); isExist {
		data["id"] = in
	}
	if in, isExist := c.GetQuery("email"); isExist {
		data["email"] = in
	}
	if in, isExist := c.GetQuery("phone"); isExist {
		data["phone"] = in
	}
	if result, err := g.svc.GetUserList(data); err != nil {
		return g.i18nMsg.GetError(err)
	} else {
		c.JSON(g.i18nMsg.GetSuccess(result))
		return nil
	}
}
func (g *UserHttpHandler) GetUser(c *gin.Context) error {
	var (
		data = new(model.User)
	)
	//將 middleware 傳入的 i18n 進行轉換
	g.i18nMsg.SetModule(MODULE)
	g.i18nMsg.SetCore(util.GetI18nData(c))
	data.ID = c.Params.ByName("id")
	if result, err := g.svc.GetUser(data); err != nil {
		return g.i18nMsg.GetError(err)
	} else {
		c.JSON(g.i18nMsg.GetSuccess(result))
	}
	return nil
}
func (g *UserHttpHandler) CreateUser(c *gin.Context) error {
	var (
		valid = new(validation.Validation)
		data  = new(model.User)
	)
	//將 middleware 傳入的 i18n 進行轉換
	g.i18nMsg.SetModule(MODULE)
	g.i18nMsg.SetCore(util.GetI18nData(c))
	//綁定參數
	var err = c.ShouldBindWith(&data, binding.JSON)
	if err != nil || data == nil {
		return g.i18nMsg.ParameterFormatError()
	}

	//驗證表單資訊是否填寫充足
	valid.Required(&data.ID, "id")
	valid.Required(&data.Email, "email")

	if err := util.NewRequestHandler().ValidValidation(valid); err != nil {
		return err
	}

	if result, err := g.svc.CreateUser(data); err != nil {
		return g.i18nMsg.CreateError(err)
	} else {
		c.JSON(g.i18nMsg.CreateSuccess(result))
		return nil
	}
}
func (g *UserHttpHandler) UpdateUser(c *gin.Context) error {
	var (
		data = new(model.User)
		err  error
	)
	//將 middleware 傳入的 i18n 進行轉換
	g.i18nMsg.SetModule(MODULE)
	g.i18nMsg.SetCore(util.GetI18nData(c))
	data.ID = c.Params.ByName("id")
	//取得 tenant
	if data, err = g.svc.GetUser(data); err != nil {
		return g.i18nMsg.GetError(err)
	}

	//綁定參數
	err = c.ShouldBindWith(data, binding.JSON)
	if err != nil || data == nil {
		return g.i18nMsg.ParameterFormatError()
	}

	//更新 tenant
	if result, err := g.svc.UpdateUser(data); err != nil {
		return g.i18nMsg.UpdateError(err)
	} else {
		c.JSON(g.i18nMsg.UpdateSuccess(result))
		return nil
	}
}
func (g *UserHttpHandler) ModifyUser(c *gin.Context) error {
	var (
		data       = new(model.User)
		updateData = new(map[string]interface{})
	)
	//將 middleware 傳入的 i18n 進行轉換
	g.i18nMsg.SetModule(MODULE)
	g.i18nMsg.SetCore(util.GetI18nData(c))
	data.ID = c.Params.ByName("id")

	//綁定參數
	err := c.ShouldBindWith(&updateData, binding.JSON)
	if err != nil || data == nil {
		return g.i18nMsg.ParameterFormatError()
	}

	if result, err := g.svc.ModifyUser(data, *updateData); err != nil {
		return g.i18nMsg.ModifyError(err)
	} else {
		c.JSON(g.i18nMsg.ModifySuccess(result))
		return nil
	}
}
func (g *UserHttpHandler) DeleteUser(c *gin.Context) error {
	var (
		data = new(model.User)
	)
	//將 middleware 傳入的 i18n 進行轉換
	g.i18nMsg.SetModule(MODULE)
	g.i18nMsg.SetCore(util.GetI18nData(c))
	data.ID = c.Params.ByName("id")

	if err := g.svc.DeleteUser(data); err != nil {
		return g.i18nMsg.DeleteError(err)
	} else {
		c.JSON(g.i18nMsg.DeleteSuccess(nil))
		return nil
	}
}

