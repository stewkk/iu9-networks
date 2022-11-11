cmake_minimum_required(VERSION 3.12)
project(beer-paint CXX)

include(GNUInstallDirs)

set (CMAKE_CXX_STANDARD 20)

find_package(Boost 1.76 REQUIRED COMPONENTS coroutine)
include_directories(${Boost_INCLUDE_DIR})

find_package(PkgConfig)
pkg_check_modules(GTKMM gtkmm-4.0)
include_directories(${GTKMM_INCLUDE_DIRS})

# Common sources
add_library(${PROJECT_NAME}_objs OBJECT
  src/tmp.hpp
)
target_link_libraries(${PROJECT_NAME}_objs PUBLIC ${Boost_LIBRARIES} ${GTKMM_LIBRARIES})


# The Service
add_executable(${PROJECT_NAME} src/main.cpp)
target_link_libraries(${PROJECT_NAME} PRIVATE ${PROJECT_NAME}_objs)


install(TARGETS ${PROJECT_NAME} DESTINATION ${CMAKE_INSTALL_BINDIR} COMPONENT ${PROJECT_NAME})
install(FILES ${CONFIGS_FILES} DESTINATION ${CMAKE_INSTALL_SYSCONFDIR}/${PROJECT_NAME} COMPONENT ${PROJECT_NAME})
