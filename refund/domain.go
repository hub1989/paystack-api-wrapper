package refund

type Response struct {
	Transaction struct {
		Id            int64  `json:"id,omitempty"`
		Domain        string `json:"domain,omitempty"`
		Reference     string `json:"reference,omitempty"`
		Amount        int    `json:"amount,omitempty"`
		PaidAt        string `json:"paid_at,omitempty"`
		Channel       string `json:"channel,omitempty"`
		Currency      string `json:"currency,omitempty"`
		Authorization struct {
			ExpMonth    interface{} `json:"exp_month,omitempty"`
			ExpYear     interface{} `json:"exp_year,omitempty"`
			AccountName interface{} `json:"account_name,omitempty"`
		} `json:"authorization,omitempty"`
		Customer struct {
			InternationalFormatPhone interface{} `json:"international_format_phone,omitempty"`
		} `json:"customer,omitempty"`
		Plan struct {
		} `json:"plan,omitempty"`
		Subaccount struct {
			Currency interface{} `json:"currency,omitempty"`
		} `json:"subaccount,omitempty"`
		Split struct {
		} `json:"split,omitempty"`
		OrderId            interface{} `json:"order_id,omitempty"`
		PaidAt1            string      `json:"paidAt,omitempty"`
		PosTransactionData interface{} `json:"pos_transaction_data,omitempty"`
		Source             interface{} `json:"source,omitempty"`
		FeesBreakdown      interface{} `json:"fees_breakdown,omitempty"`
	} `json:"transaction,omitempty"`
	Integration    int         `json:"integration,omitempty"`
	DeductedAmount int         `json:"deducted_amount,omitempty"`
	Channel        interface{} `json:"channel,omitempty"`
	MerchantNote   string      `json:"merchant_note,omitempty"`
	CustomerNote   string      `json:"customer_note,omitempty"`
	Status         string      `json:"status,omitempty"`
	RefundedBy     string      `json:"refunded_by,omitempty"`
	ExpectedAt     string      `json:"expected_at,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	Domain         string      `json:"domain,omitempty"`
	Amount         int         `json:"amount,omitempty"`
	FullyDeducted  bool        `json:"fully_deducted,omitempty"`
	Id             int         `json:"id,omitempty"`
	CreatedAt      string      `json:"createdAt,omitempty"`
	UpdatedAt      string      `json:"updatedAt,omitempty"`
}
