

all:
	go build mylang

debug:
	go build -gcflags=all="-N -l" mylang

run-all:
	go build mylang
	./mylang

run-debug:
	go build -gcflags=all="-N -l" mylang
	gdb mylang
