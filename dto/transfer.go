package dto

type TransferRequest struct {
	CpfPayee string `json:"cpf_payee"`
	Value    int    `json:"value"`
}
