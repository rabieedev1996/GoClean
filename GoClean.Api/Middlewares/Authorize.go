package Middlewares

import (
	"GoClean/GoClean.Application/Commons"
	"GoClean/GoClean.Application/Contract/Services"
	"GoClean/GoClean.Application/Model"
	GoClean_Identity "GoClean/GoClean.Identity"
	"github.com/gin-gonic/gin"
)

type JwtHelper struct {
	MessageService *Services.IMessageService
}

func (receiver JwtHelper) GINAuthorize(tokenUser *Model.TokenUser, context *gin.Context, tokenKey string, roles []string, reasons []string) {
	token := context.Request.Header["Authorization"][0]
	receiver.Authorize(tokenUser, token, tokenKey, roles, reasons)
}
func (receiver JwtHelper) Authorize(tokenUser *Model.TokenUser, token string, tokenKey string, roles []string, reasons []string) {
	resGen := Commons.NewResponseGenerator(receiver.MessageService)
	resObj := resGen.GetResponseObjectResult(Commons.RESULT_CODE_UNAUTHORIZED, nil)

	if &token == nil || token == "" {
		panic(resObj)
	}

	principals, err := GoClean_Identity.GetPrincipals(tokenKey, token)
	if err != nil {
		panic(resObj)
	}
	tokenResult := GoClean_Identity.ValidateToken(principals, roles, reasons)
	principalsList := GoClean_Identity.MapClaimsToKeyValue(*tokenResult.ClaimPrincipal)
	tokenUser.UserId = tokenResult.UserId

	//Checking TokenId From DB
	//tokenIdPrincipal := principalsList.FirstOrDefault(func(keyValue *GoClean_Identity.KeyValue) bool {
	//	return keyValue.Key == "TokenId"
	//})
	/*dBtoken := (*receiver.TokenRepository).GetTokenWithTokenId((*tokenIdPrincipal).Value)
	if dBtoken == nil || dBtoken.IsDeleted == true || dBtoken.Status != Mongo2.TOKEN_STATUS_ACTIVE {
		panic(resObj)
	}*/

	deviceIdPrincipal := principalsList.FirstOrDefault(func(keyValue *GoClean_Identity.KeyValue) bool {
		return keyValue.Key == "DeviceId"
	})

	if deviceIdPrincipal != nil {
		tokenUser.DeviceId = (*deviceIdPrincipal).Value
	}

	userCustomId := principalsList.FirstOrDefault(func(keyValue *GoClean_Identity.KeyValue) bool {
		return keyValue.Key == "UserCustomId"
	})
	if userCustomId != nil {
		tokenUser.UserCustomId = (*userCustomId).Value
	}

	if tokenResult.Status != "Valid" {
		panic(resObj)
	}
}
