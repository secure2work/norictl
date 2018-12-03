package plugin_cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "downloading,installing and uploading plugin",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("file")

		if len(path) == 0 && len(args) > 0 {
			path = args[0]
		}

		client, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)
		defer close(closeCh)

		f, err := os.Open(path)
		if err != nil {
			logrus.Fatal(err)
		}

		defer f.Close()

		so, err := ioutil.ReadAll(f)
		if err != nil {
			logrus.Fatal(err)
		}
		path = filepath.Base(path)

		reply, err := client.PluginUploadCommand(context.Background(), &commands.PluginUploadRequest{
			Name: path,
			So:   so,
		})

		if err != nil {
			logrus.Fatal(err)
			if reply != nil {
				logrus.Fatal(reply.Error)
			}
		} else {
			fmt.Printf("Plugin %q successfully uploaded\n", path)
		}
	},
}

func init() {
	PluginCmd.AddCommand(uploadCmd)
}