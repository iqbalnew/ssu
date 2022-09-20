package pb

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strings "strings"
	time "time"
)

type RemittanceTransactionORM struct {
	ACCOUNT_NUMBER                string `gorm:"type:varchar(100)"`
	AMOUNT_ORI                    string `gorm:"type:varchar(100)"`
	BENEFICIARY_ADDRESS           string `gorm:"type:varchar(100)"`
	BENEFICIARY_BANK_ADDRESS      string `gorm:"type:varchar(100)"`
	BENEFICIARY_BANK_CITY         string `gorm:"type:varchar(100)"`
	BENEFICIARY_BANK_CODE         string `gorm:"type:varchar(100)"`
	BENEFICIARY_BANK_COUNTRY_CODE string `gorm:"type:varchar(100)"`
	BENEFICIARY_BANK_NAME         string `gorm:"type:varchar(100)"`
	BENEFICIARY_CITY              string `gorm:"type:varchar(100)"`
	BENEFICIARY_COUNTRY_CODE      string `gorm:"type:varchar(100)"`
	BENEFICIARY_EMAIL             string `gorm:"type:varchar(100)"`
	BENEFICIARY_GENDER            string `gorm:"type:varchar(100)"`
	BENEFICIARY_ID_NUMBER         string `gorm:"type:varchar(100)"`
	BENEFICIARY_ID_TYPE           string `gorm:"type:varchar(100)"`
	BENEFICIARY_NAME              string `gorm:"type:varchar(100)"`
	BENEFICIARY_PHONE             string `gorm:"type:varchar(100)"`
	BENEFICIARY_PHONE_CODE        string `gorm:"type:varchar(100)"`
	BENEFICIARY_POSTCODE          string `gorm:"type:varchar(100)"`
	BENEFICIARY_RELATIONSHIP      string `gorm:"type:varchar(100)"`
	BENEFICIARY_STATE_PROVINCE    string `gorm:"type:varchar(100)"`
	BENEFICIARY_TYPE              string `gorm:"type:varchar(100)"`
	BOOK_RATE_BUY                 string `gorm:"type:varchar(100)"`
	BOOK_RATE_SELL                string `gorm:"type:varchar(100)"`
	BeneficiaryCountryName        string `gorm:"type:varchar(100)"`
	BeneficiaryEmails             string `gorm:"type:jsonb"`
	CHARGES_TYPE                  string `gorm:"type:varchar(100)"`
	COUNTERPART                   string `gorm:"type:varchar(100)"`
	CURRENCY_ORI                  string `gorm:"type:varchar(100)"`
	CheckStatusResponse           string `gorm:"type:jsonb"`
	CompanyID                     string `gorm:"type:varchar(100);not null"`
	CompletedDate                 string `gorm:"type:varchar(50);not null"`
	CreatedAt                     *time.Time
	CreatedByName                 string `gorm:"type:varchar(100)"`
	DEAL_CODE                     string `gorm:"type:varchar(100)"`
	DEBIT_ACCOUNT                 string `gorm:"type:varchar(100)"`
	DEBIT_ACCOUNT_MAIN            string `gorm:"type:varchar(100)"`
	DEBIT_AMOUNT                  string `gorm:"type:varchar(100)"`
	DataID                        uint64 `gorm:"primary_key;not null"`
	DebitAccountAlias             string `gorm:"type:varchar(100);not null"`
	DebitAccountID                uint64 `gorm:"type:bigint;not null"`
	DebitCurrencyID               string `gorm:"type:varchar(100);not null"`
	FEE_ACCOUNT_CHANNEL           string `gorm:"type:varchar(100)"`
	FEE_CREDIT_AMOUNT_CHANNEL     string `gorm:"type:varchar(100)"`
	FEE_CURRENCY_CHANNEL          string `gorm:"type:varchar(100)"`
	FEE_DEBIT_AMOUNT_CHANNEL      string `gorm:"type:varchar(100)"`
	FileAttachment                string `gorm:"type:varchar(100)"`
	INTENDED_USE                  string `gorm:"type:varchar(100)"`
	LastChecked                   string `gorm:"type:varchar(50)"`
	MakerRoleID                   uint64 `gorm:"type:bigint;not null"`
	REMARK2                       string `gorm:"type:text"`
	REMARK3                       string `gorm:"type:text"`
	ROUTING                       string `gorm:"type:varchar(100)"`
	SENDER_ADDRESS                string `gorm:"type:varchar(100)"`
	SENDER_BIRTH_CITY             string `gorm:"type:varchar(100)"`
	SENDER_BIRTH_COUNTRY          string `gorm:"type:varchar(100)"`
	SENDER_CITIZENSHIP_COUNTRY    string `gorm:"type:varchar(100)"`
	SENDER_CITY                   string `gorm:"type:varchar(100)"`
	SENDER_COUNTRY_CODE           string `gorm:"type:varchar(100)"`
	SENDER_DOB                    string `gorm:"type:varchar(100)"`
	SENDER_EMAIL                  string `gorm:"type:varchar(100)"`
	SENDER_GENDER                 string `gorm:"type:varchar(100)"`
	SENDER_ID_EXPIRED_DATE        string `gorm:"type:varchar(100)"`
	SENDER_ID_ISSUE_DATE          string `gorm:"type:varchar(100)"`
	SENDER_ID_NUMBER              string `gorm:"type:varchar(100)"`
	SENDER_ID_TYPE                string `gorm:"type:varchar(100)"`
	SENDER_NAME                   string `gorm:"type:varchar(100)"`
	SENDER_PHONE                  string `gorm:"type:varchar(100)"`
	SENDER_PHONE_CODE             string `gorm:"type:varchar(100)"`
	SENDER_POSITION               string `gorm:"type:varchar(100)"`
	SENDER_POSTCODE               string `gorm:"type:varchar(100)"`
	SENDER_PROFESSION             string `gorm:"type:varchar(100)"`
	SENDER_SOURCE_OF_FUND         string `gorm:"type:varchar(100)"`
	SENDER_STATE_PROVINCE         string `gorm:"type:varchar(100)"`
	SENDER_TYPE                   string `gorm:"type:varchar(100)"`
	SENDER_WORKING_STATUS         string `gorm:"type:varchar(100)"`
	ScheduleTransaction           string `gorm:"type:varchar(100)"`
	SelectedRoutePartner          string `gorm:"type:jsonb"`
	SendDate                      *time.Time
	TELLERID_KCBO                 string `gorm:"type:varchar(100)"`
	TICKET_NUMBER                 string `gorm:"type:varchar(100)"`
	TRANSACTION_ID                string `gorm:"type:varchar(100)"`
	TRANSACTION_PURPOSE           string `gorm:"type:varchar(100)"`
	TaskID                        uint64 `gorm:"type:bigint;not null"`
	TransactionResponse           string `gorm:"type:jsonb"`
	TransactionResponseMessage    string `gorm:"type:text"`
	TransactionStatus             string `gorm:"type:varchar(50);not null"`
	TransactionSvcId              uint64 `gorm:"default:0;not null"`
	USER_TRANSACTION              string `gorm:"type:varchar(100)"`
	USER_TRANSACTION_BRANCH_CODE  string `gorm:"type:varchar(100)"`
	UpdatedAt                     *time.Time
}

// TableName overrides the default tablename generated by GORM
func (RemittanceTransactionORM) TableName() string {
	return "remittance_transactions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *RemittanceTransaction) ToORM(ctx context.Context) (RemittanceTransactionORM, error) {
	to := RemittanceTransactionORM{}
	var err error
	if prehook, ok := interface{}(m).(RemittanceTransactionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.DataID = m.DataID
	to.AMOUNT_ORI = m.AMOUNT_ORI
	to.DEBIT_ACCOUNT = m.DEBIT_ACCOUNT
	to.DEBIT_AMOUNT = m.DEBIT_AMOUNT
	to.DEBIT_ACCOUNT_MAIN = m.DEBIT_ACCOUNT_MAIN
	to.FEE_ACCOUNT_CHANNEL = m.FEE_ACCOUNT_CHANNEL
	to.FEE_DEBIT_AMOUNT_CHANNEL = m.FEE_DEBIT_AMOUNT_CHANNEL
	to.FEE_CREDIT_AMOUNT_CHANNEL = m.FEE_CREDIT_AMOUNT_CHANNEL
	to.FEE_CURRENCY_CHANNEL = m.FEE_CURRENCY_CHANNEL
	to.TELLERID_KCBO = m.TELLERID_KCBO
	to.BOOK_RATE_BUY = m.BOOK_RATE_BUY
	to.BOOK_RATE_SELL = m.BOOK_RATE_SELL
	to.REMARK2 = m.REMARK2
	to.REMARK3 = m.REMARK3
	to.DEAL_CODE = m.DEAL_CODE
	to.COUNTERPART = m.COUNTERPART
	to.ROUTING = m.ROUTING
	to.TRANSACTION_ID = m.TRANSACTION_ID
	to.USER_TRANSACTION = m.USER_TRANSACTION
	to.USER_TRANSACTION_BRANCH_CODE = m.USER_TRANSACTION_BRANCH_CODE
	to.ACCOUNT_NUMBER = m.ACCOUNT_NUMBER
	to.CHARGES_TYPE = m.CHARGES_TYPE
	to.BENEFICIARY_BANK_CODE = m.BENEFICIARY_BANK_CODE
	to.BENEFICIARY_BANK_NAME = m.BENEFICIARY_BANK_NAME
	to.BENEFICIARY_BANK_ADDRESS = m.BENEFICIARY_BANK_ADDRESS
	to.BENEFICIARY_BANK_CITY = m.BENEFICIARY_BANK_CITY
	to.BENEFICIARY_BANK_COUNTRY_CODE = m.BENEFICIARY_BANK_COUNTRY_CODE
	to.BENEFICIARY_ID_TYPE = m.BENEFICIARY_ID_TYPE
	to.BENEFICIARY_ID_NUMBER = m.BENEFICIARY_ID_NUMBER
	to.BENEFICIARY_NAME = m.BENEFICIARY_NAME
	to.BENEFICIARY_ADDRESS = m.BENEFICIARY_ADDRESS
	to.BENEFICIARY_CITY = m.BENEFICIARY_CITY
	to.BENEFICIARY_STATE_PROVINCE = m.BENEFICIARY_STATE_PROVINCE
	to.BENEFICIARY_POSTCODE = m.BENEFICIARY_POSTCODE
	to.BENEFICIARY_COUNTRY_CODE = m.BENEFICIARY_COUNTRY_CODE
	to.BENEFICIARY_PHONE_CODE = m.BENEFICIARY_PHONE_CODE
	to.BENEFICIARY_PHONE = m.BENEFICIARY_PHONE
	to.BENEFICIARY_EMAIL = m.BENEFICIARY_EMAIL
	to.BENEFICIARY_TYPE = m.BENEFICIARY_TYPE
	to.BENEFICIARY_GENDER = m.BENEFICIARY_GENDER
	to.BENEFICIARY_RELATIONSHIP = m.BENEFICIARY_RELATIONSHIP
	to.CURRENCY_ORI = m.CURRENCY_ORI
	to.SENDER_ID_TYPE = m.SENDER_ID_TYPE
	to.SENDER_ID_NUMBER = m.SENDER_ID_NUMBER
	to.SENDER_ID_ISSUE_DATE = m.SENDER_ID_ISSUE_DATE
	to.SENDER_ID_EXPIRED_DATE = m.SENDER_ID_EXPIRED_DATE
	to.SENDER_NAME = m.SENDER_NAME
	to.SENDER_ADDRESS = m.SENDER_ADDRESS
	to.SENDER_CITY = m.SENDER_CITY
	to.SENDER_STATE_PROVINCE = m.SENDER_STATE_PROVINCE
	to.SENDER_POSTCODE = m.SENDER_POSTCODE
	to.SENDER_COUNTRY_CODE = m.SENDER_COUNTRY_CODE
	to.SENDER_DOB = m.SENDER_DOB
	to.SENDER_PHONE_CODE = m.SENDER_PHONE_CODE
	to.SENDER_PHONE = m.SENDER_PHONE
	to.SENDER_EMAIL = m.SENDER_EMAIL
	to.SENDER_TYPE = m.SENDER_TYPE
	to.SENDER_SOURCE_OF_FUND = m.SENDER_SOURCE_OF_FUND
	to.SENDER_BIRTH_CITY = m.SENDER_BIRTH_CITY
	to.SENDER_POSITION = m.SENDER_POSITION
	to.SENDER_BIRTH_COUNTRY = m.SENDER_BIRTH_COUNTRY
	to.SENDER_WORKING_STATUS = m.SENDER_WORKING_STATUS
	to.SENDER_PROFESSION = m.SENDER_PROFESSION
	to.SENDER_CITIZENSHIP_COUNTRY = m.SENDER_CITIZENSHIP_COUNTRY
	to.SENDER_GENDER = m.SENDER_GENDER
	to.TRANSACTION_PURPOSE = m.TRANSACTION_PURPOSE
	to.INTENDED_USE = m.INTENDED_USE
	to.TICKET_NUMBER = m.TICKET_NUMBER
	if m.SendDate != nil {
		t := m.SendDate.AsTime()
		to.SendDate = &t
	}
	to.CompanyID = m.CompanyID
	to.MakerRoleID = m.MakerRoleID
	to.DebitAccountID = m.DebitAccountID
	to.DebitAccountAlias = m.DebitAccountAlias
	to.DebitCurrencyID = m.DebitCurrencyID
	to.FileAttachment = m.FileAttachment
	to.ScheduleTransaction = m.ScheduleTransaction
	to.SelectedRoutePartner = m.SelectedRoutePartner
	to.BeneficiaryEmails = m.BeneficiaryEmails
	to.TransactionResponse = m.TransactionResponse
	to.CompletedDate = m.CompletedDate
	to.TransactionStatus = m.TransactionStatus
	to.TaskID = m.TaskID
	to.TransactionResponseMessage = m.TransactionResponseMessage
	to.CheckStatusResponse = m.CheckStatusResponse
	to.LastChecked = m.LastChecked
	to.BeneficiaryCountryName = m.BeneficiaryCountryName
	to.TransactionSvcId = m.TransactionSvcId
	to.CreatedByName = m.CreatedByName
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(RemittanceTransactionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *RemittanceTransactionORM) ToPB(ctx context.Context) (RemittanceTransaction, error) {
	to := RemittanceTransaction{}
	var err error
	if prehook, ok := interface{}(m).(RemittanceTransactionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.DataID = m.DataID
	to.AMOUNT_ORI = m.AMOUNT_ORI
	to.DEBIT_ACCOUNT = m.DEBIT_ACCOUNT
	to.DEBIT_AMOUNT = m.DEBIT_AMOUNT
	to.DEBIT_ACCOUNT_MAIN = m.DEBIT_ACCOUNT_MAIN
	to.FEE_ACCOUNT_CHANNEL = m.FEE_ACCOUNT_CHANNEL
	to.FEE_DEBIT_AMOUNT_CHANNEL = m.FEE_DEBIT_AMOUNT_CHANNEL
	to.FEE_CREDIT_AMOUNT_CHANNEL = m.FEE_CREDIT_AMOUNT_CHANNEL
	to.FEE_CURRENCY_CHANNEL = m.FEE_CURRENCY_CHANNEL
	to.TELLERID_KCBO = m.TELLERID_KCBO
	to.BOOK_RATE_BUY = m.BOOK_RATE_BUY
	to.BOOK_RATE_SELL = m.BOOK_RATE_SELL
	to.REMARK2 = m.REMARK2
	to.REMARK3 = m.REMARK3
	to.DEAL_CODE = m.DEAL_CODE
	to.COUNTERPART = m.COUNTERPART
	to.ROUTING = m.ROUTING
	to.TRANSACTION_ID = m.TRANSACTION_ID
	to.USER_TRANSACTION = m.USER_TRANSACTION
	to.USER_TRANSACTION_BRANCH_CODE = m.USER_TRANSACTION_BRANCH_CODE
	to.ACCOUNT_NUMBER = m.ACCOUNT_NUMBER
	to.CHARGES_TYPE = m.CHARGES_TYPE
	to.BENEFICIARY_BANK_CODE = m.BENEFICIARY_BANK_CODE
	to.BENEFICIARY_BANK_NAME = m.BENEFICIARY_BANK_NAME
	to.BENEFICIARY_BANK_ADDRESS = m.BENEFICIARY_BANK_ADDRESS
	to.BENEFICIARY_BANK_CITY = m.BENEFICIARY_BANK_CITY
	to.BENEFICIARY_BANK_COUNTRY_CODE = m.BENEFICIARY_BANK_COUNTRY_CODE
	to.BENEFICIARY_ID_TYPE = m.BENEFICIARY_ID_TYPE
	to.BENEFICIARY_ID_NUMBER = m.BENEFICIARY_ID_NUMBER
	to.BENEFICIARY_NAME = m.BENEFICIARY_NAME
	to.BENEFICIARY_ADDRESS = m.BENEFICIARY_ADDRESS
	to.BENEFICIARY_CITY = m.BENEFICIARY_CITY
	to.BENEFICIARY_STATE_PROVINCE = m.BENEFICIARY_STATE_PROVINCE
	to.BENEFICIARY_POSTCODE = m.BENEFICIARY_POSTCODE
	to.BENEFICIARY_COUNTRY_CODE = m.BENEFICIARY_COUNTRY_CODE
	to.BENEFICIARY_PHONE_CODE = m.BENEFICIARY_PHONE_CODE
	to.BENEFICIARY_PHONE = m.BENEFICIARY_PHONE
	to.BENEFICIARY_EMAIL = m.BENEFICIARY_EMAIL
	to.BENEFICIARY_TYPE = m.BENEFICIARY_TYPE
	to.BENEFICIARY_GENDER = m.BENEFICIARY_GENDER
	to.BENEFICIARY_RELATIONSHIP = m.BENEFICIARY_RELATIONSHIP
	to.CURRENCY_ORI = m.CURRENCY_ORI
	to.SENDER_ID_TYPE = m.SENDER_ID_TYPE
	to.SENDER_ID_NUMBER = m.SENDER_ID_NUMBER
	to.SENDER_ID_ISSUE_DATE = m.SENDER_ID_ISSUE_DATE
	to.SENDER_ID_EXPIRED_DATE = m.SENDER_ID_EXPIRED_DATE
	to.SENDER_NAME = m.SENDER_NAME
	to.SENDER_ADDRESS = m.SENDER_ADDRESS
	to.SENDER_CITY = m.SENDER_CITY
	to.SENDER_STATE_PROVINCE = m.SENDER_STATE_PROVINCE
	to.SENDER_POSTCODE = m.SENDER_POSTCODE
	to.SENDER_COUNTRY_CODE = m.SENDER_COUNTRY_CODE
	to.SENDER_DOB = m.SENDER_DOB
	to.SENDER_PHONE_CODE = m.SENDER_PHONE_CODE
	to.SENDER_PHONE = m.SENDER_PHONE
	to.SENDER_EMAIL = m.SENDER_EMAIL
	to.SENDER_TYPE = m.SENDER_TYPE
	to.SENDER_SOURCE_OF_FUND = m.SENDER_SOURCE_OF_FUND
	to.SENDER_BIRTH_CITY = m.SENDER_BIRTH_CITY
	to.SENDER_POSITION = m.SENDER_POSITION
	to.SENDER_BIRTH_COUNTRY = m.SENDER_BIRTH_COUNTRY
	to.SENDER_WORKING_STATUS = m.SENDER_WORKING_STATUS
	to.SENDER_PROFESSION = m.SENDER_PROFESSION
	to.SENDER_CITIZENSHIP_COUNTRY = m.SENDER_CITIZENSHIP_COUNTRY
	to.SENDER_GENDER = m.SENDER_GENDER
	to.TRANSACTION_PURPOSE = m.TRANSACTION_PURPOSE
	to.INTENDED_USE = m.INTENDED_USE
	to.TICKET_NUMBER = m.TICKET_NUMBER
	if m.SendDate != nil {
		to.SendDate = timestamppb.New(*m.SendDate)
	}
	to.CompanyID = m.CompanyID
	to.MakerRoleID = m.MakerRoleID
	to.DebitAccountID = m.DebitAccountID
	to.DebitAccountAlias = m.DebitAccountAlias
	to.DebitCurrencyID = m.DebitCurrencyID
	to.FileAttachment = m.FileAttachment
	to.ScheduleTransaction = m.ScheduleTransaction
	to.SelectedRoutePartner = m.SelectedRoutePartner
	to.BeneficiaryEmails = m.BeneficiaryEmails
	to.TransactionResponse = m.TransactionResponse
	to.CompletedDate = m.CompletedDate
	to.TransactionStatus = m.TransactionStatus
	to.TaskID = m.TaskID
	to.TransactionResponseMessage = m.TransactionResponseMessage
	to.CheckStatusResponse = m.CheckStatusResponse
	to.LastChecked = m.LastChecked
	to.BeneficiaryCountryName = m.BeneficiaryCountryName
	to.TransactionSvcId = m.TransactionSvcId
	to.CreatedByName = m.CreatedByName
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(RemittanceTransactionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type RemittanceTransaction the arg will be the target, the caller the one being converted from

// RemittanceTransactionBeforeToORM called before default ToORM code
type RemittanceTransactionWithBeforeToORM interface {
	BeforeToORM(context.Context, *RemittanceTransactionORM) error
}

// RemittanceTransactionAfterToORM called after default ToORM code
type RemittanceTransactionWithAfterToORM interface {
	AfterToORM(context.Context, *RemittanceTransactionORM) error
}

// RemittanceTransactionBeforeToPB called before default ToPB code
type RemittanceTransactionWithBeforeToPB interface {
	BeforeToPB(context.Context, *RemittanceTransaction) error
}

// RemittanceTransactionAfterToPB called after default ToPB code
type RemittanceTransactionWithAfterToPB interface {
	AfterToPB(context.Context, *RemittanceTransaction) error
}

type NostroPriorityORM struct {
	CurrencyCode string `gorm:"type:varchar(5);primary_key;not null"`
	OrderNumber  uint32 `gorm:"type:integer;not null"`
}

// TableName overrides the default tablename generated by GORM
func (NostroPriorityORM) TableName() string {
	return "nostro_priorities"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *NostroPriority) ToORM(ctx context.Context) (NostroPriorityORM, error) {
	to := NostroPriorityORM{}
	var err error
	if prehook, ok := interface{}(m).(NostroPriorityWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.CurrencyCode = m.CurrencyCode
	to.OrderNumber = m.OrderNumber
	if posthook, ok := interface{}(m).(NostroPriorityWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *NostroPriorityORM) ToPB(ctx context.Context) (NostroPriority, error) {
	to := NostroPriority{}
	var err error
	if prehook, ok := interface{}(m).(NostroPriorityWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.CurrencyCode = m.CurrencyCode
	to.OrderNumber = m.OrderNumber
	if posthook, ok := interface{}(m).(NostroPriorityWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type NostroPriority the arg will be the target, the caller the one being converted from

// NostroPriorityBeforeToORM called before default ToORM code
type NostroPriorityWithBeforeToORM interface {
	BeforeToORM(context.Context, *NostroPriorityORM) error
}

// NostroPriorityAfterToORM called after default ToORM code
type NostroPriorityWithAfterToORM interface {
	AfterToORM(context.Context, *NostroPriorityORM) error
}

// NostroPriorityBeforeToPB called before default ToPB code
type NostroPriorityWithBeforeToPB interface {
	BeforeToPB(context.Context, *NostroPriority) error
}

// NostroPriorityAfterToPB called after default ToPB code
type NostroPriorityWithAfterToPB interface {
	AfterToPB(context.Context, *NostroPriority) error
}

// DefaultCreateRemittanceTransaction executes a basic gorm create call
func DefaultCreateRemittanceTransaction(ctx context.Context, in *RemittanceTransaction, db *gorm.DB) (*RemittanceTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type RemittanceTransactionORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadRemittanceTransaction(ctx context.Context, in *RemittanceTransaction, db *gorm.DB) (*RemittanceTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.DataID == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &RemittanceTransactionORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := RemittanceTransactionORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(RemittanceTransactionORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type RemittanceTransactionORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteRemittanceTransaction(ctx context.Context, in *RemittanceTransaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.DataID == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&RemittanceTransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type RemittanceTransactionORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteRemittanceTransactionSet(ctx context.Context, in []*RemittanceTransaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.DataID == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.DataID)
	}
	if hook, ok := (interface{}(&RemittanceTransactionORM{})).(RemittanceTransactionORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("data_id in (?)", keys).Delete(&RemittanceTransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&RemittanceTransactionORM{})).(RemittanceTransactionORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type RemittanceTransactionORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*RemittanceTransaction, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*RemittanceTransaction, *gorm.DB) error
}

// DefaultStrictUpdateRemittanceTransaction clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateRemittanceTransaction(ctx context.Context, in *RemittanceTransaction, db *gorm.DB) (*RemittanceTransaction, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateRemittanceTransaction")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &RemittanceTransactionORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("data_id=?", ormObj.DataID).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type RemittanceTransactionORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchRemittanceTransaction executes a basic gorm update call with patch behavior
func DefaultPatchRemittanceTransaction(ctx context.Context, in *RemittanceTransaction, updateMask *field_mask.FieldMask, db *gorm.DB) (*RemittanceTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj RemittanceTransaction
	var err error
	if hook, ok := interface{}(&pbObj).(RemittanceTransactionWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(RemittanceTransactionWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskRemittanceTransaction(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(RemittanceTransactionWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateRemittanceTransaction(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(RemittanceTransactionWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type RemittanceTransactionWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *RemittanceTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *RemittanceTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *RemittanceTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *RemittanceTransaction, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetRemittanceTransaction executes a bulk gorm update call with patch behavior
func DefaultPatchSetRemittanceTransaction(ctx context.Context, objects []*RemittanceTransaction, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*RemittanceTransaction, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*RemittanceTransaction, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchRemittanceTransaction(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskRemittanceTransaction patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskRemittanceTransaction(ctx context.Context, patchee *RemittanceTransaction, patcher *RemittanceTransaction, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*RemittanceTransaction, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedSendDate bool
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"DataID" {
			patchee.DataID = patcher.DataID
			continue
		}
		if f == prefix+"AMOUNT_ORI" {
			patchee.AMOUNT_ORI = patcher.AMOUNT_ORI
			continue
		}
		if f == prefix+"DEBIT_ACCOUNT" {
			patchee.DEBIT_ACCOUNT = patcher.DEBIT_ACCOUNT
			continue
		}
		if f == prefix+"DEBIT_AMOUNT" {
			patchee.DEBIT_AMOUNT = patcher.DEBIT_AMOUNT
			continue
		}
		if f == prefix+"DEBIT_ACCOUNT_MAIN" {
			patchee.DEBIT_ACCOUNT_MAIN = patcher.DEBIT_ACCOUNT_MAIN
			continue
		}
		if f == prefix+"FEE_ACCOUNT_CHANNEL" {
			patchee.FEE_ACCOUNT_CHANNEL = patcher.FEE_ACCOUNT_CHANNEL
			continue
		}
		if f == prefix+"FEE_DEBIT_AMOUNT_CHANNEL" {
			patchee.FEE_DEBIT_AMOUNT_CHANNEL = patcher.FEE_DEBIT_AMOUNT_CHANNEL
			continue
		}
		if f == prefix+"FEE_CREDIT_AMOUNT_CHANNEL" {
			patchee.FEE_CREDIT_AMOUNT_CHANNEL = patcher.FEE_CREDIT_AMOUNT_CHANNEL
			continue
		}
		if f == prefix+"FEE_CURRENCY_CHANNEL" {
			patchee.FEE_CURRENCY_CHANNEL = patcher.FEE_CURRENCY_CHANNEL
			continue
		}
		if f == prefix+"TELLERID_KCBO" {
			patchee.TELLERID_KCBO = patcher.TELLERID_KCBO
			continue
		}
		if f == prefix+"BOOK_RATE_BUY" {
			patchee.BOOK_RATE_BUY = patcher.BOOK_RATE_BUY
			continue
		}
		if f == prefix+"BOOK_RATE_SELL" {
			patchee.BOOK_RATE_SELL = patcher.BOOK_RATE_SELL
			continue
		}
		if f == prefix+"REMARK2" {
			patchee.REMARK2 = patcher.REMARK2
			continue
		}
		if f == prefix+"REMARK3" {
			patchee.REMARK3 = patcher.REMARK3
			continue
		}
		if f == prefix+"DEAL_CODE" {
			patchee.DEAL_CODE = patcher.DEAL_CODE
			continue
		}
		if f == prefix+"COUNTERPART" {
			patchee.COUNTERPART = patcher.COUNTERPART
			continue
		}
		if f == prefix+"ROUTING" {
			patchee.ROUTING = patcher.ROUTING
			continue
		}
		if f == prefix+"TRANSACTION_ID" {
			patchee.TRANSACTION_ID = patcher.TRANSACTION_ID
			continue
		}
		if f == prefix+"USER_TRANSACTION" {
			patchee.USER_TRANSACTION = patcher.USER_TRANSACTION
			continue
		}
		if f == prefix+"USER_TRANSACTION_BRANCH_CODE" {
			patchee.USER_TRANSACTION_BRANCH_CODE = patcher.USER_TRANSACTION_BRANCH_CODE
			continue
		}
		if f == prefix+"ACCOUNT_NUMBER" {
			patchee.ACCOUNT_NUMBER = patcher.ACCOUNT_NUMBER
			continue
		}
		if f == prefix+"CHARGES_TYPE" {
			patchee.CHARGES_TYPE = patcher.CHARGES_TYPE
			continue
		}
		if f == prefix+"BENEFICIARY_BANK_CODE" {
			patchee.BENEFICIARY_BANK_CODE = patcher.BENEFICIARY_BANK_CODE
			continue
		}
		if f == prefix+"BENEFICIARY_BANK_NAME" {
			patchee.BENEFICIARY_BANK_NAME = patcher.BENEFICIARY_BANK_NAME
			continue
		}
		if f == prefix+"BENEFICIARY_BANK_ADDRESS" {
			patchee.BENEFICIARY_BANK_ADDRESS = patcher.BENEFICIARY_BANK_ADDRESS
			continue
		}
		if f == prefix+"BENEFICIARY_BANK_CITY" {
			patchee.BENEFICIARY_BANK_CITY = patcher.BENEFICIARY_BANK_CITY
			continue
		}
		if f == prefix+"BENEFICIARY_BANK_COUNTRY_CODE" {
			patchee.BENEFICIARY_BANK_COUNTRY_CODE = patcher.BENEFICIARY_BANK_COUNTRY_CODE
			continue
		}
		if f == prefix+"BENEFICIARY_ID_TYPE" {
			patchee.BENEFICIARY_ID_TYPE = patcher.BENEFICIARY_ID_TYPE
			continue
		}
		if f == prefix+"BENEFICIARY_ID_NUMBER" {
			patchee.BENEFICIARY_ID_NUMBER = patcher.BENEFICIARY_ID_NUMBER
			continue
		}
		if f == prefix+"BENEFICIARY_NAME" {
			patchee.BENEFICIARY_NAME = patcher.BENEFICIARY_NAME
			continue
		}
		if f == prefix+"BENEFICIARY_ADDRESS" {
			patchee.BENEFICIARY_ADDRESS = patcher.BENEFICIARY_ADDRESS
			continue
		}
		if f == prefix+"BENEFICIARY_CITY" {
			patchee.BENEFICIARY_CITY = patcher.BENEFICIARY_CITY
			continue
		}
		if f == prefix+"BENEFICIARY_STATE_PROVINCE" {
			patchee.BENEFICIARY_STATE_PROVINCE = patcher.BENEFICIARY_STATE_PROVINCE
			continue
		}
		if f == prefix+"BENEFICIARY_POSTCODE" {
			patchee.BENEFICIARY_POSTCODE = patcher.BENEFICIARY_POSTCODE
			continue
		}
		if f == prefix+"BENEFICIARY_COUNTRY_CODE" {
			patchee.BENEFICIARY_COUNTRY_CODE = patcher.BENEFICIARY_COUNTRY_CODE
			continue
		}
		if f == prefix+"BENEFICIARY_PHONE_CODE" {
			patchee.BENEFICIARY_PHONE_CODE = patcher.BENEFICIARY_PHONE_CODE
			continue
		}
		if f == prefix+"BENEFICIARY_PHONE" {
			patchee.BENEFICIARY_PHONE = patcher.BENEFICIARY_PHONE
			continue
		}
		if f == prefix+"BENEFICIARY_EMAIL" {
			patchee.BENEFICIARY_EMAIL = patcher.BENEFICIARY_EMAIL
			continue
		}
		if f == prefix+"BENEFICIARY_TYPE" {
			patchee.BENEFICIARY_TYPE = patcher.BENEFICIARY_TYPE
			continue
		}
		if f == prefix+"BENEFICIARY_GENDER" {
			patchee.BENEFICIARY_GENDER = patcher.BENEFICIARY_GENDER
			continue
		}
		if f == prefix+"BENEFICIARY_RELATIONSHIP" {
			patchee.BENEFICIARY_RELATIONSHIP = patcher.BENEFICIARY_RELATIONSHIP
			continue
		}
		if f == prefix+"CURRENCY_ORI" {
			patchee.CURRENCY_ORI = patcher.CURRENCY_ORI
			continue
		}
		if f == prefix+"SENDER_ID_TYPE" {
			patchee.SENDER_ID_TYPE = patcher.SENDER_ID_TYPE
			continue
		}
		if f == prefix+"SENDER_ID_NUMBER" {
			patchee.SENDER_ID_NUMBER = patcher.SENDER_ID_NUMBER
			continue
		}
		if f == prefix+"SENDER_ID_ISSUE_DATE" {
			patchee.SENDER_ID_ISSUE_DATE = patcher.SENDER_ID_ISSUE_DATE
			continue
		}
		if f == prefix+"SENDER_ID_EXPIRED_DATE" {
			patchee.SENDER_ID_EXPIRED_DATE = patcher.SENDER_ID_EXPIRED_DATE
			continue
		}
		if f == prefix+"SENDER_NAME" {
			patchee.SENDER_NAME = patcher.SENDER_NAME
			continue
		}
		if f == prefix+"SENDER_ADDRESS" {
			patchee.SENDER_ADDRESS = patcher.SENDER_ADDRESS
			continue
		}
		if f == prefix+"SENDER_CITY" {
			patchee.SENDER_CITY = patcher.SENDER_CITY
			continue
		}
		if f == prefix+"SENDER_STATE_PROVINCE" {
			patchee.SENDER_STATE_PROVINCE = patcher.SENDER_STATE_PROVINCE
			continue
		}
		if f == prefix+"SENDER_POSTCODE" {
			patchee.SENDER_POSTCODE = patcher.SENDER_POSTCODE
			continue
		}
		if f == prefix+"SENDER_COUNTRY_CODE" {
			patchee.SENDER_COUNTRY_CODE = patcher.SENDER_COUNTRY_CODE
			continue
		}
		if f == prefix+"SENDER_DOB" {
			patchee.SENDER_DOB = patcher.SENDER_DOB
			continue
		}
		if f == prefix+"SENDER_PHONE_CODE" {
			patchee.SENDER_PHONE_CODE = patcher.SENDER_PHONE_CODE
			continue
		}
		if f == prefix+"SENDER_PHONE" {
			patchee.SENDER_PHONE = patcher.SENDER_PHONE
			continue
		}
		if f == prefix+"SENDER_EMAIL" {
			patchee.SENDER_EMAIL = patcher.SENDER_EMAIL
			continue
		}
		if f == prefix+"SENDER_TYPE" {
			patchee.SENDER_TYPE = patcher.SENDER_TYPE
			continue
		}
		if f == prefix+"SENDER_SOURCE_OF_FUND" {
			patchee.SENDER_SOURCE_OF_FUND = patcher.SENDER_SOURCE_OF_FUND
			continue
		}
		if f == prefix+"SENDER_BIRTH_CITY" {
			patchee.SENDER_BIRTH_CITY = patcher.SENDER_BIRTH_CITY
			continue
		}
		if f == prefix+"SENDER_POSITION" {
			patchee.SENDER_POSITION = patcher.SENDER_POSITION
			continue
		}
		if f == prefix+"SENDER_BIRTH_COUNTRY" {
			patchee.SENDER_BIRTH_COUNTRY = patcher.SENDER_BIRTH_COUNTRY
			continue
		}
		if f == prefix+"SENDER_WORKING_STATUS" {
			patchee.SENDER_WORKING_STATUS = patcher.SENDER_WORKING_STATUS
			continue
		}
		if f == prefix+"SENDER_PROFESSION" {
			patchee.SENDER_PROFESSION = patcher.SENDER_PROFESSION
			continue
		}
		if f == prefix+"SENDER_CITIZENSHIP_COUNTRY" {
			patchee.SENDER_CITIZENSHIP_COUNTRY = patcher.SENDER_CITIZENSHIP_COUNTRY
			continue
		}
		if f == prefix+"SENDER_GENDER" {
			patchee.SENDER_GENDER = patcher.SENDER_GENDER
			continue
		}
		if f == prefix+"TRANSACTION_PURPOSE" {
			patchee.TRANSACTION_PURPOSE = patcher.TRANSACTION_PURPOSE
			continue
		}
		if f == prefix+"INTENDED_USE" {
			patchee.INTENDED_USE = patcher.INTENDED_USE
			continue
		}
		if f == prefix+"TICKET_NUMBER" {
			patchee.TICKET_NUMBER = patcher.TICKET_NUMBER
			continue
		}
		if !updatedSendDate && strings.HasPrefix(f, prefix+"SendDate.") {
			if patcher.SendDate == nil {
				patchee.SendDate = nil
				continue
			}
			if patchee.SendDate == nil {
				patchee.SendDate = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"SendDate."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.SendDate, patchee.SendDate, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"SendDate" {
			updatedSendDate = true
			patchee.SendDate = patcher.SendDate
			continue
		}
		if f == prefix+"CompanyID" {
			patchee.CompanyID = patcher.CompanyID
			continue
		}
		if f == prefix+"MakerRoleID" {
			patchee.MakerRoleID = patcher.MakerRoleID
			continue
		}
		if f == prefix+"DebitAccountID" {
			patchee.DebitAccountID = patcher.DebitAccountID
			continue
		}
		if f == prefix+"DebitAccountAlias" {
			patchee.DebitAccountAlias = patcher.DebitAccountAlias
			continue
		}
		if f == prefix+"DebitCurrencyID" {
			patchee.DebitCurrencyID = patcher.DebitCurrencyID
			continue
		}
		if f == prefix+"FileAttachment" {
			patchee.FileAttachment = patcher.FileAttachment
			continue
		}
		if f == prefix+"ScheduleTransaction" {
			patchee.ScheduleTransaction = patcher.ScheduleTransaction
			continue
		}
		if f == prefix+"SelectedRoutePartner" {
			patchee.SelectedRoutePartner = patcher.SelectedRoutePartner
			continue
		}
		if f == prefix+"BeneficiaryEmails" {
			patchee.BeneficiaryEmails = patcher.BeneficiaryEmails
			continue
		}
		if f == prefix+"TransactionResponse" {
			patchee.TransactionResponse = patcher.TransactionResponse
			continue
		}
		if f == prefix+"CompletedDate" {
			patchee.CompletedDate = patcher.CompletedDate
			continue
		}
		if f == prefix+"TransactionStatus" {
			patchee.TransactionStatus = patcher.TransactionStatus
			continue
		}
		if f == prefix+"TaskID" {
			patchee.TaskID = patcher.TaskID
			continue
		}
		if f == prefix+"TransactionResponseMessage" {
			patchee.TransactionResponseMessage = patcher.TransactionResponseMessage
			continue
		}
		if f == prefix+"CheckStatusResponse" {
			patchee.CheckStatusResponse = patcher.CheckStatusResponse
			continue
		}
		if f == prefix+"LastChecked" {
			patchee.LastChecked = patcher.LastChecked
			continue
		}
		if f == prefix+"BeneficiaryCountryName" {
			patchee.BeneficiaryCountryName = patcher.BeneficiaryCountryName
			continue
		}
		if f == prefix+"TransactionSvcId" {
			patchee.TransactionSvcId = patcher.TransactionSvcId
			continue
		}
		if f == prefix+"CreatedByName" {
			patchee.CreatedByName = patcher.CreatedByName
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListRemittanceTransaction executes a gorm list call
func DefaultListRemittanceTransaction(ctx context.Context, db *gorm.DB) ([]*RemittanceTransaction, error) {
	in := RemittanceTransaction{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &RemittanceTransactionORM{}, &RemittanceTransaction{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("data_id")
	ormResponse := []RemittanceTransactionORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(RemittanceTransactionORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*RemittanceTransaction{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type RemittanceTransactionORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type RemittanceTransactionORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]RemittanceTransactionORM) error
}

// DefaultCreateNostroPriority executes a basic gorm create call
func DefaultCreateNostroPriority(ctx context.Context, in *NostroPriority, db *gorm.DB) (*NostroPriority, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type NostroPriorityORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadNostroPriority(ctx context.Context, in *NostroPriority, db *gorm.DB) (*NostroPriority, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.CurrencyCode == "" {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &NostroPriorityORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := NostroPriorityORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(NostroPriorityORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type NostroPriorityORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteNostroPriority(ctx context.Context, in *NostroPriority, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.CurrencyCode == "" {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&NostroPriorityORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type NostroPriorityORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteNostroPrioritySet(ctx context.Context, in []*NostroPriority, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.CurrencyCode == "" {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.CurrencyCode)
	}
	if hook, ok := (interface{}(&NostroPriorityORM{})).(NostroPriorityORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("currency_code in (?)", keys).Delete(&NostroPriorityORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&NostroPriorityORM{})).(NostroPriorityORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type NostroPriorityORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*NostroPriority, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*NostroPriority, *gorm.DB) error
}

// DefaultStrictUpdateNostroPriority clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateNostroPriority(ctx context.Context, in *NostroPriority, db *gorm.DB) (*NostroPriority, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateNostroPriority")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &NostroPriorityORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("currency_code=?", ormObj.CurrencyCode).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type NostroPriorityORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchNostroPriority executes a basic gorm update call with patch behavior
func DefaultPatchNostroPriority(ctx context.Context, in *NostroPriority, updateMask *field_mask.FieldMask, db *gorm.DB) (*NostroPriority, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj NostroPriority
	var err error
	if hook, ok := interface{}(&pbObj).(NostroPriorityWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&pbObj).(NostroPriorityWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskNostroPriority(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(NostroPriorityWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateNostroPriority(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(NostroPriorityWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type NostroPriorityWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *NostroPriority, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *NostroPriority, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *NostroPriority, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *NostroPriority, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetNostroPriority executes a bulk gorm update call with patch behavior
func DefaultPatchSetNostroPriority(ctx context.Context, objects []*NostroPriority, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*NostroPriority, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*NostroPriority, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchNostroPriority(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskNostroPriority patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskNostroPriority(ctx context.Context, patchee *NostroPriority, patcher *NostroPriority, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*NostroPriority, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"CurrencyCode" {
			patchee.CurrencyCode = patcher.CurrencyCode
			continue
		}
		if f == prefix+"OrderNumber" {
			patchee.OrderNumber = patcher.OrderNumber
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListNostroPriority executes a gorm list call
func DefaultListNostroPriority(ctx context.Context, db *gorm.DB) ([]*NostroPriority, error) {
	in := NostroPriority{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &NostroPriorityORM{}, &NostroPriority{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("currency_code")
	ormResponse := []NostroPriorityORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(NostroPriorityORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*NostroPriority{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type NostroPriorityORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type NostroPriorityORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]NostroPriorityORM) error
}
