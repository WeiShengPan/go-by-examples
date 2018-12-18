package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"sync"
)

// urfave/cli 是一个命令行app工具
// 1. go build urfave_cli.go 生成 urfave_cli（urfave_cli.exe）
// 2. 启动: urfave_cli [global options] command [command options] [arguments...]
func main() {

	var host string
	var port int

	app := cli.NewApp()
	app.Name = "My App"
	app.Usage = "This is my command line app."
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host,H", Value: "127.0.0.1", Usage: "监听IP", Destination: &host},
		cli.IntFlag{Name: "port,p", Value: 8081, Usage: "监听端口", Destination: &port},
	}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start my server",
			Action: func(c *cli.Context) error {
				wait := sync.WaitGroup{}
				wait.Add(1)
				log.Println("My Server start on ", host, ":", port)
				wait.Wait()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
}
