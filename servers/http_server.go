/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package servers

import (
	"strconv"

	"github.com/mainflux/mainflux/controllers"
	"github.com/mainflux/mainflux/config"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)


func HttpServer(cfg config.Config) {
	// Iris config
	iris.Config.DisableBanner = true

	// set the global middlewares
	iris.Use(logger.New())

	// set the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
	})

	// register public API
	registerRoutes()

	// start the server
	iris.Listen(cfg.HttpHost + ":" + strconv.Itoa(cfg.HttpPort))

	// Use following to start HTTPS server on the same port
	//iris.ListenTLS(cfg.HttpHost + ":" + strconv.Itoa(cfg.HttpPort), "tls/mainflux.crt", "tls/mainflux.key")
}

func registerRoutes() {
	// STATUS
	iris.Get("/status", controllers.GetStatus)

	// DEVICES
	iris.Post("/devices", controllers.CreateDevice)
	iris.Get("/devices", controllers.GetDevices)

	iris.Get("/devices/:device_id", controllers.GetDevice)
	iris.Put("/devices/:device_id", controllers.UpdateDevice)

	iris.Delete("/devices/:device_id", controllers.DeleteDevice)

	// CHANNELS
	iris.Post("/channels", controllers.CreateChannel)
	iris.Get("/channels", controllers.GetChannels)

	iris.Get("/channels/:channel_id", controllers.GetChannel)
	iris.Put("/channels/:channel_id", controllers.UpdateChannel)

	iris.Delete("/channels/:channel_id", controllers.DeleteChannel)
}
