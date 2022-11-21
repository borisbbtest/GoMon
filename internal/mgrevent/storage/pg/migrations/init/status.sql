-- object: events.status | type: TABLE --
-- DROP TABLE IF EXISTS events.status CASCADE;

-- Prepended SQL commands --

-- ddl-end --

CREATE TABLE events.status (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	name text NOT NULL,
	code numeric NOT NULL,
	CONSTRAINT status_pkey PRIMARY KEY (id)
);
-- ddl-end --

-- ddl-end --

