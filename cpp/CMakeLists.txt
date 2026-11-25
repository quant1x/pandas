cmake_minimum_required(VERSION 3.20)

project(quant1x_pandas_cpp LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

if(MSVC)
    add_compile_options(/W4 /utf-8)
else()
    add_compile_options(-Wall -Wextra)
endif()

# header-only 公式库；首批不依赖 num / xtensor，后续再评估切换。
add_library(quant1x_pandas_cpp INTERFACE)
target_include_directories(quant1x_pandas_cpp INTERFACE "${CMAKE_CURRENT_SOURCE_DIR}")

add_executable(formula_test tests/formula_test.cpp)
target_link_libraries(formula_test PRIVATE quant1x_pandas_cpp)

include(CTest)
if(BUILD_TESTING)
    add_test(NAME formula_test COMMAND formula_test)
endif()
