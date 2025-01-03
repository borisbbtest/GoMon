SELECT id,
    login,
    firstname,
    lastname,
    password,
    source,
    created_at
FROM idm.users
WHERE login = $1;