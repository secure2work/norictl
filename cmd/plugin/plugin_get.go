// Copyright © 2018 Nori info@nori.io
//
// This program is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation, either version 3
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Package plugin_cmd implements commands for work with plugins
//by command prompt*/
package plugin_cmd

import (
	"fmt"
	"strings"

	"github.com/nori-io/nori-common/version"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/utils"
	commonProtoGenerated "github.com/nori-io/norictl/internal/generated/protobuf/common"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	getVerbose func() bool
)

func getCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "get [PLUGIN_ID] [OPTIONS]",
		Short: "downloading plugin",
		Long: `Get downloads the plugin, along with its dependencies.
	It then installs the plugin, like norictl plugin install.`,
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsGet()
			conn, err := connection.CurrentConnection()
			if err != nil {
				fmt.Println("%s", err)
				return
			}

			if len(args) == 0 {
				fmt.Println("PLUGIN_ID required!")
				return
			}

			pluginId := args[0]
			pluginIdSplit := strings.Split(pluginId, ":")
			versionPlugin := pluginIdSplit[1]
			_, err = version.NewVersion(versionPlugin)
			if err != nil {
				fmt.Println("Format of plugin's version is incorrect:", err)
			}

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)

			reply, err := client.PluginGetCommand(context.Background(), &protoNori.PluginGetRequest{
				Id: &commonProtoGenerated.ID{
					Id:                   pluginIdSplit[0],
					Version:              pluginIdSplit[1],
				},
				FlagVerbose:          getVerbose(),
			})

			close(closeCh)

			if err != nil {
				fmt.Println("%s", err)
				common.UI.PluginGetFailure(pluginId)
				if reply != nil {
					fmt.Println("%s", commonProtoGenerated.ErrorReply{
						Status:               false,
						Error:                err.Error(),
					})
				}
				return
			} else {
				common.UI.PluginGetSuccess(pluginId)
			}
		},
	}
}

func init() {

}

func setFlagsGet() {
	flags := utils.NewFlagBuilder(PluginCmd(), getCmd())
	flags.Bool(&getVerbose, "verbose", "-v", false, "Verbose progress and debug output")
}
