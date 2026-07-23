package models

type LimeService struct {
	ID         uint `json:"_id"`
	Product    int
	Speed      string
	Name       string
	Inactive   bool
	Service    string
	MTU        string
	CustomerID uint
}

/*
"relation_service": {
	"product": 1004,
	"speed": "",
	"name": "Hyra av fiber -",
	"inactive": false,
	"service": "Hyra av fiber",
	"mtu": "",
	"_id": 1108,
	"_timestamp": "2023-01-31T21:18:53.800000+01:00",
	"_descriptive": "Hyra av fiber -",
	"_updateduser": 1,
	"_createduser": 1,
	"_createdtime": "2022-12-15T12:06:26.160000+01:00",
	"_links": {
	  "limetype": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/service/",
		"name": "service"
	  },
	  "self": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/"
	  },
	  "relation_product": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/product/",
		"name": "product"
	  },
	  "relation_deal": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deal/",
		"name": "deal"
	  },
	  "relation_deliveries": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deliveries/",
		"name": "deliveries"
	  },
	  "relation_agreement": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/agreement/",
		"name": "agreement"
	  },
	  "new_deal": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deal/new/"
	  },
	  "new_deliveries": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deliveries/new/"
	  },
	  "new_agreement": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/agreement/new/"
	  }
	}
  },
`
*/
