package models

type Transaction struct {
}

func (w *Transaction) ToResponseTransaction() *ResponseTransaction {
	return &ResponseTransaction{}
}

func (w *Transaction) ApplyUpdateRequest(req *UpdateTransactionRequest) {

}

type ResponseTransaction struct {
}

type CreateTransactionRequest struct {
}

func ToTransaction(req *CreateTransactionRequest) *Transaction {
	return &Transaction{}
}

type UpdateTransactionRequest struct {
}

func ToResponseTransactions(transactions []*Transaction) []*ResponseTransaction {
	responseTransactions := make([]*ResponseTransaction, len(transactions))
	for i, transaction := range transactions {
		responseTransactions[i] = transaction.ToResponseTransaction()
	}
	return responseTransactions
}
