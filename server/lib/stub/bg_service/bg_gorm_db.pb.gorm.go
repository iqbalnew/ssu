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

type MappingORM struct {
	BeneficiaryID uint64     `gorm:"not null"`
	CompanyID     uint64     `gorm:"not null"`
	CreatedAt     *time.Time `gorm:"not null"`
	CreatedByID   uint64     `gorm:"not null"`
	Id            uint64     `gorm:"primary_key;not null"`
	IsMapped      bool       `gorm:"default:false;not null"`
	ThirdPartyID  uint64     `gorm:"not null"`
	UpdatedAt     *time.Time `gorm:"not null"`
	UpdatedByID   uint64     `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (MappingORM) TableName() string {
	return "mappings"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Mapping) ToORM(ctx context.Context) (MappingORM, error) {
	to := MappingORM{}
	var err error
	if prehook, ok := interface{}(m).(MappingWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.CompanyID = m.CompanyID
	to.ThirdPartyID = m.ThirdPartyID
	to.BeneficiaryID = m.BeneficiaryID
	to.IsMapped = m.IsMapped
	to.CreatedByID = m.CreatedByID
	to.UpdatedByID = m.UpdatedByID
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(MappingWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *MappingORM) ToPB(ctx context.Context) (Mapping, error) {
	to := Mapping{}
	var err error
	if prehook, ok := interface{}(m).(MappingWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.CompanyID = m.CompanyID
	to.ThirdPartyID = m.ThirdPartyID
	to.BeneficiaryID = m.BeneficiaryID
	to.IsMapped = m.IsMapped
	to.CreatedByID = m.CreatedByID
	to.UpdatedByID = m.UpdatedByID
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(MappingWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Mapping the arg will be the target, the caller the one being converted from

// MappingBeforeToORM called before default ToORM code
type MappingWithBeforeToORM interface {
	BeforeToORM(context.Context, *MappingORM) error
}

// MappingAfterToORM called after default ToORM code
type MappingWithAfterToORM interface {
	AfterToORM(context.Context, *MappingORM) error
}

// MappingBeforeToPB called before default ToPB code
type MappingWithBeforeToPB interface {
	BeforeToPB(context.Context, *Mapping) error
}

// MappingAfterToPB called after default ToPB code
type MappingWithAfterToPB interface {
	AfterToPB(context.Context, *Mapping) error
}

type CurrencyORM struct {
	Code      string     `gorm:"not null"`
	CreatedAt *time.Time `gorm:"not null"`
	Id        uint64     `gorm:"primary_key;not null"`
	NameEN    string     `gorm:"not null"`
	NameID    string     `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
}

// TableName overrides the default tablename generated by GORM
func (CurrencyORM) TableName() string {
	return "currencies"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Currency) ToORM(ctx context.Context) (CurrencyORM, error) {
	to := CurrencyORM{}
	var err error
	if prehook, ok := interface{}(m).(CurrencyWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Code = m.Code
	to.NameEN = m.NameEN
	to.NameID = m.NameID
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if posthook, ok := interface{}(m).(CurrencyWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *CurrencyORM) ToPB(ctx context.Context) (Currency, error) {
	to := Currency{}
	var err error
	if prehook, ok := interface{}(m).(CurrencyWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Code = m.Code
	to.NameEN = m.NameEN
	to.NameID = m.NameID
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if posthook, ok := interface{}(m).(CurrencyWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Currency the arg will be the target, the caller the one being converted from

// CurrencyBeforeToORM called before default ToORM code
type CurrencyWithBeforeToORM interface {
	BeforeToORM(context.Context, *CurrencyORM) error
}

// CurrencyAfterToORM called after default ToORM code
type CurrencyWithAfterToORM interface {
	AfterToORM(context.Context, *CurrencyORM) error
}

// CurrencyBeforeToPB called before default ToPB code
type CurrencyWithBeforeToPB interface {
	BeforeToPB(context.Context, *Currency) error
}

// CurrencyAfterToPB called after default ToPB code
type CurrencyWithAfterToPB interface {
	AfterToPB(context.Context, *Currency) error
}

// DefaultCreateMapping executes a basic gorm create call
func DefaultCreateMapping(ctx context.Context, in *Mapping, db *gorm.DB) (*Mapping, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type MappingORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadMapping(ctx context.Context, in *Mapping, db *gorm.DB) (*Mapping, error) {
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
	if hook, ok := interface{}(&ormObj).(MappingORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &MappingORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := MappingORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(MappingORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type MappingORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteMapping(ctx context.Context, in *Mapping, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(MappingORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&MappingORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type MappingORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteMappingSet(ctx context.Context, in []*Mapping, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&MappingORM{})).(MappingORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&MappingORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&MappingORM{})).(MappingORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type MappingORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Mapping, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Mapping, *gorm.DB) error
}

// DefaultStrictUpdateMapping clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateMapping(ctx context.Context, in *Mapping, db *gorm.DB) (*Mapping, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateMapping")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &MappingORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(MappingORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithAfterStrictUpdateSave); ok {
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

type MappingORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchMapping executes a basic gorm update call with patch behavior
func DefaultPatchMapping(ctx context.Context, in *Mapping, updateMask *field_mask.FieldMask, db *gorm.DB) (*Mapping, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Mapping
	var err error
	if hook, ok := interface{}(&pbObj).(MappingWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadMapping(ctx, &Mapping{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(MappingWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskMapping(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(MappingWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateMapping(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(MappingWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type MappingWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Mapping, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type MappingWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Mapping, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type MappingWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Mapping, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type MappingWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Mapping, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetMapping executes a bulk gorm update call with patch behavior
func DefaultPatchSetMapping(ctx context.Context, objects []*Mapping, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Mapping, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Mapping, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchMapping(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskMapping patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskMapping(ctx context.Context, patchee *Mapping, patcher *Mapping, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Mapping, error) {
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
		if f == prefix+"CompanyID" {
			patchee.CompanyID = patcher.CompanyID
			continue
		}
		if f == prefix+"ThirdPartyID" {
			patchee.ThirdPartyID = patcher.ThirdPartyID
			continue
		}
		if f == prefix+"BeneficiaryID" {
			patchee.BeneficiaryID = patcher.BeneficiaryID
			continue
		}
		if f == prefix+"IsMapped" {
			patchee.IsMapped = patcher.IsMapped
			continue
		}
		if f == prefix+"CreatedByID" {
			patchee.CreatedByID = patcher.CreatedByID
			continue
		}
		if f == prefix+"UpdatedByID" {
			patchee.UpdatedByID = patcher.UpdatedByID
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

// DefaultListMapping executes a gorm list call
func DefaultListMapping(ctx context.Context, db *gorm.DB) ([]*Mapping, error) {
	in := Mapping{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &MappingORM{}, &Mapping{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []MappingORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(MappingORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Mapping{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type MappingORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type MappingORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]MappingORM) error
}

// DefaultCreateCurrency executes a basic gorm create call
func DefaultCreateCurrency(ctx context.Context, in *Currency, db *gorm.DB) (*Currency, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type CurrencyORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadCurrency(ctx context.Context, in *Currency, db *gorm.DB) (*Currency, error) {
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
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &CurrencyORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := CurrencyORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(CurrencyORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type CurrencyORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteCurrency(ctx context.Context, in *Currency, db *gorm.DB) error {
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
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&CurrencyORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type CurrencyORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteCurrencySet(ctx context.Context, in []*Currency, db *gorm.DB) error {
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
	if hook, ok := (interface{}(&CurrencyORM{})).(CurrencyORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&CurrencyORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&CurrencyORM{})).(CurrencyORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type CurrencyORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Currency, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Currency, *gorm.DB) error
}

// DefaultStrictUpdateCurrency clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateCurrency(ctx context.Context, in *Currency, db *gorm.DB) (*Currency, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateCurrency")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &CurrencyORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithAfterStrictUpdateSave); ok {
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

type CurrencyORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchCurrency executes a basic gorm update call with patch behavior
func DefaultPatchCurrency(ctx context.Context, in *Currency, updateMask *field_mask.FieldMask, db *gorm.DB) (*Currency, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Currency
	var err error
	if hook, ok := interface{}(&pbObj).(CurrencyWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadCurrency(ctx, &Currency{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(CurrencyWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskCurrency(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(CurrencyWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateCurrency(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(CurrencyWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type CurrencyWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Currency, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CurrencyWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Currency, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CurrencyWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Currency, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CurrencyWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Currency, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetCurrency executes a bulk gorm update call with patch behavior
func DefaultPatchSetCurrency(ctx context.Context, objects []*Currency, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Currency, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Currency, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchCurrency(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskCurrency patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskCurrency(ctx context.Context, patchee *Currency, patcher *Currency, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Currency, error) {
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
		if f == prefix+"Code" {
			patchee.Code = patcher.Code
			continue
		}
		if f == prefix+"NameEN" {
			patchee.NameEN = patcher.NameEN
			continue
		}
		if f == prefix+"NameID" {
			patchee.NameID = patcher.NameID
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

// DefaultListCurrency executes a gorm list call
func DefaultListCurrency(ctx context.Context, db *gorm.DB) ([]*Currency, error) {
	in := Currency{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &CurrencyORM{}, &Currency{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []CurrencyORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CurrencyORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Currency{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type CurrencyORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CurrencyORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]CurrencyORM) error
}
