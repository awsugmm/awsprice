// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	resource "github.com/aws-user-group-myanmar-aws-ugm/awsprice/resource"
	"github.com/spf13/cobra"
)

// getec2Cmd represents the getec2 command
var getec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "A brief price description of AWS EC2 instances",
	Long:  `Example: awsprice get ec2 --os Linux --type t2.small --tenancy Shared`,
	Run: func(cmd *cobra.Command, args []string) {
		tenancy := cmd.Flag("tenancy").Value.String()
		itype := cmd.Flag("type").Value.String()
		opsys := cmd.Flag("os").Value.String()
		resource.GetEc2Price(itype, opsys, tenancy)
	},
}

func init() {
	getCmd.AddCommand(getec2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//getec2Cmd.Flags().StringP("ec2", "", viper.GetString("EC2"), "set a ec2")
	// getec2Cmd.PersistentFlags().String("foo", "", "A help for foo")
	getec2Cmd.PersistentFlags().String("type", "t2.small", "instance Type")
	getec2Cmd.PersistentFlags().String("os", "Linux", "Operating System")
	getec2Cmd.PersistentFlags().String("tenancy", "Shared", "Tenancy")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getec2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
