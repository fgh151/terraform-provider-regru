package provider

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type RequestParams struct {
	Username          string
	Password          string
	OutputContentType string
	Domains           []RequestDomain
}

type RequestDomain struct {
	Dname string
}

type RegruProvider struct {
	Username string
	Password string
}

type DnsRecord struct {
	Domain         string
	Host           string `json:"host"`
	Type           string `json:"type"`
	Value          string `json:"value"`
	Ttl            int    `json:"ttl"`
	Subdomain      string `json:"subdomain"`
	ExternalId     string `json:"external_id"`
	AdditionalInfo string `json:"additional_info"`
}

type domainsRequestParam struct {
	Dname string `json:"dname"`
}

func (r RegruProvider) AddRecord(record DnsRecord) (error, []byte) {
	var endpoint = ""

	params := map[string]interface{}{
		"username":  r.Username,
		"password":  r.Password,
		"subdomain": record.Subdomain,
		"text":      record.Value,
		"domains":   []domainsRequestParam{{Dname: record.Domain}},
	}

	switch record.Type {
	case "A":
		endpoint = "https://api.reg.ru/api/regru2/zone/add_alias"
		params["ipaddr"] = record.Value
		break
	case "AAAA":
		endpoint = "https://api.reg.ru/api/regru2/zone/add_aaaa"
		params["ipaddr"] = record.Value
		break
	case "CNAME":
		endpoint = "https://api.reg.ru/api/regru2/zone/add_cname"
		params["canonical_name"] = record.Value
		break
	case "MX":
		endpoint = "https://api.reg.ru/api/regru2/zone/add_mx"
		params["mail_server"] = record.Value
		params["priority"] = record.Ttl
		break
	case "NS":
		endpoint = "https://api.reg.ru/api/regru2/zone/add_ns"
		params["dns_server"] = record.Value
		params["priority"] = record.Ttl
		break
	case "TXT":
		endpoint = "https://api.reg.ru/api/regru2/zone/add_txt"
		params["text"] = record.Value
		break
	default:
		endpoint = "https://api.reg.ru/api/regru2/zone/add_txt"
		params["text"] = record.Value
		break
	}

	req, err := http.NewRequest("GET", endpoint, nil)

	q := req.URL.Query()

	b, err := json.Marshal(params)

	q.Add("input_data", string(b))
	q.Add("input_format", "json")

	req.URL.RawQuery = q.Encode()

	c := &http.Client{}
	resp, err := c.Do(req)

	body, err := ioutil.ReadAll(resp.Body)

	return err, body
}

type RegruRecord struct {
	Content string
	Prio    int
	Rectype string
	State   string
	Subname string
}

type RegruDomain struct {
	Dname string
	Reult string
	Rrs   []RegruRecord
}

type RegruAnswer struct {
	Domains []RegruDomain
}

type RegruResponse struct {
	Answer RegruAnswer
	Result string
}

//https://www.reg.ru/support/help/api2#zone_get_resource_records

func (r RegruProvider) GetRecords(domain string) []DnsRecord {

	req, err := http.NewRequest("GET", "https://api.reg.ru/api/regru2/zone/get_resource_records", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()

	q.Add("input_data", string(r.crateParams(domain)))
	q.Add("input_format", "json")

	req.URL.RawQuery = q.Encode()

	c := &http.Client{}

	resp, _ := c.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var rResp RegruResponse
	json.Unmarshal(body, &rResp)

	var returnAr []DnsRecord

	for _, d := range rResp.Answer.Domains {
		if d.Dname == domain {
			for _, rrs := range d.Rrs {
				returnAr = append(returnAr, DnsRecord{
					Host:  rrs.Subname,
					Type:  rrs.Rectype,
					Value: rrs.Content,
					Ttl:   10,
				})
			}
		}
	}

	return returnAr
}

func (r RegruProvider) DeleteRecord(record DnsRecord) {
	req, err := http.NewRequest("GET", "https://api.reg.ru/api/regru2/zone/remove_record", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()

	params := struct {
		Username          string
		Password          string
		Subdomain         string
		Content           string
		RecordType        string
		OutputContentType string
	}{
		Username:          r.Username,
		Password:          r.Password,
		Subdomain:         record.Host,
		Content:           record.Value,
		RecordType:        record.Type,
		OutputContentType: "plain",
	}

	b, _ := json.Marshal(params)

	q.Add("input_data", string(b))
	q.Add("input_format", "json")

	req.URL.RawQuery = q.Encode()

	c := &http.Client{}
	c.Do(req)
}

func (r RegruProvider) crateParams(domain string) []byte {

	var d []RequestDomain

	p := r.getRequestParams()

	d = append(p.Domains, RequestDomain{Dname: domain})

	p.Domains = d

	b, _ := json.Marshal(p)

	return b
}

func (r RegruProvider) getRequestParams() RequestParams {

	user := r.Username
	pass := r.Password

	if strings.HasPrefix(pass, "ENV_") {
		pass = os.Getenv(pass)
	}

	if strings.HasPrefix(user, "ENV_") {
		user = os.Getenv(user)
	}

	return RequestParams{
		Username:          user,
		Password:          pass,
		OutputContentType: "json",
	}
}
