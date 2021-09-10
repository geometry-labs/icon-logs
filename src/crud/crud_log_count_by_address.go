package crud

import (
	"errors"
	"strings"
	"sync"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/geometry-labs/icon-logs/models"
)

// LogCountByAddressModel - type for logCountByAddress table model
type LogCountByAddressModel struct {
	db        *gorm.DB
	model     *models.LogCountByAddress
	modelORM  *models.LogCountByAddressORM
	WriteChan chan *models.LogCountByAddress
}

var logCountByAddressModel *LogCountByAddressModel
var logCountByAddressModelOnce sync.Once

// GetLogCountByAddressModel - create and/or return the logCountByAddresss table model
func GetLogCountByAddressModel() *LogCountByAddressModel {
	logCountByAddressModelOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		logCountByAddressModel = &LogCountByAddressModel{
			db:        dbConn,
			model:     &models.LogCountByAddress{},
			WriteChan: make(chan *models.LogCountByAddress, 1),
		}

		err := logCountByAddressModel.Migrate()
		if err != nil {
			zap.S().Fatal("LogCountByAddressModel: Unable migrate postgres table: ", err.Error())
		}
	})

	return logCountByAddressModel
}

// Migrate - migrate logCountByAddresss table
func (m *LogCountByAddressModel) Migrate() error {
	// Only using LogCountByAddressRawORM (ORM version of the proto generated struct) to create the TABLE
	zap.S().Info("Migrating LogCountByAddressModel....")
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

// Insert - Insert logCountByAddress into table
func (m *LogCountByAddressModel) Insert(logCountByAddress *models.LogCountByAddress) error {

	err := backoff.Retry(func() error {
		query := m.db.Create(logCountByAddress)
		if query.Error != nil && !strings.Contains(query.Error.Error(), "duplicate key value violates unique constraint") {
			zap.S().Warn("POSTGRES Insert Error : ", query.Error.Error())
			return query.Error
		}

		return nil
	}, backoff.NewExponentialBackOff())

	return err
}

func (m *LogCountByAddressModel) Update(logCountByAddress *models.LogCountByAddress) error {

	db := m.db
	//computeCount := false

	// Set table
	db = db.Model(&models.LogCountByAddress{})

	// Transaction Hash
	db = db.Where("transaction_hash = ?", logCountByAddress.TransactionHash)

	// Log Index
	db = db.Where("log_index = ?", logCountByAddress.LogIndex)

	// Update
	db = db.Save(logCountByAddress)

	return db.Error
}

func (m *LogCountByAddressModel) SelectLargestCountByAddress(address string) (uint64, error) {

	db := m.db
	//computeCount := false

	// Set table
	db = db.Model(&models.LogCountByAddress{})

	// Address
	db = db.Where("address = ?", address)

	// Get max count
	count := uint64(0)
	row := db.Select("max(count)").Row()
	row.Scan(&count)

	return count, db.Error
}

func (m *LogCountByAddressModel) SelectOne(transactionHash string, logIndex uint64) (*models.LogCountByAddress, error) {

	db := m.db
	//computeCount := false

	// Set table
	db = db.Model(&models.LogCountByAddress{})

	// Transaction Hash
	db = db.Where("transaction_hash = ?", transactionHash)

	// Log Index
	db = db.Where("log_index = ?", logIndex)

	// Select
	logCountByAddress := &models.LogCountByAddress{}
	db = db.First(logCountByAddress)

	return logCountByAddress, db.Error
}

// StartLogCountByAddressLoader starts loader
func StartLogCountByAddressLoader() {
	go func() {

		for {
			// Read transaction
			newLogCountByAddress := <-GetLogCountByAddressModel().WriteChan

			// Read current state
			curCount, err := GetLogCountByAddressModel().SelectLargestCountByAddress(
				newLogCountByAddress.Address,
			)
			if err != nil {
				zap.S().Fatal(err.Error())
			}

			// Add count
			newLogCountByAddress.Count += curCount

			// Insert
			_, err = GetLogCountByAddressModel().SelectOne(
				newLogCountByAddress.TransactionHash,
				newLogCountByAddress.LogIndex,
			)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Insert
				err = GetLogCountByAddressModel().Insert(newLogCountByAddress)
				if err != nil {
					zap.S().Fatal(err.Error())
				}
			} else if err != nil {
				// Error
				zap.S().Fatal(err.Error())
			}

			zap.S().Debugf("Loader LogCountByAddress: Loaded in postgres table LogCountByAddresss, Address: %s", newLogCountByAddress.Address)
		}
	}()
}
