SELECT "name",
    description,
    "update",
    created,
    created_by,
    "type"
FROM cmdb.cis
WHERE "name" = $1;