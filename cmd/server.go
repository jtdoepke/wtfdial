package cmd

import (
	"net/http"
	"net/rpc"

	"github.com/jtdoepke/wtfdial/wtf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverCmd represents the server subcommand
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the WTF server",
	Run:   runServer,
}

// In addition to `main()`, Go also looks for a special function name
// `init()`. This function is called on program start, and can be used to
// initialize stuff in the package. In this case it is setting up the command
// line argument parsing. Each source file can have an `init()` function.
func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().StringP("listen", "l", ":"+defaultPort, "Hostname/port that the server listens on.")
}

func runServer(cmd *cobra.Command, args []string) {
	port := cmd.Flag("listen").Value.String()

	// Create the WTF service object.
	service := wtf.New()

	// Register the service with rpc so it's methods are exposed
	// as a JSON-RPC API.
	err := rpc.Register(service)
	if err != nil {
		log.Fatalf("Format of service service isn't correct. %s", err)
	}

	// Register rpc as the default request handler for the built-in HTTP server.
	rpc.HandleHTTP()

	// Start the HTTP server.
	log.Println("Serving RPC handler")
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
