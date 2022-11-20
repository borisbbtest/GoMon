package storagepg

import (
	"context"
	"time"

	"github.com/borisbbtest/GoMon/internal/metrics/models"
	"github.com/borisbbtest/GoMon/internal/metrics/utils"
)

// Получаем один из одну запись по uuid. В будущем  будет релизов сервис виде воркера который будет возвращать данные из БД

func (hook *StoreDBinPostgreSQL) GetMetricsDuration(ctx context.Context, startTime time.Time, endTime time.Time) (err error, res *models.Metrics) {
	query, err := SQLFileInit.ReadFile("migrations/core/select_events_duration.sql")
	if err != nil {
		utils.Log.Error().Msgf("Get Events", err)
		return err, nil
	}
	conn, err := hook.pgp.NewConn()
	if err != nil {
		utils.Log.Error().Msgf("Get Events", err)
		return err, nil
	}
	// получаем данные из БД

	rows, err := conn.PostgresPool.Query(ctx, string(query), startTime, endTime)
	if err != nil {
		utils.Log.Error().Msgf(err.Error())
	}

	orders := []*models.PGMetric{}
	for rows.Next() {
		k := models.PGMetric{}
		err = rows.Scan(
			&k.Id,
			&k.Name,
			&k.Value,
			&k.LoadTime,
			&k.Source,
			&k.RelarionCi,
			&k.SourceTime,
		)

		if err != nil {
			utils.Log.Error().Msgf(err.Error())
			return err, nil
		}
		orders = append(orders, &k)
	}
	if err != nil {
		utils.Log.Error().Msgf(err.Error())
		return err, nil
	}
	//utils.Log.Debug().Msgf("%s", orders)
	// Создаем ссылочный объект для экономии памяти
	res = &models.Metrics{EventsPG: orders}

	return
}
