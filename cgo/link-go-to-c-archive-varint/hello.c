#include <stdio.h>
#include <stdint.h>
#include "libhello.h"

int main() {
    printf("This is a C Application.\n");
    GoString name = {(char*)"Jane", 4};
    SayHello(name);
    GoSlice buf = {(void*)"Jane", 4, 4};
    SayHelloByte(buf);
    SayBye();


    int32_t id = 1, row = 2;
    GoInterface gi = {&id, &row};

    GoSlice slice = DocIdEncode(gi);
    printf("DocIdEncode:");
    for (int i = 0; i < slice.len; i++) {
        char* p = (char*)slice.data;
        printf("%d\t", p[i]);
    }
    return 0;
}
