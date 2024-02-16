package services

import (
	"context"
	"errors"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/COMTECH-63/fitness-management/repositories"
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		// Handle error (e.g., log, return an error, etc.)
		childSpan.Finish()
		return err
	}

	// Set the hashed password in the account model
	account.Password = string(hashedPassword)

	err = s.accountRepository.CreateAccount(ctx, childSpan, account)
	childSpan.Finish()

	return err
}

func (s accountService) UpdateAccount(ctx context.Context, span *sentry.Span, id int, accountDto *AccountDto) error {
	childSpan := span.StartChild("UpdateAccountService")

	// Retrieve the existing account from the repository
	existingAccount, err := s.accountRepository.GetAccountByID(ctx, childSpan, id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// Handle the case where the account does not exist
			childSpan.Finish()
			return errors.New("account not found")
		}
		childSpan.Finish()
		return err
	}

	// Update fields from the DTO
	existingAccount.Username = accountDto.Username

	// Check if a new password is provided
	if accountDto.Password != "" {
		// Hash the new password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accountDto.Password), bcrypt.DefaultCost)
		if err != nil {
			childSpan.Finish()
			return err
		}
		existingAccount.Password = string(hashedPassword)
	}

	// Update the account in the repository
	err = s.accountRepository.UpdateAccount(ctx, childSpan, id, &existingAccount)
	childSpan.Finish()

	return err
}

func (s accountService) DeleteAccount(ctx context.Context, span *sentry.Span, id int) error {
	childSpan := span.StartChild("DeleteAccountService")
	err := s.accountRepository.DeleteAccount(ctx, childSpan, id)
	childSpan.Finish()

	return err
}
