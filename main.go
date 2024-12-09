package main

import (
	"fmt"
	store "keyscape/store"
	util "keyscape/utils"
)

func main() {

	s := store.NewStore()
	cmd := store.Command{
		CmdType: store.SET,
		Key:     "lol",
		Value:   "lol2",
	}

	encodedCmd, err := store.EncodeCommand(cmd)
	if err != nil {
		println(err.Error())
		return
	}

	res, err := s.Apply(encodedCmd)
	if err != nil {
		util.DebugError(err, fmt.Sprintf("%s", cmd.CmdType))
		return
	}

	cmd = store.Command{
		CmdType: store.GET,
		Key:     "lol",
		Value:   "lol2",
	}

	encodedCmd, err = store.EncodeCommand(cmd)
	if err != nil {
		println(err.Error())
		return
	}

	res, err = s.Apply(encodedCmd)
	if err != nil {
		util.DebugError(err, fmt.Sprintf("%s", cmd.CmdType))
		return
	}

	println(string(res[:]))
}
