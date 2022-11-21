CREATE SCHEMA IF NOT EXISTS cmdb;
CREATE TABLE IF NOT EXISTS cmdb.cis (
    "name" text NOT NULL,
    description text,
    "update" timestamptz,
    created timestamptz,
    created_by text,
    "type" text,
    CONSTRAINT cis_pkey PRIMARY KEY ("name")
);