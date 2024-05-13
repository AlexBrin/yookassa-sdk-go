package yoopayment

// PaymentSubject
// The attribute of the subject of calculation (tag in 54 FZ — 1212) is what payment is accepted for, for example, a product or service.
//
// List of possible values:
//
//	for Receipts from YooKassa: https://yookassa.ru/developers/payment-acceptance/receipts/54fz/yoomoney/parameters-values#payment-subject
//	for third-party online sales registers: https://yookassa.ru/developers/payment-acceptance/receipts/54fz/other-services/parameters-values#payment-subject
type PaymentSubject string

const (
	SubjectCommodity                      PaymentSubject = "commodity"
	SubjectJob                            PaymentSubject = "job"
	SubjectService                        PaymentSubject = "service"
	SubjectPayment                        PaymentSubject = "payment"
	SubjectCasino                         PaymentSubject = "casino"
	SubjectGamblingBet                    PaymentSubject = "gambling_bet"
	SubjectGamblingPrize                  PaymentSubject = "gambling_prize"
	SubjectLottery                        PaymentSubject = "lottery"
	SubjectLotteryPrize                   PaymentSubject = "lottery_prize"
	SubjectIntellectualActivity           PaymentSubject = "intellectual_activity"
	SubjectAgentCommission                PaymentSubject = "agent_commission"
	SubjectPropertyRight                  PaymentSubject = "property_right"
	SubjectNonOperatingGain               PaymentSubject = "non_operating_gain"
	SubjectInsurancePremium               PaymentSubject = "insurance_premium"
	SubjectResortFee                      PaymentSubject = "resort_fee"
	SubjectMarked                         PaymentSubject = "market"
	SubjectNonMarked                      PaymentSubject = "non_marked"
	SubjectFine                           PaymentSubject = "fine"
	SubjectTax                            PaymentSubject = "tax"
	SubjectLien                           PaymentSubject = "lien"
	SubjectCost                           PaymentSubject = "cost"
	SubjectAgentWithdrawals               PaymentSubject = "agent_withdrawals"
	SubjectPensionInsuranceWithoutPayouts PaymentSubject = "pension_insurance_without_payouts"
	SubjectPensionInsuranceWithPayouts    PaymentSubject = "pension_insurance_with_payouts"
	SubjectHealthInsuranceWithoutPayouts  PaymentSubject = "health_insurance_without_payouts"
	SubjectHealthInsuranceWithPayouts     PaymentSubject = "health_insurance_with_payouts"
	SubjectAnother                        PaymentSubject = "another"
)
