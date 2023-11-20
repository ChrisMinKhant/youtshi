package controller

import (
	"fmt"
	"net/http"
)

func SecondHandleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is second handler")
}
