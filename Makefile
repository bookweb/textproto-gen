.PHONY: install-tools install-gotools build install
install-tools:
	go install github.com/autotag-dev/autotag@v1.4.1
	go install github.com/bookweb/textproto-gen/cmd/protoc-gen-textproto@v0.0.0-20251021161649-c827e42fb7b1

install-gotools:
	go install github.com/go-delve/delve/cmd/dlv@v1.25.2

build:
	go build -o cmd/protoc-gen-textproto/protoc-gen-textproto cmd/protoc-gen-textproto/main.go

install:
	cd cmd/protoc-gen-textproto && go install

tag-first:
	git tag v0.0.1 -m'create project'

tag:
	autotag -b master > .VERSION

tag-dev:
	autotag -p dev -b develop > .VERSION

tag-stg:
	autotag -p next -b release-next > .VERSION

gen-proto:
	./third_party/protoc/bin/protoc \
		--proto_path=proto \
		--go_out=protopb \
		--go_opt=paths=source_relative \
		proto/service.proto \
		proto/message.proto \
		proto/error.proto
		
gen-proto-test:
	./third_party/protoc/bin/protoc \
		--plugin=protoc-gen-textproto=./cmd/protoc-gen-textproto/protoc-gen-textproto \
		--proto_path=_testproto \
		--textproto_out=_testdata \
		_testproto/*.proto

.PHONY: get-protoc

get-protoc:
	mkdir -p third_party
	rm -rf third_party/protoc
	curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v21.3/protoc-21.3-linux-x86_64.zip
	unzip -d third_party/protoc protoc-21.3-linux-x86_64.zip
	rm -f protoc-21.3-linux-x86_64.zip
