package serviceprovider

import "log"

/*
* As a name ' handler repository ' this is where all
* the binded handler are stored.
* Handlers are binded to their respective path which is like mapping.
 */

var ServiceMap = make(map[string]any)

func RegisterService(name string, registeredSerive any) {
	log.Printf("Registering Service : %v, %v", name, registeredSerive)
	ServiceMap[name] = registeredSerive
}

func GetService(serviceName string) any {
	return ServiceMap[serviceName]
}
