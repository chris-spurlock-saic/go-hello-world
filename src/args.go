package main

import (
	"os"
	"strings"
)

type Args struct {
	Port string
}

func GetArgs() *Args {
	args := Args{
		Port:      "80"}
	for flag, arg := range readArgs() {
		switch flag {
		case "-port":
			args.Port = arg
		}
	}
	port := os.Getenv("PORT")
	if port != "" {
		args.Port = port
	}
	return &args
}

func readArgs() map[string]string {
	args := make(map[string]string)
	if len(os.Args) < 2 {
		return args
	}
	flag := ""
	for i, val := range os.Args {
		if i == 0 {
			continue
		}
		arg := strings.TrimSpace(val)
		if flag != "" {
			args[flag] = arg
			flag = ""
			continue
		}
		if strings.HasPrefix(arg, "-") {
			flag = arg
		}
	}
	return args
}
