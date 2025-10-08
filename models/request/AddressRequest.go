package request 

type AddressRequest struct {
	FullName      string `json:"fullname"`
	Phone         string `json:"phone"`
	Province      string `json:"province"`
	District      string `json:"district"`
	Village       string `json:"village"`
	DetailAddress string `json:"detailaddress"`
	TypeAddress   bool   `json:"typeaddress"`
}