cmake_minimum_required(VERSION 3.20)
project(openim-sdk-cpp)
set(CMAKE_CXX_STANDARD 11)
include_directories(${CMAKE_SOURCE_DIR}/gen)
link_directories(${CMAKE_SOURCE_DIR}/gen)
file(GLOB INCLUDES "include/*.h")
file(GLOB SOURCES "src/*.cpp")

add_library(openimsdk_wrapper STATIC ${SOURCES} ${INCLUDES})
target_include_directories(openimsdk_wrapper PUBLIC include)
target_link_libraries(openimsdk_wrapper openimsdk)
add_subdirectory(examples)

