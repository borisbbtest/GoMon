SELECT id,
    login,
    config,
    created,
    duration
FROM idm.sessions
WHERE login = $1
    and id = $2;