#include "textflag.h"
#include "go_asm.h"

// copied from go_tls.h
#ifdef GOARCH_amd64
#define get_tls(r)	MOVQ TLS, r
#define g(r)	0(r)(TLS*1)
#endif

// func setg(gg *g)
// set g. for use by needm.
TEXT ·setg<ABIInternal>(SB), NOSPLIT, $0-8
	//get_tls(CX)
	//MOVQ AX, g(CX)
	MOVQ AX, R14   // set the g register
	RET
