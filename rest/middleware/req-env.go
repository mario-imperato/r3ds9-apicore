package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mario-imperato/r3ds9-apicore/rest"
	"net/http"
)

type AuthInfo struct {
	Sid    string `json:"sid,omitempty" yaml:"sid,omitempty"`
	UserId string `json:"user,omitempty" yaml:"user,omitempty"`
}

type ReqEnv struct {
	ReqType  rest.ReqType
	Domain   string
	Site     string
	Lang     string
	AuthInfo AuthInfo
}

func (re ReqEnv) IsMalformed() bool {
	return re.ReqType == rest.ReqTypeMalformed
}

func (re ReqEnv) IsInvalid() bool {
	return re.ReqType == rest.ReqTypeInvalid
}

func (re ReqEnv) IsValid() bool {
	return re.ReqType != rest.ReqTypeInvalid && re.ReqType != rest.ReqTypeMalformed
}

func (re ReqEnv) String() string {
	return fmt.Sprintf("[%s] domain: %s, site: %s, lang: %s, appId: %s, extra-path: %s", re.ReqType, re.Domain, re.Site, re.Lang)
}

func (re ReqEnv) Path() string {
	return RequestPath4(re.Domain, re.Site, re.Lang)
}

func (re ReqEnv) AbortRequest(c *gin.Context, redirectUrl string) {
	c.String(http.StatusBadRequest, "Bad Request")
	c.Abort()
}

func RequestPath4(domain, site, lang string) string {
	return fmt.Sprintf("/%s/%s/%s", domain, site, lang)
}
