package cmd

import (
	"fmt"
	"github.com/MruV-RP/mruv-pb-go/texturestudio"
	"github.com/MruV-RP/mruv-service-template/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if we crash the go code, we get the file name and line number in log
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		logrus.Infoln("Starting listener for gRPC API")
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port")))
		if err != nil {
			logrus.Fatalf("Failed to listen: %v", err)
		}
		defer func() {
			err := listener.Close()
			if err != nil {
				logrus.Infoln("Errors caused by closing listener:", err)
			}
		}()
		logrus.Infoln("Listener started on", listener.Addr())

		//server options
		var opts []grpc.ServerOption
		if viper.GetBool("ssl.enabled") {
			creds, sslErr := credentials.NewServerTLSFromFile(viper.GetString("ssl.cert"), viper.GetString("ssl.key"))
			if sslErr != nil {
				logrus.Fatalln("Failed loading certificates:", sslErr)
				return
			}
			opts = append(opts, grpc.Creds(creds))
		}

		//serve gRPC services
		s := grpc.NewServer(opts...)
		reflection.Register(s)
		defer s.Stop()

		// Register services
		texturestudio.RegisterTextureStudioServiceServer(s, api.NewServer())

		go func() {
			logrus.Println("Starting server.")

			if err := s.Serve(listener); err != nil {
				logrus.Fatalln("Failed to serve", err)
			}
		}()

		// Wait for CTRL+C to exit
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)

		<-ch // Block until signal is received
		logrus.Infoln("\nStopping the server.")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
