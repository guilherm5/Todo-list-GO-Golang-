package autentication

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var srv *server.Server
var clientStore *store.ClientStore

func Config() {
	manager := manage.NewDefaultManager()

	// armazenamento de token
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// armazenamento de cliente
	clientStore = store.NewClientStore()
	manager.MapClientStorage(clientStore)

	srv = server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	refreshTokenCfg := &manage.RefreshingConfig{
		AccessTokenExp:     time.Hour,
		RefreshTokenExp:    1 * time.Hour,
		IsGenerateRefresh:  true,
		IsResetRefreshTime: false,
	}

	manager.SetRefreshTokenCfg(refreshTokenCfg)
}

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		srv.HandleTokenRequest(c.Writer, c.Request)
	}
}

func Credentials() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientId := uuid.New().String()
		clientSecret := uuid.New().String()
		err := clientStore.Set(clientId, &models.Client{
			ID:     clientId,
			Secret: clientSecret,
			Domain: "http://localhost:9094",
		})
		if err != nil {
			fmt.Println(err.Error())
		}

		c.Header("Content-Type", "application/json")
		c.JSON(200, gin.H{"clientId": clientId, "clientSecret": clientSecret})
	}
}

func MiddlewareAutentication() gin.HandlerFunc {

	return func(c *gin.Context) {
		_, err := srv.ValidationBearerToken(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})

		}
	}

}
