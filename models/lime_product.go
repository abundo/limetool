package models

type LimeProduct struct {
	ID   uint `json:"_id"`
	Name string
}

/*
"relation_product": {
	"name": "Svartfiber",
	"initial": "LF",
	"_id": 1004,
	"_timestamp": "2023-01-23T15:01:22.537000+01:00",
	"_descriptive": "Svartfiber",
	"_updateduser": 1,
	"_createduser": 1,
	"_createdtime": "2022-12-15T11:54:52.190000+01:00",
	"_links": {
		"limetype": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/product/",
		"name": "product"
		},
		"self": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/"
		},
		"relation_service": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/service/",
		"name": "service"
		},
		"relation_deal": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deal/",
		"name": "deal"
		},
		"relation_agreement": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/agreement/",
		"name": "agreement"
		},
		"relation_deliveries": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deliveries/",
		"name": "deliveries"
		},
		"new_service": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/service/new/"
		},
		"new_deal": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deal/new/"
		},
		"new_agreement": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/agreement/new/"
		},
		"new_deliveries": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deliveries/new/"
		}
	}
}

*/
