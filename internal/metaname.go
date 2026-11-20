package internal

/**
enum of api methods
structs for each methods (with json tags)
wrap function that creates payloads
*/
import (
	"net"
)

type Contact struct {
	name               string
	email              string
	phone_country_code string
	phone_area_code    *string
	phone_local_number string
	organisation       *string
	address_line1      string
	address_line2      *string
	city               string
	region             *string
	postal_code        string
	country_code       string
}

type Contacts struct {
	registrant *Contact
	admin      *Contact
	technical  *Contact
}

type NameServer struct {
	name        string
	ip4_address net.TCPAddr
	ip6_address net.TCPAddr
}

type Record struct {
	name       string
	recordType string
	aux        *int // must be set only for record types of MX & SRV
	ttl        int
	data       string
}

type JsonRpcPayload struct {
	jsonrpc string
	Method  string
}

func validMethods() []string {
	return []string{
		"price",
		"check_availability",
		"register_domain_name",
		"import_nz_domain_name",
		"import_other_domain_name",
		"domain_names",
		"domain_name",
		"update_domain_name",
		"update_contacts",
		"update_name_servers",
		"dns_zones",
		"dns_zone",
		"configure_zone",
		"create_dns_record",
		"update_dns_record",
		"delete_dns_record",
		"update_ds_records",
		"renew_domain_name",
		"domain_name_is_locked",
		"lock_domain_name",
		"unlock_domain_name",
		"reset_domain_name_secret",
		"account_balance",
		"transactions",
	}
}
