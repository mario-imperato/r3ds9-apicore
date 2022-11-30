package user

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-gin/httpsrv"
	"github.com/gin-gonic/gin"
	"github.com/mario-imperato/r3ds9-apicommon/definitions"
	"github.com/mario-imperato/r3ds9-apicore/rest/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
)

func init() {
	log.Info().Msg("api endpoints init function")
	ra := httpsrv.GetApp()
	ra.RegisterGFactory(registerGroups)
}

func registerGroups(srvCtx httpsrv.ServerContext) []httpsrv.G {

	const semLogContext = "/api/site/resource/register-groups"

	var apiKey string
	if v, ok := srvCtx.GetConfig("api-key"); !ok {
		log.Error().Msg(semLogContext + " - api-key is missing")
	} else {
		if apiKey, ok = v.(string); !ok {
			log.Error().Msg(semLogContext + " - api-key is not a string")
		}
	}

	log.Info().Str("api-key", apiKey).Msg(semLogContext)

	gs := make([]httpsrv.G, 0, 2)

	gs = append(gs, httpsrv.G{
		Name:        "Api Host Site",
		Path:        definitions.ApiCorePathGroupHostSite,
		Middlewares: []httpsrv.H{middleware.RequestEnvResolver(apiKey)},
		Resources: []httpsrv.R{
			{
				Name: "HostSite",
				// Path:          "host",
				Method:        http.MethodGet,
				RouteHandlers: []httpsrv.H{getHostSite()},
			},
		},
	})

	gs = append(gs, httpsrv.G{
		Name:        "Api Site",
		Path:        definitions.ApiCorePathGroupSite,
		Middlewares: []httpsrv.H{middleware.RequestEnvResolver(apiKey)},
		Resources: []httpsrv.R{
			{
				Name:          "Site",
				Path:          ":site",
				Method:        http.MethodGet,
				RouteHandlers: []httpsrv.H{getSite()},
			},
		},
	})

	return gs
}

func getHostSite() httpsrv.H {
	return func(c *gin.Context) {
		const semLogContext = "/api/site/resource/get-host-site"
		c.String(http.StatusOK, "host site Invoked: "+c.Param(definitions.HostSiteUrlParam))
	}
}

func getSite() httpsrv.H {
	return func(c *gin.Context) {
		const semLogContext = "/api/site/resource/get-site"
		c.String(http.StatusOK, "site Invoked: "+c.Param("site"))
	}
}
