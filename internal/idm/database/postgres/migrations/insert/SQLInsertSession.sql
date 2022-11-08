insert into idm.sessions (
        id,
        login,
        created,
        duration
    )
values ($1, $2, $3, $4);