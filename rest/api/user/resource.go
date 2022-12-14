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

	const semLogContext = "/api/user/resource/register-groups"

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
		Name:        "Api User",
		Path:        definitions.ApiCorePathGroupWhoAmI,
		Middlewares: []httpsrv.H{middleware.RequestEnvResolver(apiKey)},
		Resources: []httpsrv.R{
			{
				Name: "WhoAmI",
				//Path:          definitions.ApiCorePathPatternWhoAmI,
				Method:        http.MethodGet,
				RouteHandlers: []httpsrv.H{whoAmi()},
			},
		},
	})

	return gs
}

func whoAmi() httpsrv.H {
	return func(c *gin.Context) {
		const semLogContext = "/api/user/resource/whoAmi"
		c.String(http.StatusOK, "WhoAmi Invoked!")
	}
}
