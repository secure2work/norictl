// Copyright © 2018 Secure2Work info@secure2work.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package connection_cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "github.com/secure2work/norictl/client/consts"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows detailed information about specific connection.",
	Long:  `Shows detailed information about specific connection.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

func init() {
	showCmd.Flags().StringP(CONN_SHOW_FORMAT, CONN_SHOW_FORMAT_SHORT, "table", "Data representation template: json or table")

	viper.BindPFlag(CONN_SHOW_FORMAT_VIPER, showCmd.Flags().Lookup(CONN_SHOW_FORMAT))
	ConnectionCmd.AddCommand(showCmd)
}
