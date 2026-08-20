package limetool

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abundo/limetool/models"
)

func TestLimePersonJSONRoundTrip(t *testing.T) {
	raw := []byte(`{
		"_embedded": {
			"limeobjects": [
				{
					"firstname": "Anders",
					"inactive": false,
					"phone": "0920-28 43 34",
					"mobilephone": "070-111",
					"email": "anders.edstrom@norrbotten.se",
					"company": 1014,
					"lastname": "Edström",
					"name": "Anders Edström",
					"position": "Nätverkstekniker",
					"_id": 1011
				}
			]
		},
		"_links": {}
	}`)

	var resp models.PersonResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Embedded.Persons) != 1 {
		t.Fatalf("got %d persons, want 1", len(resp.Embedded.Persons))
	}
	p := resp.Embedded.Persons[0]
	if p.ID != 1011 {
		t.Errorf("ID = %d, want 1011", p.ID)
	}
	if p.Name != "Anders Edström" {
		t.Errorf("Name = %q", p.Name)
	}
	if p.Email != "anders.edstrom@norrbotten.se" {
		t.Errorf("Email = %q", p.Email)
	}
	if p.Phone != "0920-28 43 34" {
		t.Errorf("Phone = %q", p.Phone)
	}
	if p.Mobilephone != "070-111" {
		t.Errorf("Mobilephone = %q", p.Mobilephone)
	}
	if p.Company != 1014 {
		t.Errorf("Company = %d, want 1014", p.Company)
	}
	if p.Inactive {
		t.Error("Inactive = true, want false")
	}
}

func TestLimePersonDisplayNameAndPhone(t *testing.T) {
	named := models.LimePerson{Name: "Full Name", Firstname: "A", Lastname: "B", Phone: "1", Mobilephone: "2"}
	if named.DisplayName() != "Full Name" {
		t.Errorf("DisplayName = %q, want Full Name", named.DisplayName())
	}
	if named.PhoneNumber() != "1" {
		t.Errorf("PhoneNumber = %q, want desk phone", named.PhoneNumber())
	}

	fallback := models.LimePerson{Firstname: "Ada", Lastname: "Lovelace", Mobilephone: "070"}
	if fallback.DisplayName() != "Ada Lovelace" {
		t.Errorf("DisplayName = %q, want Ada Lovelace", fallback.DisplayName())
	}
	if fallback.PhoneNumber() != "070" {
		t.Errorf("PhoneNumber = %q, want mobile fallback", fallback.PhoneNumber())
	}
}

func TestGetPersonsAppendsActiveAndInactive(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/hal+json")
		_, _ = w.Write([]byte(`{
			"_embedded": {
				"limeobjects": [
					{"_id": 1, "name": "Active Person", "email": "a@example.com", "inactive": false, "company": 9},
					{"_id": 2, "name": "Gone Person", "email": "b@example.com", "inactive": true, "company": 9}
				]
			},
			"_links": {}
		}`))
	}))
	defer srv.Close()

	lime := NewLime(srv.URL, "test-key")
	company := &models.LimeCompany{
		ID:   9,
		Name: "Acme",
		Links: map[string]models.LinkType{
			"relation_person": {Href: srv.URL + "/person/", Name: "person"},
		},
	}
	if err := lime.GetPersons(company, true); err != nil {
		t.Fatalf("GetPersons: %v", err)
	}
	if len(company.Persons) != 2 {
		t.Fatalf("got %d persons, want 2", len(company.Persons))
	}
	if company.Persons[0].Name != "Active Person" || company.Persons[1].Inactive != true {
		t.Errorf("unexpected persons: %+v", company.Persons)
	}
}

func TestGetPersonsEmptyListIsSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/hal+json")
		_, _ = w.Write([]byte(`{"_embedded": {"limeobjects": []}, "_links": {}}`))
	}))
	defer srv.Close()

	lime := NewLime(srv.URL, "test-key")
	company := &models.LimeCompany{
		Links: map[string]models.LinkType{
			"relation_person": {Href: srv.URL + "/person/", Name: "person"},
		},
	}
	if err := lime.GetPersons(company, true); err != nil {
		t.Fatalf("GetPersons: %v", err)
	}
	if len(company.Persons) != 0 {
		t.Errorf("got %d persons, want 0", len(company.Persons))
	}
}

func TestGetPersonsMissingRelationIsSuccess(t *testing.T) {
	lime := NewLime("http://unused", "test-key")
	company := &models.LimeCompany{Name: "no-people-link"}
	if err := lime.GetPersons(company, true); err != nil {
		t.Fatalf("GetPersons: %v", err)
	}
	if len(company.Persons) != 0 {
		t.Errorf("got %d persons, want 0", len(company.Persons))
	}
}
