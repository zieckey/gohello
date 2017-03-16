
FLATBUFFERS=/home/weizili/local/flatbuffers
FLATC=$(FLATBUFFERS)/bin/flatc

test.out : main.cc test_generated.h
	g++ $< -std=c++11 -I $(FLATBUFFERS)/include -o $@

fb : test.fb
	$(FLATC) -c test.fb

test_generated.h : fb

clean :
	rm -rf test_generated.h test.out
