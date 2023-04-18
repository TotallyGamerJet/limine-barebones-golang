#include "textflag.h"

// Go doesn't have inline assembly
TEXT main·hcf(SB), NOSPLIT, $0-0
    CLI
loop:
    HLT
    JMP loop
    RET
