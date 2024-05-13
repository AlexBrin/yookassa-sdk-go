package yoopayment

type PaymentMode string

const (
	PaymentModeFullPrepayment    PaymentMode = "full_prepayment"
	PaymentModePartialPrepayment PaymentMode = "partial_prepayment"
	PaymentModeAdvance           PaymentMode = "advance"
	PaymentModeFullPayment       PaymentMode = "full_payment"
	PaymentModePartialPayment    PaymentMode = "partial_payment"
	PaymentModeCredit            PaymentMode = "credit"
	PaymentModeCreditPayment     PaymentMode = "credit_payment"
)
