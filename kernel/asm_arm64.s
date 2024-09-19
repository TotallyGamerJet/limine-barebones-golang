#include "textflag.h"

TEXT _start(SB),NOSPLIT|NOFRAME,$0-0
	MOVD main·_start(SB), R26
	MOVD (R26), R2
	CALL (R2)
	RET
