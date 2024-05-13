package yoopayment

type MarkQuantity struct {
	// The numerator is the number of goods sold from one consumer package (tag b 54 FZ — 1293).
	// Cannot exceed denominator.
	Numerator int `json:"numerator"`

	// The denominator is the total number of goods in consumer packaging (tag in 54 FZ — 1294).
	Denominator int `json:"denominator"`
}
