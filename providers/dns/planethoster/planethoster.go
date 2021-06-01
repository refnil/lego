import (
    "fmt"
    "net/http"

	"github.com/go-acme/lego/v4/platform/config/env"
)

// Environment variables names.
const (
    BASE_URL = "https://api.planethoster.net/reseller-api/"
	envNamespace = "PLANETHOSTER_"

    EnvKey = envNamespace + "KEY"
	EnvUser = envNamespace + "USER"
)

type DNSProvider struct {
	apikey string
    apiuser string
}

type Record struct {
    hostname string
    address string
    recordtype string
}

func NewDNSProvider() (*DNSProviderBestDNS, error) {
	apikey := env.GetOrFile(EnvKey)
    apiuser := env.GetOrFile(EnvUser)
    return &DNSProvider{apikey:apikey, apiuser:apiuser}, nil
}

// Present creates a TXT record to fulfill the dns-01 challenge.
func (d *DNSProvider) Present(domain, token, keyAuth string) error {
	fqdn, value := dns01.GetRecord(domain, keyAuth)

    records, err := d.GetRecord()
	if err != nil {
		return fmt.Errorf("planethoster: %w", err)
	}

    challenge_record := Record {hostname=domain, recordtype="txt", address=value}

    records = append(records, challenge_record)

    err := d.SetRecord(records)
	if err != nil {
		return fmt.Errorf("planethoster: %w", err)
	}

	return nil
}

// CleanUp removes the TXT record matching the specified parameters.
func (d *DNSProvider) CleanUp(domain, token, keyAuth string) error {
	fqdn, _ := dns01.GetRecord(domain, keyAuth)

    current_records, err := d.GetRecord()
	if err != nil {
		return fmt.Errorf("planethoster: %w", err)
	}

    target_records = []Record{}

    for i := range current_records{
        cur_rec := current_records[i]
        if !(cur_rec.hostname == domain && cur_rec.recordtype == "txt") {
            target_records = append(target_records, cur_rec)
        }
    }


    err := d.SetRecord(records)
	if err != nil {
		return fmt.Errorf("planethoster: %w", err)
	}

	return nil
}

func (d *DNSProvider) GetRecord(domain) ([]Record, error) {
    client := &http.Client{}

    req, err := http.NewRequest("GET", BASE_URL+"/get-ph-dns-records", nil)
    req.Header.Add("X-API-KEY", d.apikey)
    req.Header.Add("X-API-USER", d.apiuser)
    resp, err := client.Do(req)
}

func (d *DNSProvider) SetRecord(domain, records []Record) error {
}
