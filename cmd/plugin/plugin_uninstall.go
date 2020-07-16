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

	"github.com/nori-io/nori-common/v2/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	protoGenerated "github.com/nori-io/norictl/pkg/proto"
)

var (
	uninstallAll      bool
	uninstallDependent bool
)

func uninstallCmd() *cobra.Command {

	cmd:=&cobra.Command{
		Use:   "uninstall [PLUGIN_ID] [OPTIONS]",
		Short: "Uninstall plugin or plugins.",
		Run: func(cmd *cobra.Command, args []string) {
			pluginId := viper.GetString("id")
			if len(pluginId) == 0 && len(args) > 0 {
				pluginId = args[0]
			}
			pluginIdSplit := strings.Split(pluginId, ":")
			versionPlugin := pluginIdSplit[1]
			_, err := version.NewVersion(versionPlugin)
			if err != nil {
				fmt.Println("Format of plugin's version is incorrect:", err)
			}

			cli, closeCh := client.NewClient(
				viper.GetString("grpc-address"),
				viper.GetString("ca"),
				viper.GetString("ServerHostOverride"),
			)

			reply, err := cli.PluginUninstallCommand(context.Background(), &protoGenerated.PluginUninstallRequest{
				Id: &protoGenerated.ID{
					PluginId: pluginIdSplit[0],
					Version:  pluginIdSplit[1],
				},
				FlagAll:       uninstallAll,
				FlagDependent: uninstallDependent,
			})
			defer close(closeCh)
			if err != nil {
				if reply != nil {
					common.UI.PluginUninstallFailure(pluginId)
					fmt.Println("%s", protoGenerated.Error{
						Code:    reply.GetCode(),
						Message: reply.GetMessage(),
					})
				}
				fmt.Println("%s", err)
			}
			common.UI.PluginUninstallSuccess(pluginId)
		},
	}
	cmd.Flags().BoolVarP(&uninstallAll, "all", "a", true, "Uninstall all installed plugins")
	cmd.Flags().BoolVarP(&uninstallDependent, "dependent", "d", true, "Uninstall plugin and depend plugins")
	return cmd
}



