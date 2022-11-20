-- object: events.events | type: TABLE --
-- DROP TABLE IF EXISTS events.events CASCADE;

-- Prepended SQL commands --


-- ddl-end --

CREATE TABLE metrics.item_metrics (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	name text NOT NULL,
	value bytea NOT NULL,
	loadtime timestamp NOT NULL,
	source_from_systems text NOT NULL,
	related_ci text NOT NULL,
	source_time timestamptz NOT NULL,
	type numeric NOT NULL,
	CONSTRAINT item_metrics_pk PRIMARY KEY (id,source_time)
);

-- ddl-end --
-- ddl-end --

