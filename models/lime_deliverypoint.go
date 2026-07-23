package models

type LimeDeliverypoint struct {
	ID       uint `json:"_id"`
	Address  string
	Streetno string
	Zipcode  string
	City     string
}

type LimeDeliverypoint2 struct {
	ID       uint `json:"_id"`
	Address  string
	Streetno string
	Zipcode  string
	City     string
}

/*
"relation_deliverypoint": {
	"fulladdress": "Flera Leveranser ,",
	"zipcode": "",
	"city": "",
	"address": "Flera Leveranser",
	"streetno": "",
	"_id": 1021,
	"_timestamp": "2023-06-19T16:07:30.120000+02:00",
	"_descriptive": "Flera Leveranser ,",
	"_updateduser": 1,
	"_createduser": 1,
	"_createdtime": "2023-06-19T16:07:30.080000+02:00",
	"_links": {
		"limetype": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliverypoint/",
			"name": "deliverypoint"
		  },
		  "self": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/"
		  },
		  "relation_deliveries": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deliveries/",
			"name": "deliveries"
		  },
		  "relation_deliveries2": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deliveries2/",
			"name": "deliveries2"
		  },
		  "relation_agreement": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/agreement/",
			"name": "agreement"
		  },
		  "relation_agreement2": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/agreement2/",
			"name": "agreement2"
		  },
		  "relation_deal2": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deal2/",
			"name": "deal2"
		  },
		  "relation_deal": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deal/",
			"name": "deal"
		  },
		  "new_deliveries": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deliveries/new/"
		  },
		  "new_deliveries2": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deliveries2/new/"
		  },
		  "new_agreement": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/agreement/new/"
		  },
		  "new_agreement2": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/agreement2/new/"
		  },
		  "new_deal2": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deal2/new/"
		  },
		  "new_deal": {
			"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deal/new/"
		  }
		}
	  },

*/
