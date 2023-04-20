#include "textflag.h"
#include "go_asm.h"

TEXT _start(SB), NOSPLIT|TOPFRAME, $0-0
	MOVQ $runtime·g0(SB), DI
	LEAQ (-65536)(SP), BX
	MOVQ BX, g_stackguard0(DI)
	MOVQ BX, g_stackguard1(DI)
	MOVQ BX, (g_stack+stack_lo)(DI)
	MOVQ SP, (g_stack+stack_hi)(DI)
	MOVQ DI, R14                    // set the g register

	CALL main·_start<ABIInternal>(SB)

	CLI

loop:
	HLT
	JMP loop
	RET
