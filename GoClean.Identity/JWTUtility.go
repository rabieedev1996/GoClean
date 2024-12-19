package GoClean_Identity

import (
	Rasa_Common "GoClean/GoClean.Common"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type KeyValue struct {
	Key   string
	Value string
}
type TokenResult struct {
	Status         string
	TokenId        string
	UserId         string
	ClaimPrincipal *jwt.MapClaims
}

const (
	TokenValidateResultNotValid       = "NotValid"
	TokenValidateResultNotValidExpire = "NotValidExpire"
	TokenValidateResultNotValidRole   = "NotValidRole"
	TokenValidateResultNotValidReason = "NotValidReason"
)

const (
	TOKEN_REASON_OTP_LOGIN                   = "TOKEN_REASON_OTP_LOGIN"
	TOKEN_REASON_FORGET_PASSWORD_OTP         = "TOKEN_REASON_FORGET_PASSWORD_OTP"
	TOKEN_REASON_FORGET_PASSWORD_CHANGE_PASS = "TOKEN_REASON_FORGET_PASSWORD_CHANGE_PASS"
	TOKEN_REASON_REGISTER                    = "TOKEN_REASON_REGISTER"
)

const (
	TOKEN_ROLE_TEMP_TOKEN      = "TOKEN_ROLE_TEMP_TOKEN"
	TOKEN_ROLE_VALIDATE_OTP    = "TOKEN_ROLE_VALIDATE_OTP"
	TOKEN_ROLE_CHANGE_PASSWORD = "TOKEN_ROLE_CHANGE_PASSWORD"
	TOKEN_ROLE_SUPER_ADMIN     = "TOKEN_ROLE_SUPER_ADMIN"
	TOKEN_ROLE_ADMIN           = "TOKEN_ROLE_ADMIN"
	TOKEN_ROLE_USER_ADMIN      = "TOKEN_ROLE_USER_ADMIN"
	TOKEN_ROLE_AUTHOR          = "TOKEN_ROLE_AUTHOR"
	TOKEN_ROLE_SYSTEM_ADMIN    = "TOKEN_ROLE_SYSTEM_ADMIN"
	TOKEN_ROLE_GENERAL_TOKEN   = "TOKEN_ROLE_GENERAL_TOKEN"
	TOKEN_ROLE_RSA_OPR         = "TOKEN_ROLE_RSA_OPR"
)

func GenerateToken(key string, userId string, ttl int64, reason string, roles []string, tokenId string, otherClaims Rasa_Common.Slice[KeyValue]) (string, error) {
	// ایجاد claims
	claims := jwt.MapClaims{}
	claims["UserId"] = userId
	claims["Expire"] = time.Now().Add(time.Millisecond * time.Duration(ttl)).Format("2006-01-02 15:04:05")

	if otherClaims != nil {
		for _, customClaim := range otherClaims {
			claims[customClaim.Key] = customClaim.Value
		}
	}

	if tokenId != "" {
		claims["TokenId"] = tokenId
	}

	if roles != nil && len(roles) > 0 {
		var roleList []string

		for _, role := range roles {
			roleList = append(roleList, role)
		}

		claims["Role"] = roleList
	}

	if reason != "" {
		claims["Reason"] = reason
	}

	securityKey := []byte(key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// اضافه کردن تاریخ انقضا
	token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(time.Millisecond * time.Duration(ttl)).Unix()

	// امضای توکن
	signedToken, err := token.SignedString(securityKey)
	if err != nil {
		return "", err
	}

	// بازگرداندن توکن
	return signedToken, nil
}

func GetPrincipals(key string, tokenStr string) (*jwt.MapClaims, error) {
	tokenStr = strings.Replace(tokenStr, "bearer ", "", -1)
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//claims := mapClaimsToKeyValue(claims)
		return &claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func ValidateToken(principal *jwt.MapClaims, validRoles []string, validReasons []string) TokenResult {

	tokenUserId, ok := (*principal)["UserId"].(string)
	if !ok {
		return TokenResult{Status: TokenValidateResultNotValid}
	}

	tokenId, _ := (*principal)["TokenId"].(string)
	expireTimeStr, _ := (*principal)["Expire"].(string)

	expireTime, err := time.Parse("2006-01-02 15:04:05", expireTimeStr)
	if err != nil || time.Now().After(expireTime) {
		return TokenResult{Status: TokenValidateResultNotValidExpire}
	}

	if len(validRoles) > 0 {
		tokenRoles := []string{}
		if roles, ok := (*principal)["Role"].([]interface{}); ok {
			for _, role := range roles {
				tokenRoles = append(tokenRoles, role.(string))
			}
		}

		validRoleFound := false
		for _, role := range tokenRoles {
			for _, validRole := range validRoles {
				if role == validRole {
					validRoleFound = true
					break
				}
			}
		}

		if !validRoleFound {
			return TokenResult{Status: TokenValidateResultNotValidRole}
		}
	}

	if len(validReasons) > 0 {
		tokenReason, _ := (*principal)["Reason"].(string)
		reasonValid := false
		for _, validReason := range validReasons {
			if tokenReason == validReason {
				reasonValid = true
				break
			}
		}

		if !reasonValid {
			return TokenResult{Status: TokenValidateResultNotValidReason}
		}
	}

	return TokenResult{Status: "Valid", TokenId: tokenId, UserId: tokenUserId, ClaimPrincipal: principal}
}

func MapClaimsToKeyValue(claims jwt.MapClaims) Rasa_Common.Slice[*KeyValue] {
	var result Rasa_Common.Slice[*KeyValue]
	for key, value := range claims {
		strValue, ok := value.(string)
		if !ok {
			strValue = fmt.Sprintf("%v", value)
		}
		result = append(result, &KeyValue{
			Key:   key,
			Value: strValue,
		})
	}
	return result
}
