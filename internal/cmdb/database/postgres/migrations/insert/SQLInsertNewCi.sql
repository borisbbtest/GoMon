insert into cmdb.cis (
        "name",
        description,
        "update",
        created,
        created_by,
        "type"
    )
values ($1, $2, $3, $4, $5, $6);