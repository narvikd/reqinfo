package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func main() {
	a := fiber.New(fiber.Config{
		AppName:           "Request Info",
		EnablePrintRoutes: false,
		IdleTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	})

	a.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(ctx.Request().String())
	})

	errListen := a.Listen("0.0.0.0:3003")
	if errListen != nil {
		log.Fatalln("server can't be started:", errListen)
	}
}
