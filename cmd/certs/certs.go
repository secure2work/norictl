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

package certs_cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CertsCmd = &cobra.Command{
	Use:   "certs",
	Short: "certs",
}

func init() {
	CertsCmd.PersistentFlags().String("pem", "server.pem", "path to pem file")
	CertsCmd.PersistentFlags().String("key", "server.key", "path to key file")
	CertsCmd.PersistentFlags().String("passkey", "", "secret passkey")

	viper.BindPFlag("pem", CertsCmd.PersistentFlags().Lookup("pem"))
	viper.BindPFlag("key", CertsCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("passkey", CertsCmd.PersistentFlags().Lookup("passkey"))
}
