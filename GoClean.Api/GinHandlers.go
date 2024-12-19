package main

import (
	"GoClean/GoClean.Api/Middlewares"
	"GoClean/GoClean.Application/Commons"
	"GoClean/GoClean.Application/Contract/Services"
	"GoClean/GoClean.Application/Features/Api/Sample"
	"GoClean/GoClean.Application/Model"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindJsonToModel[T any](c *gin.Context) T {
	var result T
	if err := c.ShouldBindJSON(&result); err != nil {
	}
	//c.BindJSON(&result)
	return result
}

func BindToModel[T any](c *gin.Context) T {
	var result T
	if err := c.ShouldBind(&result); err != nil {
	}
	//c.BindJSON(&result)
	return result
}

// @Param requestBody body Sample.SampleBusinessServiceInput true "Sample data"
// @Success 200 {object}  Commons.ResponseModel[Sample.SampleBusinessServiceOutput]
// @Router /samplepost [post]
func Gin_SampleService_Post(router *gin.Engine) {
	router.POST("/samplepost", func(c *gin.Context) {

		container := GetDIContainer()
		input := BindToModel[Sample.SampleBusinessServiceInput](c)
		err := container.Invoke(func(handler *Sample.SampleBusinessServiceHandler, tokenUser *Model.TokenUser, messageService *Services.IMessageService, jwtHelper *Middlewares.JwtHelper) {

			//for authorization and fill tokenUser from authorization header
			//jwtHelper.GINAuthorize(tokenUser, c, GetProjectConfigs().TokeKey, []string{GoClean_Identity.TOKEN_ROLE_GENERAL_TOKEN}, []string{})
			result := handler.Handle(&input)
			restResult := SuccessResult(Commons.RESULT_CODE_SUCCESS, result, messageService)
			c.JSON(restResult.HTTPStatusCode, restResult.ResponseDetil)
		})
		fmt.Print(err)
	})
}

// @Success 200 {object}  Commons.ResponseModel[Sample.SampleBusinessServiceOutput]
// @Router /sampleget [get]
func Gin_SampleService_Get(router *gin.Engine) {
	router.GET("/sampleget", func(c *gin.Context) {

		container := GetDIContainer()
		input := Sample.SampleBusinessServiceInput{
			Field: c.Query("Field"),
		}
		err := container.Invoke(func(handler *Sample.SampleBusinessServiceHandler, tokenUser *Model.TokenUser, messageService *Services.IMessageService, jwtHelper *Middlewares.JwtHelper) {

			//for authorization and fill tokenUser from authorization header
			//jwtHelper.GINAuthorize(tokenUser, c, GetProjectConfigs().TokeKey, []string{GoClean_Identity.TOKEN_ROLE_GENERAL_TOKEN}, []string{})
			result := handler.Handle(&input)
			restResult := SuccessResult(Commons.RESULT_CODE_SUCCESS, result, messageService)
			c.JSON(restResult.HTTPStatusCode, restResult.ResponseDetil)
		})
		fmt.Print(err)
	})
}

func RegisterGin() *gin.Engine {

	router := gin.New()
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}
	router.Use(cors.New(config))
	router.Use(RecoveryMiddleware())
	Gin_SampleService_Get(router)
	Gin_SampleService_Post(router)
	return router

}
func SuccessResult(resultCode int, tdata any, message *Services.IMessageService) Commons.ResponseModel {
	resGen := Commons.NewResponseGenerator(message)
	return resGen.GetResponseObjectResult(resultCode, nil, tdata)
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered from panic (SocketHandler):", err)
				switch e := err.(type) {
				case Commons.ResponseModel:
					c.JSON(e.HTTPStatusCode, e.ResponseDetil)
				case *Commons.ResponseModel:
					c.JSON(e.HTTPStatusCode, e.ResponseDetil)
				default:
					c.JSON(http.StatusInternalServerError, Commons.ResponseDetail{
						ResultCode:    Commons.RESULT_CODE_EXCEPTION,
						Data:          nil,
						Message:       "Internal Server Error",
						IsSuccess:     false,
						ErrorMessages: nil,
					})
				}
				c.Abort()
			}
		}()

		c.Next()
	}
}
