package testrunner

import (
	"github.com/dogmatiq/dogmatest"
	dogmatiqapp "github.com/koden-km/dogma-app-setup"
)

// Runner is a test runner for the example app.
var Runner = dogmatest.New(&dogmatiqapp.App{})
