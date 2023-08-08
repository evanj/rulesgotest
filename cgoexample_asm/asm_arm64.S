#include "asm_macros.h"

#if EXAMPLE_C_DEFINE

.global _example_asm_func
.p2align 2 /* ld: warning: arm64 function not 4-byte aligned */
_example_asm_func:
    mov x0, #44
    ret

#endif
