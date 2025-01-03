select id,name,value,loadtime,source_from_systems,related_ci,source_time,type
from   metrics.item_metrics  where loadtime >= $1 and loadtime <= $2 order by loadtime;