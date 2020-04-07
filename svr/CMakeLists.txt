cmake_minimum_required (VERSION 3.8)

project(mo LANGUAGES C CXX)

set(CMAKE_C_STANDARD 11)
set(CMAKE_CXX_STANDARD 11)

message(STATUS "platform: ${CMAKE_SYSTEM_NAME} ")

if(CMAKE_SYSTEM_NAME MATCHES "Linux")
    set(ME_COM
        ME_COM_COMPILER=1
        ME_COM_LIB=1
        ME_COM_MATRIXSSL=0
        ME_COM_MBEDTLS=1
        ME_COM_MPR=1
        ME_COM_NANOSSL=0
        ME_COM_OPENSSL=0
        ME_COM_OSDEP=1
        ME_COM_PCRE=1
        ME_COM_SSL=1
        ME_COM_VXWORKS=0
        -D_REENTRANT
        -DPIC
    )
    set(sys_libs
        -lrt
        -ldl
        -lpthread
        -lm
    )
elseif(CMAKE_SYSTEM_NAME MATCHES "Windows")

elseif(CMAKE_SYSTEM_NAME MATCHES "FreeBSD")
    set(sys_libs
        -ldl
        -lpthread
        -lm
    )
    message(WARNING "${CMAKE_SYSTEM_NAME} Not Supported")
elseif(CMAKE_SYSTEM_NAME MATCHES "Darwin")
    set(sys_libs
        -ldl
        -lpthread
        -lm
    )
    message(WARNING "${CMAKE_SYSTEM_NAME} Not Supported")
elseif(CMAKE_SYSTEM_NAME MATCHES "VxWorks")
    message(WARNING "${CMAKE_SYSTEM_NAME} Not Supported")
else()
    message(WARNING "${CMAKE_SYSTEM_NAME} Not Supported")
endif(CMAKE_SYSTEM_NAME MATCHES "Linux")
if(CMAKE_SYSTEM_NAME MATCHES "Windows")

else(CMAKE_SYSTEM_NAME MATCHES "Windows")
    add_compile_options(-fPIC -w)
    file(COPY platform/me.${CMAKE_SYSTEM_NAME} DESTINATION inc)
    file(RENAME inc/me.${CMAKE_SYSTEM_NAME} inc/me.h)
endif(CMAKE_SYSTEM_NAME MATCHES "Windows")


include_directories(deps/http/)
include_directories(deps/mbedtls/)
include_directories(deps/mpr/)
include_directories(deps/mpr-mbedtls/)
include_directories(deps/mpr-openssl/)
include_directories(deps/osdep/)
include_directories(deps/pcre/)
include_directories(deps/appweb/)
include_directories(inc/)

set(LIBRARY_OUTPUT_PATH build/lib) 
set(EXECUTABLE_OUTPUT_PATH build/out)

add_library(pcre ./deps/pcre/pcre.c)
add_library(mpr ./deps/mpr/mprLib.c)
add_library(mpr-mbedtls ./deps/mpr-mbedtls/mpr-mbedtls.c)
add_library(http ./deps/http/httpLib.c)
add_library(mbedtls ./deps/mbedtls/mbedtls.c)
add_library(appweb ./deps/appweb/config.c)

link_directories(${LIBRARY_OUTPUT_PATH})

set(self_build_libs
    http
    pcre
    mpr
    mpr-mbedtls
    mbedtls
    appweb
)

add_executable (mo "src/server.cc")

target_compile_definitions(pcre
    PRIVATE
    ${ME_COM}
)

target_compile_definitions(mpr
    PRIVATE
    ${ME_COM}
    MBEDTLS_USER_CONFIG_FILE="embedtls.h"        
)

target_compile_definitions(mpr-mbedtls
    PRIVATE
    ${ME_COM}
    MBEDTLS_USER_CONFIG_FILE="embedtls.h"        
)

target_compile_definitions(http
    PRIVATE
    ${ME_COM}
    MBEDTLS_USER_CONFIG_FILE="embedtls.h"        
)

target_compile_definitions(mbedtls
    PRIVATE
    ${ME_COM}
    MBEDTLS_USER_CONFIG_FILE="embedtls.h"        
)

target_compile_definitions(appweb
    PRIVATE
    ${ME_COM}
)

target_link_libraries(mo
    ${sys_libs}
    ${self_build_libs}
)


##########################################################
#add_library(mpr-openssl ./deps/mpr-openssl/mpr-openssl.c)

#target_compile_definitions(mpr-openssl
#    PRIVATE
#    ${ME_COM}
#)
##########################################################