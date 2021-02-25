package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"UCMF-SERVER/UCMFClient/"
	"UCMF-SERVER/UCMFServer/"
)

var UDR = &udr_service.UDR{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "udr"
	appLog.Infoln(app.Name)
	appLog.Infoln("UDR version: ", version.GetVersion())
	app.Usage = "-free5gccfg common configuration file -udrcfg udr configuration file"
	app.Action = action
	app.Flags = UDR.GetCliCmd()
	if err := app.Run(os.Args); err != nil {
		logger.AppLog.Warnf("Error args: %v", err)
	}
}

func action(c *cli.Context) {
	app.AppInitializeWillInitialize(c.String("free5gccfg"))
	UDR.Initialize(c)
	UDR.Start()
}