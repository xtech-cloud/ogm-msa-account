module ogm-msa-account

go 1.16

require (
	github.com/asim/go-micro/plugins/config/encoder/yaml/v3 v3.0.0-20210721080634-e1bc7e302871
	github.com/asim/go-micro/plugins/config/source/etcd/v3 v3.0.0-20210721080634-e1bc7e302871
	github.com/asim/go-micro/plugins/logger/logrus/v3 v3.0.0-20210721080634-e1bc7e302871
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.0.0-20210721080634-e1bc7e302871
	github.com/asim/go-micro/plugins/server/grpc/v3 v3.0.0-20210721080634-e1bc7e302871
	github.com/asim/go-micro/v3 v3.5.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/xtech-cloud/ogm-msp-account v1.14.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	google.golang.org/genproto v0.0.0-20210721163202-f1cecdd8b78a // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.12
)
