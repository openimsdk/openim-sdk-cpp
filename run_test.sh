#!/bin/bash


# if arg 1 is c,or is none,means test c sdk
# if arg 1 is cc,means test cpp sdk

rm ./test
if [ "$1" == "c" ] || [ "$1" == "" ]; then
    echo "test c sdk"
    gcc -o test  ./test.c ./openimsdk.so
    ./test
    exit 0
elif [ "$1" == "cc" ]; then
    echo "test cpp sdk"
    g++ -o test  ./test.cc ./openimsdk.so ./openimsdkcc.so
    ./test
    exit 0
fi