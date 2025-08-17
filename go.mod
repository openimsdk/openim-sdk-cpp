module github.com/openimsdk/openim-sdk-cpp

go 1.24.0

require github.com/openimsdk/openim-sdk-core/v3 v3.8.3-patch.10-win7-go1.20-compat

require (
	github.com/coder/websocket v1.8.13 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/jinzhu/copier v0.4.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/openimsdk/protocol v0.0.73-alpha.12 // indirect
	github.com/openimsdk/tools v0.0.50-alpha.80 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/goleak v1.3.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/image v0.26.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/grpc v1.68.0 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gorm.io/driver/sqlite v1.5.5 // indirect
	gorm.io/gorm v1.25.10 // indirect
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.60.1
	google.golang.org/protobuf => google.golang.org/protobuf v1.33.0
)
