
TARGET = $(shell basename `pwd`)

all : 
	$(MAKE) -C libprint
	go build
	./${TARGET}

clean :
	rm -rf ./${TARGET}
