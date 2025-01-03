select id,title,description,source,status,created,update,key,key_close,assigned,severity,auto_runner,relation_ci
from  events.events where update >= $1 and update <= $2 order by update;