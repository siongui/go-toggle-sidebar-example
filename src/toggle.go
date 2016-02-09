package main

import "github.com/gopherjs/gopherjs/js"

var la *js.Object
var ra *js.Object
var lh *js.Object
var rh *js.Object

func showSidebar(event *js.Object) {
	la.Get("classList").Call("toggle", "invisible")
	ra.Get("classList").Call("toggle", "invisible")
	rh.Get("classList").Call("toggle", "invisible")
	lh.Get("style").Set("width", "50%")
}

func hideSidebar(event *js.Object) {
	la.Get("classList").Call("toggle", "invisible")
	ra.Get("classList").Call("toggle", "invisible")
	rh.Get("classList").Call("toggle", "invisible")
	lh.Get("style").Set("width", "100%")
}

func main() {
	la = js.Global.Get("document").Call("getElementById", "left-arrow")
	ra = js.Global.Get("document").Call("getElementById", "right-arrow")
	lh = js.Global.Get("document").Call("getElementById", "left-half")
	rh = js.Global.Get("document").Call("getElementById", "right-half")

	la.Call("addEventListener", "click", showSidebar, false)
	ra.Call("addEventListener", "click", hideSidebar, false)
}
