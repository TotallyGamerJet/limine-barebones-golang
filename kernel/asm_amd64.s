#include "textflag.h"

TEXT _start(SB),NOSPLIT|NOFRAME,$0-0
	MOVQ main·_start(SB), DX
	MOVQ (DX), CX
	CALL CX
	RET
