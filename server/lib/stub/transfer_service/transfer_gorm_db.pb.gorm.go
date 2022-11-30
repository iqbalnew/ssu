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

type InternalTransferTransactionORM struct {
	CreatedAt            *time.Time `gorm:"not null"`
	Data                 string     `gorm:"type:jsonb"`
	Id                   uint64     `gorm:"primary_key;not null"`
	TaskID               uint64
	TransactionID        string
	TransactionServiceID uint64
	UpdatedAt            *time.Time `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (InternalTransferTransactionORM) TableName() string {
	return "internal_transfer_transactions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *InternalTransferTransaction) ToORM(ctx context.Context) (InternalTransferTransactionORM, error) {
	to := InternalTransferTransactionORM{}
	var err error
	if prehook, ok := interface{}(m).(InternalTransferTransactionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.TaskID = m.TaskID
	to.TransactionID = m.TransactionID
	to.TransactionServiceID = m.TransactionServiceID
	to.Data = m.Data
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(InternalTransferTransactionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *InternalTransferTransactionORM) ToPB(ctx context.Context) (InternalTransferTransaction, error) {
	to := InternalTransferTransaction{}
	var err error
	if prehook, ok := interface{}(m).(InternalTransferTransactionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.TaskID = m.TaskID
	to.TransactionID = m.TransactionID
	to.TransactionServiceID = m.TransactionServiceID
	to.Data = m.Data
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(InternalTransferTransactionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type InternalTransferTransaction the arg will be the target, the caller the one being converted from

// InternalTransferTransactionBeforeToORM called before default ToORM code
type InternalTransferTransactionWithBeforeToORM interface {
	BeforeToORM(context.Context, *InternalTransferTransactionORM) error
}

// InternalTransferTransactionAfterToORM called after default ToORM code
type InternalTransferTransactionWithAfterToORM interface {
	AfterToORM(context.Context, *InternalTransferTransactionORM) error
}

// InternalTransferTransactionBeforeToPB called before default ToPB code
type InternalTransferTransactionWithBeforeToPB interface {
	BeforeToPB(context.Context, *InternalTransferTransaction) error
}

// InternalTransferTransactionAfterToPB called after default ToPB code
type InternalTransferTransactionWithAfterToPB interface {
	AfterToPB(context.Context, *InternalTransferTransaction) error
}

type InternalTransferSingleTemplateORM struct {
	CreatedAt    *time.Time `gorm:"not null"`
	Data         string     `gorm:"type:jsonb"`
	Id           uint64     `gorm:"primary_key;not null"`
	TemplateName string
	UpdatedAt    *time.Time `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (InternalTransferSingleTemplateORM) TableName() string {
	return "internal_single_templates"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *InternalTransferSingleTemplate) ToORM(ctx context.Context) (InternalTransferSingleTemplateORM, error) {
	to := InternalTransferSingleTemplateORM{}
	var err error
	if prehook, ok := interface{}(m).(InternalTransferSingleTemplateWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Data = m.Data
	to.TemplateName = m.TemplateName
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(InternalTransferSingleTemplateWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *InternalTransferSingleTemplateORM) ToPB(ctx context.Context) (InternalTransferSingleTemplate, error) {
	to := InternalTransferSingleTemplate{}
	var err error
	if prehook, ok := interface{}(m).(InternalTransferSingleTemplateWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Data = m.Data
	to.TemplateName = m.TemplateName
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(InternalTransferSingleTemplateWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type InternalTransferSingleTemplate the arg will be the target, the caller the one being converted from

// InternalTransferSingleTemplateBeforeToORM called before default ToORM code
type InternalTransferSingleTemplateWithBeforeToORM interface {
	BeforeToORM(context.Context, *InternalTransferSingleTemplateORM) error
}

// InternalTransferSingleTemplateAfterToORM called after default ToORM code
type InternalTransferSingleTemplateWithAfterToORM interface {
	AfterToORM(context.Context, *InternalTransferSingleTemplateORM) error
}

// InternalTransferSingleTemplateBeforeToPB called before default ToPB code
type InternalTransferSingleTemplateWithBeforeToPB interface {
	BeforeToPB(context.Context, *InternalTransferSingleTemplate) error
}

// InternalTransferSingleTemplateAfterToPB called after default ToPB code
type InternalTransferSingleTemplateWithAfterToPB interface {
	AfterToPB(context.Context, *InternalTransferSingleTemplate) error
}

// DefaultCreateInternalTransferTransaction executes a basic gorm create call
func DefaultCreateInternalTransferTransaction(ctx context.Context, in *InternalTransferTransaction, db *gorm.DB) (*InternalTransferTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type InternalTransferTransactionORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadInternalTransferTransaction(ctx context.Context, in *InternalTransferTransaction, db *gorm.DB) (*InternalTransferTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &InternalTransferTransactionORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := InternalTransferTransactionORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(InternalTransferTransactionORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type InternalTransferTransactionORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteInternalTransferTransaction(ctx context.Context, in *InternalTransferTransaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&InternalTransferTransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type InternalTransferTransactionORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteInternalTransferTransactionSet(ctx context.Context, in []*InternalTransferTransaction, db *gorm.DB) error {
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
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&InternalTransferTransactionORM{})).(InternalTransferTransactionORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&InternalTransferTransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&InternalTransferTransactionORM{})).(InternalTransferTransactionORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type InternalTransferTransactionORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*InternalTransferTransaction, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*InternalTransferTransaction, *gorm.DB) error
}

// DefaultStrictUpdateInternalTransferTransaction clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateInternalTransferTransaction(ctx context.Context, in *InternalTransferTransaction, db *gorm.DB) (*InternalTransferTransaction, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateInternalTransferTransaction")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &InternalTransferTransactionORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithAfterStrictUpdateSave); ok {
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

type InternalTransferTransactionORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchInternalTransferTransaction executes a basic gorm update call with patch behavior
func DefaultPatchInternalTransferTransaction(ctx context.Context, in *InternalTransferTransaction, updateMask *field_mask.FieldMask, db *gorm.DB) (*InternalTransferTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj InternalTransferTransaction
	var err error
	if hook, ok := interface{}(&pbObj).(InternalTransferTransactionWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadInternalTransferTransaction(ctx, &InternalTransferTransaction{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(InternalTransferTransactionWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskInternalTransferTransaction(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(InternalTransferTransactionWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateInternalTransferTransaction(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(InternalTransferTransactionWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type InternalTransferTransactionWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *InternalTransferTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *InternalTransferTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *InternalTransferTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *InternalTransferTransaction, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetInternalTransferTransaction executes a bulk gorm update call with patch behavior
func DefaultPatchSetInternalTransferTransaction(ctx context.Context, objects []*InternalTransferTransaction, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*InternalTransferTransaction, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*InternalTransferTransaction, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchInternalTransferTransaction(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskInternalTransferTransaction patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskInternalTransferTransaction(ctx context.Context, patchee *InternalTransferTransaction, patcher *InternalTransferTransaction, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*InternalTransferTransaction, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"TaskID" {
			patchee.TaskID = patcher.TaskID
			continue
		}
		if f == prefix+"TransactionID" {
			patchee.TransactionID = patcher.TransactionID
			continue
		}
		if f == prefix+"TransactionServiceID" {
			patchee.TransactionServiceID = patcher.TransactionServiceID
			continue
		}
		if f == prefix+"Data" {
			patchee.Data = patcher.Data
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

// DefaultListInternalTransferTransaction executes a gorm list call
func DefaultListInternalTransferTransaction(ctx context.Context, db *gorm.DB) ([]*InternalTransferTransaction, error) {
	in := InternalTransferTransaction{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &InternalTransferTransactionORM{}, &InternalTransferTransaction{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []InternalTransferTransactionORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferTransactionORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*InternalTransferTransaction{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type InternalTransferTransactionORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferTransactionORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]InternalTransferTransactionORM) error
}

// DefaultCreateInternalTransferSingleTemplate executes a basic gorm create call
func DefaultCreateInternalTransferSingleTemplate(ctx context.Context, in *InternalTransferSingleTemplate, db *gorm.DB) (*InternalTransferSingleTemplate, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type InternalTransferSingleTemplateORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadInternalTransferSingleTemplate(ctx context.Context, in *InternalTransferSingleTemplate, db *gorm.DB) (*InternalTransferSingleTemplate, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &InternalTransferSingleTemplateORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := InternalTransferSingleTemplateORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(InternalTransferSingleTemplateORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type InternalTransferSingleTemplateORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteInternalTransferSingleTemplate(ctx context.Context, in *InternalTransferSingleTemplate, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&InternalTransferSingleTemplateORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type InternalTransferSingleTemplateORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteInternalTransferSingleTemplateSet(ctx context.Context, in []*InternalTransferSingleTemplate, db *gorm.DB) error {
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
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&InternalTransferSingleTemplateORM{})).(InternalTransferSingleTemplateORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&InternalTransferSingleTemplateORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&InternalTransferSingleTemplateORM{})).(InternalTransferSingleTemplateORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type InternalTransferSingleTemplateORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*InternalTransferSingleTemplate, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*InternalTransferSingleTemplate, *gorm.DB) error
}

// DefaultStrictUpdateInternalTransferSingleTemplate clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateInternalTransferSingleTemplate(ctx context.Context, in *InternalTransferSingleTemplate, db *gorm.DB) (*InternalTransferSingleTemplate, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateInternalTransferSingleTemplate")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &InternalTransferSingleTemplateORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithAfterStrictUpdateSave); ok {
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

type InternalTransferSingleTemplateORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchInternalTransferSingleTemplate executes a basic gorm update call with patch behavior
func DefaultPatchInternalTransferSingleTemplate(ctx context.Context, in *InternalTransferSingleTemplate, updateMask *field_mask.FieldMask, db *gorm.DB) (*InternalTransferSingleTemplate, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj InternalTransferSingleTemplate
	var err error
	if hook, ok := interface{}(&pbObj).(InternalTransferSingleTemplateWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadInternalTransferSingleTemplate(ctx, &InternalTransferSingleTemplate{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(InternalTransferSingleTemplateWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskInternalTransferSingleTemplate(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(InternalTransferSingleTemplateWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateInternalTransferSingleTemplate(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(InternalTransferSingleTemplateWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type InternalTransferSingleTemplateWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *InternalTransferSingleTemplate, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *InternalTransferSingleTemplate, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *InternalTransferSingleTemplate, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *InternalTransferSingleTemplate, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetInternalTransferSingleTemplate executes a bulk gorm update call with patch behavior
func DefaultPatchSetInternalTransferSingleTemplate(ctx context.Context, objects []*InternalTransferSingleTemplate, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*InternalTransferSingleTemplate, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*InternalTransferSingleTemplate, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchInternalTransferSingleTemplate(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskInternalTransferSingleTemplate patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskInternalTransferSingleTemplate(ctx context.Context, patchee *InternalTransferSingleTemplate, patcher *InternalTransferSingleTemplate, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*InternalTransferSingleTemplate, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Data" {
			patchee.Data = patcher.Data
			continue
		}
		if f == prefix+"TemplateName" {
			patchee.TemplateName = patcher.TemplateName
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

// DefaultListInternalTransferSingleTemplate executes a gorm list call
func DefaultListInternalTransferSingleTemplate(ctx context.Context, db *gorm.DB) ([]*InternalTransferSingleTemplate, error) {
	in := InternalTransferSingleTemplate{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &InternalTransferSingleTemplateORM{}, &InternalTransferSingleTemplate{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []InternalTransferSingleTemplateORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(InternalTransferSingleTemplateORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*InternalTransferSingleTemplate{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type InternalTransferSingleTemplateORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type InternalTransferSingleTemplateORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]InternalTransferSingleTemplateORM) error
}
