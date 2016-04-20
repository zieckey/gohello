#include "myprint.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* retmalloc(int len, int *retlen)
{
    static const char* s = "0123456789";
    char* p = malloc(len);
    if (len <= strlen(s)) {
        memcpy(p, s, len);
    } else {
        memset(p, 'a', len);
    }
    *retlen = len;
    return p;
}
