package main

import (
	"io/ioutil"
	"log"
	"os"
)

var _ = log.Print // for debugging, remove

const DOCKIT_RUN_PATH string = "/var/run/dockit-containers"

func hasPid(svcName string) bool {
	if _, err := os.Stat(DOCKIT_RUN_PATH + "/" + svcName); err == nil {
		return true
	}
	return false
}

func getPid(svcName string) (string, error) {
	b, err := ioutil.ReadFile(DOCKIT_RUN_PATH + "/" + svcName)
	if err != nil {
		return "", err
	}

	return string(b[:]), nil
}

func setPid(svcName string, id string) error {
	b := []byte(id)

	err := ioutil.WriteFile(DOCKIT_RUN_PATH+"/"+svcName, b, 0644)

	return err
}
func removePid(svcName string) error {
	return os.Remove(DOCKIT_RUN_PATH + "/" + svcName)
}
