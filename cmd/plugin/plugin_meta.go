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

package plugin_cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/client"
	"github.com/nori-io/norictl/client/connection"
	"github.com/nori-io/norictl/client/utils"
	protoNori "github.com/nori-io/norictl/internal/generated/protobuf/plugin"
)

var (
	metaDeps       func() bool
	metaDepsStatus func() bool
)

var metaCmd = &cobra.Command{
	Use:   "norictl plugin meta [PLUGIN_ID] [OPTIONS]",
	Short: "Show plugin meta data.",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := connection.CurrentConnection()
		if err != nil {
			log.Fatal(err)
		}

		if len(args) == 0 {
			log.Fatal("PLUGIN_ID required!")
		}

		pluginId := args[0]

		client, closeCh := client.NewClient(
			conn.HostPort(),
			conn.CertPath,
			"",
		)

		meta := &protoNori.PluginMetaRequest{
			ID: &protoNori.PluginID{
				MetaId:               pluginId,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			FlagDeps:             false,
			FlagDepsStatus:       false,
			FlagDependent:        false,
			FlagDependentStatus:  false,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}

		reply, err := client.PluginMetaCommand(context.Background(), meta)
		defer close(closeCh)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Meta info for plugin \"%s\":\n%s\n", pluginId, reply.Json)
	},
}

func init() {
	flags := utils.NewFlagBuilder(PluginCmd, metaCmd)
	flags.Bool(&metaDeps, "deps", "d", false, "Show only plugin dependencies")
	flags.Bool(&metaDepsStatus, "deps-status", "s", false, "Show plugin dependencies with dependent plugin status (downloaded, installed, not found etc)")
}
