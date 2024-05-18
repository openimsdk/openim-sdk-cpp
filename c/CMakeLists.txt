cmake_minimum_required(VERSION 3.12)
project(test_project)

Set source code files
set(SOURCE_FILES test.c)

Set dynamic library file path
set(LIBRARY_PATH ../shared/windows)

Set dynamic library file name
set(LIBRARY_NAME libopenimsdk)

Add executable
add_executable(test.exe ${SOURCE_FILES})

Add library directory for linking
link_directories(${LIBRARY_PATH})

Add library for linking
target_link_libraries(test.exe ${LIBRARY_NAME})