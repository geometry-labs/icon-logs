// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kafka_job.proto

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

type KafkaJobORM struct {
	JobId       string `gorm:"primary_key"`
	Partition   uint64 `gorm:"primary_key"`
	StopOffset  uint64
	Topic       string `gorm:"primary_key"`
	WorkerGroup string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (KafkaJobORM) TableName() string {
	return "kafka_jobs"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *KafkaJob) ToORM(ctx context.Context) (KafkaJobORM, error) {
	to := KafkaJobORM{}
	var err error
	if prehook, ok := interface{}(m).(KafkaJobWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.JobId = m.JobId
	to.WorkerGroup = m.WorkerGroup
	to.Topic = m.Topic
	to.Partition = m.Partition
	to.StopOffset = m.StopOffset
	if posthook, ok := interface{}(m).(KafkaJobWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *KafkaJobORM) ToPB(ctx context.Context) (KafkaJob, error) {
	to := KafkaJob{}
	var err error
	if prehook, ok := interface{}(m).(KafkaJobWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.JobId = m.JobId
	to.WorkerGroup = m.WorkerGroup
	to.Topic = m.Topic
	to.Partition = m.Partition
	to.StopOffset = m.StopOffset
	if posthook, ok := interface{}(m).(KafkaJobWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type KafkaJob the arg will be the target, the caller the one being converted from

// KafkaJobBeforeToORM called before default ToORM code
type KafkaJobWithBeforeToORM interface {
	BeforeToORM(context.Context, *KafkaJobORM) error
}

// KafkaJobAfterToORM called after default ToORM code
type KafkaJobWithAfterToORM interface {
	AfterToORM(context.Context, *KafkaJobORM) error
}

// KafkaJobBeforeToPB called before default ToPB code
type KafkaJobWithBeforeToPB interface {
	BeforeToPB(context.Context, *KafkaJob) error
}

// KafkaJobAfterToPB called after default ToPB code
type KafkaJobWithAfterToPB interface {
	AfterToPB(context.Context, *KafkaJob) error
}

// DefaultCreateKafkaJob executes a basic gorm create call
func DefaultCreateKafkaJob(ctx context.Context, in *KafkaJob, db *gorm1.DB) (*KafkaJob, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type KafkaJobORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type KafkaJobORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskKafkaJob patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskKafkaJob(ctx context.Context, patchee *KafkaJob, patcher *KafkaJob, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*KafkaJob, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"JobId" {
			patchee.JobId = patcher.JobId
			continue
		}
		if f == prefix+"WorkerGroup" {
			patchee.WorkerGroup = patcher.WorkerGroup
			continue
		}
		if f == prefix+"Topic" {
			patchee.Topic = patcher.Topic
			continue
		}
		if f == prefix+"Partition" {
			patchee.Partition = patcher.Partition
			continue
		}
		if f == prefix+"StopOffset" {
			patchee.StopOffset = patcher.StopOffset
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListKafkaJob executes a gorm list call
func DefaultListKafkaJob(ctx context.Context, db *gorm1.DB) ([]*KafkaJob, error) {
	in := KafkaJob{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &KafkaJobORM{}, &KafkaJob{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("job_id")
	ormResponse := []KafkaJobORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(KafkaJobORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*KafkaJob{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type KafkaJobORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type KafkaJobORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type KafkaJobORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]KafkaJobORM) error
}
