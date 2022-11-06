CMAKE_COMMON_FLAGS ?= -DCMAKE_EXPORT_COMPILE_COMMANDS=ON
CMAKE_DEBUG_FLAGS ?= -DSANITIZE='addr ub'
CMAKE_RELEASE_FLAGS ?=
CMAKE_OS_FLAGS ?=
NPROCS ?= $(shell nproc)

-include Makefile.local

all: test-debug test-release

build_debug/Makefile:
	@git submodule update --init
	@mkdir -p build_debug
	@cd build_debug && \
      cmake -DCMAKE_BUILD_TYPE=Debug $(CMAKE_COMMON_FLAGS) $(CMAKE_DEBUG_FLAGS) $(CMAKE_OS_FLAGS) $(CMAKE_OPTIONS) ..

# Release cmake configuration
build_release/Makefile:
	@git submodule update --init
	@mkdir -p build_release
	@cd build_release && \
      cmake -DCMAKE_BUILD_TYPE=Release $(CMAKE_COMMON_FLAGS) $(CMAKE_RELEASE_FLAGS) $(CMAKE_OS_FLAGS) $(CMAKE_OPTIONS) ..

# build using cmake
build-impl-%: build_%/Makefile
	@cmake --build build_$* -j $(NPROCS)

service-impl-start-%: build-impl-%
	@cd ./build_$* && ./beer-paint

# clean
clean-impl-%:
	cd build_$* && $(MAKE) clean

# dist-clean
.PHONY: dist-clean
dist-clean:
	@rm -rf build_*

# format
.PHONY: format
format:
	@find src -name '*pp' -type f | xargs clang-format -i

.PHONY: cmake-debug build-debug test-debug clean-debug cmake-release build-release test-release clean-release install install-debug

install: build-release
	@cd build_release && \
		cmake --install . -v

install-debug: build-debug
	@cd build_debug && \
		cmake --install . -v

# Explicitly specifying the targets to help shell with completions
cmake-debug: build_debug/Makefile
cmake-release: build_release/Makefile

build-debug: build-impl-debug
build-release: build-impl-release

test-debug: test-impl-debug
test-release: test-impl-release

service-start-debug: service-impl-start-debug
service-start-release: service-impl-start-release

clean-debug: clean-impl-debug
clean-release: clean-impl-release
