package main

import (
	"fmt"
	"strings"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/lima-vm/lima/pkg/hostagent/dns"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "host-resolver-dns [flags]",
		Short: "DNS Server using host system resolver",
		Long:  "See https://github.com/dee-kryvenko/host-resolver-dns/blob/main/README.md for details",
		RunE: func(cmd *cobra.Command, args []string) error {
			udp, err := cmd.Flags().GetInt("udp")
			if err != nil {
				return err
			}

			tcp, err := cmd.Flags().GetInt("tcp")
			if err != nil {
				return err
			}

			ipv6, err := cmd.Flags().GetBool("ipv6")
			if err != nil {
				return err
			}

			_hosts, err := cmd.Flags().GetStringSlice("hosts")
			if err != nil {
				return err
			}

			headless, err := cmd.Flags().GetBool("headless")
			if err != nil {
				return err
			}

			hosts := make(map[string]string)
			for _, host := range _hosts {
				h := strings.SplitN(host, "=", 2)
				if len(h) != 2 {
					return fmt.Errorf("Invalid host entry %q", host)
				}
				hosts[h[0]] = h[1]
			}

			c := make(chan struct{})
			if !headless {
				go func() {
					<-c
					systray.Quit()
				}()
			}

			var server *dns.Server

			server, err = dns.Start(udp, tcp, ipv6, hosts)
			defer server.Shutdown()

			fmt.Printf("DNS started (udp=%v, tcp=%v, ipv6=%v, hosts=%s)\n", udp, tcp, ipv6, hosts)

			if headless {
				<-c
			} else {
				systray.Run(func() {
					systray.SetIcon(icon.Data)
					systray.SetTooltip("DNS Server using host system resolver")
					mQuit := systray.AddMenuItem("Exit", "Stop this DNS server")
					mQuit.SetIcon(icon.Data)
					mQuit.ClickedCh = c
				}, func() {
					if server != nil {
						fmt.Printf("Shutting down...\n")
						server.Shutdown()
					}
				})
			}

			return err
		},
	}
)

func init() {
	rootCmd.Flags().IntP("udp", "u", 16237, "UDP port number")
	rootCmd.Flags().IntP("tcp", "t", 16237, "TCP port number")
	rootCmd.Flags().BoolP("ipv6", "6", false, "Enable IPv6")
	rootCmd.Flags().StringSlice("hosts", []string{}, "Additional hosts to resolve")
	rootCmd.Flags().Bool("headless", false, "Enable headless mode")
}
