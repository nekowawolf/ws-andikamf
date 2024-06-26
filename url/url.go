package url

import (
	"github.com/nekowawolf/ws-andikamf/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage)
	page.Get("/presensi", controller.GetPresensi)
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Post("/insert", controller.InsertDataPresensi)
	page.Get("/docs/*", swagger.HandlerDefault)
}
