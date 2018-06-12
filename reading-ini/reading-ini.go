package main

import (
	"github.com/go-ini/ini"
	"fmt"
	"os"
)

func main() {
	cfg, err := ini.Load("/home/abhishek/research/ini-golang/cfg.ini")
	if err != nil {
		fmt.Printf("err reading file: %v", err)
		os.Exit(1)
	}
	server := cfg.Section("server").Key("protocol").String()
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())
	fmt.Println(server)
	fmt.Println("Server Protocol", cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
}
