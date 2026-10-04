package main

import (
	"fmt"
	"errors"
	"net/url"
	"strconv"
)

func buildRDSConnectionString(parameters map[string]string) (string, error) {
	host, ok := parameters["host"]
	if !ok || len(host) == 0 {
		return "", errors.New("host parameter is missing or empty")
	}
	port, ok := parameters["port"]
	if !ok || len(port) == 0 {
		return "", errors.New("host parameter is missing or empty")
	}
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return "", err
	}
	if portNum < 1 || portNum > 65336 {
		return "", errors.New("port number must be within 1-65336 range")
	}
	name, ok := parameters["name"]
	if !ok || len(name) == 0 {
		return "", errors.New("name parameter is missing or empty")
	}
	user, ok := parameters["user"]
	if !ok || len(user) == 0 {
		return "", errors.New("user parameter is missing or empty")
	}
	password, ok := parameters["password"]
	if !ok || len(password) == 0 {
		return "", errors.New("passwor parameter is missing or empty")
	}
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require", user, password, host, portNum, name)
	u, err := url.Parse(connString)
	if err != nil {
		return "", nil
	}
	
	return u.String(), nil
}
