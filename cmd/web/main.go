package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KARTIKrocks/go-booking-system/pkg/config"
	"github.com/KARTIKrocks/go-booking-system/pkg/handlers"
	"github.com/KARTIKrocks/go-booking-system/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {


	// change this to true  when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // it is true for production

	app.Session = session
	
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s",PORT  )

	srv := &http.Server{
		Addr: PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

