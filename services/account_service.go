package services

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/COMTECH-63/fitness-management/repositories"
	"github.com/getsentry/sentry-go"
)

type (
	accountService struct {
		accountRepository repositories.AccountRepository
	}
)

func NewAccountService(
	accountRepo repositories.AccountRepository,
) AccountService {
	return &accountService{
		accountRepository: accountRepo,
	}
}

func (s accountService) GetAccounts(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error) {
	childSpan := span.StartChild("GetAccountsService")
	result, err := s.accountRepository.GetAccountPaginate(ctx, childSpan, paginate, search)
	childSpan.Finish()
	return result, err
}

func (s accountService) GetAccount(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error) {
	childSpan := span.StartChild("GetAccountService")
	account, err := s.accountRepository.GetAccountByID(ctx, childSpan, id)
	childSpan.Finish()

	return map[string]interface{}{"data": account}, err
}

func (s accountService) CreateAccount(ctx context.Context, span *sentry.Span, accountDto *AccountDto) error {
	childSpan := span.StartChild("CreateAccountService")
	account := new(models.Account)

	account.Username = accountDto.Username
	account.Password = accountDto.Password

	err := s.accountRepository.CreateAccount(ctx, childSpan, account)
	childSpan.Finish()

	return err
}

func (s accountService) UpdateAccount(ctx context.Context, span *sentry.Span, id int, accountDto *AccountDto) error {
	childSpan := span.StartChild("UpdateAccountService")
	account := new(models.Account)

	account.Username = accountDto.Username
	account.Password = accountDto.Password

	err := s.accountRepository.UpdateAccount(ctx, childSpan, id, account)
	childSpan.Finish()

	return err
}

func (s accountService) DeleteAccount(ctx context.Context, span *sentry.Span, id int) error {
	childSpan := span.StartChild("DeleteAccountService")
	err := s.accountRepository.DeleteAccount(ctx, childSpan, id)
	childSpan.Finish()

	return err
}
