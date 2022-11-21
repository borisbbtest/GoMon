DELETE FROM idm.sessions
WHERE login = $1
    and id = $2;