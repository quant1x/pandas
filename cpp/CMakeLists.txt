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
# 头文件与 Go 的 formula/*.go 同目录（仓库根的 formula/），包含根是仓库根
target_include_directories(quant1x_pandas_cpp INTERFACE "${CMAKE_CURRENT_SOURCE_DIR}/..")

add_executable(formula_test tests/formula_test.cpp)
target_link_libraries(formula_test PRIVATE quant1x_pandas_cpp)

add_executable(formula_fixture_test tests/fixture_test.cpp)
target_link_libraries(formula_fixture_test PRIVATE quant1x_pandas_cpp)
target_compile_definitions(formula_fixture_test PRIVATE
        PANDAS_FORMULA_FIXTURES="${CMAKE_CURRENT_SOURCE_DIR}/../tests/formula/cases.csv")

# M8 特征 fixture：输入 OHLCV 与期望特征同表
add_executable(features_fixture_test tests/features_test.cpp)
target_link_libraries(features_fixture_test PRIVATE quant1x_pandas_cpp)
target_compile_definitions(features_fixture_test PRIVATE
        PANDAS_FEATURE_FIXTURES="${CMAKE_CURRENT_SOURCE_DIR}/../tests/features/features.csv"
        PANDAS_FEATURE_MODEL="${CMAKE_CURRENT_SOURCE_DIR}/../tests/features/model.csv"
        PANDAS_FEATURE_MLP="${CMAKE_CURRENT_SOURCE_DIR}/../tests/features/model_mlp.csv")


# M8：ONNX 推理对拍（需要 onnxruntime SDK；缺省不构建）
option(PANDAS_WITH_ONNX "构建 ONNX 推理对拍（需 onnxruntime SDK 的 include / lib / dll）" OFF)
set(PANDAS_ORT_INCLUDE_DIR "" CACHE PATH "onnxruntime 头文件目录（含 onnxruntime_c_api.h）")
set(PANDAS_ORT_LIB "" CACHE FILEPATH "onnxruntime 导入库")
set(PANDAS_ORT_DLL "" CACHE FILEPATH "onnxruntime 动态库（构建后拷到可执行文件旁）")
if(PANDAS_WITH_ONNX)
    add_executable(onnx_fixture_test tests/onnx_test.cpp)
    target_include_directories(onnx_fixture_test PRIVATE ${PANDAS_ORT_INCLUDE_DIR})
    target_link_libraries(onnx_fixture_test PRIVATE quant1x_pandas_cpp ${PANDAS_ORT_LIB})
    target_compile_definitions(onnx_fixture_test PRIVATE
            PANDAS_FEATURE_FIXTURES="${CMAKE_CURRENT_SOURCE_DIR}/../tests/features/features.csv"
            PANDAS_ONNX_MODEL="${CMAKE_CURRENT_SOURCE_DIR}/../tests/features/onnx/mlp.onnx")
    add_custom_command(TARGET onnx_fixture_test POST_BUILD
            COMMAND ${CMAKE_COMMAND} -E copy_if_different
            ${PANDAS_ORT_DLL} $<TARGET_FILE_DIR:onnx_fixture_test>)
endif()

include(CTest)
if(BUILD_TESTING)
    add_test(NAME formula_test COMMAND formula_test)
    add_test(NAME formula_fixture_test COMMAND formula_fixture_test)
    add_test(NAME features_fixture_test COMMAND features_fixture_test)
    if(PANDAS_WITH_ONNX)
        add_test(NAME onnx_fixture_test COMMAND onnx_fixture_test)
    endif()
endif()
