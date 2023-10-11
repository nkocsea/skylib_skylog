package main

import (
	"github.com/nkocsea/skylib_skylog/skylog"
)

type A struct {
	Test string `json:"test"`
}

func main() {
	skylog.SetLogFile("app")
	a := A{Test: "test"}
	skylog.Info("Log Info", a)
	skylog.Error("Log Error", a)
	skylog.DetailInfo("Log DetailInfo", a)
	skylog.DetailError("Log DetailError", a)
}
