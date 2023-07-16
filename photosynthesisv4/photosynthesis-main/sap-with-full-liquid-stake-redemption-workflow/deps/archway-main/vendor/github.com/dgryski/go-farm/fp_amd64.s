// Code generated by command: go run asm.go -out=fp_amd64.s -go111=false. DO NOT EDIT.

// +build amd64,!purego

#include "textflag.h"

// func Fingerprint64(s []byte) uint64
TEXT ·Fingerprint64(SB), NOSPLIT, $0-32
	MOVQ  s_base+0(FP), CX
	MOVQ  s_len+8(FP), AX
	CMPQ  AX, $0x10
	JG    check32
	CMPQ  AX, $0x08
	JL    check4
	MOVQ  (CX), DX
	MOVQ  AX, BX
	SUBQ  $0x08, BX
	ADDQ  CX, BX
	MOVQ  (BX), BX
	MOVQ  $0x9ae16a3b2f90404f, BP
	ADDQ  BP, DX
	SHLQ  $0x01, AX
	ADDQ  BP, AX
	MOVQ  BX, BP
	RORQ  $0x25, BP
	IMULQ AX, BP
	ADDQ  DX, BP
	RORQ  $0x19, DX
	ADDQ  BX, DX
	IMULQ AX, DX
	XORQ  DX, BP
	IMULQ AX, BP
	MOVQ  BP, BX
	SHRQ  $0x2f, BX
	XORQ  BP, BX
	XORQ  BX, DX
	IMULQ AX, DX
	MOVQ  DX, BX
	SHRQ  $0x2f, BX
	XORQ  DX, BX
	IMULQ AX, BX
	MOVQ  BX, ret+24(FP)
	RET

check4:
	CMPQ  AX, $0x04
	JL    check0
	MOVQ  $0x9ae16a3b2f90404f, DX
	MOVQ  AX, BX
	SHLQ  $0x01, BX
	ADDQ  DX, BX
	MOVL  (CX), SI
	SHLQ  $0x03, SI
	ADDQ  AX, SI
	SUBQ  $0x04, AX
	ADDQ  AX, CX
	MOVL  (CX), DI
	XORQ  DI, SI
	IMULQ BX, SI
	MOVQ  SI, DX
	SHRQ  $0x2f, DX
	XORQ  SI, DX
	XORQ  DX, DI
	IMULQ BX, DI
	MOVQ  DI, DX
	SHRQ  $0x2f, DX
	XORQ  DI, DX
	IMULQ BX, DX
	MOVQ  DX, ret+24(FP)
	RET

check0:
	TESTQ   AX, AX
	JZ      empty
	MOVBQZX (CX), DX
	MOVQ    AX, BX
	SHRQ    $0x01, BX
	ADDQ    CX, BX
	MOVBQZX (BX), BP
	MOVQ    AX, BX
	SUBQ    $0x01, BX
	ADDQ    CX, BX
	MOVBQZX (BX), BX
	SHLQ    $0x08, BP
	ADDQ    BP, DX
	SHLQ    $0x02, BX
	ADDQ    BX, AX
	MOVQ    $0xc3a5c85c97cb3127, BX
	IMULQ   BX, AX
	MOVQ    $0x9ae16a3b2f90404f, BX
	IMULQ   BX, DX
	XORQ    DX, AX
	MOVQ    AX, DX
	SHRQ    $0x2f, DX
	XORQ    AX, DX
	IMULQ   BX, DX
	MOVQ    DX, ret+24(FP)
	RET

empty:
	MOVQ $0x9ae16a3b2f90404f, DX
	MOVQ DX, ret+24(FP)
	RET

check32:
	CMPQ  AX, $0x20
	JG    check64
	MOVQ  AX, DX
	SHLQ  $0x01, DX
	MOVQ  $0x9ae16a3b2f90404f, BX
	ADDQ  BX, DX
	MOVQ  (CX), BP
	MOVQ  $0xb492b66fbe98f273, SI
	IMULQ SI, BP
	MOVQ  8(CX), SI
	MOVQ  AX, DI
	SUBQ  $0x10, DI
	ADDQ  CX, DI
	MOVQ  8(DI), R12
	IMULQ DX, R12
	MOVQ  (DI), DI
	IMULQ BX, DI
	MOVQ  BP, R13
	ADDQ  SI, R13
	RORQ  $0x2b, R13
	ADDQ  DI, R13
	MOVQ  R12, DI
	RORQ  $0x1e, DI
	ADDQ  DI, R13
	ADDQ  R12, BP
	ADDQ  BX, SI
	RORQ  $0x12, SI
	ADDQ  SI, BP
	XORQ  BP, R13
	IMULQ DX, R13
	MOVQ  R13, BX
	SHRQ  $0x2f, BX
	XORQ  R13, BX
	XORQ  BX, BP
	IMULQ DX, BP
	MOVQ  BP, BX
	SHRQ  $0x2f, BX
	XORQ  BP, BX
	IMULQ DX, BX
	MOVQ  BX, ret+24(FP)
	RET

check64:
	CMPQ  AX, $0x40
	JG    long
	MOVQ  AX, DX
	SHLQ  $0x01, DX
	MOVQ  $0x9ae16a3b2f90404f, BX
	ADDQ  BX, DX
	MOVQ  (CX), BP
	IMULQ BX, BP
	MOVQ  8(CX), SI
	MOVQ  AX, DI
	SUBQ  $0x10, DI
	ADDQ  CX, DI
	MOVQ  8(DI), R12
	IMULQ DX, R12
	MOVQ  (DI), DI
	IMULQ BX, DI
	MOVQ  BP, R13
	ADDQ  SI, R13
	RORQ  $0x2b, R13
	ADDQ  DI, R13
	MOVQ  R12, DI
	RORQ  $0x1e, DI
	ADDQ  DI, R13
	ADDQ  BP, R12
	ADDQ  BX, SI
	RORQ  $0x12, SI
	ADDQ  SI, R12
	MOVQ  R13, BX
	XORQ  R12, BX
	IMULQ DX, BX
	MOVQ  BX, SI
	SHRQ  $0x2f, SI
	XORQ  BX, SI
	XORQ  SI, R12
	IMULQ DX, R12
	MOVQ  R12, BX
	SHRQ  $0x2f, BX
	XORQ  R12, BX
	IMULQ DX, BX
	MOVQ  16(CX), SI
	IMULQ DX, SI
	MOVQ  24(CX), DI
	MOVQ  AX, R12
	SUBQ  $0x20, R12
	ADDQ  CX, R12
	MOVQ  (R12), R14
	ADDQ  R13, R14
	IMULQ DX, R14
	MOVQ  8(R12), R12
	ADDQ  BX, R12
	IMULQ DX, R12
	MOVQ  SI, BX
	ADDQ  DI, BX
	RORQ  $0x2b, BX
	ADDQ  R12, BX
	MOVQ  R14, R12
	RORQ  $0x1e, R12
	ADDQ  R12, BX
	ADDQ  R14, SI
	ADDQ  BP, DI
	RORQ  $0x12, DI
	ADDQ  DI, SI
	XORQ  SI, BX
	IMULQ DX, BX
	MOVQ  BX, BP
	SHRQ  $0x2f, BP
	XORQ  BX, BP
	XORQ  BP, SI
	IMULQ DX, SI
	MOVQ  SI, BX
	SHRQ  $0x2f, BX
	XORQ  SI, BX
	IMULQ DX, BX
	MOVQ  BX, ret+24(FP)
	RET

long:
	XORQ R8, R8
	XORQ R9, R9
	XORQ R10, R10
	XORQ R11, R11
	MOVQ $0x01529cba0ca458ff, DX
	ADDQ (CX), DX
	MOVQ $0x226bb95b4e64b6d4, BX
	MOVQ $0x134a747f856d0526, BP
	MOVQ AX, SI
	SUBQ $0x01, SI
	MOVQ $0xffffffffffffffc0, DI
	ANDQ DI, SI
	MOVQ AX, DI
	SUBQ $0x01, DI
	ANDQ $0x3f, DI
	SUBQ $0x3f, DI
	ADDQ SI, DI
	MOVQ DI, SI
	ADDQ CX, SI
	MOVQ AX, DI

loop:
	MOVQ  $0xb492b66fbe98f273, R12
	ADDQ  BX, DX
	ADDQ  R8, DX
	ADDQ  8(CX), DX
	RORQ  $0x25, DX
	IMULQ R12, DX
	ADDQ  R9, BX
	ADDQ  48(CX), BX
	RORQ  $0x2a, BX
	IMULQ R12, BX
	XORQ  R11, DX
	ADDQ  R8, BX
	ADDQ  40(CX), BX
	ADDQ  R10, BP
	RORQ  $0x21, BP
	IMULQ R12, BP
	IMULQ R12, R9
	MOVQ  DX, R8
	ADDQ  R10, R8
	ADDQ  (CX), R9
	ADDQ  R9, R8
	ADDQ  24(CX), R8
	RORQ  $0x15, R8
	MOVQ  R9, R10
	ADDQ  8(CX), R9
	ADDQ  16(CX), R9
	MOVQ  R9, R13
	RORQ  $0x2c, R13
	ADDQ  R13, R8
	ADDQ  24(CX), R9
	ADDQ  R10, R8
	XCHGQ R9, R8
	ADDQ  BP, R11
	MOVQ  BX, R10
	ADDQ  16(CX), R10
	ADDQ  32(CX), R11
	ADDQ  R11, R10
	ADDQ  56(CX), R10
	RORQ  $0x15, R10
	MOVQ  R11, R13
	ADDQ  40(CX), R11
	ADDQ  48(CX), R11
	MOVQ  R11, R14
	RORQ  $0x2c, R14
	ADDQ  R14, R10
	ADDQ  56(CX), R11
	ADDQ  R13, R10
	XCHGQ R11, R10
	XCHGQ BP, DX
	ADDQ  $0x40, CX
	SUBQ  $0x40, DI
	CMPQ  DI, $0x40
	JG    loop
	MOVQ  SI, CX
	MOVQ  BP, DI
	ANDQ  $0xff, DI
	SHLQ  $0x01, DI
	ADDQ  R12, DI
	MOVQ  SI, CX
	SUBQ  $0x01, AX
	ANDQ  $0x3f, AX
	ADDQ  AX, R10
	ADDQ  R10, R8
	ADDQ  R8, R10
	ADDQ  BX, DX
	ADDQ  R8, DX
	ADDQ  8(CX), DX
	RORQ  $0x25, DX
	IMULQ DI, DX
	ADDQ  R9, BX
	ADDQ  48(CX), BX
	RORQ  $0x2a, BX
	IMULQ DI, BX
	MOVQ  $0x00000009, AX
	IMULQ R11, AX
	XORQ  AX, DX
	MOVQ  $0x00000009, AX
	IMULQ R8, AX
	ADDQ  AX, BX
	ADDQ  40(CX), BX
	ADDQ  R10, BP
	RORQ  $0x21, BP
	IMULQ DI, BP
	IMULQ DI, R9
	MOVQ  DX, R8
	ADDQ  R10, R8
	ADDQ  (CX), R9
	ADDQ  R9, R8
	ADDQ  24(CX), R8
	RORQ  $0x15, R8
	MOVQ  R9, AX
	ADDQ  8(CX), R9
	ADDQ  16(CX), R9
	MOVQ  R9, SI
	RORQ  $0x2c, SI
	ADDQ  SI, R8
	ADDQ  24(CX), R9
	ADDQ  AX, R8
	XCHGQ R9, R8
	ADDQ  BP, R11
	MOVQ  BX, R10
	ADDQ  16(CX), R10
	ADDQ  32(CX), R11
	ADDQ  R11, R10
	ADDQ  56(CX), R10
	RORQ  $0x15, R10
	MOVQ  R11, AX
	ADDQ  40(CX), R11
	ADDQ  48(CX), R11
	MOVQ  R11, SI
	RORQ  $0x2c, SI
	ADDQ  SI, R10
	ADDQ  56(CX), R11
	ADDQ  AX, R10
	XCHGQ R11, R10
	XCHGQ BP, DX
	XORQ  R10, R8
	IMULQ DI, R8
	MOVQ  R8, AX
	SHRQ  $0x2f, AX
	XORQ  R8, AX
	XORQ  AX, R10
	IMULQ DI, R10
	MOVQ  R10, AX
	SHRQ  $0x2f, AX
	XORQ  R10, AX
	IMULQ DI, AX
	ADDQ  BP, AX
	MOVQ  BX, CX
	SHRQ  $0x2f, CX
	XORQ  BX, CX
	MOVQ  $0xc3a5c85c97cb3127, BX
	IMULQ BX, CX
	ADDQ  CX, AX
	XORQ  R11, R9
	IMULQ DI, R9
	MOVQ  R9, CX
	SHRQ  $0x2f, CX
	XORQ  R9, CX
	XORQ  CX, R11
	IMULQ DI, R11
	MOVQ  R11, CX
	SHRQ  $0x2f, CX
	XORQ  R11, CX
	IMULQ DI, CX
	ADDQ  DX, CX
	XORQ  CX, AX
	IMULQ DI, AX
	MOVQ  AX, DX
	SHRQ  $0x2f, DX
	XORQ  AX, DX
	XORQ  DX, CX
	IMULQ DI, CX
	MOVQ  CX, AX
	SHRQ  $0x2f, AX
	XORQ  CX, AX
	IMULQ DI, AX
	MOVQ  AX, ret+24(FP)
	RET

// func Fingerprint32(s []byte) uint32
TEXT ·Fingerprint32(SB), NOSPLIT, $0-28
	MOVQ    s_base+0(FP), AX
	MOVQ    s_len+8(FP), CX
	CMPQ    CX, $0x18
	JG      long
	CMPQ    CX, $0x0c
	JG      hash_13_24
	CMPQ    CX, $0x04
	JG      hash_5_12
	XORL    DX, DX
	MOVL    $0x00000009, BX
	TESTQ   CX, CX
	JZ      done
	MOVQ    CX, BP
	MOVL    $0xcc9e2d51, DI
	IMULL   DI, DX
	MOVBLSX (AX), SI
	ADDL    SI, DX
	XORL    DX, BX
	SUBQ    $0x01, BP
	TESTQ   BP, BP
	JZ      done
	IMULL   DI, DX
	MOVBLSX 1(AX), SI
	ADDL    SI, DX
	XORL    DX, BX
	SUBQ    $0x01, BP
	TESTQ   BP, BP
	JZ      done
	IMULL   DI, DX
	MOVBLSX 2(AX), SI
	ADDL    SI, DX
	XORL    DX, BX
	SUBQ    $0x01, BP
	TESTQ   BP, BP
	JZ      done
	IMULL   DI, DX
	MOVBLSX 3(AX), SI
	ADDL    SI, DX
	XORL    DX, BX
	SUBQ    $0x01, BP
	TESTQ   BP, BP
	JZ      done

done:
	MOVL  CX, BP
	MOVL  $0xcc9e2d51, SI
	IMULL SI, BP
	RORL  $0x11, BP
	MOVL  $0x1b873593, SI
	IMULL SI, BP
	XORL  BP, BX
	RORL  $0x13, BX
	LEAL  (BX)(BX*4), BP
	LEAL  3864292196(BP), BX
	MOVL  $0xcc9e2d51, BP
	IMULL BP, DX
	RORL  $0x11, DX
	MOVL  $0x1b873593, BP
	IMULL BP, DX
	XORL  DX, BX
	RORL  $0x13, BX
	LEAL  (BX)(BX*4), DX
	LEAL  3864292196(DX), BX
	MOVL  BX, DX
	SHRL  $0x10, DX
	XORL  DX, BX
	MOVL  $0x85ebca6b, DX
	IMULL DX, BX
	MOVL  BX, DX
	SHRL  $0x0d, DX
	XORL  DX, BX
	MOVL  $0xc2b2ae35, DX
	IMULL DX, BX
	MOVL  BX, DX
	SHRL  $0x10, DX
	XORL  DX, BX
	MOVL  BX, ret+24(FP)
	RET

hash_5_12:
	MOVL  CX, DX
	MOVL  DX, BX
	SHLL  $0x02, BX
	ADDL  DX, BX
	MOVL  $0x00000009, BP
	MOVL  BX, SI
	ADDL  (AX), DX
	MOVQ  CX, DI
	SUBQ  $0x04, DI
	ADDQ  AX, DI
	ADDL  (DI), BX
	MOVQ  CX, DI
	SHRQ  $0x01, DI
	ANDQ  $0x04, DI
	ADDQ  AX, DI
	ADDL  (DI), BP
	MOVL  $0xcc9e2d51, DI
	IMULL DI, DX
	RORL  $0x11, DX
	MOVL  $0x1b873593, DI
	IMULL DI, DX
	XORL  DX, SI
	RORL  $0x13, SI
	LEAL  (SI)(SI*4), DX
	LEAL  3864292196(DX), SI
	MOVL  $0xcc9e2d51, DX
	IMULL DX, BX
	RORL  $0x11, BX
	MOVL  $0x1b873593, DX
	IMULL DX, BX
	XORL  BX, SI
	RORL  $0x13, SI
	LEAL  (SI)(SI*4), BX
	LEAL  3864292196(BX), SI
	MOVL  $0xcc9e2d51, DX
	IMULL DX, BP
	RORL  $0x11, BP
	MOVL  $0x1b873593, DX
	IMULL DX, BP
	XORL  BP, SI
	RORL  $0x13, SI
	LEAL  (SI)(SI*4), BP
	LEAL  3864292196(BP), SI
	MOVL  SI, DX
	SHRL  $0x10, DX
	XORL  DX, SI
	MOVL  $0x85ebca6b, DX
	IMULL DX, SI
	MOVL  SI, DX
	SHRL  $0x0d, DX
	XORL  DX, SI
	MOVL  $0xc2b2ae35, DX
	IMULL DX, SI
	MOVL  SI, DX
	SHRL  $0x10, DX
	XORL  DX, SI
	MOVL  SI, ret+24(FP)
	RET

hash_13_24:
	MOVQ  CX, DX
	SHRQ  $0x01, DX
	ADDQ  AX, DX
	MOVL  -4(DX), BX
	MOVL  4(AX), BP
	MOVQ  CX, SI
	ADDQ  AX, SI
	MOVL  -8(SI), DI
	MOVL  (DX), DX
	MOVL  (AX), R8
	MOVL  -4(SI), SI
	MOVL  $0xcc9e2d51, R9
	IMULL DX, R9
	ADDL  CX, R9
	RORL  $0x0c, BX
	ADDL  SI, BX
	MOVL  DI, R10
	MOVL  $0xcc9e2d51, R11
	IMULL R11, R10
	RORL  $0x11, R10
	MOVL  $0x1b873593, R11
	IMULL R11, R10
	XORL  R10, R9
	RORL  $0x13, R9
	LEAL  (R9)(R9*4), R10
	LEAL  3864292196(R10), R9
	ADDL  BX, R9
	RORL  $0x03, BX
	ADDL  DI, BX
	MOVL  $0xcc9e2d51, DI
	IMULL DI, R8
	RORL  $0x11, R8
	MOVL  $0x1b873593, DI
	IMULL DI, R8
	XORL  R8, R9
	RORL  $0x13, R9
	LEAL  (R9)(R9*4), R8
	LEAL  3864292196(R8), R9
	ADDL  BX, R9
	ADDL  SI, BX
	RORL  $0x0c, BX
	ADDL  DX, BX
	MOVL  $0xcc9e2d51, DX
	IMULL DX, BP
	RORL  $0x11, BP
	MOVL  $0x1b873593, DX
	IMULL DX, BP
	XORL  BP, R9
	RORL  $0x13, R9
	LEAL  (R9)(R9*4), BP
	LEAL  3864292196(BP), R9
	ADDL  BX, R9
	MOVL  R9, DX
	SHRL  $0x10, DX
	XORL  DX, R9
	MOVL  $0x85ebca6b, DX
	IMULL DX, R9
	MOVL  R9, DX
	SHRL  $0x0d, DX
	XORL  DX, R9
	MOVL  $0xc2b2ae35, DX
	IMULL DX, R9
	MOVL  R9, DX
	SHRL  $0x10, DX
	XORL  DX, R9
	MOVL  R9, ret+24(FP)
	RET

long:
	MOVL       CX, DX
	MOVL       $0xcc9e2d51, BX
	IMULL      DX, BX
	MOVL       BX, BP
	MOVQ       CX, SI
	ADDQ       AX, SI
	MOVL       $0xcc9e2d51, DI
	MOVL       $0x1b873593, R8
	MOVL       -4(SI), R9
	IMULL      DI, R9
	RORL       $0x11, R9
	IMULL      R8, R9
	XORL       R9, DX
	RORL       $0x13, DX
	MOVL       DX, R9
	SHLL       $0x02, R9
	ADDL       R9, DX
	ADDL       $0xe6546b64, DX
	MOVL       -8(SI), R9
	IMULL      DI, R9
	RORL       $0x11, R9
	IMULL      R8, R9
	XORL       R9, BX
	RORL       $0x13, BX
	MOVL       BX, R9
	SHLL       $0x02, R9
	ADDL       R9, BX
	ADDL       $0xe6546b64, BX
	MOVL       -16(SI), R9
	IMULL      DI, R9
	RORL       $0x11, R9
	IMULL      R8, R9
	XORL       R9, DX
	RORL       $0x13, DX
	MOVL       DX, R9
	SHLL       $0x02, R9
	ADDL       R9, DX
	ADDL       $0xe6546b64, DX
	MOVL       -12(SI), R9
	IMULL      DI, R9
	RORL       $0x11, R9
	IMULL      R8, R9
	XORL       R9, BX
	RORL       $0x13, BX
	MOVL       BX, R9
	SHLL       $0x02, R9
	ADDL       R9, BX
	ADDL       $0xe6546b64, BX
	PREFETCHT0 (AX)
	MOVL       -20(SI), SI
	IMULL      DI, SI
	RORL       $0x11, SI
	IMULL      R8, SI
	ADDL       SI, BP
	RORL       $0x13, BP
	ADDL       $0x71, BP

loop80:
	CMPQ       CX, $0x64
	JL         loop20
	PREFETCHT0 20(AX)
	MOVL       (AX), SI
	ADDL       SI, DX
	MOVL       4(AX), DI
	ADDL       DI, BX
	MOVL       8(AX), R8
	ADDL       R8, BP
	MOVL       12(AX), R9
	MOVL       R9, R11
	MOVL       $0xcc9e2d51, R10
	IMULL      R10, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, R10
	IMULL      R10, R11
	XORL       R11, DX
	RORL       $0x13, DX
	LEAL       (DX)(DX*4), R11
	LEAL       3864292196(R11), DX
	MOVL       16(AX), R10
	ADDL       R10, DX
	MOVL       R8, R11
	MOVL       $0xcc9e2d51, R8
	IMULL      R8, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, R8
	IMULL      R8, R11
	XORL       R11, BX
	RORL       $0x13, BX
	LEAL       (BX)(BX*4), R11
	LEAL       3864292196(R11), BX
	ADDL       SI, BX
	MOVL       $0xcc9e2d51, SI
	IMULL      SI, R10
	MOVL       R10, R11
	ADDL       DI, R11
	MOVL       $0xcc9e2d51, SI
	IMULL      SI, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, SI
	IMULL      SI, R11
	XORL       R11, BP
	RORL       $0x13, BP
	LEAL       (BP)(BP*4), R11
	LEAL       3864292196(R11), BP
	ADDL       R9, BP
	ADDL       BX, BP
	ADDL       BP, BX
	PREFETCHT0 40(AX)
	MOVL       20(AX), SI
	ADDL       SI, DX
	MOVL       24(AX), DI
	ADDL       DI, BX
	MOVL       28(AX), R8
	ADDL       R8, BP
	MOVL       32(AX), R9
	MOVL       R9, R11
	MOVL       $0xcc9e2d51, R10
	IMULL      R10, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, R10
	IMULL      R10, R11
	XORL       R11, DX
	RORL       $0x13, DX
	LEAL       (DX)(DX*4), R11
	LEAL       3864292196(R11), DX
	MOVL       36(AX), R10
	ADDL       R10, DX
	MOVL       R8, R11
	MOVL       $0xcc9e2d51, R8
	IMULL      R8, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, R8
	IMULL      R8, R11
	XORL       R11, BX
	RORL       $0x13, BX
	LEAL       (BX)(BX*4), R11
	LEAL       3864292196(R11), BX
	ADDL       SI, BX
	MOVL       $0xcc9e2d51, SI
	IMULL      SI, R10
	MOVL       R10, R11
	ADDL       DI, R11
	MOVL       $0xcc9e2d51, SI
	IMULL      SI, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, SI
	IMULL      SI, R11
	XORL       R11, BP
	RORL       $0x13, BP
	LEAL       (BP)(BP*4), R11
	LEAL       3864292196(R11), BP
	ADDL       R9, BP
	ADDL       BX, BP
	ADDL       BP, BX
	PREFETCHT0 60(AX)
	MOVL       40(AX), SI
	ADDL       SI, DX
	MOVL       44(AX), DI
	ADDL       DI, BX
	MOVL       48(AX), R8
	ADDL       R8, BP
	MOVL       52(AX), R9
	MOVL       R9, R11
	MOVL       $0xcc9e2d51, R10
	IMULL      R10, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, R10
	IMULL      R10, R11
	XORL       R11, DX
	RORL       $0x13, DX
	LEAL       (DX)(DX*4), R11
	LEAL       3864292196(R11), DX
	MOVL       56(AX), R10
	ADDL       R10, DX
	MOVL       R8, R11
	MOVL       $0xcc9e2d51, R8
	IMULL      R8, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, R8
	IMULL      R8, R11
	XORL       R11, BX
	RORL       $0x13, BX
	LEAL       (BX)(BX*4), R11
	LEAL       3864292196(R11), BX
	ADDL       SI, BX
	MOVL       $0xcc9e2d51, SI
	IMULL      SI, R10
	MOVL       R10, R11
	ADDL       DI, R11
	MOVL       $0xcc9e2d51, SI
	IMULL      SI, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, SI
	IMULL      SI, R11
	XORL       R11, BP
	RORL       $0x13, BP
	LEAL       (BP)(BP*4), R11
	LEAL       3864292196(R11), BP
	ADDL       R9, BP
	ADDL       BX, BP
	ADDL       BP, BX
	PREFETCHT0 80(AX)
	MOVL       60(AX), SI
	ADDL       SI, DX
	MOVL       64(AX), DI
	ADDL       DI, BX
	MOVL       68(AX), R8
	ADDL       R8, BP
	MOVL       72(AX), R9
	MOVL       R9, R11
	MOVL       $0xcc9e2d51, R10
	IMULL      R10, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, R10
	IMULL      R10, R11
	XORL       R11, DX
	RORL       $0x13, DX
	LEAL       (DX)(DX*4), R11
	LEAL       3864292196(R11), DX
	MOVL       76(AX), R10
	ADDL       R10, DX
	MOVL       R8, R11
	MOVL       $0xcc9e2d51, R8
	IMULL      R8, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, R8
	IMULL      R8, R11
	XORL       R11, BX
	RORL       $0x13, BX
	LEAL       (BX)(BX*4), R11
	LEAL       3864292196(R11), BX
	ADDL       SI, BX
	MOVL       $0xcc9e2d51, SI
	IMULL      SI, R10
	MOVL       R10, R11
	ADDL       DI, R11
	MOVL       $0xcc9e2d51, SI
	IMULL      SI, R11
	RORL       $0x11, R11
	MOVL       $0x1b873593, SI
	IMULL      SI, R11
	XORL       R11, BP
	RORL       $0x13, BP
	LEAL       (BP)(BP*4), R11
	LEAL       3864292196(R11), BP
	ADDL       R9, BP
	ADDL       BX, BP
	ADDL       BP, BX
	ADDQ       $0x50, AX
	SUBQ       $0x50, CX
	JMP        loop80

loop20:
	CMPQ  CX, $0x14
	JLE   after
	MOVL  (AX), SI
	ADDL  SI, DX
	MOVL  4(AX), DI
	ADDL  DI, BX
	MOVL  8(AX), R8
	ADDL  R8, BP
	MOVL  12(AX), R9
	MOVL  R9, R11
	MOVL  $0xcc9e2d51, R10
	IMULL R10, R11
	RORL  $0x11, R11
	MOVL  $0x1b873593, R10
	IMULL R10, R11
	XORL  R11, DX
	RORL  $0x13, DX
	LEAL  (DX)(DX*4), R11
	LEAL  3864292196(R11), DX
	MOVL  16(AX), R10
	ADDL  R10, DX
	MOVL  R8, R11
	MOVL  $0xcc9e2d51, R8
	IMULL R8, R11
	RORL  $0x11, R11
	MOVL  $0x1b873593, R8
	IMULL R8, R11
	XORL  R11, BX
	RORL  $0x13, BX
	LEAL  (BX)(BX*4), R11
	LEAL  3864292196(R11), BX
	ADDL  SI, BX
	MOVL  $0xcc9e2d51, SI
	IMULL SI, R10
	MOVL  R10, R11
	ADDL  DI, R11
	MOVL  $0xcc9e2d51, SI
	IMULL SI, R11
	RORL  $0x11, R11
	MOVL  $0x1b873593, SI
	IMULL SI, R11
	XORL  R11, BP
	RORL  $0x13, BP
	LEAL  (BP)(BP*4), R11
	LEAL  3864292196(R11), BP
	ADDL  R9, BP
	ADDL  BX, BP
	ADDL  BP, BX
	ADDQ  $0x14, AX
	SUBQ  $0x14, CX
	JMP   loop20

after:
	MOVL  $0xcc9e2d51, AX
	RORL  $0x0b, BX
	IMULL AX, BX
	RORL  $0x11, BX
	IMULL AX, BX
	RORL  $0x0b, BP
	IMULL AX, BP
	RORL  $0x11, BP
	IMULL AX, BP
	ADDL  BX, DX
	RORL  $0x13, DX
	MOVL  DX, CX
	SHLL  $0x02, CX
	ADDL  CX, DX
	ADDL  $0xe6546b64, DX
	RORL  $0x11, DX
	IMULL AX, DX
	ADDL  BP, DX
	RORL  $0x13, DX
	MOVL  DX, CX
	SHLL  $0x02, CX
	ADDL  CX, DX
	ADDL  $0xe6546b64, DX
	RORL  $0x11, DX
	IMULL AX, DX
	MOVL  DX, ret+24(FP)
	RET
