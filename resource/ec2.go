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

package resource

import (
	"os"
	"strings"

	"github.com/chronojam/aws-pricing-api/types/schema"
	"github.com/olekukonko/tablewriter"
)

func GetEc2Price(instantType string, opsys string, tenancy string) {
	ec2 := &schema.AmazonEC2{}
	err := ec2.Refresh()
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Description", "USD"})
	table.SetRowLine(true)
	data := []*schema.AmazonEC2_Product{}
	for _, price := range ec2.Products {
		if price.Attributes.InstanceType == instantType && price.Attributes.OperatingSystem == opsys && price.Attributes.Tenancy == tenancy {
			data = append(data, price)
		}
	}

	for _, p := range data {
		for _, term := range ec2.Terms {
			if term.Sku == p.Sku {
				for _, priceData := range term.PriceDimensions {
					if strings.Contains(priceData.Description, "On Demand") {
						x := []string{}
						v := append(x, priceData.Description, priceData.PricePerUnit.USD)
						table.Append(v)
					}
				}
			}
		}
	}
	table.Render()
}
