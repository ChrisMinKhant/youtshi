package serviceprovider

/*
* As a name ' service repository ' this is where all
* the binded services are stored.
 */

var ServiceMap = make(map[string]any)

func RegisterService(serviceName string, registeredService any) {
	ServiceMap[serviceName] = registeredService
}

func GetService(serviceName string) any {
	return ServiceMap[serviceName]
}
