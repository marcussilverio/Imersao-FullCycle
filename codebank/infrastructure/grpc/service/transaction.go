package service

import (
	"context"
	"github.com/marcussilverio/codebank/infrastructure/grpc/pb"
	"github.com/marcussilverio/codebank/usecase"
)

type TransactionService struct{
	ProcesTransactionUseCase usecase.UseCaseTransaction
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, in *pb.PaymentRequest) (*empty.Empty, error) {
	transactionDto:= dto.Transaction{
		Name: in.GetCreditCard().GetName(),
		Number: in.GetCreditCard().GetNumber(),
		ExpirationMonth: in.GetCreditCard().getExpirationMonth(),
		ExpirationYear: in.GetCreditCard().getExpirationYear(),
		Amount: in.GetAmount(),
		CVV: in.GetCreditCard().GetCvv(),
		Store: in.GetStore(),
		Description: in.GetDescription(),
	}
	transaction, err := t.ProcesTransactionUseCase.ProcesTransaction(transactionDto)
	if err != nil {
		return &empty.Empty{}, status.Error(code.FailedPrecondition, err.Error())
	}
	if transaction.Status != "approved" {
		return &empty.Empty{}, status.Error(code.FailedPrecondition, msg:"Transaction rejected by the bank.")
	} 
	return &empty.Empty{}, nil
}

