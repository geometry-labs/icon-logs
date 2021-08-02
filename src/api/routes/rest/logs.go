package rest

import (
	"context"
	"encoding/json"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/geometry-labs/icon-logs/config"
	"github.com/geometry-labs/icon-logs/crud"
)

type LogsQuery struct {
	Limit int64 `query:"limit"`
	Skip  int64 `query:"skip"`

	Hash string `query:"hash"`
	From string `query:"from"`
	To   string `query:"to"`
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
// @Param hash query string false "find by hash"
// @Param from query string false "find by from address"
// @Param to query string false "find by to address"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logs, err := crud.GetLogModel().Select(
		ctx,
		params.Limit,
		params.Skip,
		params.Hash,
		params.From,
		params.To,
	)
  if err != nil {
    c.Status(500)
    zap.S().Errorf("ERROR: %s", err.Error())
    return c.SendString(`{"error": "unable to query logs"}`)
  }

	if len(logs) == 0 {
		// No Content
		c.Status(204)
	}

	body, _ := json.Marshal(&logs)
	return c.SendString(string(body))
}
