package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HTMLRendering(app *fiber.App) {

	var (
		siteDomain      = "kewMusics.ru"
		siteDescription = "An example template"
		fileServer      = "http://localhost:3001"
		metaDescription = "metaDescription"
	)

	app.Get("/", func(c *fiber.Ctx) error {
		//var musics []models.Music
		//database.MusicDB.Find(&musics) // Fetch audio files from DB

		// Render index template
		return c.Render("index", fiber.Map{
			"siteDomain":      siteDomain,
			"fileServer":      fileServer,
			"pageTitle":       "pageTitle",
			"siteDescription": siteDescription,
			"metaDescription": "metaDescription",
			"disClaimer":      "Наш сайт не нарушает не кокой законы.",
			//"newMusics":       musics,
			//"trendMusics":     musics,
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"siteDomain":      siteDomain,
			"fileServer":      fileServer,
			"metaDescription": "metaDescription",
		})
	})

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{
			"siteDomain":      siteDomain,
			"fileServer":      fileServer,
			"metaDescription": "metaDescription",
		})
	})

	app.Get("/profile", func(c *fiber.Ctx) error {
		return c.Render("profile", fiber.Map{
			"siteDomain":      siteDomain,
			"fileServer":      fileServer,
			"metaDescription": metaDescription,
		})
	})

	app.Get("/admin/dashboard", func(c *fiber.Ctx) error {
		return c.Render("dashboard", fiber.Map{
			"siteDomain": siteDomain,
			"fileServer": fileServer,
		})
	})

	app.Get("/rules", func(c *fiber.Ctx) error { // Согласие с правила на сайте
		return c.Render("rules", fiber.Map{
			"siteDomain": siteDomain,
			"fileServer": fileServer,
		})
	})

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", fiber.Map{
			"siteDomain": siteDomain,
			"fileServer": fileServer,
		})
	})

	app.Get("/createBlog", func(c *fiber.Ctx) error {
		return c.Render("createBlog", fiber.Map{
			"siteDomain": siteDomain,
			"fileServer": fileServer,
		})
	})

	//app.Get("/musics", func(c *fiber.Ctx) error {
	//	return c.Render("getMusics", fiber.Map{
	//		"siteDomain": siteDomain,
	//		"fileServer": fileServer,
	//	})
	//})

}
