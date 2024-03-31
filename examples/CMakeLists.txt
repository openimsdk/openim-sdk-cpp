cmake_minimum_required(VERSION 3.20)

file(GLOB EXAMPLE_INCLUDES "*.h")
file(GLOB EXAMPLE_SOURCES "*.cpp")
add_executable(example ${EXAMPLE_SOURCES} ${EXAMPLE_INCLUDES})
target_link_libraries(example openimsdk_wrapper)


add_custom_command(TARGET example PRE_BUILD
		COMMAND ${CMAKE_COMMAND} -E copy
		${CMAKE_SOURCE_DIR}/gen/openimsdk.dll  $<TARGET_FILE_DIR:example>/)