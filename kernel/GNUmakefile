# Nuke built-in rules and variables.
override MAKEFLAGS += -rR

# This is the name that our final kernel executable will have.
# Change as needed.
override KERNEL := kernel.elf

# How should these be defined?
override OBJCOPY := x86_64-elf-objcopy
override LD := x86_64-elf-ld

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
    -extld $(LD) \
    -extldflags '$(LDFLAGS)'"

# Use "find" to glob all *.go files in the tree and obtain the
# object and header dependency file names.
override GOFILES := $(shell find -L . -type f -name '*.go')
override ASMFILES := $(shell find -L . -type f -name '*.s')
override OBJ := $(GOFILES:.go=.o) $(ASFILES:.S=.o)

# Default target.
.PHONY: all
all: $(KERNEL)

# Link rules for the final kernel executable.
$(KERNEL): $(OBJ)
	$(LD) $(OBJ) $(LDFLAGS) -o $@

# Compilation rules for *.go and *.s files.
%.o: %.go %.s
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $@
	# Go symbols are not global by default so globalize it
	$(OBJCOPY) --globalize-symbol=_start go.o $@
	rm go.o
	rm trivial.c

# Remove object files and the final executable.
.PHONY: clean
clean:
	rm -rf $(KERNEL) $(OBJ) $(HEADER_DEPS)

.PHONY: distclean
distclean: clean
	rm -f limine.h
