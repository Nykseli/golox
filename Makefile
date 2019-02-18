

all:
	go build -tags gloxrun mylang

compiler:
	go build -tags gloxcompiler mylang

vm:
	go build -tags gloxvm mylang

debug:
	go build -gcflags=all="-N -l" mylang

run-all:
	go build mylang
	./mylang

run-debug:
	go build -gcflags=all="-N -l" mylang
	gdb mylang
