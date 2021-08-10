package ec2tools

import "os"

func GetHostName() string {
	host, _ := os.Hostname()
	return host
}

func GetEnv() []string {
	return os.Environ()
}
