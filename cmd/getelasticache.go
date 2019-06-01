// Copyright © 2019 AWS User Group Myanmar <awsug.mm@gmail.com>
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

var getelasticacheCmd = &cobra.Command{
	Use:   "escache",
	Short: "A brief price description of AWS Elastic Cache",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		resource.GetElasticCache()
	},
}

func init() {
	getCmd.AddCommand(getelasticacheCmd)
}
