package models

type LimeAgreementQoS struct {
	ID   uint   `json:"id"`
	Key  string `json:"key"`
	Text string `json:"text"`
}

type LimeAgreementStatus struct {
	ID   uint   `json:"id"`
	Key  string `json:"key"`
	Text string `json:"text"`
}

type LimeAgreement struct {
	ID         uint                 `json:"_id"`
	MonthlyFee int                  `json:"monthly_fee"`
	OnetimeFee int                  `json:"onetime_fee"`
	QOS        *LimeAgreementQoS    `json:"qos" gorm:"-"`
	Status     *LimeAgreementStatus `json:"status" gorm:"-"`
}

// ----------------------------------------------------------------------------------------

/*
"relation_agreement": {
	"qos": {
	  "id": 538201,
	  "key": "sn0",
	  "text": "SN0"
	},
	"placeholder2": "",
	"other_servicedescription": "",
	"agreement_status": {
	  "id": 552601,
	  "key": "active",
	  "text": "Active"
	},
	"invoice_row3": "356",
	"end_reason": "",
	"internal_comment": "Region Norrbotten PÄS-Sbyn",
	"endclient_contactphone": "",
	"endclient_contactemail": "",
	"placeholder": "",
	"endclient_contactname": "",
	"monthly_fee": 17748,
	"specific_term": "",
	"placeholder3": "",
	"deliverypoint2": null,
	"deliverynode": null,
	"order_value": 1079880,
	"endclient": "",
	"deliverynode2": null,
	"agreementno": "20-73",
	"qinq_option": {
	  "id": 553901,
	  "key": "nej",
	  "text": "No"
	},
	"invoice_row2": "",
	"service": 1111,
	"agreement_period": 60,
	"company": 1014,
	"placeholder4": "",
	"placeholder5": "",
	"person": null,
	"interface": {
	  "id": 525901,
	  "key": "empty",
	  "text": ""
	},
	"deal": null,
	"agreement_name": "",
	"agreement_startdate": "2022-01-30T00:00:00+01:00",
	"agreement_enddate": "2027-12-31T00:00:00+01:00",
	"onetime_fee": 15000,
	"client_end_date": null,
	"qinq": false,
	"connection_number": "",
	"invoice_reference": "",
	"invoicing_stopped": false,
	"fiber_length": 0.0,
	"deliverypoint": null,
	"estimated_delivery": null,
	"product": 1001,
	"client_end_yn": false,
	"coworker": 1007,
	"endclient_phone": "",
	"last_invoice_date": null,
	"iru_fee": 0,
	"invoice_row1": "",
	"_id": 1609,
	"_timestamp": "2023-09-25T21:41:06.740000+02:00",
	"_descriptive": "20-73",
	"_updateduser": 1,
	"_createduser": 1,
	"_createdtime": "2023-06-21T10:28:16.963000+02:00",
	"_links": {
	  "limetype": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
		"name": "agreement"
	  },
	  "self": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/"
	  },
	  "relation_document": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/document/",
		"name": "document"
	  },
	  "relation_deliveries": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliveries/",
		"name": "deliveries"
	  },
	  "relation_deliverypoint2": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliverypoint2/",
		"name": "deliverypoint2"
	  },
	  "relation_deliverynode": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliverynode/",
		"name": "deliverynode"
	  },
	  "relation_deliverynode2": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliverynode2/",
		"name": "deliverynode2"
	  },
	  "relation_service": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/service/",
		"name": "service"
	  },
	  "relation_history": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/history/",
		"name": "history"
	  },
	  "relation_company": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/company/",
		"name": "company"
	  },
	  "relation_person": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/person/",
		"name": "person"
	  },
	  "relation_deal": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deal/",
		"name": "deal"
	  },
	  "relation_invoice": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/invoice/",
		"name": "invoice"
	  },
	  "relation_deliverypoint": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliverypoint/",
		"name": "deliverypoint"
	  },
	  "relation_product": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/product/",
		"name": "product"
	  },
	  "relation_coworker": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/coworker/",
		"name": "coworker"
	  },
	  "relation_signing": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/signing/",
		"name": "signing"
	  },
	  "new_document": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/document/new/"
	  },
	  "new_deliveries": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliveries/new/"
	  },
	  "new_history": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/history/new/"
	  },
	  "new_invoice": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/invoice/new/"
	  },
	  "new_signing": {
		"href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/signing/new/"
	  }
	}
  },

*/
