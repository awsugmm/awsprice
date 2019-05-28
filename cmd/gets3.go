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

// gets3Cmd represents the gets3 command
var gets3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "A brief price description of Amazon S3",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//usageType := cmd.Flag("usage-type").Value.String()
		//location := cmd.Flag("location").Value.String()
		resource.GetS3Price()
	},
}

func init() {
	getCmd.AddCommand(gets3Cmd)

	// Here you will define your flags and configuration settings.
	//gets3Cmd.PersistentFlags().String("usage-type", "", "Usage Type")
	//gets3Cmd.PersistentFlags().String("location", "", "Location")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gets3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gets3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
