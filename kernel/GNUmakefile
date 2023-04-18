# Nuke built-in rules and variables.
override MAKEFLAGS += -rR

# This is the name that our final kernel executable will have.
# Change as needed.
override KERNEL := kernel.elf

# Convenience macro to reliably declare user overridable variables.
define DEFAULT_VAR =
    ifeq ($(origin $1),default)
        override $(1) := $(2)
    endif
    ifeq ($(origin $1),undefined)
        override $(1) := $(2)
    endif
endef

# It is highly recommended to use a custom built cross toolchain to build a kernel.
# We are only using "ld" as a placeholder here. It may work by using
# the host system's toolchain, but this is not guaranteed.
$(eval $(call DEFAULT_VAR,LD,ld))

# User controllable Go flags.
$(eval $(call DEFAULT_VAR,GOFLAGS,))

# User controllable linker flags. We set none by default.
$(eval $(call DEFAULT_VAR,LDFLAGS,))

# Internal linker flags that should not be changed by the user.
override LDFLAGS += \
    -nostdlib \
    -static \
    -m elf_x86_64 \
    -z max-page-size=0x1000 \
    -T linker.ld

# Check if the linker supports -no-pie and enable it if it does.
ifeq ($(shell $(LD) --help 2>&1 | grep 'no-pie' >/dev/null 2>&1; echo $$?),0)
    override LDFLAGS += -no-pie
endif

override GOFLAGS += \
	# tell the compiler this package is a part of the stdlib so that Cgo pragmas are allowed
	# TODO: change this to the official package once moved
	-gcflags=github.com/totallygamerjet/limine-barebones-golang/kernel=-std \
    -ldflags=" \
        -linkmode external \
        -extld '$(LD)' \
        -extldflags '$(LDFLAGS)' \
    "

# Use "find" to glob all *.go and *.S files in the tree and obtain the
# object and header dependency file names.
override GOFILES := $(shell find -L . -type f -name '*.go')

# Default target.
.PHONY: all
all: $(KERNEL)

# Compilation rules for *.go files.
$(KERNEL): $(GOFILES)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $@

# Remove object files and the final executable.
.PHONY: clean
clean:
	rm -rf $(KERNEL)
