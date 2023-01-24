/*
Copyright © 2022 kruemelmann

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/kruemelmann/pomodo/web"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start serving the screenvideo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		printAsciiart()

		fmt.Printf("pomodo is serving under the ips:\n\n")
		ips := getIPAddresses()
		for _, v := range ips {
			fmt.Printf("\thttp://%s:%d\n", v, port)
		}
		web.StartWebserver(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntP("port", "p", 9200, "Port to run pomodo gui on (default: 9200)")
}

func getIPAddresses() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("unable to get interfaces %v", err)
	}
	iplist := []string{}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatalf("unable to find interfaces Addrs%v", err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.To4() != nil {
				iplist = append(iplist, ip.String())
			}
		}
	}
	return iplist
}

func printAsciiart() {
	fmt.Println("                                  _")
	fmt.Println("                                 | |")
	fmt.Println("  _ __   ___  _ __ ___   ___   __| | ___")
	fmt.Println(" | '_ \\ / _ \\| '_ ` _ \\ / _ \\ / _` |/ _ \\")
	fmt.Println(" | |_) | (_) | | | | | | (_) | (_| | (_) |")
	fmt.Println(" | .__/ \\___/|_| |_| |_|\\___/ \\__,_|\\___/")
	fmt.Println(" | |")
	fmt.Println(" |_|")
}
