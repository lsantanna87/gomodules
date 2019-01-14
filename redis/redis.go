package redis

import (
	"log"

	"github.com/mediocregopher/radix/v3"
)

var (
	Client *radix.Cluster
)

const get string = "GET"
const set string = "SET"

type fnExecCommand func() radix.CmdAction

func exec(fn fnExecCommand) error {

	err := Client.Do(fn())

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func Get(key string) (string, error) {

	var ret string

	err := exec(func() radix.CmdAction {
		return radix.Cmd(&ret, get, key)
	})

	return ret, err
}

func Set(key string, value string) error {

	err := exec(func() radix.CmdAction {
		return radix.Cmd(nil, set, key, value)
	})

	return err
}
