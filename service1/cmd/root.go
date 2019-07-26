// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"net"

	"github.com/ddsgok/gql"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server/info"
	"github.com/ddspog/hashlab-hiring-backapply/service1/server/protocols"
	"github.com/ddspog/hashlab-hiring-backapply/service1/today"
	"github.com/ddsgok/log"
	"github.com/ddsgok/str"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// rootCmd represents the base command when called without any subcommands
var (
	// rootCmd represents the base command when called without any subcommands.
	rootCmd = &cobra.Command{
		Use:   "service1",
		Short: str.Fmt("Server responsible of %s micro service", info.Name),
		Long: str.Fmt(`%s - Program containing micro service that allows discount consult for a simple DB.
	Start the server at localhost, in desired port. User can specify the port with --port flag, or in 
	SERVICE1_PORT environment variable. The host can also be specified through --host flag or SERVICE1_HOST 
	environment	variable. It connects to a GraphQL service`, bold(info.Name)),
		Run: func(cmd *cobra.Command, args []string) {
			Start()
		},
	}
	// verbose and debug stores global bool flag related to level of
	// information printed on log.
	verbose, debug bool

	// port and host stores the desired port and host to use for server.
	port, host string

	// graphqlsrv stores the choosen GraphQL engine to use on this server.
	graphqlsrv string
)

// Start the service.
func Start() {
	// First message.
	log.Print("info: Starting %s.", info.Name)

	// Connect and starts the database.
	log.Print("info: - Connecting on GrapgQL engine gql=%s", graphqlsrv)

	gql.ConnectAt(graphqlsrv)
	defer gql.Disconnect()

	log.Print("info: - Setting net connection.")

	lis, err := net.Listen("tcp", str.Fmt("%s:%s", host, port))
	if err != nil {
		log.Fatal("error: failed to listen: %v", err)
	}

	log.Print("info: - Creating GRPC server.")

	gsrv := grpc.NewServer()
	protocols.RegisterDiscountServer(gsrv, server.New())

	if info.TodayDate != "" {
		log.Print("info: - Mocking Today value to %s.", info.TodayDate)
		today.MockNow(info.TodayDate)
		defer today.ResetMocks()
	}

	log.Print("message: Server started at %s:%s", host, port)
	if err := gsrv.Serve(lis); err != nil {
		log.Fatal("error: failed starting http server: %v", err)
	}

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("fatal: couldn't execute root command err=%v", err)
	}

	// Print if successful, to show change in terminal.
	str.New("\n").Print()
}

// init root command
func init() {
	cobra.OnInitialize(initLog)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "if set, will print more detailed information")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "if set, will print debug information")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", info.Port, "port to load the server")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "", info.Host, "host to load the server")
	rootCmd.PersistentFlags().StringVarP(&graphqlsrv, "graphql-server", "g", info.GraphQLServer, "GraphQL engine server to load data")
}

// initConfig setups log context depending on verbose or debug.
func initLog() {
	switch {
	case verbose:
		log.SetLogContext(log.VerboseContext)
	case debug:
		log.SetLogContext(log.DebuggingContext)
	}
}
