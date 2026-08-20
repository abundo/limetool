package limetool

//
// Lib for communicating with the Lime CRM REST API
//

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"log/slog"

	"github.com/abundo/limetool/models"
)

var Debug bool
var Verbose bool
var companiesCount int
var deliveriesCount int
var cache = true

type Lime struct {
	APIURL string
	APIKey string
}

func NewLime(APIURL string, APIKey string) *Lime {
	lime := new(Lime)
	lime.APIURL = APIURL
	lime.APIKey = APIKey
	return lime
}

// Handler for Lime REST API
func (lime *Lime) limeAPI(apiURL string, refresh bool) ([]byte, error) {
	slog.Debug("Query API", "URL", apiURL)
	var filename string
	if cache && !refresh {
		h := sha256.New()
		h.Write([]byte(apiURL))
		hash := h.Sum(nil)
		filename = fmt.Sprintf("/tmp/lime-%x", hash)
		bodyBytes, err := os.ReadFile(filename)
		if err == nil {
			slog.Debug("Using cached data", "from", filename)
			return bodyBytes, nil
		}
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		slog.Error("Error creating request:", "err", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/hal+json") // Indicate we expect a JSON response
	req.Header.Set("x-api-Key", lime.APIKey)         // Add the custom API key header
	// Send the request using the client
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error sending request to API endpoint", "err", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { // Check if the status code is 200 OK
		// Read the body even on error for potential error messages from the API
		bodyBytes, _ := io.ReadAll(resp.Body)
		slog.Error("API request failed with status ", "code", resp.StatusCode, "body", string(bodyBytes))
		return nil, err
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response", "body", err)
		return nil, err
	}
	if cache {
		err = os.WriteFile(filename, bodyBytes, 0644)
	}
	return bodyBytes, nil
}

// func Writefile(filename string, data []byte) {
// 	err := os.WriteFile(filename, data, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// sync deliveries for one customer
func (lime *Lime) GetDeliveries(company *models.LimeCompany, refresh bool) error {

	var deliveriesURL string
	for key, link := range company.Links {
		if key == "relation_deliveries" {
			deliveriesURL = link.Href
		}
	}

	deliveriesURL = deliveriesURL + "?_embed=product&_embed=service&_embed=agreement&_embed=deliverypoint&_embed=deliverypoint2"

	for deliveriesURL != "" {
		bodyBytes, err := lime.limeAPI(deliveriesURL, refresh)
		if err != nil {
			return err
		}

		var deliveriesData models.DeliveryResponse
		err = json.Unmarshal(bodyBytes, &deliveriesData)
		if err != nil {
			return err
		}
		// delivery := deliveriesData.Embedded.Delivery[0]

		// unpack _embedded data
		for _, d := range deliveriesData.Embedded.Delivery {
			slog.Debug("ConnectionNumber", "data", d.ConnectionNumber)
			deliveriesCount++
			for relation, raw := range d.Embedded {
				switch relation {
				case "relation_agreement":
					{
						var agreement models.LimeAgreement
						err = json.Unmarshal(raw, &agreement)
						if err != nil {
							slog.Error("Error unmarshalling JSON", "response", err)
							return err
						}
						d.Agreeement = &agreement
						slog.Debug("agrement", "data", agreement)
					}
				case "relation_deliverypoint":
					{
						var deliveryendpoint models.LimeDeliverypoint
						slog.Debug("deliveryendpoint", "raw", raw)
						err = json.Unmarshal(raw, &deliveryendpoint)
						if err != nil {
							return err
						}
						d.DeliveryPoint = &deliveryendpoint
						slog.Debug("deliveryendpoint", "data", deliveryendpoint)
					}
				case "relation_deliverypoint2":
					{
						var deliveryendpoint2 models.LimeDeliverypoint2
						err = json.Unmarshal(raw, &deliveryendpoint2)
						if err != nil {
							slog.Error("Error unmarshalling JSON", "response", err)
							return err
						}
						d.DeliveryPoint2 = &deliveryendpoint2
						slog.Debug("deliveryendpoint2", "data", deliveryendpoint2)
					}
				case "relation_product":
					{
						var product models.LimeProduct
						err = json.Unmarshal(raw, &product)
						if err != nil {
							slog.Error("Error unmarshalling JSON", "response", err)
							return err
						}
						d.Product = &product
						slog.Debug("product", "data", product)
					}
				case "relation_service":
					{
						var service models.LimeService
						err = json.Unmarshal(raw, &service)
						if err != nil {
							slog.Error("Error unmarshalling JSON", "response", err)
							return err
						}
						d.Service = &service
						slog.Debug("service", "data", service)
					}
				default:
					slog.Warn("Unknown", "relation", relation)
				}
			}
			company.Deliveries = append(company.Deliveries, d)
		}

		link, ok := deliveriesData.Links["next"]
		if ok {
			deliveriesURL = link.Href
		} else {
			deliveriesURL = ""
		}
	}
	return nil
}

// GetPersons loads the persons related to a company via relation_person.
// An empty list is success: some companies have no people in Lime.
func (lime *Lime) GetPersons(company *models.LimeCompany, refresh bool) error {
	var personsURL string
	for key, link := range company.Links {
		if key == "relation_person" {
			personsURL = link.Href
		}
	}
	if personsURL == "" {
		return nil
	}

	for personsURL != "" {
		bodyBytes, err := lime.limeAPI(personsURL, refresh)
		if err != nil {
			return err
		}

		var personsData models.PersonResponse
		err = json.Unmarshal(bodyBytes, &personsData)
		if err != nil {
			return err
		}

		for _, p := range personsData.Embedded.Persons {
			slog.Debug("Person", "name", p.Name, "company", company.Name)
			company.Persons = append(company.Persons, p)
		}

		link, ok := personsData.Links["next"]
		if ok {
			personsURL = link.Href
		} else {
			personsURL = ""
		}
	}
	return nil
}

// Fetch data from LIME
// note: Limit in api, companyname is "starts-with"
// response is list of companies, with their corresponding data
func (lime *Lime) GetCompanies(companyName string, refresh bool) ([]*models.LimeCompany, error) {
	var companies []*models.LimeCompany

	URL := lime.APIURL + "/company"
	if companyName != "" {
		URL = URL + "/?name=" + url.QueryEscape(companyName)
	}
	for URL != "" {
		bodyBytes, err := lime.limeAPI(URL, refresh)
		if err != nil {
			return nil, err
		}
		var apiData models.CompanyResponse
		err = json.Unmarshal(bodyBytes, &apiData)
		if err != nil {
			return nil, err
		}

		if len(apiData.Embedded.Customers) == 0 {
			return nil, errors.New("No customers returned")
		}

		if companyName != "" && len(apiData.Embedded.Customers) > 1 {
			return nil, errors.New("Search for customer returned more than one customer")
		}

		for _, company := range apiData.Embedded.Customers {
			companiesCount++
			slog.Debug("Company", "name", company.Name)
			err := lime.GetDeliveries(&company, refresh)
			if err != nil {
				return nil, err
			}
			companies = append(companies, &company)
		}

		// go to next if there are any more data
		link, ok := apiData.Links["next"]
		if ok {
			URL = link.Href
		} else {
			URL = ""
		}
	}
	slog.Debug("Done", "companiesCount", companiesCount, "deliveriescount", deliveriesCount)
	return companies, nil
}
