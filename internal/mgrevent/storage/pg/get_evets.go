package storagepg

import (
	"context"
	"time"

	"github.com/borisbbtest/GoMon/internal/mgrevent/models"
	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

// GetEvent - Получаем один из одну запись по uuid. В будущем  будет релизов сервис виде воркера который будет возвращать данные из БД
func (hook *StoreDBinPostgreSQL) GetEvent(ctx context.Context, eve *mgrevent.Event) (err error, res *models.Events) {
	query, err := SQLFileInit.ReadFile("migrations/core/select_event.sql")
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
	var k models.PGEvent
	err = conn.PostgresPool.QueryRow(ctx, string(query), eve.Uuid).Scan(
		&k.Id,
		&k.Title,
		&k.Description,
		&k.Source,
		&k.Status,
		&k.Created,
		&k.Update,
		&k.Key,
		&k.KeyClose,
		&k.Assigned,
		&k.Severity,
		&k.AutoRunner,
		&k.RelarionCi,
	)

	if err != nil {
		utils.Log.Error().Msgf(err.Error())
	}
	// Создаем ссылочный объект для экономии памяти
	res = &models.Events{EventsPG: append(make([]*models.PGEvent, 0), &k)}

	return
}

// GetEventDuration - получает список событий в промежуток времени
func (hook *StoreDBinPostgreSQL) GetEventDuration(ctx context.Context, startTime time.Time, endTime time.Time) (err error, res *models.Events) {
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

	orders := []*models.PGEvent{}
	for rows.Next() {
		k := models.PGEvent{}
		err = rows.Scan(
			&k.Id,
			&k.Title,
			&k.Description,
			&k.Source,
			&k.Status,
			&k.Created,
			&k.Update,
			&k.Key,
			&k.KeyClose,
			&k.Assigned,
			&k.Severity,
			&k.AutoRunner,
			&k.RelarionCi,
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
	res = &models.Events{EventsPG: orders}

	return
}
