package common


func GetHostIP() (hostIP string) {
	hostIP = "192.168.1.0" // TODO: get from env
	return hostIP
}

func GetHostPort() (hostPort string) {
	hostPort = "8092"
	return hostPort
}

func GetAgentPort() (agentPort string) {
	agentPort = "8093"
	return agentPort
}
