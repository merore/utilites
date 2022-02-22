OUT?=out

default: build

build: out_clean
	CGO_ENABLED=0 go build -o ${OUT} ./...

out_clean:
	if [ -d ${OUT} ]; then rm -rf ${OUT}; fi
	mkdir ${OUT}
