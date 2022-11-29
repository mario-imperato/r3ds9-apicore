package middleware

import (
	"context"
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-gin/httpsrv"
	"github.com/gin-gonic/gin"
	"github.com/mario-imperato/r3ds9-apicommon/linkedservices"
	"github.com/mario-imperato/r3ds9-apicommon/linkedservices/mongodb"
	"github.com/mario-imperato/r3ds9-apicore/rest"
	"github.com/rs/zerolog/log"
	"net/http"
)

func RequestEnvResolver() httpsrv.H {

	const semLogContext = "middleware/request-api-env-resolver"
	return func(c *gin.Context) {
		env, ok := extractApiEnvFromContext(c)
		if !ok {
			c.AbortWithError(http.StatusBadRequest, errors.New("bad request"))
			return
		}

		lks, err := linkedservices.GetMongoDbService(context.Background(), "r3ds9")
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			systemError(c, err)
			return
		}

		env, ok = validateApiRequestFromStore(lks, env)
		if !ok {
			log.Error().Err(err).Msg(semLogContext)
			c.AbortWithError(http.StatusBadRequest, errors.New("bad request"))
			return
		}

		SetRequestEnvironmentInContext(c, env)
	}
}

func extractApiEnvFromContext(c *gin.Context) (ReqEnv, bool) {

	env := ReqEnv{
		Domain: c.Param("domain"),
		Site:   c.Param("site"),
		Lang:   c.Param("lang"),
	}

	env.ReqType = resolveApiRequestType(env.Domain, env.Site, env.Lang)
	return env, env.IsValid()
}

// resolveRequestType syntax level resolution. doesn't check if domain or site do really exist.
func resolveApiRequestType(domain, site, lang string) rest.ReqType {

	at := rest.ReqTypeDomains
	if domain != rest.UnassignedTargetDomain {
		if site != rest.UnassignedTargetSite {
			at = rest.ReqTypeSite
		} else {
			at = rest.ReqTypeDomain
		}
	} else if site != rest.UnassignedTargetSite {
		at = rest.ReqTypeMalformed
	}

	return at
}

func validateApiRequestFromStore(lks *mongodb.MDbLinkedService, env ReqEnv) (ReqEnv, bool) {

	const semLogContext = "middleware/request-api-env-resolver/validate-env-from-store"
	return env, true
}
