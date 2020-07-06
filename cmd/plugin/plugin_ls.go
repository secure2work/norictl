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
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/nori-io/norictl/cmd/common"
	"github.com/nori-io/norictl/internal/client"
	"github.com/nori-io/norictl/internal/client/connection"
	"github.com/nori-io/norictl/internal/client/utils"
	protoGenerated "github.com/nori-io/norictl/internal/generated/protobuf"
)

var (
	listAll         func() bool
	listError       func() bool
	listInactive    func() bool
	listInstallable func() bool
	listInstalled   func() bool
	listRunning     func() bool
)

func lsCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "ls [OPTIONS]",
		Aliases: []string{"list"},
		Short:   "Shows list of plugins on remote Nori node.",
		Run: func(cmd *cobra.Command, args []string) {
			setFlagsLs()
			conn, err := connection.CurrentConnection()
			if err != nil {
				fmt.Println("%s", err)
			}

			client, closeCh := client.NewClient(
				conn.HostPort(),
				conn.CertPath,
				"",
			)

			reply, err := client.PluginListCommand(context.Background(), &protoGenerated.PluginListRequest{
				FlagAll:              listAll(),
				FlagError:            listError(),
				FlagInstalled:        listInstalled(),
				FlagRunning:          listRunning(),
				FlagInstallable:      listInstallable(),
				FlagInactive:         listInactive(),

			})
			close(closeCh)
			if err != nil {
				if reply != nil {
					fmt.Println("%s", protoGenerated.ErrorReply{
						Status:               false,
						Error:                err.Error(),
					})
				}
				fmt.Println("%s", err)
			}

			list := []*protoGenerated.PluginListWithStatus{{
				Id:              nil,
				Author:          nil,
				Dependencies:    nil,
				Description:     nil,
				Core:            nil,
				Interface:       "",
				License:         nil,
				Links:           nil,
				Repository:      nil,
				Tags:            nil,
				FlagAll:         true,
				FlagError:       false,
				FlagInstalled:   false,
				FlagRunning:     false,
				FlagInstallable: false,
				FlagInactive:    false,
			}}

			filter := func(list []*protoGenerated.PluginListWithStatus, f func(p protoGenerated.PluginListWithStatus) bool) []*protoGenerated.PluginListWithStatus {
				newList := make([]*protoGenerated.PluginListWithStatus, 0)
				plugins := make([][]string, len(list))
				for _, l := range list {
					if f(*l) {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.Id.String(), l.Author.String()})
					}
				}
				common.UI.PluginsAll(plugins)
				return newList
			}

			if listAll() {
				list = filter(list, func(p protoGenerated.PluginListWithStatus) bool {
					newList := make([]*protoGenerated.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.Id.String(), l.Author.String()})
					}
					common.UI.PluginsAll(plugins)
					return p.FlagAll
				})
			}

			if listError() {
				list = filter(list, func(p protoGenerated.PluginListWithStatus) bool {
					newList := make([]*protoGenerated.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.Id.String(), l.Author.String()})

					}
					common.UI.PluginsError(plugins)
					return p.FlagError
				})
			}

			if listInactive() {
				list = filter(list, func(p protoGenerated.PluginListWithStatus) bool {
					newList := make([]*protoGenerated.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.Id.String(), l.Author.String()})
					}
					common.UI.PluginsInactive(plugins)
					return p.FlagInactive
				})
			}

			if listInstallable() {
				list = filter(list, func(p protoGenerated.PluginListWithStatus) bool {
					newList := make([]*protoGenerated.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.Id.String(), l.Author.String()})

					}
					common.UI.PluginsInstallable(plugins)
					return p.FlagInstallable
				})
			}

			if listInstalled() {
				list = filter(list, func(p protoGenerated.PluginListWithStatus) bool {
					newList := make([]*protoGenerated.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.Id.String(), l.Author.String()})

					}
					common.UI.PluginsInstalled(plugins)
					return p.FlagInstalled
				})
			}

			if listRunning() {
				list = filter(list, func(p protoGenerated.PluginListWithStatus) bool {
					newList := make([]*protoGenerated.PluginListWithStatus, 0)
					plugins := make([][]string, len(list))
					for _, l := range list {
						newList = append(newList, l)
						plugins = append(plugins, []string{l.Id.String(), l.Author.String()})
					}
					common.UI.PluginsRunning(plugins)
					return p.FlagRunning
				})
			}

		},
	}
}

func init() {
}

func setFlagsLs() {
	flags := utils.NewFlagBuilder(PluginCmd(), lsCmd())
	flags.Bool(&listAll, "all", "--all", false, "Show all plugins")                                          // TODO
	flags.Bool(&listError, "error", "-e", false, "Show plugins with errors (not implement)")                 // TODO
	flags.Bool(&listInactive, "inactive", "--inactive", false, "Show plugins that are not running")          // TODO
	flags.Bool(&listInstallable, "installable", "--installable", false, "Show plugins that need to install") // TODO
	flags.Bool(&listInstalled, "installed", "-i", false, "Show only installed plugins")
	flags.Bool(&listRunning, "running", "-r", false, "Show only running plugins")
}
