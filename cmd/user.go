/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"github.com/spf13/cobra"
)

// var (
// 	port          string
// 	httpPort      string
// 	swaggerSwitch bool
// 	gatewaySwitch bool
// )

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "用户服务",
	Long:  `用户服务`,
	Run: func(cmd *cobra.Command, args []string) {
		if gatewaySwitch {
			if err := gatewayRun(); err != nil {
				log.Fatal("cannot start server: ", err)
			}
		} else {
			if err := grpcRun(); err != nil {
				log.Fatal("cannot start server: ", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
