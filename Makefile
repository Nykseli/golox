

all: runner compiler vm


runner:
	go build -tags gloxrun mylang
	mv mylang gloxrun

compiler:
	go build -tags gloxcompiler mylang
	mv mylang gloxc

vm:
	go build -tags gloxvm mylang
	mv mylang gloxvm

debug:
	go build -gcflags=all="-N -l" mylang

run-all:
	go build mylang
	./mylang

run-debug:
	go build -gcflags=all="-N -l" mylang
	gdb mylang
