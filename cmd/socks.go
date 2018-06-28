// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var socks = &cobra.Command{
	Use:   "socks",
	Short: "socks",
	Long:  `socks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("socks called")
	},
}

var socksDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy SOCKS Proxy",
	Long:  `Deploy SOCKS Proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("socks deploy called")
	},
}

var socksDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy Specific SOCKS Proxy",
	Long:  `file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("socks destroy called")
	},
}

var socksList = &cobra.Command{
	Use:   "list",
	Short: "List available SOCKS Proxies",
	Long:  `List available SOCKS Proxies`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("socks list called")
	},
}

func init() {
	rootCmd.AddCommand(socks)
	socks.AddCommand(socksDeploy, socksDestroy, socksList)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
