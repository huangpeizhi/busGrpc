package main

import (
	"blog"
	"os"

	"github.com/go-playground/log"
	tjconfig "github.com/tj/go-config"
)

func main() {
	initConfigure()

	f, err := NewFactory()
	if err != nil {
		log.WithFields(log.F("func", "main.NewFactory")).Fatal(err)
	}

	//加载Fence数据
	if opts.LoadFence {
		log.WithFields(log.F("func", "main.initFence")).Info("begin load fence data")
		err := f.initFence()
		if err != nil {
			log.WithFields(log.F("func", "main.initFence")).Fatal(err)
		}
		log.WithFields(log.F("func", "main.initFence")).Info("end load fence data, exit.")
		os.Exit(0)
	}

	for i := 0; i < opts.PCount; i++ {
		go f.AnalysisLoop()
	}
	go f.CollectGpsLoop()

	gRpcRun()
}

func initConfigure() {
	opts = defaultOpts
	tjconfig.MustResolve(opts)

	cLog := &blog.BLogHandler{
		Buffer:     0,
		DateFormat: "2006-01-02 15:04:05.000",
		FuncName:   "func",
		ISep:       " :: ",
	}

	if opts.Debug {
		log.RegisterHandler(cLog, blog.AllLevels...)
	} else {
		log.RegisterHandler(cLog, blog.RunLevels...)
	}
}
