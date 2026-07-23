package main

//
// CLI for communicating with the Lime CRM REST API
//

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/GiGurra/boa/pkg/boa"
	"github.com/abundo/limetool"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Configuration file
type ConfigLime struct {
	APIURL string `yaml:"apiurl" boa:"noflag"`
	APIkey string `yaml:"apikey" boa:"noflag"`
}

type CommonParams struct {
	Debug    bool   `descr:"Enable verbose debug logging" short:"d" optional:"true"`
	Loglevel string `descr:"Set log level" alts:"error,warn,info,debug" default:"info" strict:"true"`
}
type Params struct {
	ConfigFile string `configfile:"true" optional:"true" default:"/etc/limetool/limetool.yaml" short:"f"`
	CommonParams
	Lime    ConfigLime `yaml:"lime"`
	Refresh bool       `descr:"Refresh data in cache from Lime" short:"r" optional:"true"`
	Company []string   `descr:"Company name to fetch (repeatable). If not specified, all companies are fetched" short:"c" optional:"true"`
}

func SetupLog(c CommonParams) error {
	level := c.Loglevel
	if c.Debug {
		level = "debug"
	}
	var logLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	slog.SetLogLoggerLevel(logLevel)
	return nil
}

// print structures JSON formatted
func Pprint(data any) {
	s, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(s))
}

func main() {
	boa.RegisterConfigFormat(".yaml", yaml.Unmarshal)
	boa.CmdT[boa.NoParams]{
		Use:   "limetool",
		Short: "limetool",
		SubCmds: boa.SubCmds(
			boa.CmdT[Params]{
				Use:   "get",
				Short: "Get librenms company data",
				RunFuncE: func(p *Params, cmd *cobra.Command, args []string) error {
					SetupLog(p.CommonParams)
					companies := p.Company
					if len(companies) == 0 {
						companies = []string{""}
					}
					lime := limetool.NewLime(p.Lime.APIURL, p.Lime.APIkey)
					for _, companyName := range companies {
						companies, err := lime.GetCompanies(companyName, p.Refresh)
						if err != nil {
							return err
						}
						Pprint(companies)
					}
					return nil
				},
			},
			boa.CmdT[Params]{
				Use:   "show-config",
				Short: "Show configuration",
				RunFuncE: func(p *Params, cmd *cobra.Command, args []string) error {
					SetupLog(p.CommonParams)
					Pprint(p)
					return nil
				},
			},
		),
	}.Run()
}
