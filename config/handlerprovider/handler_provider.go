package handlerprovider

import "v1/handler"

/*
* ' BindHandler() ' is the place where you can
* bind all the handlers which implement Handler interface.
**/

func BindHandler() {

	/*
	* Bind Handler
	**/

	RegisterHandler("/register", &handler.UserHandler{})

}
