
CC := gcc
CFLAGS := -g -c -fPIC -Wshadow -Wcast-qual -Wcast-align -Wwrite-strings -Wsign-compare -Winvalid-pch -fms-extensions -Wall
LDFLAGS := -L . -lhttp

SRCS := $(wildcard *.c)
OBJS := $(patsubst %.c, %.o, $(SRCS))
DEPS := $(patsubst %.o, %.d, $(OBJS))

http :	libhttp.so $(OBJS)
	$(CC) $(OBJS) $(LDFLAGS) -o $@

libhttp.so: http.go
	go build -buildmode=c-shared -o $@ http.go

run :

%.o : %.c
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm -rf libhttp.* *.o http 
