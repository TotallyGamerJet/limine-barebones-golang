# Nuke built-in rules and variables.
override MAKEFLAGS += -rR

# This is the name that our final kernel executable will have.
# Change as needed.
override KERNEL := kernel.elf

# Internal linker flags that should not be changed by the user.
override LDFLAGS += \
    -nostdlib \
    -static \
    -m elf_x86_64 \
    -z max-page-size=0x1000 \
    -T linker.ld

override GOFLAGS := \
    -ldflags="-tmpdir ./ \
    -linkmode external \
    -extld x86_64-elf-ld \
    -extldflags '$(LDFLAGS)'"

# Check if the linker supports -no-pie and enable it if it does.
ifeq ($(shell $(LD) --help 2>&1 | grep 'no-pie' >/dev/null 2>&1; echo $$?),0)
    override LDFLAGS += -no-pie
endif

# Use "find" to glob all *.go files in the tree and obtain the
# object and header dependency file names.
override GOFILES := $(shell find -L . -type f -name '*.go')
override OBJ := $(GOFILES:.go=.o)

# Default target.
.PHONY: all
all: $(KERNEL)

# Link rules for the final kernel executable.
$(KERNEL): $(OBJ)
	x86_64-elf-ld $(OBJ) $(LDFLAGS) -o $@

# Compilation rules for *.go files.
%.o: %.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $@
	x86_64-elf-objcopy --globalize-symbol=_start go.o $@
	rm go.o
	rm trivial.c

# Remove object files and the final executable.
.PHONY: clean
clean:
	rm -rf $(KERNEL) $(OBJ) $(HEADER_DEPS)

.PHONY: distclean
distclean: clean
	rm -f limine.h
