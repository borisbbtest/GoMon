select id,name,value,loadtime,source_from_systems,related_ci,source_time
from   metrics.item_metrics  where source_from_systems >= $1 and source_from_systems <= $2 order by update;