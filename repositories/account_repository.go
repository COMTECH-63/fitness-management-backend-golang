package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/COMTECH-63/fitness-management/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return accountRepository{db: db}
}

func (r accountRepository) GetAccountPaginate(ctx context.Context, pagination database.Pagination, search string) (*database.Pagination, error) {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "GetAccountPaginateRepository", trace.WithAttributes(attribute.String("repository", "GetAccountPaginate"), attribute.String("search", search)))
		accounts     []models.Account
		err          error
	)

	// Pagination query
	if search != "" {
		if err = r.db.Scopes(database.Paginate(accounts, &pagination, r.db)).
			Where(`username LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Or(`password LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Find(&accounts).Error; err != nil {
			log.Println(err)
			return nil, errors.New("GetAccountPaginateError")
		}
	} else {
		if err = r.db.Scopes(database.Paginate(accounts, &pagination, r.db)).
			Preload("Users").Find(&accounts).Error; err != nil {
			return nil, err
		}
	}

	// Set data
	pagination.Data = accounts

	childSpan.End()

	return &pagination, nil
}

func (r accountRepository) GetAccountByID(ctx context.Context, id int) (models.Account, error) {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "GetAccountByIDRepository", trace.WithAttributes(attribute.String("repository", "GetAccountByID")))
		account      models.Account
		err          error
	)

	// Query
	if err = r.db.Preload("Users").Find(&account, id).Error; err != nil {
		return account, err
	}

	childSpan.End()

	return account, nil
}

func (r accountRepository) CreateAccount(ctx context.Context, account *models.Account) error {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "CreateAccountRepository", trace.WithAttributes(attribute.String("repository", "CreateAccount")))
		err          error
	)

	// Execute
	if err = r.db.Create(&account).Error; err != nil {
		return err
	}

	childSpan.End()

	return nil
}

func (r accountRepository) UpdateAccount(ctx context.Context, id int, account *models.Account) error {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "UpdateAccountRepository", trace.WithAttributes(attribute.String("repository", "UpdateAccount")))
		existAccount *models.Account
		err          error
	)

	// Get model
	r.db.First(&existAccount)

	// Set attributes
	existAccount.Username = account.Username
	existAccount.Password = account.Password

	// Execute
	if err = r.db.Save(&existAccount).Error; err != nil {
		return err
	}

	childSpan.End()

	return nil
}

func (r accountRepository) DeleteAccount(ctx context.Context, id int) error {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "DeleteAccountRepository", trace.WithAttributes(attribute.String("repository", "DeleteAccount")))
		err          error
	)

	// Execute
	if err = r.db.Delete(&models.Account{}, id).Error; err != nil {
		return err
	}

	childSpan.End()

	return nil
}
