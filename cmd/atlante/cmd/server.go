package cmd

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-spatial/maptoolkit/atlante/queuer"
	"github.com/prometheus/common/log"

	"github.com/dimfeld/httptreemux"

	"github.com/go-spatial/maptoolkit/atlante/config"
	"github.com/go-spatial/maptoolkit/atlante/server"
	cmdconfig "github.com/go-spatial/maptoolkit/cmd/atlante/config"
	"github.com/spf13/cobra"
)

var (
	// Server is the command to start up the api server
	Server = &cobra.Command{
		Use:     "serve",
		Short:   "Use atlante as an api server",
		Aliases: []string{"server"},
		Long:    `Use atlante as an api server. Grids are served up as sheets/info/mgdid/:mdgid or /sheets/info/:lat/:lng`,
		RunE:    serverCmdRunE,
	}

	// port that server should start up on, but default we will use :8080
	port = ":8080"
)

func init() {
	Server.Flags().StringVar(&port, "port", ":8080", "port to start the server on")
}

func serverCmdRunE(cmd *cobra.Command, args []string) error {

	aURL, err := url.Parse(configFile)
	if err != nil {
		return err
	}
	conf, err := config.LoadAndValidate(aURL)
	if err != nil {
		return err
	}

	a, err := cmdconfig.LoadConfig(conf, dpi, cmd.Flag("dpi").Changed)
	if err != nil {
		return ErrExitWith{
			Err:       err,
			Msg:       "error loading config",
			ExitCode:  1,
			ShowUsage: true,
		}
	}

	// Shadow port and then check to see if it changed and the config
	// has a value we should use instead
	port := port
	if !cmd.Flag("port").Changed && conf.Webserver.Port != "" {
		port = string(conf.Webserver.Port)
	}

	// Need to initialize the server
	srv := server.Server{
		Hostname: string(conf.Webserver.HostName),
		Port:     port,
		Scheme:   string(conf.Webserver.Scheme),
		Headers:  make(map[string]string),
		Atlante:  a,
	}

	// Now we need to look to see if a queue has been configured
	if conf.Webserver.Queue != nil {
		qType, _ := conf.Webserver.Queue.String(queuer.ConfigKeyType, nil)
		if qType != "none" || qType != "" {
			// Configure the queue
			srv.Queue, err = queuer.For(qType, conf.Webserver.Queue)
			if err != nil {
				if _, ok := err.(queuer.ErrUnknownProvider); ok {
					log.Infoln("known queue providers:")
					for _, p := range queuer.Registered() {
						log.Infoln("\t", p)
					}
				}
				return err
			}
		}
	}

	for name, value := range conf.Webserver.Headers {
		// cast to string
		val := fmt.Sprintf("%v", value)
		if val == "" {
			fmt.Fprintln(cmd.OutOrStderr(), "warning, webserver.header (%v) has no configured value, ignoring", name)
		}
		srv.Headers[name] = val
	}

	router := httptreemux.New()

	srv.RegisterRoutes(router)

	fmt.Fprintf(cmd.OutOrStdout(), "starting up server on %v\n", port)
	err = http.ListenAndServe(srv.Port, router)
	switch err {
	case nil:
		fmt.Fprintf(cmd.OutOrStderr(), "shutting down")
		return nil
	case http.ErrServerClosed:
		fmt.Fprintf(cmd.OutOrStderr(), "http server closed")
		return nil
	default:
		return ErrExitWith{
			Err:      err,
			Msg:      "Failed to start up server",
			ExitCode: 1,
		}
	}
}