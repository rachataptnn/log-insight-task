package dataParser

import (
	"errors"
	"strconv"
	"strings"
)

func GetLatencyInMs(latencyStr string) (int, error) {
	split := strings.Split(latencyStr, "ms")
	if len(split) != 2 {
		return 0, errors.New("not found ms in latency")
	}
	latency, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, err
	}

	return latency, nil
}

func GetRoute(uri string) (route string, err error) {
	split := strings.Split(uri, "/")
	if len(split) < 4 {
		return "", errors.New("URI has less than 3 parts")
	}
	route = split[3] // need to check the correct index

	return route, nil
}
