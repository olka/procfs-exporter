[![Build Status](https://travis-ci.org/olka/proc-metrics-exporter.svg?branch=master)](https://travis-ci.org/olka/proc-metrics-exporter)

Basic procfs exporter (/proc/stat /proc/loadavg) exposed on 9100 port

CPU usage calculates in similar to [htop](https://github.com/hishamhm/htop) way 

Output gzipped with [NYTimes' gziphandler](https://github.com/NYTimes/gziphandler)

Test coverage: 72.5% of statements