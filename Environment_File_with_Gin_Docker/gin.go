package main
//go get github.com/cespare/reflex
//reflex -g 'gin.go' -s -- sh -c 'go run gin.go'

import (
	`fmt`
	"net/http"
	`os`

	"ginServer/utils"

	"github.com/gin-gonic/gin"
)

func helloGin(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

//set default gin config, return a *gin.Engine
func configureGin() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	r := gin.Default()
	return r
}

func main() {
	//load env variables into os.environment
	if nil != utils.ReadEnv(".env"){
		return
	}

	r := configureGin()

	r.GET("/", helloGin)

	port:= ":" + os.Getenv("SRV_PORT")
	fmt.Println("Listening on " + port)

	_ = r.Run(port)

}

