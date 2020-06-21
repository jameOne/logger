/*
Package logger defines exported wrapper functions
allowing a single file access to all of the logging
utilites.
*/
package logger

import (
	"log"

	"github.com/jameOne/psm/logger/action"
	"github.com/jameOne/psm/logger/error"
	"github.com/jameOne/psm/logger/input"
	"github.com/jameOne/psm/logger/output"
	"github.com/jameOne/psm/logger/update"
	"github.com/jameOne/psm/logger/yield"
)

type actionLoggerSymbol struct {
	s action.Logger
}

// Action logs the type `action.Logger` defined in logutil.
func Action(a action.Logger) {
	log.SetPrefix("A:")
	log.Printf("%s:", a.Action())
}

type errorLoggerSymbol struct {
	s error.Logger
}

// Error logs the type `error.Error` defined in logutil.
func Error(e error.Logger) {
	log.SetPrefix("E:")
	log.Printf("%s:", e.Error())
}

type inputLoggerSymbol struct {
	s input.Logger
}

// Input logs the type `input.IO` defined in logutil.
func Input(i input.Logger) {
	log.SetPrefix("I:")
	log.Printf("%s:", i.IO())
}

type outputLoggerSymbol struct {
	s output.Logger
}

// Output logs the type `output.Output` defined in logutil.
func Output(o output.Logger) {
	// Doc
	log.SetPrefix("O:")
	log.Printf("%s:", o.Output())
}

type updateLoggerSymbol struct {
	s update.Logger
}

// Update logs the type `update.Update` defined in logutil.
func Update(u update.Logger) {
	log.SetPrefix("U:")
	log.Printf("%s:", u.Update())
}

type yieldLoggerSymbol struct {
	s yield.Logger
}

// Yield logs the type `yield.Yield` defined in logutil.
func Yield(y yield.Logger) {
	log.SetPrefix("Y:")
	log.Fatalf("%s:", y.Yield())
}
