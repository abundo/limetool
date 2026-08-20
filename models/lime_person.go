package models

import "strings"

type LimePerson struct {
	ID          uint   `json:"_id"`
	Name        string `json:"name"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Mobilephone string `json:"mobilephone"`
	Position    string `json:"position"`
	Company     uint   `json:"company"`
	Inactive    bool   `json:"inactive"`

	Links map[string]LinkType `json:"_links" gorm:"-"`
}

// DisplayName is Lime's combined name, falling back to first+last when empty.
func (p LimePerson) DisplayName() string {
	if p.Name != "" {
		return p.Name
	}
	return strings.TrimSpace(p.Firstname + " " + p.Lastname)
}

// PhoneNumber prefers the desk phone and falls back to mobile.
func (p LimePerson) PhoneNumber() string {
	if p.Phone != "" {
		return p.Phone
	}
	return p.Mobilephone
}

type PersonLimeObjects struct {
	Persons []LimePerson `json:"limeobjects"`
}

type PersonResponse struct {
	Embedded PersonLimeObjects   `json:"_embedded"`
	Links    map[string]LinkType `json:"_links"`
}

// ----------------------------------------------------------------------------------------

/*
{
   "_links":{
      "self":{
         "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/company/1014/person/"
      }
   },
   "_embedded":{
      "limeobjects":[
         {
            "firstname":"Anders",
            "inactive":false,
            "phone":"0920-28 43 34",
            "mobilephone":"",
            "email":"anders.edstrom@norrbotten.se",
            "company":1014,
            "lastname":"Edström",
            "name":"Anders Edström",
            "position":"Nätverkstekniker",
            "expireddate":null,
            "anonymizeddate":null,
            "_id":1011,
            "_timestamp":"2023-05-17T11:21:12.507000+02:00",
            "_descriptive":"Anders Edström",
            "_updateduser":3701,
            "_createduser":3701,
            "_createdtime":"2023-05-17T11:21:12.423000+02:00",
            "_links":{
               "limetype":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limetype/person/",
                  "name":"person"
               },
               "self":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/person/1011/"
               },
               "relation_company":{
                  "href":"https://itnorrbotten.lime-crm.com/informationsteknik%20i%20norrbotten%20ab/api/v1/limeobject/person/1011/company/",
                  "name":"company"
               }
            }
         }
      ]
   }
}
*/
