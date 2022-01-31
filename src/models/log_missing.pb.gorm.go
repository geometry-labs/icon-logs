// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log_missing.proto

package models

import (
	context "context"
	fmt "fmt"
	
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"

	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type LogMissingORM struct {
	BlockNumber     uint64
	TransactionHash string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (LogMissingORM) TableName() string {
	return "log_missings"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *LogMissing) ToORM(ctx context.Context) (LogMissingORM, error) {
	to := LogMissingORM{}
	var err error
	if prehook, ok := interface{}(m).(LogMissingWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.BlockNumber = m.BlockNumber
	if posthook, ok := interface{}(m).(LogMissingWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *LogMissingORM) ToPB(ctx context.Context) (LogMissing, error) {
	to := LogMissing{}
	var err error
	if prehook, ok := interface{}(m).(LogMissingWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.BlockNumber = m.BlockNumber
	if posthook, ok := interface{}(m).(LogMissingWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type LogMissing the arg will be the target, the caller the one being converted from

// LogMissingBeforeToORM called before default ToORM code
type LogMissingWithBeforeToORM interface {
	BeforeToORM(context.Context, *LogMissingORM) error
}

// LogMissingAfterToORM called after default ToORM code
type LogMissingWithAfterToORM interface {
	AfterToORM(context.Context, *LogMissingORM) error
}

// LogMissingBeforeToPB called before default ToPB code
type LogMissingWithBeforeToPB interface {
	BeforeToPB(context.Context, *LogMissing) error
}

// LogMissingAfterToPB called after default ToPB code
type LogMissingWithAfterToPB interface {
	AfterToPB(context.Context, *LogMissing) error
}

// DefaultCreateLogMissing executes a basic gorm create call
func DefaultCreateLogMissing(ctx context.Context, in *LogMissing, db *gorm1.DB) (*LogMissing, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogMissingORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogMissingORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type LogMissingORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogMissingORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskLogMissing patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskLogMissing(ctx context.Context, patchee *LogMissing, patcher *LogMissing, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*LogMissing, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"TransactionHash" {
			patchee.TransactionHash = patcher.TransactionHash
			continue
		}
		if f == prefix+"BlockNumber" {
			patchee.BlockNumber = patcher.BlockNumber
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListLogMissing executes a gorm list call
func DefaultListLogMissing(ctx context.Context, db *gorm1.DB) ([]*LogMissing, error) {
	in := LogMissing{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogMissingORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &LogMissingORM{}, &LogMissing{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogMissingORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("transaction_hash")
	ormResponse := []LogMissingORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogMissingORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*LogMissing{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type LogMissingORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogMissingORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogMissingORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]LogMissingORM) error
}
