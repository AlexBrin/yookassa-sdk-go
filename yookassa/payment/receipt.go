package yoopayment

import yoocommon "github.com/AlexBrin/yookassa-sdk-go/yookassa/common"

type Receipt struct {
	// User information
	// It is necessary to specify at least contact information:
	// for Receipts from YooKassa — e—mail (customer.email),
	// in other cases - e-mail (customer.email) or phone number (customer.phone).
	Customer *Customer `json:"customer,omitempty"`

	Items []ReceiptItem `json:"items"`
}

type ReceiptItem struct {
	Description string           `json:"description"`
	Amount      yoocommon.Amount `json:"amount"`
	VatCode     int              `json:"vat_code"`
	Quantity    int              `json:"quantity"`
	Measure     string           `json:"measure,omitempty"`

	// The fractional amount of the marked product (tag in 54 FZ — 1291).
	// A mandatory parameter if these conditions are met at the same time:
	//  * you use Receipts from Kassa or an online sales register updated to FFD 1.2;
	//  * the product must be marked;
	//  * the measure field has the value piece.
	//
	// Example:
	// 	you sell pencils by the piece.
	//	They are supplied in packs of 100 pieces with one marking code.
	//	When selling one pencil, you need to transfer 1 to numerator, and 100 to denominator.
	MarkQuantity *MarkQuantity `json:"mark_quantity,omitempty"`

	// The attribute of the subject of calculation (tag in 54 FZ — 1212)
	// is what payment is accepted for, for example, a product or service.
	//
	// List of possible values:
	// https://yookassa.ru/developers/payment-acceptance/receipts/54fz/yoomoney/parameters-values#payment-subject
	PaymentSubject *PaymentSubject `json:"payment_subject,omitempty"`

	// The sign of the payment method (tag in 54 FZ — 1214) reflects the type of payment and the fact of transfer of the goods. Example: the buyer pays for the product in full and receives it immediately. In this case, you need to pass the full_payment value (full calculation).
	//
	// List of possible values:
	//  for Receipts from YooKassa: https://yookassa.ru/developers/payment-acceptance/receipts/54fz/yoomoney/parameters-values#payment-mode
	//  for third-party online sales registers: https://yookassa.ru/developers/payment-acceptance/receipts/54fz/other-services/parameters-values#payment-mode
	PaymentMode *PaymentMode `json:"payment_mode,omitempty"`

	// The code of the country of origin of the goods according to the
	// all-Russian classifier of the countries of the world
	// (OK (MK (ISO 3166) 004-97) 025-2001). The tag in 54 FZ — 1230.
	//
	// Example: RU.
	// You can transfer it if you use the online sales register Orange Data, Kit Invest.
	//CountryOfOriginCode string `json:"country_of_origin_code,omitempty"`

	// The number of the customs declaration (from 1 to 32 characters). The tag in 54 FZ — 1231.
	// Can be transferred if you use the online sales register Orange Data, Kit Invest
	//CustomsDeclarationNumber string `json:"customs_declaration_number,omitempty"`

	// The amount of the excise tax of the goods, including kopecks (tag in 54 FZ — 1229).
	// A decimal number with an accuracy of 2 digits after the dot.
	// You can transfer it if you use the online sales register Orange Data, Kit Invest.
	//Excise string `json:"excise,omitempty"`

	// The product code (tag in 54 FZ — 1162) is a unique number that is assigned to
	// a copy of the product during labeling.
	// Format: a number in hexadecimal representation with spaces.
	// The maximum length is 32 bytes.
	//
	// Example: 00 00 00 01 00 21 FA 41 00 23 05 41 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 12 00 AB 00.
	// A mandatory parameter if these conditions are met at the same time:
	//  * you are using an online sales register updated to FVD 1.05;
	//  * the product must be labeled.
	//ProductCode string `json:"product_code,omitempty"`
}
