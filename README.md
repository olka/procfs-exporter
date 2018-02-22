[![Build Status](https://travis-ci.org/olka/procfs-exporter.svg?branch=master)](https://travis-ci.org/olka/procfs-exporter)

Basic procfs exporter (/proc/stat /proc/loadavg) exposed on 9100 port

CPU usage calculates in similar to [htop](https://github.com/hishamhm/htop) way 

Output gzipped with [NYTimes' gziphandler](https://github.com/NYTimes/gziphandler)

Test coverage: ~60% of statements (check Travis for more details)

Performance test during 3 min on dual-core G3460@3.50GHz (load average: 1,03, 0,49, 0,28) 

|  Req/s |	Min |	50th pct |	90th pct |	99th pct |	Max  |	Mean |	Std Dev |
|--------|------|------------|-----------|-----------|-------|-------|----------|
|   1491 | 	2 	|      10 	 |      70 	 |     297 	 | 	3698 |	 31  |	 88     |
