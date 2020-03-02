
.PHONY: build
build: 
	mkdir -p ./bin
	go build -o ./bin/

.PHONY: install
install: 
	go install

.PHONY: call
call:
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signup '{"username":"user001", "password":"11112222"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signin '{"strategy":1, "username":"user", "password":"22223333"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signin '{"strategy":1, "username":"user001", "password":"11112222"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signin '{"strategy":1, "username":"user001", "password":"22223333"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.Signout '{"accessToken":"sssssssss"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Auth.ResetPasswd '{"accessToken":"sssssssss", "password":"22221111"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Profile.Update '{"accessToken":"sssssssss", "profile":"sdasdsada"}'
	MICRO_REGISTRY=consul micro call omo.msa.account Profile.Query '{"accessToken":"sssssssss"}'

.PHONY: tcall
tcall:
	mkdir -p ./bin
	go build -o ./bin/ ./client
	./bin/client

.PHONY: docker
docker:
	docker build . -t omo-msa-startkit:latest
