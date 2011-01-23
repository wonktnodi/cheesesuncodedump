#include <stdio.h>
#include <stdlib.h>

void PrintExit(void) {
    printf("exiting\n");
}

int GetStuff() {
    atexit(PrintExit);
    return 42;
}

