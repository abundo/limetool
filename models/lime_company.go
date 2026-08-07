package models

type LimeCompany struct {
	ID   uint   `json:"_id"`
	Name string `json:"name"`

	EmailInvoice string
	Phone        string

	Postaladdress1 string
	Postaladdress2 string
	Postalcity     string
	Postalzipcode  string
	Country        string

	Visitingaddress1 string
	Visitingaddress2 string
	VistingZipcode   string
	Visitingcity     string

	Invoiceaddress1 string
	InvoiceAddress2 string
	Invoicezipcode  string
	Invoicecity     string

	Fullinvoiceaddress  string
	Fullvisitingaddress string
	Fullpostaladdress   string

	WWW      string
	Coworker int

	ERP_id           string
	ERP_firstsynced  string
	ERP_errormessage string

	Registrationno     string
	Latestsalescontact string

	NOC_email string
	NOC_phone string

	Deliveries []LimeDelivery `gorm:"-"`

	Links map[string]LinkType `json:"_links" gorm:"-"`
}

type CompanyLimeObjects struct {
	Customers []LimeCompany `json:"limeobjects"`
}

type CompanyResponse struct {
	Embedded CompanyLimeObjects  `json:"_embedded"`
	Links    map[string]LinkType `json:"_links"`
}

// ----------------------------------------------------------------------------------------

/*

{
   "_links":{
      "self":{
         "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/?name=Region+Norrbotten"
      }
   },
   "_embedded":{
      "limeobjects":[
         {
            "name":"Region Norrbotten",
            "phone":"0920-284000",
            "www":"",
            "postaladdress1":"Box 511",
            "visitingaddress1":"",
            "registrationno":"232100-0230",
            "noc_phone":"",
            "noc_email":"",
            "coworker":1007,
            "postalzipcode":"961 28",
            "postalcity":"BODEN",
            "visitingzipcode":"",
            "visitingcity":"",
            "postaladdress2":"",
            "visitingaddress2":"",
            "country":"Sverige",
            "buyingstatus":{
               "id":108001,
               "key":"empty",
               "text":""
            },
            "fullpostaladdress":"Box 511\n961 28 BODEN",
            "fullvisitingaddress":"",
            "invoiceaddress1":"Box 511",
            "invoiceaddress2":"",
            "invoicezipcode":"961 28",
            "invoicecity":"BODEN",
            "fullinvoiceaddress":"Box 511\n961 28 BODEN",
            "emailinvoice":"",
            "erp_id":"1",
            "erp_status":{
               "id":373501,
               "key":"synced",
               "text":"Synced"
            },
            "erp_turnover_yearnow":null,
            "erp_turnover_lastyeartodate":null,
            "erp_turnover_lastyear":null,
            "erp_salestrend":{
               "id":375801,
               "key":"empty",
               "text":""
            },
            "erp_ytdgrowth":null,
            "erp_errormessage":"",
            "erp_firstsynced":null,
            "erp_lastsynced":null,
            "classification":{
               "id":414301,
               "key":"empty",
               "text":""
            },
            "potential":{
               "id":415301,
               "key":"empty",
               "text":""
            },
            "latestsalescontact":null,
            "_id":1014,
            "_timestamp":"2023-09-06T08:04:11.857000+02:00",
            "_descriptive":"Region Norrbotten",
            "_updateduser":3701,
            "_createduser":3101,
            "_createdtime":"2023-02-13T12:58:26.270000+01:00",
            "_links":{
               "limetype":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/company/",
                  "name":"company"
               },
               "self":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/"
               },
               "relation_document":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/document/",
                  "name":"document"
               },
               "relation_person":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/person/",
                  "name":"person"
               },
               "relation_coworker":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/coworker/",
                  "name":"coworker"
               },
               "relation_deal":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/deal/",
                  "name":"deal"
               },
               "relation_todo":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/todo/",
                  "name":"todo"
               },
               "relation_history":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/history/",
                  "name":"history"
               },
               "relation_lead":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/lead/",
                  "name":"lead"
               },
               "relation_invoice":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/invoice/",
                  "name":"invoice"
               },
               "relation_invoicerow":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/invoicerow/",
                  "name":"invoicerow"
               },
               "relation_agreement":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/agreement/",
                  "name":"agreement"
               },
               "relation_signing":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/signing/",
                  "name":"signing"
               },
               "relation_deliveries":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/deliveries/",
                  "name":"deliveries"
               },
               "new_document":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/document/new/"
               },
               "new_person":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/person/new/"
               },
               "new_deal":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/deal/new/"
               },
               "new_todo":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/todo/new/"
               },
               "new_history":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/history/new/"
               },
               "new_lead":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/lead/new/"
               },
               "new_invoice":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/invoice/new/"
               },
               "new_invoicerow":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/invoicerow/new/"
               },
               "new_agreement":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/agreement/new/"
               },
               "new_signing":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/signing/new/"
               },
               "new_deliveries":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/deliveries/new/"
               }
            }
         }
      ]
   }
}
*/
