package main

import (
	"fmt"
	"reflect"
)

type WasRun struct {
	WasRun bool
	Tc     TestCase
}

func NewWasRun(name string) WasRun {
	return WasRun{Tc: NewTestCase(name)}
}

func (w *WasRun) TestMethod() {
	fmt.Println("callsed")
	w.WasRun = true
}

type TestCase struct {
	name string
}

func NewTestCase(method string) TestCase {
	return TestCase{name: method}
}

// Memo
// In textbook, Run() is defined in TestCase.
func (w *WasRun) Run() {
	method := reflect.ValueOf(w).MethodByName(w.Tc.name)
	method.Call(nil)
}
