# weizili@360.cn 

CC=gcc
CXX=g++
AR=ar
ARFLAGS=cru
CURRENT_DIR=$(shell echo `pwd`)
SOURCE_DIR=$(CURRENT_DIR)/src
CFLAGS=$(TOP_CFLAGS) -g -c -O3 -fPIC \
	   -Wall -Wno-unused-parameter -Wno-unused-function -Wunused-variable -Wunused-value \
	   -Wshadow -Wcast-qual -Wcast-align -Wwrite-strings -Wsign-compare \
	   -Winvalid-pch -fms-extensions -Wfloat-equal -Wextra \
	   -I .. \
	   -I include \
	   -I src \
	   -I/usr/local/include  -MMD
CPPFLAGS= $(CFLAGS) -Woverloaded-virtual -Wsign-promo -fno-gnu-keywords
LDFLAGS=$(TOP_LDFLAGS) -L/usr/local/lib -lpthread 

LIBNAME=pcap

SRCS := $(wildcard src/*.c)
OBJS := $(patsubst %.c, %.o, $(SRCS))
DEPS := $(patsubst %.o, %.d, $(OBJS))

TARGET_SO=lib$(LIBNAME).so
TARGET_A=lib$(LIBNAME).a

all : $(TARGET_SO) $(TARGET_A)

$(TARGET_A) : $(OBJS)
	$(AR) $(ARFLAGS) $@ $(OBJS) 
	ranlib $@			   

$(TARGET_SO) : $(OBJS)
	$(CC) $(LDFLAGS) $(OBJS) -shared -o $@
	

netcore:
	$(MAKE) -f Makefile.netcoreRouter

check : $(TARGET_A) $(TARGET_SO)	
	$(MAKE) check -C test
	$(MAKE) check -C python

%.o : %.c
	$(CC) $(CFLAGS) $< -o $@
	
-include $(DEPS)

clean:
	rm -rf */*.o */*.d $(OBJS) $(DEPS) $(TARGET_SO) $(TARGET_A)

