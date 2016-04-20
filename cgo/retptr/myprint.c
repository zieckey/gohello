#include "myprint.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* retmalloc(int len, int *retlen)
{
    static const char* s = "0123456789";
    char* p = malloc(len);
    memcpy(p, s, len);
    *retlen = len;
    return p;
}
