module go-code

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	// github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	go.etcd.io/bbolt v1.3.4 => github.com/coreos/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/ClickHouse/clickhouse-go v1.4.3
	github.com/Shopify/sarama v1.19.0
	github.com/antchfx/xpath v1.1.10 // indirect
	github.com/aws/aws-sdk-go v1.27.0
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/emirpasic/gods v1.12.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-redis/redis/v8 v8.11.0
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/consul/api v1.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jmoiron/sqlx v1.2.0
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/oschwald/geoip2-golang v1.4.0
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.4.0
	github.com/zouyx/agollo/v4 v4.0.4
	gitlab.xunlei.cn/pub/jaeger v0.4.2 // indirect
	go.etcd.io/bbolt v1.3.4 // indirect
	//go.etcd.io/etcd v0.5.0-alpha.5.0.20200306183522-221f0cc107cb
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.33.2 // indirect
	google.golang.org/grpc/examples v0.0.0-20201119211538-40076094f63b // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
)
