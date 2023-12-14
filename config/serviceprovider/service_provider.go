package serviceprovider

/*
* ' BindService() ' is the place where you can
* bind all the services which implement service interface.
**/

func BindService() {

	/*
	* Bind Service
	**/

	RegisterService("TestService", nil)
}
