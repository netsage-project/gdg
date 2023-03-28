package main

import (
	_ "embed"
	"sync"

	"github.com/esnet/gdg/cmd"

	applogger "github.com/esnet/gdg/log"
)

//go:embed conf/importer-example.yml
var defaultConfig string

var doOnce sync.Once

func main() {
	cmd.DefaultConfig = defaultConfig
	cmd.Execute()
}

func init() {
	doOnce.Do(func() {
		applogger.InitializeAppLogger()
	})
}
