bridge.so:
	/usr/bin/clang++ -o libbridge.so listeners_bridge.cpp  -std=c++20 -O3 -Wall -Wextra -fPIC -shared