package nutshell

import(
	"log"
	"encoding/json"
	"net/http"
	"bytes"
	//"io"
	"io/ioutil"
	//"bufio"
)


type NutshellConfig struct{
	User string `jason:"user"`
	ApiKey string  `json:"apiKey"`
}

var config NutshellConfig

func Init(nutConfig NutshellConfig){
	config = nutConfig
}  

type APIS struct{
	Id int 			`json:"id"`
	Method string 		`json:"method"`
	Params interface{} 	`json:"params"`
}

type ContactCallS struct {
	Contact ContactS `json:"contact,omitempty"`
}

type ContactS struct {
	Name        string       `json:"name,omitempty"`
	//Owner       OwnerS       `json:"owner,omitempty"`
	Description string       `json:"description,omitempty"`
	Phone       []string     `json:"phone,omitempty"`
	Email       []string     `json:"email,omitempty"`
	Address     []AddressS   `json:"address,omitempty"`
	Lead        []LeadS      `json:"lead,omitempty"`
	//CustomField CustomFieldS `json:"customField,omitempty"`
	TerritoryId int          `json:"territoryId,omitempty"`
}

type OwnerS struct {
	EntityType string `json:"enityType,omitempty"`
	Id         int    `json:"id,omitempty"`
}

type AddressS struct {
	Address_1  string `json:"address_1,omitempty"`
	Address_2  string `json:"address_2,omitempty"`
	Address_3  string `json:"address_3,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode int    `json:"postalCode,omitempty"`
	Country    string `json:"country,omitempty"`
}

type LeadS struct {
	Relationship string `json:"relationship,omitempty"`
	Id           int    `json:"id,omitempty"`
}

type CustomFieldS struct {
	Id int `json:"id,omitempty"`
}


//APICall takes in APIS and marshals it into a JSON sent by HTTP to nutshell
func ApiCall(call APIS) {

	//Marshal APIS into a JSON for HTTP Body
	var buf bytes.Buffer
	d := json.NewEncoder(&buf)
	err := d.Encode(call)
	if(err != nil){
		log.Printf("ERROR with json marshaling: %s\n", err)
	}

	req, err := http.NewRequest("POST", "https://app01.nutshell.com/api/v1/json", bytes.NewReader(buf.Bytes()))
	if(err != nil){
		log.Printf("Error creating HTTP request: %s\n", err)
		return
	}

	req.SetBasicAuth(config.User, config.ApiKey)
		
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("HTTP Responce Error: %s\n", err)
		return
	}
	
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERROR with reading HTTP response body: %s\n", err)
	}
	log.Printf("Responce:\n%s\n", responseBody)
}
