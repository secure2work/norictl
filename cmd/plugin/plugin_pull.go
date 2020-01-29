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

	"github.com/nori-io/nori-common/v2/logger"
	"github.com/nori-io/nori-common/version"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/utils"
	"github.com/nori-io/norictl/internal/generated/protobuf/common_messages"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin_messages"
)

var (
	pullDeps func() bool
)

func pullCmd(log logger.Logger) *cobra.Command {

	return &cobra.Command{
		Use:   "norictl plugin pull [PLUGIN_ID] [OPTIONS]",
		Short: "downloading plugin",
		Long:  `Pull downloads the plugin, with or without it's dependencies.`,
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsPull(log)
			conn, err := connection.CurrentConnection()
			if err != nil {
				log.Fatal("%s", err)
			}

			if len(args) == 0 {
				log.Fatal("PLUGIN_ID required!")
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

			reply, err := client.PluginPullCommand(context.Background(), &protoNori.PluginPullRequest{
				Id: &common_messages.ID{
					Id:                   pluginIdSplit[0],
					Version:              pluginIdSplit[1],
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				FlagDeps:             pullDeps(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			})

			close(closeCh)
			if err != nil {
				log.Fatal("%s", err)
				common.UI.PullFailure(pluginId)
				if reply != nil {
					log.Fatal("%s", common_messages.ErrorReply{
						Status:               false,
						Error:                err.Error(),
						XXX_NoUnkeyedLiteral: struct{}{},
						XXX_unrecognized:     nil,
						XXX_sizecache:        0,
					})
				}
			} else {
				common.UI.PullSuccess(pluginId)
			}
		},
	}
}

func init() {
}

func setFlagsPull(log logger.Logger) {
	flags := utils.NewFlagBuilder(PluginCmd(log), pullCmd(log))
	flags.Bool(&pullDeps, "deps", "-d", false, "	Download plugin with it's dependencies")
}
