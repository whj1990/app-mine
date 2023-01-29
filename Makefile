.PHONY: init
init:
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
	go install github.com/cloudwego/thriftgo@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: thrift
thrift:
	kitex -module github.com/whj1990/app-mine -service app-mine mine.thrift

.PHONY: wire
wire:
	wire