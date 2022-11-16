package storage

import (
	"github.com/borisbbtest/GoMon/internal/mgrevent/configs"
	"github.com/borisbbtest/GoMon/internal/mgrevent/database/postgres"
	"github.com/borisbbtest/GoMon/internal/mgrevent/utils"
)

type StoreDBinPostgreSQL struct {
	pgp     postgres.Plugin
	connStr string
}

func NewPostgreSQLStorage(connStr *configs.MainConfig) (res *StoreDBinPostgreSQL, err error) {
	res = &StoreDBinPostgreSQL{}
	res.connStr = connStr.DatabaseURI
	res.pgp.Start(connStr.DatabaseURI)
	conn, err := res.pgp.NewConn()
	if err != nil {
		utils.Log.Error().Msgf("selectOrdersHandler - c: ", err)
		return nil, err
	}
	query := `
				    CREATE TABLE IF NOT EXISTS public."Users"
					(
						"Id" "serial",
						"Login" "text" UNIQUE NOT NULL,
						"Password" "text" NOT NULL,
						CONSTRAINT "Id" PRIMARY KEY ("Id")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Users" 	OWNER to postgres;

					COMMENT ON TABLE public."Users"  IS 'This table was created for storage data about persons users in within inside project';

					CREATE TABLE IF NOT EXISTS public."Orders"
					(
						"Number" "text" NOT NULL,
						"Status" "text" NOT NULL,
						"Person" "numeric" NOT NULL,
						"Accrual" "numeric",
						"Uploaded_at" "timestamptz" NOT NULL,
						CONSTRAINT "Number" PRIMARY KEY ("Number")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Orders" 	OWNER to postgres;

					COMMENT ON TABLE public."Orders"  IS 'This table was created for storage data about orders';

					CREATE TABLE IF NOT EXISTS public."Wallet"
					(
						"Order" "text" NOT NULL,
						"Person" "numeric" NOT NULL,
						"Sum" "numeric",
						"Uploaded_at" "timestamptz" NOT NULL
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Wallet" 	OWNER to postgres;

					COMMENT ON TABLE public."Wallet"  IS 'This table was created for storage data about Wallet';

					CREATE TABLE IF NOT EXISTS public."Balance"
					(
						"Person"           "numeric",
						"Withdrawn"        "numeric",
						"CurrentAccrual"   "numeric",
					    CONSTRAINT "Person" PRIMARY KEY ("Person")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Balance" 	OWNER to postgres;

					COMMENT ON TABLE public."Balance"  IS 'This table was created for storage data about Balance';

			`

	if _, err := conn.PostgresPool.Exec(connStr.Ctx, query); err != nil {
		return nil, err
	}
	return
}

func (hook *StoreDBinPostgreSQL) Close() {
	hook.pgp.Stop()
}
