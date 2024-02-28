## **PROMETHEUS-DISPATCHER-EXPORTER**
#

Export dispatcher metrics (time series) & statistics (long time data).  
For now, only [Thinkbox Deadline](https://aws.amazon.com/fr/thinkbox-deadline/) is supported.  

Metrics are exposed to [Prometheus](https://prometheus.io/).  
Statistics are pushed in a [PostGreSQL](https://www.postgresql.org/) instance.  
#
## Grafana dashboards
Dashboards can be found in the `dashboards` directory.  

1. Metrics :
![img](/.assets/metrics.png)
2. Stats :
![img](/.assets/stats.png)

#
## Usage
### Environment variables
* `PROMETHEUS_ENABLED`: enable prometheus metrics (default true)
* `PROMETHEUS_PORT`: prometheus listen port (default 9101)
* `POSTGRESQL_ENABLED`: enable postgresql statistic (default false)
* `POSTGRESQL_HOST`: postgresql host address
* `POSTGRESQL_USER`: postgresql user
* `POSTGRESQL_PASSWORD`: postgresql password
* `POSTGRESQL_PORT`: postgresql port (default 5432)
* `POSTGRESQL_DB`: postgresql daabase name (default dispatcher)
* `HEARTBEAT`: sleeping time between each api call in seconds (default 300)
* `DISPATCHER_API`: dispatcher API address

### Command line arguments
```
Usage of prometheus-dispatcher-exporter:
  -dispatcher-api string
        Dispatcher API host
  -heartbeat int
        Application heartbeat (seconds) (default 300)
  -port int
        Prometheus exporter listening port (default 9101)
  -postgresql
        Launch postgresql exporter
  -postgresql-dbname string
        Postgresql database name (default "dispatcher")
  -postgresql-host string
        Postgresql host
  -postgresql-password string
        Postgresql password
  -postgresql-port int
        Postgresql port (default 5432)
  -postgresql-user string
        Postgresql user
  -prometheus
        Launch prometheus exporter (default true)
  -version
        Farm exporter version
```
***Note** : each argument gets its default value from environment variable.*  

#
## docker compose
To launch docker compose, fill the `env` file and execute :  
```
docker compose --env-file env up -d
```
#
## PosgreSQL requirements
If you want to use long term statistics, you must have a PostgreSQL server.  

* You have to create a database called like `${POSTGRESQL_DB}` value (`dispatcher` by default)
* You have to create a user and grant privileges to this database.
* All environment variable prefixed with `POSTGRESQL` should have a value.

#
## Build docker image
1. Set your values in `env` file.
2. Build `prometheus-dispatcher-exporter binary`
```
export VERSION=<current_version>; go build -ldflags="-X 'main.Version=${VERSION}'" -o prometheus-dispatcher-exporter_${VERSION}
```
`VERSION` must be the same as in `env` file.
3. Build docker image  
```
docker compose --env-file env build
```