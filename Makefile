APP_NAME := omo-msa-account
BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )

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
	rm -rf /tmp/msa-account.db

.PHONY: call
call:
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signup '{"username":"user001", "password":"11112222"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signin '{"strategy":1, "username":"user", "password":"22223333"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signin '{"strategy":1, "username":"user001", "password":"222333444"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signin '{"strategy":1, "username":"user001", "password":"11112222"}' 
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signout '{"accessToken":"sssssssss"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.ResetPasswd '{"accessToken":"sssssssss", "password":"22221111"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Profile.Update '{"accessToken":"sssssssss", "profile":"sdasdsada"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Profile.Query '{"accessToken":"sssssssss"}'

.PHONY: tcall
tcall:
	mkdir -p ./bin
	go build -o ./bin/ ./client
	./bin/client

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

.PHONY: docker
docker:
	docker build . -t omo-msa-startkit:latest
