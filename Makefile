APP_NAME := ogm-account
BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )
TOKEN := $(shell cat /tmp/msa-token)

.PHONY: build
build:
	go build -ldflags \
		"\
		-X 'main.BuildVersion=${BUILD_VERSION}' \
		-X 'main.BuildTime=${BUILD_TIME}' \
		-X 'main.CommitID=${COMMIT_SHA1}' \
		"\
		-o ./bin/${APP_NAME}

.PHONY: run
run:
	./bin/${APP_NAME}

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm -rf /tmp/ogm-account.db

.PHONY: call
call:
	gomu --registry=etcd --client=grpc call xtc.ogm.account Healthy.Echo '{"msg":"hello"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Auth.Signup '{"username":"user001", "password":"11112222"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Auth.Signin '{"strategy":1, "username":"user", "password":"22223333"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Auth.Signin '{"strategy":1, "username":"user001", "password":"222333444"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Auth.Signin '{"strategy":1, "username":"user001", "password":"11112222"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Auth.Signout '{"accessToken":"${TOKEN}", "strategy":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Auth.ResetPasswd '{"accessToken":"${TOKEN}", "password":"22221111", "strategy":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Auth.ResetPasswd '{"accessToken":"${TOKEN}", "password":"11112222", "strategy":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Profile.Update '{"accessToken":"${TOKEN}", "profile":"sdasdsada", "strategy":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Profile.Query '{"accessToken":"${TOKEN}", "strategy":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.account Query.List '{"count":10}'

.PHONY: post
post:
	curl -X POST -d '{"msg":"hello"}' localhost/ogm/account/Healthy/Echo

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

.PHONY: docker
docker:
	docker build -t xtechcloud/${APP_NAME}:${BUILD_VERSION} .
	docker rm -f ${APP_NAME}
	docker run --restart=always --name=${APP_NAME} --net=host -v /data/${APP_NAME}:/ogm -e MSA_REGISTRY_ADDRESS='localhost:2379' -e MSA_CONFIG_DEFINE='{"source":"file","prefix":"/ogm/config","key":"${APP_NAME}.yaml"}' -d xtechcloud/${APP_NAME}:${BUILD_VERSION}
	docker logs -f ${APP_NAME}
