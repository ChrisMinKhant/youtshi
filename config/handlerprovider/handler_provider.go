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

	RegisterHandler("/users", &handler.UserHandler{})
	RegisterHandler("/notify", &handler.NotifyHandler{})

}
