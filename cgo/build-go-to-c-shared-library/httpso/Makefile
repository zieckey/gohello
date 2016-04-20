
CC := gcc
CFLAGS := -g -c -fPIC -Wshadow -Wcast-qual -Wcast-align -Wwrite-strings -Wsign-compare -Winvalid-pch -fms-extensions -Wall

backdoor.so: main.go
	go build -buildmode=c-shared -o $@

run :
	LD_PRELOAD=./backdoor.so top

%.o : %.c
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm -rf libhello.* *.o hello
