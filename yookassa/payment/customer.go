package yoopayment

type Customer struct {
	// For a legal entity — the name of the organization,
	// for sole proprietors and individuals — full name.
	// If an individual does not have an INN, passport data is transmitted
	// in the same parameter. Maximum of 256 characters.
	// You can transfer it if you use Receipts from YooKassa
	// or the online sales register Orange Data, Atol Online.
	FullName string `json:"full_name,omitempty"`

	// The user's INN (10 or 12 digits)
	// If an individual does not have an INN, it is necessary to pass
	// the passport data in the full_name parameter.
	// You can transfer it if you use Receipts from YUKASA or the online sales register
	// Orange Data, Atol Online
	INN string `json:"inn,omitempty"`
	// The user's email address for sending the receipt.
	// A mandatory parameter if you use Receipts from KasaKasa
	// or if you use another solution
	// (third-party online sales register, self-employed receipts)
	// and do not transfer the phone
	Email string `json:"email,omitempty"`

	// The user's phone number for sending the receipt. It is specified in the ITU-T E.164 format,
	// for example 79000000000.
	// Required parameter if no email is passed
	Phone string `json:"phone,omitempty"`
}
