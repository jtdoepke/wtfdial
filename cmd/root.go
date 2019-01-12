package cmd

import (
	"errors"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus" // Alias logrus as log
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jtdoepke/wtfdial/wtf"
)

// This is the default port the server and client
// both try to use on localhost.
var defaultPort = "9837"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wtfdial username wtf",
	Short: "Express your WTF",
	Long:  "A client and server to aggregate team WTF levels.",
	Args:  cobra.ExactArgs(1),
	Run:   runClient,
}

// In addition to `main()`, Go also looks for a special function name
// `init()`. This function is called on program start, and can be used to
// initialize stuff in the package. In this case it is setting up the command
// line argument parsing. Each source file can have an `init()` function.
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wtfdial.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().String("host", "localhost:"+defaultPort, "Hostname of WTF server")
	rootCmd.Flags().StringP("username", "u", "", "Your user name")
}

func runClient(cmd *cobra.Command, args []string) {
	// Parse the command line arguments.
	serverHost := cmd.Flag("host").Value.String()
	username := cmd.Flag("username").Value.String()
	if username == "" {
		log.Fatalln("Must set -u flag with username")
	}
	level, err := ParseWTFLevelArg(args[0])
	if err != nil {
		log.Fatalln(err)
	}
	if level < 0 || level > 1 {
		log.Fatalln("WTF level must be between 0 and 1")
	}

	// Make connection to rpc server.
	client, err := rpc.DialHTTP("tcp", serverHost)
	if err != nil {
		log.Fatalln("Error connecting to server.", err)
	}
	defer client.Close() // Close connection at the end of this function.

	// Make request object
	req := wtf.SetLevelRequest{
		Username: username,
		Level:    level,
	}

	// Call remote procedure to set user level.
	err = client.Call("WTFService.SetLevel", req, nil)
	if err != nil {
		log.Fatalln("Error setting WTF level", err)
	}

	// Call remote procedure to get avg wtf level.
	var result float64                              // Store the result here.
	err = client.Call("WTFService.Avg", 0, &result) // Pass a pointer to `result`, so the result can be stored there.
	if err != nil {
		log.Fatalln("Error getting WTF level average", err)
	}
	//we got our result in result
	log.Printf("Average WTF level: %v \n", result)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".wtfdial" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".wtfdial")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// ErrParseWTFLevel is returned by `ParseWTFLevelArg()` if the
// WTF level string cannot be parsed.
var ErrParseWTFLevel = errors.New("cannot parse WTF level arg")

// ParseWTFLevelArg parses various forms of WTF level string.
func ParseWTFLevelArg(wtfLevel string) (float64, error) {
	if strings.ContainsRune(wtfLevel, '/') {
		parts := strings.Split(wtfLevel, "/")
		if len(parts) != 2 {
			return 0, ErrParseWTFLevel
		}
		num, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return 0, ErrParseWTFLevel
		}
		denom, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return 0, ErrParseWTFLevel
		}
		return float64(num) / float64(denom), nil
	}
	wtfLevel = strings.TrimSuffix(wtfLevel, "%") // In case of percentage, just strip off the '%' and let the next if take care of it.
	if lvl, err := strconv.ParseInt(wtfLevel, 10, 64); err == nil {
		return float64(lvl) / 100, nil
	}
	if lvl, err := strconv.ParseFloat(wtfLevel, 64); err == nil {
		return lvl, nil
	}
	return 0, ErrParseWTFLevel
}
