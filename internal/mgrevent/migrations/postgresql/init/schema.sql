-- object: events | type: SCHEMA --
-- DROP SCHEMA IF EXISTS events CASCADE;
CREATE SCHEMA events;
-- ddl-end --
ALTER SCHEMA events OWNER TO postgres;
-- ddl-end --
CREATE EXTENSION "uuid-ossp";