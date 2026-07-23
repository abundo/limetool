package models

import (
	"encoding/json"
	"time"
)

type LimeDeliveryStatus struct {
	ID   uint `json:"_id"`
	Key  string
	Text string
	// "delivery_status": {
	//   "id": 525701,
	//   "key": "done",
	//   "text": "Done"
	// },
}

type LimeDelivery struct {
	ID      uint `json:"_id"`
	Updated bool `json:"updated"`

	Name    string `json:"name"`
	Company int    `json:"company"`
	Comment string `json:"comment"`
	Person  int    `json:"person"`

	DeliveryDate         time.Time `json:"delivery_date"`
	ExpectedDeliveryDate time.Time `json:"expected_delivery"`
	OrderDate            time.Time `json:"order_date"`
	StartedDate          time.Time `json:"started"`

	CodePortA          string `json:"code_port"`
	CodePortB          string `json:"code_port_b"`
	ConnectionNumber   string `json:"connection_number"`
	ConnectionDetailsA string `json:"connection_details_a"`
	ConnectionDetailsB string `json:"connection_details_b"`

	ODF_equipmentA string `json:"odf_equipment"`
	ODF_equipmentB string `json:"odf_equipment_b"`

	ElectrictySubscript string `json:"electricity_subscript"`
	Waydescription      string `json:"waydescription"`
	Alarm               string `json:"alarm"`
	KeyAccess           string `json:"key_access"`
	Space               string `json:"space"`
	StandA              string `json:"stand"`
	StandB              string `json:"stand_b"`
	RoomA               string `json:"room"`
	RoomB               string `json:"room_b"`

	PowerIntake bool `json:"power_intake"`
	PowerPlant  bool `json:"power_plant"`
	Deal        int  `json:"deal"`

	Agreeement       *LimeAgreement      `json:"-"`
	AgreementID      int                 `json:"agreement"`
	DeliveryPoint    *LimeDeliverypoint  `json:"-"`
	DeliveryPointID  int                 `json:"deliverypoint"`
	DeliveryPoint2   *LimeDeliverypoint2 `json:"-"`
	DeliveryPoint2ID int                 `json:"deliverypoint2"`
	Product          *LimeProduct        `json:"-"`
	ProductID        int                 `json:"product"`
	Service          *LimeService        `json:"-"`
	ServiceID        int                 `json:"service"`

	DeliveryStatus *LimeDeliveryStatus        `json:"_delivery_status"`
	Links          map[string]LinkType        `json:"_links"`
	Embedded       map[string]json.RawMessage `json:"_embedded"`
}

type DeliveryObjects struct {
	Delivery []LimeDelivery `json:"limeobjects"`
}

type DeliveryResponse struct {
	Links    map[string]LinkType `json:"_links"`
	Embedded DeliveryObjects     `json:"_embedded"`
}

/*

{
  "_links": {
    "self": {
      "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/deliveries/?_embed=product&_embed=service&_embed=agreement&_embed=deliverypoint&_embed=deliverypoint2"
    },
    "next": {
      "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/deliveries/?_embed=product&_embed=service&_embed=agreement&_embed=deliverypoint&_embed=deliverypoint2&_offset=10"
    }
  },
  "_embedded": {
    "limeobjects": [
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "Kompletteras senare",
        "agreement": 1554,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2021-01-01T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2021-01-01T00:00:00+01:00",
        "started": "2021-01-01T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": null,
        "comment": "Region Norrbotten Lokala accesser",
        "product": 1004,
        "service": 1108,
        "deliverypoint2": null,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2021-01-01T00:00:00+01:00",
        "_id": 1351,
        "_timestamp": "2025-04-11T16:26:12.630000+02:00",
        "_descriptive": "Kompletteras senare",
        "_updateduser": 4201,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:08:45.877000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "296",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "20-73",
            "connection_number": "",
            "agreement_startdate": "2021-01-01T00:00:00+01:00",
            "agreement_enddate": "2025-12-31T00:00:00+01:00",
            "agreement_period": 60,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": null,
            "service": 1108,
            "onetime_fee": 0,
            "monthly_fee": 116761,
            "agreement_name": "",
            "deliverypoint2": null,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1004,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": null,
            "deliverynode2": null,
            "internal_comment": "Region Norrbotten Lokala accesser",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 7005660,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
            "_id": 1554,
            "_timestamp": "2025-06-11T13:18:39.090000+02:00",
            "_descriptive": "20-73",
            "_updateduser": 3701,
            "_createduser": 1,
            "_createdtime": "2023-06-21T10:28:04.277000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1554/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "Hyra av fiber -",
            "product": 1004,
            "service": "Hyra av fiber",
            "speed": "",
            "mtu": "",
            "inactive": false,
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
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deliveries/new/"
              }
            }
          },
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
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "Kompletteras senare",
        "agreement": 1609,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2022-01-01T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2022-01-30T00:00:00+01:00",
        "started": "2022-01-01T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": null,
        "comment": "Region Norrbotten PÄS-Sbyn",
        "product": 1001,
        "service": 1111,
        "deliverypoint2": null,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2022-01-01T00:00:00+01:00",
        "_id": 1362,
        "_timestamp": "2023-06-21T16:08:48.917000+02:00",
        "_descriptive": "Kompletteras senare",
        "_updateduser": 1,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:08:48.877000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1362/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "356",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "20-73",
            "connection_number": "",
            "agreement_startdate": "2022-01-30T00:00:00+01:00",
            "agreement_enddate": "2027-12-31T00:00:00+01:00",
            "agreement_period": 60,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": null,
            "service": 1111,
            "onetime_fee": 15000,
            "monthly_fee": 17748,
            "agreement_name": "",
            "deliverypoint2": null,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1001,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": null,
            "deliverynode2": null,
            "internal_comment": "Region Norrbotten PÄS-Sbyn",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 1079880,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
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
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1609/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "EVPN -",
            "product": 1001,
            "service": "EVPN",
            "speed": "",
            "mtu": "",
            "inactive": false,
            "_id": 1111,
            "_timestamp": "2023-06-20T14:54:24.217000+02:00",
            "_descriptive": "EVPN -",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2023-06-20T14:54:24.150000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/service/",
                "name": "service"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1111/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1111/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1111/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1111/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1111/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1111/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1111/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1111/deliveries/new/"
              }
            }
          },
          "relation_product": {
            "name": "Kapacitet",
            "initial": "CN",
            "_id": 1001,
            "_timestamp": "2023-01-23T15:01:03.250000+01:00",
            "_descriptive": "Kapacitet",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T11:54:11.717000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/product/",
                "name": "product"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/",
                "name": "service"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "Kompletteras senare",
        "agreement": 1617,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2021-01-01T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2022-05-01T00:00:00+02:00",
        "started": "2021-01-01T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": null,
        "comment": "Region Norrbotten Singelfiber Björksgatan",
        "product": 1004,
        "service": 1108,
        "deliverypoint2": null,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2021-01-01T00:00:00+01:00",
        "_id": 1368,
        "_timestamp": "2025-04-11T16:26:00.770000+02:00",
        "_descriptive": "Kompletteras senare",
        "_updateduser": 4201,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:08:50.627000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1368/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "365",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "20-73",
            "connection_number": "",
            "agreement_startdate": "2022-05-01T00:00:00+02:00",
            "agreement_enddate": "2025-12-31T00:00:00+01:00",
            "agreement_period": 60,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": null,
            "service": 1108,
            "onetime_fee": 45000,
            "monthly_fee": 1469,
            "agreement_name": "",
            "deliverypoint2": null,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1004,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": null,
            "deliverynode2": null,
            "internal_comment": "Region Norrboten Singelfiber Björksgatan",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 133140,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
            "_id": 1617,
            "_timestamp": "2023-09-25T21:41:06.917000+02:00",
            "_descriptive": "20-73",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2023-06-21T10:28:19.040000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1617/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "Hyra av fiber -",
            "product": 1004,
            "service": "Hyra av fiber",
            "speed": "",
            "mtu": "",
            "inactive": false,
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
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deliveries/new/"
              }
            }
          },
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
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "Kompletteras senare",
        "agreement": 1622,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2021-01-01T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2022-06-01T00:00:00+02:00",
        "started": "2021-01-01T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": null,
        "comment": "Region Norrbotten Lokala accesser Piteå",
        "product": 1004,
        "service": 1108,
        "deliverypoint2": null,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2021-01-01T00:00:00+01:00",
        "_id": 1370,
        "_timestamp": "2025-04-11T16:25:52.893000+02:00",
        "_descriptive": "Kompletteras senare",
        "_updateduser": 4201,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:08:51.237000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1370/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "373",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "20-73",
            "connection_number": "",
            "agreement_startdate": "2022-06-01T00:00:00+02:00",
            "agreement_enddate": "2025-12-31T00:00:00+01:00",
            "agreement_period": 60,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": null,
            "service": 1108,
            "onetime_fee": 0,
            "monthly_fee": 16515,
            "agreement_name": "",
            "deliverypoint2": null,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1004,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": null,
            "deliverynode2": null,
            "internal_comment": "Region Norrboten Lokala accesser Piteå",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 990900,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
            "_id": 1622,
            "_timestamp": "2023-09-25T21:41:07.083000+02:00",
            "_descriptive": "20-73",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2023-06-21T10:28:20.360000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1622/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "Hyra av fiber -",
            "product": 1004,
            "service": "Hyra av fiber",
            "speed": "",
            "mtu": "",
            "inactive": false,
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
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1108/deliveries/new/"
              }
            }
          },
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
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1004/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "Flera Leveranser",
        "agreement": 1279,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2006-01-01T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2005-01-21T00:00:00+01:00",
        "started": "2006-01-01T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": 1021,
        "comment": "Region Norrbotten WAN tjänst",
        "product": 1001,
        "service": 1042,
        "deliverypoint2": null,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2006-01-01T00:00:00+01:00",
        "_id": 1386,
        "_timestamp": "2025-04-20T00:38:04.200000+02:00",
        "_descriptive": "Flera Leveranser",
        "_updateduser": 4201,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:08:55.707000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1386/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "17",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "S50",
            "connection_number": "",
            "agreement_startdate": "2005-01-21T00:00:00+01:00",
            "agreement_enddate": "2028-05-01T00:00:00+02:00",
            "agreement_period": 34,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": 1021,
            "service": 1042,
            "onetime_fee": 0,
            "monthly_fee": 346925,
            "agreement_name": "",
            "deliverypoint2": 1021,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1001,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": 1037,
            "deliverynode2": 1037,
            "internal_comment": "Region Norrbotten WAN tjänst",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 11795450,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
            "_id": 1279,
            "_timestamp": "2025-06-11T13:30:50.287000+02:00",
            "_descriptive": "S50",
            "_updateduser": 3701,
            "_createduser": 1,
            "_createdtime": "2023-06-21T10:26:56.467000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1279/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "L3VPN - 1Gb",
            "product": 1001,
            "service": "L3VPN",
            "speed": "1Gb",
            "mtu": "1508 - 9216",
            "inactive": false,
            "_id": 1042,
            "_timestamp": "2023-01-31T21:18:51.453000+01:00",
            "_descriptive": "L3VPN - 1Gb",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T12:06:17.173000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/service/",
                "name": "service"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1042/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1042/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1042/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1042/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1042/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1042/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1042/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1042/deliveries/new/"
              }
            }
          },
          "relation_deliverypoint": {
            "address": "Flera Leveranser",
            "streetno": "",
            "zipcode": "",
            "city": "",
            "fulladdress": "Flera Leveranser ,",
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
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deal/",
                "name": "deal"
              },
              "relation_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deal2/",
                "name": "deal2"
              },
              "relation_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/agreement2/",
                "name": "agreement2"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deliveries/",
                "name": "deliveries"
              },
              "relation_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deliveries2/",
                "name": "deliveries2"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deal/new/"
              },
              "new_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deal2/new/"
              },
              "new_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/agreement2/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deliveries/new/"
              },
              "new_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1021/deliveries2/new/"
              }
            }
          },
          "relation_product": {
            "name": "Kapacitet",
            "initial": "CN",
            "_id": 1001,
            "_timestamp": "2023-01-23T15:01:03.250000+01:00",
            "_descriptive": "Kapacitet",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T11:54:11.717000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/product/",
                "name": "product"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/",
                "name": "service"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "CN00288",
        "agreement": 1448,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2017-11-13T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2017-11-13T00:00:00+01:00",
        "started": "2017-11-13T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": null,
        "comment": "Gällivare Flygplats, 1 Gig kapacitet.",
        "product": 1001,
        "service": 1017,
        "deliverypoint2": null,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2017-11-13T00:00:00+01:00",
        "_id": 1488,
        "_timestamp": "2023-06-21T16:09:27.653000+02:00",
        "_descriptive": "CN00288",
        "_updateduser": 1,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:09:27.603000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1488/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "196",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "N/A",
            "connection_number": "",
            "agreement_startdate": "2017-11-13T00:00:00+01:00",
            "agreement_enddate": "2026-06-01T00:00:00+02:00",
            "agreement_period": 36,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": null,
            "service": 1017,
            "onetime_fee": 350,
            "monthly_fee": 3900,
            "agreement_name": "",
            "deliverypoint2": null,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1001,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": null,
            "deliverynode2": null,
            "internal_comment": "Gällivare Flygplats, 1 Gig kapacitet.",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 140750,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
            "_id": 1448,
            "_timestamp": "2023-09-25T21:41:03.840000+02:00",
            "_descriptive": "N/A",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2023-06-21T10:27:40.340000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1448/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "EVPN - 1Gb",
            "product": 1001,
            "service": "EVPN",
            "speed": "1Gb",
            "mtu": "<1508",
            "inactive": false,
            "_id": 1017,
            "_timestamp": "2023-01-31T21:18:51.453000+01:00",
            "_descriptive": "EVPN - 1Gb",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T12:06:12.687000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/service/",
                "name": "service"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1017/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1017/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1017/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1017/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1017/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1017/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1017/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1017/deliveries/new/"
              }
            }
          },
          "relation_product": {
            "name": "Kapacitet",
            "initial": "CN",
            "_id": 1001,
            "_timestamp": "2023-01-23T15:01:03.250000+01:00",
            "_descriptive": "Kapacitet",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T11:54:11.717000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/product/",
                "name": "product"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/",
                "name": "service"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "CN00029-1",
        "agreement": 1453,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2018-01-01T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2018-01-01T00:00:00+01:00",
        "started": "2018-01-01T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": 1219,
        "comment": "Internet 1000Mb",
        "product": 1001,
        "service": 1067,
        "deliverypoint2": null,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2018-01-01T00:00:00+01:00",
        "_id": 1494,
        "_timestamp": "2025-04-20T00:47:10.283000+02:00",
        "_descriptive": "CN00029-1",
        "_updateduser": 4201,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:09:29.097000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1494/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "201",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "N/A",
            "connection_number": "",
            "agreement_startdate": "2018-01-01T00:00:00+01:00",
            "agreement_enddate": "2020-12-31T00:00:00+01:00",
            "agreement_period": 36,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": null,
            "service": 1067,
            "onetime_fee": 0,
            "monthly_fee": 6900,
            "agreement_name": "",
            "deliverypoint2": null,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1001,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": 1038,
            "deliverynode2": null,
            "internal_comment": "Internet 1000Mb",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 248400,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
            "_id": 1453,
            "_timestamp": "2023-09-25T21:41:03.993000+02:00",
            "_descriptive": "N/A",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2023-06-21T10:27:41.350000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1453/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "Polarix - 1Gb",
            "product": 1001,
            "service": "Polarix",
            "speed": "1Gb",
            "mtu": "9216",
            "inactive": false,
            "_id": 1067,
            "_timestamp": "2023-01-31T21:18:51.453000+01:00",
            "_descriptive": "Polarix - 1Gb",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T12:06:21.133000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/service/",
                "name": "service"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1067/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1067/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1067/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1067/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1067/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1067/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1067/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1067/deliveries/new/"
              }
            }
          },
          "relation_deliverypoint": {
            "address": "Sjukhusvägen",
            "streetno": "10",
            "zipcode": "95442",
            "city": "Luleå",
            "fulladdress": "Sjukhusvägen 10, Luleå",
            "_id": 1219,
            "_timestamp": "2023-12-22T13:28:12.183000+01:00",
            "_descriptive": "Sjukhusvägen 10, Luleå",
            "_updateduser": 3701,
            "_createduser": 3701,
            "_createdtime": "2023-12-22T13:28:12.147000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliverypoint/",
                "name": "deliverypoint"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/deal/",
                "name": "deal"
              },
              "relation_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/deal2/",
                "name": "deal2"
              },
              "relation_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/agreement2/",
                "name": "agreement2"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/deliveries/",
                "name": "deliveries"
              },
              "relation_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/deliveries2/",
                "name": "deliveries2"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/deal/new/"
              },
              "new_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/deal2/new/"
              },
              "new_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/agreement2/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/deliveries/new/"
              },
              "new_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1219/deliveries2/new/"
              }
            }
          },
          "relation_product": {
            "name": "Kapacitet",
            "initial": "CN",
            "_id": 1001,
            "_timestamp": "2023-01-23T15:01:03.250000+01:00",
            "_descriptive": "Kapacitet",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T11:54:11.717000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/product/",
                "name": "product"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/",
                "name": "service"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "CN00342",
        "agreement": 1485,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2019-01-30T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2019-01-30T00:00:00+01:00",
        "started": "2019-01-30T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": null,
        "comment": "SuS-Pite Älvdals sjukhus, 10 Gig",
        "product": 1002,
        "service": 1084,
        "deliverypoint2": 1116,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2019-01-30T00:00:00+01:00",
        "_id": 1516,
        "_timestamp": "2025-04-16T14:14:11.777000+02:00",
        "_descriptive": "CN00342",
        "_updateduser": 4201,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:09:35.137000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1516/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "233",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "20-73",
            "connection_number": "",
            "agreement_startdate": "2019-01-30T00:00:00+01:00",
            "agreement_enddate": "2022-01-29T00:00:00+01:00",
            "agreement_period": 36,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": null,
            "service": 1084,
            "onetime_fee": 40000,
            "monthly_fee": 3000,
            "agreement_name": "",
            "deliverypoint2": null,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1002,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": null,
            "deliverynode2": 1127,
            "internal_comment": "SuS-Pite Älvdals sjukhus, 10 Gig",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 148000,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
            "_id": 1485,
            "_timestamp": "2023-09-25T21:41:04.503000+02:00",
            "_descriptive": "20-73",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2023-06-21T10:27:48.357000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1485/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "Våglängdstjänst - 10Gb",
            "product": 1002,
            "service": "Våglängdstjänst",
            "speed": "10Gb",
            "mtu": "",
            "inactive": false,
            "_id": 1084,
            "_timestamp": "2023-01-31T21:18:51.453000+01:00",
            "_descriptive": "Våglängdstjänst - 10Gb",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T12:06:23.667000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/service/",
                "name": "service"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1084/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1084/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1084/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1084/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1084/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1084/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1084/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1084/deliveries/new/"
              }
            }
          },
          "relation_deliverypoint2": {
            "address": "Sjukhus",
            "streetno": "",
            "zipcode": "",
            "city": "Piteå",
            "fulladdress": "Sjukhus , Piteå",
            "_id": 1116,
            "_timestamp": "2025-04-11T10:40:30.217000+02:00",
            "_descriptive": "Sjukhus , Piteå",
            "_updateduser": 4201,
            "_createduser": 1,
            "_createdtime": "2023-06-19T16:07:44.723000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliverypoint/",
                "name": "deliverypoint"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/deal/",
                "name": "deal"
              },
              "relation_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/deal2/",
                "name": "deal2"
              },
              "relation_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/agreement2/",
                "name": "agreement2"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/deliveries/",
                "name": "deliveries"
              },
              "relation_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/deliveries2/",
                "name": "deliveries2"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/deal/new/"
              },
              "new_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/deal2/new/"
              },
              "new_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/agreement2/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/deliveries/new/"
              },
              "new_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1116/deliveries2/new/"
              }
            }
          },
          "relation_product": {
            "name": "Våglängd",
            "initial": "VL",
            "_id": 1002,
            "_timestamp": "2022-12-22T16:07:30.320000+01:00",
            "_descriptive": "Våglängd",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T11:54:25.693000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/product/",
                "name": "product"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/service/",
                "name": "service"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1002/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": null,
        "connection_number": "CN0007-144",
        "agreement": 1491,
        "deal": null,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2019-03-15T00:00:00+01:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2019-03-15T00:00:00+01:00",
        "started": "2019-03-15T00:00:00+01:00",
        "company": 1014,
        "deliverypoint": null,
        "comment": "100 Mb inkl QinQ Blodcentral Vårdhögskolan Boden",
        "product": 1001,
        "service": 1008,
        "deliverypoint2": null,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2019-03-15T00:00:00+01:00",
        "_id": 1521,
        "_timestamp": "2025-04-29T09:19:13.513000+02:00",
        "_descriptive": "CN0007-144",
        "_updateduser": 4201,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:09:36.357000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1521/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "239",
            "iru_fee": 0,
            "company": 1014,
            "deal": null,
            "person": null,
            "agreementno": "N/A",
            "connection_number": "",
            "agreement_startdate": "2019-03-15T00:00:00+01:00",
            "agreement_enddate": "2026-06-01T00:00:00+02:00",
            "agreement_period": 60,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": null,
            "service": 1008,
            "onetime_fee": 3500,
            "monthly_fee": 1950,
            "agreement_name": "",
            "deliverypoint2": null,
            "estimated_delivery": null,
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1001,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": null,
            "deliverynode2": null,
            "internal_comment": "100 Mb inkl QinQ Blodcentral Vårdhögskolan Boden",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 120500,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "",
            "_id": 1491,
            "_timestamp": "2023-09-25T21:41:04.503000+02:00",
            "_descriptive": "N/A",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2023-06-21T10:27:49.480000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1491/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "EVPN - 100Mb",
            "product": 1001,
            "service": "EVPN",
            "speed": "100Mb",
            "mtu": "<1508",
            "inactive": false,
            "_id": 1008,
            "_timestamp": "2024-02-02T11:09:21.087000+01:00",
            "_descriptive": "EVPN - 100Mb",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T12:06:11.260000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/service/",
                "name": "service"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1008/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1008/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1008/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1008/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1008/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1008/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1008/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1008/deliveries/new/"
              }
            }
          },
          "relation_product": {
            "name": "Kapacitet",
            "initial": "CN",
            "_id": 1001,
            "_timestamp": "2023-01-23T15:01:03.250000+01:00",
            "_descriptive": "Kapacitet",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T11:54:11.717000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/product/",
                "name": "product"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/",
                "name": "service"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/new/"
              }
            }
          }
        }
      },
      {
        "name": "",
        "room_b": "",
        "stand_b": "",
        "odf_equipment_b": "",
        "code_port_b": "",
        "person": 1011,
        "connection_number": "CN1799",
        "agreement": 1762,
        "deal": 1012,
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "delivery_date": "2023-09-22T00:00:00+02:00",
        "space": "",
        "stand": "",
        "odf_equipment": "",
        "code_port": "",
        "connection_details_a": "Kommentar",
        "connection_details_b": "",
        "waydescription": "",
        "key_access": "",
        "electricty_subscript": "",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "alarm": "",
        "power_plant": false,
        "power_intake": false,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "room": "",
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "order_date": "2023-08-17T16:29:17.353000+02:00",
        "started": "2023-09-22T00:00:00+02:00",
        "company": 1014,
        "deliverypoint": 1086,
        "comment": "",
        "product": 1001,
        "service": 1071,
        "deliverypoint2": 1007,
        "deliverynode": null,
        "deliverynode2": null,
        "expected_delivery": "2023-07-01T00:00:00+02:00",
        "_id": 1799,
        "_timestamp": "2025-04-11T12:54:56.143000+02:00",
        "_descriptive": "CN1799",
        "_updateduser": 4201,
        "_createduser": 3701,
        "_createdtime": "2023-08-17T16:29:44.507000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/history/",
            "name": "history"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/todo/",
            "name": "todo"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/person/",
            "name": "person"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/agreement/",
            "name": "agreement"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/document/",
            "name": "document"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/deal/",
            "name": "deal"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/company/",
            "name": "company"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/product/",
            "name": "product"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/service/",
            "name": "service"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/deliverynode/",
            "name": "deliverynode"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/deliverynode2/",
            "name": "deliverynode2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/history/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/todo/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1799/document/new/"
          }
        },
        "_embedded": {
          "relation_agreement": {
            "invoice_reference": "Samlingsfaktura Regionen",
            "invoice_row1": "",
            "invoice_row2": "",
            "invoice_row3": "",
            "iru_fee": 0,
            "company": 1014,
            "deal": 1012,
            "person": 1017,
            "agreementno": "1762",
            "connection_number": "",
            "agreement_startdate": "2023-07-01T00:00:00+02:00",
            "agreement_enddate": "2028-06-30T00:00:00+02:00",
            "agreement_period": 60,
            "placeholder": "",
            "placeholder2": "",
            "placeholder5": "",
            "placeholder3": "",
            "placeholder4": "",
            "coworker": 1007,
            "deliverypoint": 1026,
            "service": 1071,
            "onetime_fee": 0,
            "monthly_fee": 13997,
            "agreement_name": "",
            "deliverypoint2": 1007,
            "estimated_delivery": "2023-07-01T00:00:00+02:00",
            "endclient": "",
            "endclient_contactname": "",
            "endclient_contactemail": "",
            "endclient_contactphone": "",
            "endclient_phone": "",
            "product": 1001,
            "interface": {
              "id": 525901,
              "key": "empty",
              "text": ""
            },
            "deliverynode": null,
            "deliverynode2": null,
            "internal_comment": "Detta avtal är en utökning av befintlig intenettjänst, därmed ersätter detta avtal tidigare avtal.",
            "qinq": false,
            "qos": {
              "id": 538201,
              "key": "sn0",
              "text": "SN0"
            },
            "specific_term": "Detta avtal är en utökning av befintlig intenettjänst, därmed ersätter detta avtal tidigare avtal.",
            "agreement_status": {
              "id": 552601,
              "key": "active",
              "text": "Active"
            },
            "client_end_yn": false,
            "client_end_date": null,
            "last_invoice_date": null,
            "invoicing_stopped": false,
            "order_value": 839820,
            "end_reason": "",
            "qinq_option": {
              "id": 553901,
              "key": "nej",
              "text": "No"
            },
            "fiber_length": 0.0,
            "other_servicedescription": "Detta avtal är en utökning av befintlig intenettjänst, därmed ersätter detta avtal tidigare avtal.",
            "_id": 1762,
            "_timestamp": "2023-09-22T10:58:55.393000+02:00",
            "_descriptive": "1762",
            "_updateduser": 3701,
            "_createduser": 3701,
            "_createdtime": "2023-08-17T16:28:37.163000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/agreement/",
                "name": "agreement"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/"
              },
              "relation_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/invoice/",
                "name": "invoice"
              },
              "relation_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/document/",
                "name": "document"
              },
              "relation_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/history/",
                "name": "history"
              },
              "relation_company": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/company/",
                "name": "company"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/deal/",
                "name": "deal"
              },
              "relation_person": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/person/",
                "name": "person"
              },
              "relation_coworker": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/coworker/",
                "name": "coworker"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/deliveries/",
                "name": "deliveries"
              },
              "relation_deliverypoint": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/deliverypoint/",
                "name": "deliverypoint"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/service/",
                "name": "service"
              },
              "relation_deliverypoint2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/deliverypoint2/",
                "name": "deliverypoint2"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/product/",
                "name": "product"
              },
              "relation_deliverynode": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/deliverynode/",
                "name": "deliverynode"
              },
              "relation_deliverynode2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/deliverynode2/",
                "name": "deliverynode2"
              },
              "relation_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/signing/",
                "name": "signing"
              },
              "new_invoice": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/invoice/new/"
              },
              "new_document": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/document/new/"
              },
              "new_history": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/history/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/deliveries/new/"
              },
              "new_signing": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/agreement/1762/signing/new/"
              }
            }
          },
          "relation_service": {
            "name": "Polarix - 5Gb",
            "product": 1001,
            "service": "Polarix",
            "speed": "5Gb",
            "mtu": "9216",
            "inactive": false,
            "_id": 1071,
            "_timestamp": "2023-01-31T21:18:51.453000+01:00",
            "_descriptive": "Polarix - 5Gb",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T12:06:21.643000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/service/",
                "name": "service"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1071/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1071/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1071/deal/",
                "name": "deal"
              },
              "relation_product": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1071/product/",
                "name": "product"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1071/deliveries/",
                "name": "deliveries"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1071/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1071/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/service/1071/deliveries/new/"
              }
            }
          },
          "relation_deliverypoint2": {
            "address": "Robertsviksgatan",
            "streetno": "7",
            "zipcode": "97241",
            "city": "Luleå",
            "fulladdress": "Robertsviksgatan 7, Luleå",
            "_id": 1007,
            "_timestamp": "2023-05-17T11:28:06.020000+02:00",
            "_descriptive": "Robertsviksgatan 7, Luleå",
            "_updateduser": 3701,
            "_createduser": 3701,
            "_createdtime": "2023-05-17T11:28:06.007000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliverypoint/",
                "name": "deliverypoint"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/deal/",
                "name": "deal"
              },
              "relation_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/deal2/",
                "name": "deal2"
              },
              "relation_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/agreement2/",
                "name": "agreement2"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/deliveries/",
                "name": "deliveries"
              },
              "relation_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/deliveries2/",
                "name": "deliveries2"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/deal/new/"
              },
              "new_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/deal2/new/"
              },
              "new_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/agreement2/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/deliveries/new/"
              },
              "new_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1007/deliveries2/new/"
              }
            }
          },
          "relation_deliverypoint": {
            "address": "Rådstugatan",
            "streetno": "11",
            "zipcode": "972 38",
            "city": "Luleå",
            "fulladdress": "Rådstugatan 11, Luleå",
            "_id": 1086,
            "_timestamp": "2025-04-11T10:41:56.837000+02:00",
            "_descriptive": "Rådstugatan 11, Luleå",
            "_updateduser": 4201,
            "_createduser": 1,
            "_createdtime": "2023-06-19T16:07:39.927000+02:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliverypoint/",
                "name": "deliverypoint"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/deal/",
                "name": "deal"
              },
              "relation_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/deal2/",
                "name": "deal2"
              },
              "relation_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/agreement2/",
                "name": "agreement2"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/deliveries/",
                "name": "deliveries"
              },
              "relation_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/deliveries2/",
                "name": "deliveries2"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/deal/new/"
              },
              "new_deal2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/deal2/new/"
              },
              "new_agreement2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/agreement2/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/deliveries/new/"
              },
              "new_deliveries2": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliverypoint/1086/deliveries2/new/"
              }
            }
          },
          "relation_product": {
            "name": "Kapacitet",
            "initial": "CN",
            "_id": 1001,
            "_timestamp": "2023-01-23T15:01:03.250000+01:00",
            "_descriptive": "Kapacitet",
            "_updateduser": 1,
            "_createduser": 1,
            "_createdtime": "2022-12-15T11:54:11.717000+01:00",
            "_links": {
              "limetype": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/product/",
                "name": "product"
              },
              "self": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/"
              },
              "relation_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/",
                "name": "service"
              },
              "relation_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/",
                "name": "agreement"
              },
              "relation_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/",
                "name": "deal"
              },
              "relation_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/",
                "name": "deliveries"
              },
              "new_service": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/service/new/"
              },
              "new_agreement": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/agreement/new/"
              },
              "new_deal": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deal/new/"
              },
              "new_deliveries": {
                "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/product/1001/deliveries/new/"
              }
            }
          }
        }
      }
    ]
  }
}



"_embedded": {
    "limeobjects": [
      {
        "name": "",
        "client_wavelenght": {
          "id": 513701,
          "key": "empty",
          "text": "N/A"
        },
        "expected_delivery": "2021-01-01T00:00:00+01:00",
        "type_of_node": {
          "id": 519901,
          "key": "519901",
          "text": ""
        },
        "stand_b": "",
        "comment": "Region Norrbotten Lokala accesser",
        "deliverynode2": null,
        "electricty_subscript": "",
        "signal_format": {
          "id": 513901,
          "key": "513901",
          "text": ""
        },
        "code_port": "",
        "waydescription": "",
        "alarm": "",
        "service": 1108,
        "deliverypoint": null,
        "person": null,
        "key_access": "",
        "space": "",
        "code_port_b": "",
        "company": 1014,
        "connection_details_b": "",
        "delivery_date": "2021-01-01T00:00:00+01:00",
        "odf_equipment_b": "",
        "stand": "",
        "room": "",
        "client_fiber": {
          "id": 513501,
          "key": "empty",
          "text": "N/A"
        },
        "room_b": "",
        "connection_number": "Kompletteras senare",
        "power_intake": false,
        "order_date": "2021-01-01T00:00:00+01:00",
        "agreement": 1554,
        "climate_ventilation": {
          "id": 520401,
          "key": "520401",
          "text": ""
        },
        "delivery_status": {
          "id": 525701,
          "key": "done",
          "text": "Done"
        },
        "deliverynode": null,
        "started": "2021-01-01T00:00:00+01:00",
        "power_plant": false,
        "product": 1004,
        "odf_equipment": "",
        "deal": null,
        "deliverypoint2": null,
        "_id": 1351,
        "_timestamp": "2025-04-11T16:26:12.630000+02:00",
        "_descriptive": "Kompletteras senare",
        "_updateduser": 4201,
        "_createduser": 1,
        "_createdtime": "2023-06-21T16:08:45.877000+02:00",
        "_links": {
          "limetype": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/deliveries/",
            "name": "deliveries"
          },
          "self": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/"
          },
          "relation_deliverynode2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deliverynode2/",
            "name": "deliverynode2"
          },
          "relation_service": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/service/",
            "name": "service"
          },
          "relation_deliverypoint": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deliverypoint/",
            "name": "deliverypoint"
          },
          "relation_person": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/person/",
            "name": "person"
          },
          "relation_company": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/company/",
            "name": "company"
          },
          "relation_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/history/",
            "name": "history"
          },
          "relation_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/document/",
            "name": "document"
          },
          "relation_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/todo/",
            "name": "todo"
          },
          "relation_agreement": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/agreement/",
            "name": "agreement"
          },
          "relation_deliverynode": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deliverynode/",
            "name": "deliverynode"
          },
          "relation_product": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/product/",
            "name": "product"
          },
          "relation_deal": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deal/",
            "name": "deal"
          },
          "relation_deliverypoint2": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/deliverypoint2/",
            "name": "deliverypoint2"
          },
          "new_history": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/history/new/"
          },
          "new_document": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/document/new/"
          },
          "new_todo": {
            "href": "https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/deliveries/1351/todo/new/"
          }
        },
*/
