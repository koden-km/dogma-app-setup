package dogmatiqapp

import (
	"github.com/dogmatiq/dogma"
	"github.com/koden-km/dogma-app-setup/customer"
)

// App is an implementation of dogma.Application for the practice app.
type App struct {
	customerAggregate customer.Aggregate
}

// Configure configures the Dogma engine for this application.
func (a *App) Configure(c dogma.ApplicationConfigurer) {
	c.Name("dogmaticapp")
	c.RegisterAggregate(a.customerAggregate)
}
