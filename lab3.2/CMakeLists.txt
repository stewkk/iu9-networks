cmake_minimum_required(VERSION 3.12)
project(beer-paint CXX)

include(GNUInstallDirs)


set(BOOST_URL https://boostorg.jfrog.io/artifactory/main/release/1.80.0/source/boost_1_80_0.7z)
set(BOOST_URL_SHA256 d12a2af721e22dbfb984ef061ee4c4ab2387f1904f6d65bc5acebaa34d6366ec)
add_subdirectory(boost-cmake)

# Common sources
add_library(${PROJECT_NAME}_objs OBJECT
  src/painting/model.hpp
)
target_link_libraries(${PROJECT_NAME}_objs PUBLIC Boost::coroutine)


# The Service
add_executable(${PROJECT_NAME} src/main.cpp)
target_link_libraries(${PROJECT_NAME} PRIVATE ${PROJECT_NAME}_objs)


install(TARGETS ${PROJECT_NAME} DESTINATION ${CMAKE_INSTALL_BINDIR} COMPONENT ${PROJECT_NAME})
install(FILES ${CONFIGS_FILES} DESTINATION ${CMAKE_INSTALL_SYSCONFDIR}/${PROJECT_NAME} COMPONENT ${PROJECT_NAME})
