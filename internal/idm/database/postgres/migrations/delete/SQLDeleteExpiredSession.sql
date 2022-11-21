DELETE FROM idm.sessions
WHERE duration < current_timestamp - (5 || ' minutes')::interval;