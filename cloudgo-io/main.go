package main

import (
	"os"

	"github.com/gitgiter/ServiceComputing/cloudgo-io/services"

	"github.com/kataras/iris"
	"github.com/spf13/pflag"
)

const (
	// PORT 8080 (default)
	PORT string = "8080"
)

func main() {

	// get and set server listening port

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := pflag.StringP("port", "p", PORT, "http listening port")
	pflag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// get iris http server app
	app := iris.Default()

	// set logger level
	app.Logger().SetLevel("debug")

	services.StartServices(app)

	// listen and serve on http://localhost:port

	// configuring by file is more convenient
	app.Run(iris.Addr(":"+port), iris.WithConfiguration(iris.TOML("./configs/main.tml")))

	// or you can config like this
	// app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.Configuration{
	// 	DisableInterruptHandler:           false,
	// 	DisablePathCorrection:             false,
	// 	EnablePathEscape:                  false,
	// 	FireMethodNotAllowed:              false,
	// 	DisableBodyConsumptionOnUnmarshal: false,
	// 	DisableAutoFireStatusCode:         false,
	// 	TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
	// 	Charset:                           "UTF-8",
	// }))
}
