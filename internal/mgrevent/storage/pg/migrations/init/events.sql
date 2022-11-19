-- object: events.events | type: TABLE --
-- DROP TABLE IF EXISTS events.events CASCADE;

-- Prepended SQL commands --


-- ddl-end --

CREATE TABLE IF NOT EXISTS events.events (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	title text,
	description text,
	source text NOT NULL,
	status numeric NOT NULL,
	created timestamp NOT NULL,
	update timestamp NOT NULL,
	key text,
	key_close text,
	assigned text[],
	severity numeric NOT NULL,
	auto_runner text,
	relation_ci text[],
	CONSTRAINT events_pkey PRIMARY KEY (id)
);
-- ddl-end --

-- ddl-end --

