package main
//go get github.com/cespare/reflex
//reflex -g 'gin.go' -s -- sh -c 'go run gin.go'

import (
	"net/http"
	`os`
	`sync`

	"ginServer/utils"

	"github.com/gin-gonic/gin"
)

const HttpPort 		=	"8080"
const EnvFile 		=	".env"
const SetupSteps	=	3

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
func getPort() string{
	port, envPort:= os.LookupEnv("SRV_PORT")
	if envPort {return ":" + port} else {return ":" + HttpPort}
}
func main() {
	var r *gin.Engine

	//Create goroutines to setup server
	var setup = make(chan bool, SetupSteps)
	var wg sync.WaitGroup

	wg.Add(SetupSteps)
	//load env variables into os.environment
	go func() {
		defer wg.Done()
		if nil != utils.ReadEnv(EnvFile){
			setup <- false
		}else {
			setup <- true
		}
	}()
	//configure Gin and setup route
	go func() {
		defer wg.Done()
		r = configureGin()
		setup <- true
		go func() {
			defer wg.Done()
			r.GET("/", helloGin)
			setup <- true
		}()
	}()

	wg.Wait()
	close(setup)

	for x := range setup{
		println(x)
		if !x {
				println("Error in setup")
			}
	}

	println("Listening on " + getPort())
	_ = r.Run(getPort())

}



