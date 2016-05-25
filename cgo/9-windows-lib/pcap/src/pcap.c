#include "../include/pcap.h"
#include <stdio.h>

int pcap_print(const char* s) {
    printf("%s:%d %s <%s>\n", __FILE__, __LINE__, __FUNCTION__, s);
    return 0;
}
