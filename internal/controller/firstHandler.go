package controller

import (
	"fmt"
	"net/http"
	"v1/util"
)

func HandleTest(w http.ResponseWriter, r *http.Request) {
	value := util.GetEvnValue("SECOND_TEST_KEY")

	fmt.Fprint(w, value)
}
