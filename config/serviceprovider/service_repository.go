package serviceprovider

/*
* This is where all the names and associated services are
* actually exist.
 */
var ServiceMap = make(map[string]any)

func RegisterService(name string, registeredSerive any) {
	ServiceMap[name] = registeredSerive
}

func GetService(serviceName string) any {
	return ServiceMap[serviceName]
}
