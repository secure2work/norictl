package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"

	"github.com/secure2work/nori/core/grpc"
	"github.com/secure2work/nori/proto"
	"github.com/secure2work/norictl/client"
)

var uploadCertsCmd = &cobra.Command{
	Use:   "upload",
	Short: "uploading certs",
	Run: func(cmd *cobra.Command, args []string) {
		pem := viper.GetString("pem")
		key := viper.GetString("key")
		passkey := viper.GetString("passkey")

		if len(pem) == 0 || len(key) == 0 {
			logrus.Fatal("Required pem and key files")
		}

		if len(passkey) == 0 {
			logrus.Fatal("Required passkey")
		}

		pemFile, err := os.Open(pem)
		if err != nil {
			logrus.Fatal(err)
		}
		defer pemFile.Close()

		keyFile, err := os.Open(key)
		if err != nil {
			logrus.Fatal(err)
		}
		defer keyFile.Close()

		pemBytes, err := ioutil.ReadAll(pemFile)
		if err != nil {
			logrus.Fatal(err)
		}

		keyBytes, err := ioutil.ReadAll(keyFile)
		if err != nil {
			logrus.Fatal(err)
		}

		client, closeCh := client.NewClient(
			viper.GetString("grpc-address"),
			viper.GetString("ca"),
			viper.GetString("ServerHostOverride"),
		)

		// encrypting...
		pk, err := grpc.PasskeyFromString(passkey)
		if err != nil {
			logrus.Fatal(err)
		}

		var hmac []byte

		pemBytes, hmac, err = pk.Encrypt(pemBytes)
		if err != nil {
			logrus.Fatal(err)
		}

		hmacLen := len(hmac)

		pemBs := append([]byte{byte(hmacLen)}, hmac...)
		pemBs = append(pemBs, pemBytes...)

		keyBytes, hmac, err = pk.Encrypt(keyBytes)
		if err != nil {
			logrus.Fatal(err)
		}

		keyBs := append([]byte{byte(hmacLen)}, hmac...)
		keyBs = append(keyBs, keyBytes...)

		reply, err := client.UploadCertsCommand(context.Background(), &commands.UploadCertsRequest{
			Pem: pemBs,
			Key: keyBs,
		})

		close(closeCh)

		if err != nil {
			logrus.Fatal(err)
			if reply != nil {
				logrus.Fatal(reply.Error)
			}
		} else {
			fmt.Printf("Certificates %q and %q successfully uploaded\n", pem, key)
		}
	},
}

func init() {
	certsCmd.AddCommand(uploadCertsCmd)
}
