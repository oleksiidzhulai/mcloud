// Copyright © 2018 Oleksii Dzhulai <nixlike@gmail.com>
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
	"github.com/spf13/cobra"
)

// createNetworkCmd represents the createNetwork command
var createNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		mcloud.Projects[cmd.Flag("project").Value.String()].addNetwork(cmd.Flag("name").Value.String())
	},
}

func init() {
	createCmd.AddCommand(createNetworkCmd)
	createNetworkCmd.Flags().String("project", "", "Project where we create the networks")
	createNetworkCmd.MarkFlagRequired("project")
}
