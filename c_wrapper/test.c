// gcc -o test.exe -lc_wrapper.dll test.c

#include <stdio.h>

#include "c_wrapper.h"

void main(int argc, char **argv)
{
    Init_SDK();
}