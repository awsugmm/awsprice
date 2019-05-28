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
