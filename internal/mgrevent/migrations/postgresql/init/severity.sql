-- object: events.severity | type: TABLE --
-- DROP TABLE IF EXISTS events.severity CASCADE;

-- Prepended SQL commands --

-- ddl-end --

CREATE TABLE events.severity (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	name text NOT NULL,
	code numeric NOT NULL,
	CONSTRAINT severity_pkey PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE events.severity OWNER TO postgres;
-- ddl-end --

