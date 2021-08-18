package rest

import (
	"encoding/json"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/geometry-labs/icon-logs/config"
	"github.com/geometry-labs/icon-logs/crud"
	"github.com/geometry-labs/icon-logs/models"
)

type LogsQuery struct {
	Limit int `query:"limit"`
	Skip  int `query:"skip"`

	TransactionHash string `query:"transaction_hash"`
	ScoreAddress    string `query:"score_address"`
}

func LogsAddHandlers(app *fiber.App) {

	prefix := config.Config.RestPrefix + "/logs"

	app.Get(prefix+"/", handlerGetLogs)
}

// Logs
// @Summary Get Logs
// @Description get historical logs
// @Tags Logs
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param transaction_hash query string false "find by transaction hash"
// @Param score_address query string false "find by score address"
// @Router /api/v1/logs [get]
// @Success 200 {object} []models.Log
// @Failure 422 {object} map[string]interface{}
func handlerGetLogs(c *fiber.Ctx) error {
	params := new(LogsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Logs Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 1
	}

	// Get Logs
	logs, err := crud.GetLogModel().Select(
		params.Limit,
		params.Skip,
		params.TransactionHash,
		params.ScoreAddress,
	)
	if err != nil {
		zap.S().Warnf("Logs CRUD ERROR: %s", err.Error())
		c.Status(500)
		return c.SendString(`{"error": "could not retrieve logs"}`)
	}

	if len(logs) == 0 {
		// No Content
		c.Status(204)
	}

	// Set X-TOTAL-COUNT
	counter, err := crud.GetLogCountModel().Select()
	if err != nil {
		counter = models.LogCount{
			Count: 0,
			Id:    0,
		}
		zap.S().Warn("Could not retrieve log count: ", err.Error())
	}
	c.Append("X-TOTAL-COUNT", strconv.FormatUint(counter.Count, 10))

	body, _ := json.Marshal(&logs)
	return c.SendString(string(body))
}
