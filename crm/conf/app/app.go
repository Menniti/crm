package app

import (
	mcache "github.com/go-macaron/cache"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/jade"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"github.com/menniti/crm/conf"
	"github.com/menniti/crm/handler"
	"github.com/menniti/crm/lib/cache"
	"github.com/menniti/crm/lib/contx"
	"github.com/menniti/crm/lib/cors"
	"github.com/menniti/crm/lib/template"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/macaron.v1"
)

//SetupMiddlewares configures the middlewares using in each web request
func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(gzip.Gziper())
	app.Use(toolbox.Toolboxer(app, toolbox.Options{
		HealthCheckers: []toolbox.HealthChecker{
			new(handler.AppChecker),
		},
	}))
	app.Use(macaron.Static("public"))
	app.Use(i18n.I18n(i18n.Options{
		Directory: "locale",
		Langs:     []string{"pt-BR", "en-US"},
		Names:     []string{"Português do Brasil", "American English"},
	}))
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	app.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	//Cache in memory
	app.Use(mcache.Cacher(
		cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
	))
	/*
		Redis Cache
		Add this lib to import session: _ "github.com/go-macaron/cache/redis"
		Later replaces the cache in memory instructions for the lines below
		optCache := mcache.Options{
				Adapter:       conf.Cfg.Section("").Key("cache_adapter").Value(),
				AdapterConfig: conf.Cfg.Section("").Key("cache_adapter_config").Value(),
			}
		app.Use(mcache.Cacher(optCache))
	*/
	app.Use(session.Sessioner())
	app.Use(contx.Contexter())
	app.Use(cors.Cors())
}

//SetupRoutes defines the routes the Web Application will respond
func SetupRoutes(app *macaron.Macaron) {
	app.Get("", func() string {
		return "Mercurius Works!"
	})

	//HealthChecker
	app.Get("/health", handler.HealthCheck)

	//Prometheus metrics
	app.Get("/metrics", prometheus.Handler())

	//User
	app.Group("/user", func() {
		app.Get("/:id", handler.GetUser)
		app.Get("/all", handler.GetAllUsers)
		app.Post("/new", handler.NewUser)
		app.Put("/update/:id", handler.UpdateUser)
		app.Delete("/delete/:id", handler.DeleteUser)
	})

	//Company
	app.Group("/company", func() {
		app.Get("/:id", handler.GetCompany)
		app.Get("/all", handler.GetAllCompanies)
		app.Post("/new", handler.NewCompany)
		app.Put("/update/:id", handler.UpdateCompany)
		app.Delete("/delete/:id", handler.DeleteCompany)
	})

	app.Group("/deal", func() {
		app.Get("/:id", handler.GetDeal)
		app.Get("/all", handler.GetAllDeals)
		app.Post("/new", handler.NewDeal)
		app.Put("/update/:id", handler.UpdateDeal)
		app.Delete("/delete/:id", handler.DeleteDeal)
	})

	app.Group("/client", func() {
		app.Get("/:id", handler.GetClient)
		app.Get("/all", handler.GetAllClients)
		app.Post("/new", handler.NewClient)
		app.Put("/update/:id", handler.UpdateClient)
		app.Delete("/delete/:id", handler.DeleteClient)
	})

	//Include this struct after import session
	type Teste struct {
		Status string
	}

	//An example to test DB connection

	app.Get("/teste", func() string {
		//	db, err := conf.GetDB()
		//	if err != nil {
		//		return err.Error()
		//	}
		//	err = db.Ping()
		//	if err != nil {
		//		return err.Error()
		//	}
		col, err := conf.GetMongoCollection("teste")
		if err != nil {
			return err.Error()
		}
		defer col.Database.Session.Close()
		teste := Teste{Status: "OK"}
		err = col.Insert(teste)
		return "Mercurius Works!"
	})

}
