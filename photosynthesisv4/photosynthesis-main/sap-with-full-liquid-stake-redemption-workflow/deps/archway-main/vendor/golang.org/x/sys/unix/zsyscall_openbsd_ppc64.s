// go run mkasm.go openbsd ppc64
// Code generated by the command above; DO NOT EDIT.

#include "textflag.h"

TEXT libc_getgroups_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getgroups(SB)
	RET
GLOBL	·libc_getgroups_trampoline_addr(SB), RODATA, $8
DATA	·libc_getgroups_trampoline_addr(SB)/8, $libc_getgroups_trampoline<>(SB)

TEXT libc_setgroups_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setgroups(SB)
	RET
GLOBL	·libc_setgroups_trampoline_addr(SB), RODATA, $8
DATA	·libc_setgroups_trampoline_addr(SB)/8, $libc_setgroups_trampoline<>(SB)

TEXT libc_wait4_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_wait4(SB)
	RET
GLOBL	·libc_wait4_trampoline_addr(SB), RODATA, $8
DATA	·libc_wait4_trampoline_addr(SB)/8, $libc_wait4_trampoline<>(SB)

TEXT libc_accept_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_accept(SB)
	RET
GLOBL	·libc_accept_trampoline_addr(SB), RODATA, $8
DATA	·libc_accept_trampoline_addr(SB)/8, $libc_accept_trampoline<>(SB)

TEXT libc_bind_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_bind(SB)
	RET
GLOBL	·libc_bind_trampoline_addr(SB), RODATA, $8
DATA	·libc_bind_trampoline_addr(SB)/8, $libc_bind_trampoline<>(SB)

TEXT libc_connect_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_connect(SB)
	RET
GLOBL	·libc_connect_trampoline_addr(SB), RODATA, $8
DATA	·libc_connect_trampoline_addr(SB)/8, $libc_connect_trampoline<>(SB)

TEXT libc_socket_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_socket(SB)
	RET
GLOBL	·libc_socket_trampoline_addr(SB), RODATA, $8
DATA	·libc_socket_trampoline_addr(SB)/8, $libc_socket_trampoline<>(SB)

TEXT libc_getsockopt_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getsockopt(SB)
	RET
GLOBL	·libc_getsockopt_trampoline_addr(SB), RODATA, $8
DATA	·libc_getsockopt_trampoline_addr(SB)/8, $libc_getsockopt_trampoline<>(SB)

TEXT libc_setsockopt_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setsockopt(SB)
	RET
GLOBL	·libc_setsockopt_trampoline_addr(SB), RODATA, $8
DATA	·libc_setsockopt_trampoline_addr(SB)/8, $libc_setsockopt_trampoline<>(SB)

TEXT libc_getpeername_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getpeername(SB)
	RET
GLOBL	·libc_getpeername_trampoline_addr(SB), RODATA, $8
DATA	·libc_getpeername_trampoline_addr(SB)/8, $libc_getpeername_trampoline<>(SB)

TEXT libc_getsockname_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getsockname(SB)
	RET
GLOBL	·libc_getsockname_trampoline_addr(SB), RODATA, $8
DATA	·libc_getsockname_trampoline_addr(SB)/8, $libc_getsockname_trampoline<>(SB)

TEXT libc_shutdown_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_shutdown(SB)
	RET
GLOBL	·libc_shutdown_trampoline_addr(SB), RODATA, $8
DATA	·libc_shutdown_trampoline_addr(SB)/8, $libc_shutdown_trampoline<>(SB)

TEXT libc_socketpair_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_socketpair(SB)
	RET
GLOBL	·libc_socketpair_trampoline_addr(SB), RODATA, $8
DATA	·libc_socketpair_trampoline_addr(SB)/8, $libc_socketpair_trampoline<>(SB)

TEXT libc_recvfrom_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_recvfrom(SB)
	RET
GLOBL	·libc_recvfrom_trampoline_addr(SB), RODATA, $8
DATA	·libc_recvfrom_trampoline_addr(SB)/8, $libc_recvfrom_trampoline<>(SB)

TEXT libc_sendto_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_sendto(SB)
	RET
GLOBL	·libc_sendto_trampoline_addr(SB), RODATA, $8
DATA	·libc_sendto_trampoline_addr(SB)/8, $libc_sendto_trampoline<>(SB)

TEXT libc_recvmsg_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_recvmsg(SB)
	RET
GLOBL	·libc_recvmsg_trampoline_addr(SB), RODATA, $8
DATA	·libc_recvmsg_trampoline_addr(SB)/8, $libc_recvmsg_trampoline<>(SB)

TEXT libc_sendmsg_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_sendmsg(SB)
	RET
GLOBL	·libc_sendmsg_trampoline_addr(SB), RODATA, $8
DATA	·libc_sendmsg_trampoline_addr(SB)/8, $libc_sendmsg_trampoline<>(SB)

TEXT libc_kevent_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_kevent(SB)
	RET
GLOBL	·libc_kevent_trampoline_addr(SB), RODATA, $8
DATA	·libc_kevent_trampoline_addr(SB)/8, $libc_kevent_trampoline<>(SB)

TEXT libc_utimes_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_utimes(SB)
	RET
GLOBL	·libc_utimes_trampoline_addr(SB), RODATA, $8
DATA	·libc_utimes_trampoline_addr(SB)/8, $libc_utimes_trampoline<>(SB)

TEXT libc_futimes_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_futimes(SB)
	RET
GLOBL	·libc_futimes_trampoline_addr(SB), RODATA, $8
DATA	·libc_futimes_trampoline_addr(SB)/8, $libc_futimes_trampoline<>(SB)

TEXT libc_poll_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_poll(SB)
	RET
GLOBL	·libc_poll_trampoline_addr(SB), RODATA, $8
DATA	·libc_poll_trampoline_addr(SB)/8, $libc_poll_trampoline<>(SB)

TEXT libc_madvise_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_madvise(SB)
	RET
GLOBL	·libc_madvise_trampoline_addr(SB), RODATA, $8
DATA	·libc_madvise_trampoline_addr(SB)/8, $libc_madvise_trampoline<>(SB)

TEXT libc_mlock_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mlock(SB)
	RET
GLOBL	·libc_mlock_trampoline_addr(SB), RODATA, $8
DATA	·libc_mlock_trampoline_addr(SB)/8, $libc_mlock_trampoline<>(SB)

TEXT libc_mlockall_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mlockall(SB)
	RET
GLOBL	·libc_mlockall_trampoline_addr(SB), RODATA, $8
DATA	·libc_mlockall_trampoline_addr(SB)/8, $libc_mlockall_trampoline<>(SB)

TEXT libc_mprotect_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mprotect(SB)
	RET
GLOBL	·libc_mprotect_trampoline_addr(SB), RODATA, $8
DATA	·libc_mprotect_trampoline_addr(SB)/8, $libc_mprotect_trampoline<>(SB)

TEXT libc_msync_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_msync(SB)
	RET
GLOBL	·libc_msync_trampoline_addr(SB), RODATA, $8
DATA	·libc_msync_trampoline_addr(SB)/8, $libc_msync_trampoline<>(SB)

TEXT libc_munlock_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_munlock(SB)
	RET
GLOBL	·libc_munlock_trampoline_addr(SB), RODATA, $8
DATA	·libc_munlock_trampoline_addr(SB)/8, $libc_munlock_trampoline<>(SB)

TEXT libc_munlockall_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_munlockall(SB)
	RET
GLOBL	·libc_munlockall_trampoline_addr(SB), RODATA, $8
DATA	·libc_munlockall_trampoline_addr(SB)/8, $libc_munlockall_trampoline<>(SB)

TEXT libc_pipe2_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_pipe2(SB)
	RET
GLOBL	·libc_pipe2_trampoline_addr(SB), RODATA, $8
DATA	·libc_pipe2_trampoline_addr(SB)/8, $libc_pipe2_trampoline<>(SB)

TEXT libc_getdents_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getdents(SB)
	RET
GLOBL	·libc_getdents_trampoline_addr(SB), RODATA, $8
DATA	·libc_getdents_trampoline_addr(SB)/8, $libc_getdents_trampoline<>(SB)

TEXT libc_getcwd_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getcwd(SB)
	RET
GLOBL	·libc_getcwd_trampoline_addr(SB), RODATA, $8
DATA	·libc_getcwd_trampoline_addr(SB)/8, $libc_getcwd_trampoline<>(SB)

TEXT libc_ioctl_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_ioctl(SB)
	RET
GLOBL	·libc_ioctl_trampoline_addr(SB), RODATA, $8
DATA	·libc_ioctl_trampoline_addr(SB)/8, $libc_ioctl_trampoline<>(SB)

TEXT libc_sysctl_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_sysctl(SB)
	RET
GLOBL	·libc_sysctl_trampoline_addr(SB), RODATA, $8
DATA	·libc_sysctl_trampoline_addr(SB)/8, $libc_sysctl_trampoline<>(SB)

TEXT libc_ppoll_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_ppoll(SB)
	RET
GLOBL	·libc_ppoll_trampoline_addr(SB), RODATA, $8
DATA	·libc_ppoll_trampoline_addr(SB)/8, $libc_ppoll_trampoline<>(SB)

TEXT libc_access_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_access(SB)
	RET
GLOBL	·libc_access_trampoline_addr(SB), RODATA, $8
DATA	·libc_access_trampoline_addr(SB)/8, $libc_access_trampoline<>(SB)

TEXT libc_adjtime_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_adjtime(SB)
	RET
GLOBL	·libc_adjtime_trampoline_addr(SB), RODATA, $8
DATA	·libc_adjtime_trampoline_addr(SB)/8, $libc_adjtime_trampoline<>(SB)

TEXT libc_chdir_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_chdir(SB)
	RET
GLOBL	·libc_chdir_trampoline_addr(SB), RODATA, $8
DATA	·libc_chdir_trampoline_addr(SB)/8, $libc_chdir_trampoline<>(SB)

TEXT libc_chflags_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_chflags(SB)
	RET
GLOBL	·libc_chflags_trampoline_addr(SB), RODATA, $8
DATA	·libc_chflags_trampoline_addr(SB)/8, $libc_chflags_trampoline<>(SB)

TEXT libc_chmod_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_chmod(SB)
	RET
GLOBL	·libc_chmod_trampoline_addr(SB), RODATA, $8
DATA	·libc_chmod_trampoline_addr(SB)/8, $libc_chmod_trampoline<>(SB)

TEXT libc_chown_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_chown(SB)
	RET
GLOBL	·libc_chown_trampoline_addr(SB), RODATA, $8
DATA	·libc_chown_trampoline_addr(SB)/8, $libc_chown_trampoline<>(SB)

TEXT libc_chroot_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_chroot(SB)
	RET
GLOBL	·libc_chroot_trampoline_addr(SB), RODATA, $8
DATA	·libc_chroot_trampoline_addr(SB)/8, $libc_chroot_trampoline<>(SB)

TEXT libc_close_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_close(SB)
	RET
GLOBL	·libc_close_trampoline_addr(SB), RODATA, $8
DATA	·libc_close_trampoline_addr(SB)/8, $libc_close_trampoline<>(SB)

TEXT libc_dup_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_dup(SB)
	RET
GLOBL	·libc_dup_trampoline_addr(SB), RODATA, $8
DATA	·libc_dup_trampoline_addr(SB)/8, $libc_dup_trampoline<>(SB)

TEXT libc_dup2_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_dup2(SB)
	RET
GLOBL	·libc_dup2_trampoline_addr(SB), RODATA, $8
DATA	·libc_dup2_trampoline_addr(SB)/8, $libc_dup2_trampoline<>(SB)

TEXT libc_dup3_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_dup3(SB)
	RET
GLOBL	·libc_dup3_trampoline_addr(SB), RODATA, $8
DATA	·libc_dup3_trampoline_addr(SB)/8, $libc_dup3_trampoline<>(SB)

TEXT libc_exit_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_exit(SB)
	RET
GLOBL	·libc_exit_trampoline_addr(SB), RODATA, $8
DATA	·libc_exit_trampoline_addr(SB)/8, $libc_exit_trampoline<>(SB)

TEXT libc_faccessat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_faccessat(SB)
	RET
GLOBL	·libc_faccessat_trampoline_addr(SB), RODATA, $8
DATA	·libc_faccessat_trampoline_addr(SB)/8, $libc_faccessat_trampoline<>(SB)

TEXT libc_fchdir_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fchdir(SB)
	RET
GLOBL	·libc_fchdir_trampoline_addr(SB), RODATA, $8
DATA	·libc_fchdir_trampoline_addr(SB)/8, $libc_fchdir_trampoline<>(SB)

TEXT libc_fchflags_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fchflags(SB)
	RET
GLOBL	·libc_fchflags_trampoline_addr(SB), RODATA, $8
DATA	·libc_fchflags_trampoline_addr(SB)/8, $libc_fchflags_trampoline<>(SB)

TEXT libc_fchmod_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fchmod(SB)
	RET
GLOBL	·libc_fchmod_trampoline_addr(SB), RODATA, $8
DATA	·libc_fchmod_trampoline_addr(SB)/8, $libc_fchmod_trampoline<>(SB)

TEXT libc_fchmodat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fchmodat(SB)
	RET
GLOBL	·libc_fchmodat_trampoline_addr(SB), RODATA, $8
DATA	·libc_fchmodat_trampoline_addr(SB)/8, $libc_fchmodat_trampoline<>(SB)

TEXT libc_fchown_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fchown(SB)
	RET
GLOBL	·libc_fchown_trampoline_addr(SB), RODATA, $8
DATA	·libc_fchown_trampoline_addr(SB)/8, $libc_fchown_trampoline<>(SB)

TEXT libc_fchownat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fchownat(SB)
	RET
GLOBL	·libc_fchownat_trampoline_addr(SB), RODATA, $8
DATA	·libc_fchownat_trampoline_addr(SB)/8, $libc_fchownat_trampoline<>(SB)

TEXT libc_flock_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_flock(SB)
	RET
GLOBL	·libc_flock_trampoline_addr(SB), RODATA, $8
DATA	·libc_flock_trampoline_addr(SB)/8, $libc_flock_trampoline<>(SB)

TEXT libc_fpathconf_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fpathconf(SB)
	RET
GLOBL	·libc_fpathconf_trampoline_addr(SB), RODATA, $8
DATA	·libc_fpathconf_trampoline_addr(SB)/8, $libc_fpathconf_trampoline<>(SB)

TEXT libc_fstat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fstat(SB)
	RET
GLOBL	·libc_fstat_trampoline_addr(SB), RODATA, $8
DATA	·libc_fstat_trampoline_addr(SB)/8, $libc_fstat_trampoline<>(SB)

TEXT libc_fstatat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fstatat(SB)
	RET
GLOBL	·libc_fstatat_trampoline_addr(SB), RODATA, $8
DATA	·libc_fstatat_trampoline_addr(SB)/8, $libc_fstatat_trampoline<>(SB)

TEXT libc_fstatfs_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fstatfs(SB)
	RET
GLOBL	·libc_fstatfs_trampoline_addr(SB), RODATA, $8
DATA	·libc_fstatfs_trampoline_addr(SB)/8, $libc_fstatfs_trampoline<>(SB)

TEXT libc_fsync_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_fsync(SB)
	RET
GLOBL	·libc_fsync_trampoline_addr(SB), RODATA, $8
DATA	·libc_fsync_trampoline_addr(SB)/8, $libc_fsync_trampoline<>(SB)

TEXT libc_ftruncate_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_ftruncate(SB)
	RET
GLOBL	·libc_ftruncate_trampoline_addr(SB), RODATA, $8
DATA	·libc_ftruncate_trampoline_addr(SB)/8, $libc_ftruncate_trampoline<>(SB)

TEXT libc_getegid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getegid(SB)
	RET
GLOBL	·libc_getegid_trampoline_addr(SB), RODATA, $8
DATA	·libc_getegid_trampoline_addr(SB)/8, $libc_getegid_trampoline<>(SB)

TEXT libc_geteuid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_geteuid(SB)
	RET
GLOBL	·libc_geteuid_trampoline_addr(SB), RODATA, $8
DATA	·libc_geteuid_trampoline_addr(SB)/8, $libc_geteuid_trampoline<>(SB)

TEXT libc_getgid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getgid(SB)
	RET
GLOBL	·libc_getgid_trampoline_addr(SB), RODATA, $8
DATA	·libc_getgid_trampoline_addr(SB)/8, $libc_getgid_trampoline<>(SB)

TEXT libc_getpgid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getpgid(SB)
	RET
GLOBL	·libc_getpgid_trampoline_addr(SB), RODATA, $8
DATA	·libc_getpgid_trampoline_addr(SB)/8, $libc_getpgid_trampoline<>(SB)

TEXT libc_getpgrp_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getpgrp(SB)
	RET
GLOBL	·libc_getpgrp_trampoline_addr(SB), RODATA, $8
DATA	·libc_getpgrp_trampoline_addr(SB)/8, $libc_getpgrp_trampoline<>(SB)

TEXT libc_getpid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getpid(SB)
	RET
GLOBL	·libc_getpid_trampoline_addr(SB), RODATA, $8
DATA	·libc_getpid_trampoline_addr(SB)/8, $libc_getpid_trampoline<>(SB)

TEXT libc_getppid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getppid(SB)
	RET
GLOBL	·libc_getppid_trampoline_addr(SB), RODATA, $8
DATA	·libc_getppid_trampoline_addr(SB)/8, $libc_getppid_trampoline<>(SB)

TEXT libc_getpriority_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getpriority(SB)
	RET
GLOBL	·libc_getpriority_trampoline_addr(SB), RODATA, $8
DATA	·libc_getpriority_trampoline_addr(SB)/8, $libc_getpriority_trampoline<>(SB)

TEXT libc_getrlimit_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getrlimit(SB)
	RET
GLOBL	·libc_getrlimit_trampoline_addr(SB), RODATA, $8
DATA	·libc_getrlimit_trampoline_addr(SB)/8, $libc_getrlimit_trampoline<>(SB)

TEXT libc_getrtable_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getrtable(SB)
	RET
GLOBL	·libc_getrtable_trampoline_addr(SB), RODATA, $8
DATA	·libc_getrtable_trampoline_addr(SB)/8, $libc_getrtable_trampoline<>(SB)

TEXT libc_getrusage_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getrusage(SB)
	RET
GLOBL	·libc_getrusage_trampoline_addr(SB), RODATA, $8
DATA	·libc_getrusage_trampoline_addr(SB)/8, $libc_getrusage_trampoline<>(SB)

TEXT libc_getsid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getsid(SB)
	RET
GLOBL	·libc_getsid_trampoline_addr(SB), RODATA, $8
DATA	·libc_getsid_trampoline_addr(SB)/8, $libc_getsid_trampoline<>(SB)

TEXT libc_gettimeofday_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_gettimeofday(SB)
	RET
GLOBL	·libc_gettimeofday_trampoline_addr(SB), RODATA, $8
DATA	·libc_gettimeofday_trampoline_addr(SB)/8, $libc_gettimeofday_trampoline<>(SB)

TEXT libc_getuid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_getuid(SB)
	RET
GLOBL	·libc_getuid_trampoline_addr(SB), RODATA, $8
DATA	·libc_getuid_trampoline_addr(SB)/8, $libc_getuid_trampoline<>(SB)

TEXT libc_issetugid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_issetugid(SB)
	RET
GLOBL	·libc_issetugid_trampoline_addr(SB), RODATA, $8
DATA	·libc_issetugid_trampoline_addr(SB)/8, $libc_issetugid_trampoline<>(SB)

TEXT libc_kill_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_kill(SB)
	RET
GLOBL	·libc_kill_trampoline_addr(SB), RODATA, $8
DATA	·libc_kill_trampoline_addr(SB)/8, $libc_kill_trampoline<>(SB)

TEXT libc_kqueue_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_kqueue(SB)
	RET
GLOBL	·libc_kqueue_trampoline_addr(SB), RODATA, $8
DATA	·libc_kqueue_trampoline_addr(SB)/8, $libc_kqueue_trampoline<>(SB)

TEXT libc_lchown_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_lchown(SB)
	RET
GLOBL	·libc_lchown_trampoline_addr(SB), RODATA, $8
DATA	·libc_lchown_trampoline_addr(SB)/8, $libc_lchown_trampoline<>(SB)

TEXT libc_link_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_link(SB)
	RET
GLOBL	·libc_link_trampoline_addr(SB), RODATA, $8
DATA	·libc_link_trampoline_addr(SB)/8, $libc_link_trampoline<>(SB)

TEXT libc_linkat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_linkat(SB)
	RET
GLOBL	·libc_linkat_trampoline_addr(SB), RODATA, $8
DATA	·libc_linkat_trampoline_addr(SB)/8, $libc_linkat_trampoline<>(SB)

TEXT libc_listen_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_listen(SB)
	RET
GLOBL	·libc_listen_trampoline_addr(SB), RODATA, $8
DATA	·libc_listen_trampoline_addr(SB)/8, $libc_listen_trampoline<>(SB)

TEXT libc_lstat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_lstat(SB)
	RET
GLOBL	·libc_lstat_trampoline_addr(SB), RODATA, $8
DATA	·libc_lstat_trampoline_addr(SB)/8, $libc_lstat_trampoline<>(SB)

TEXT libc_mkdir_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mkdir(SB)
	RET
GLOBL	·libc_mkdir_trampoline_addr(SB), RODATA, $8
DATA	·libc_mkdir_trampoline_addr(SB)/8, $libc_mkdir_trampoline<>(SB)

TEXT libc_mkdirat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mkdirat(SB)
	RET
GLOBL	·libc_mkdirat_trampoline_addr(SB), RODATA, $8
DATA	·libc_mkdirat_trampoline_addr(SB)/8, $libc_mkdirat_trampoline<>(SB)

TEXT libc_mkfifo_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mkfifo(SB)
	RET
GLOBL	·libc_mkfifo_trampoline_addr(SB), RODATA, $8
DATA	·libc_mkfifo_trampoline_addr(SB)/8, $libc_mkfifo_trampoline<>(SB)

TEXT libc_mkfifoat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mkfifoat(SB)
	RET
GLOBL	·libc_mkfifoat_trampoline_addr(SB), RODATA, $8
DATA	·libc_mkfifoat_trampoline_addr(SB)/8, $libc_mkfifoat_trampoline<>(SB)

TEXT libc_mknod_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mknod(SB)
	RET
GLOBL	·libc_mknod_trampoline_addr(SB), RODATA, $8
DATA	·libc_mknod_trampoline_addr(SB)/8, $libc_mknod_trampoline<>(SB)

TEXT libc_mknodat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mknodat(SB)
	RET
GLOBL	·libc_mknodat_trampoline_addr(SB), RODATA, $8
DATA	·libc_mknodat_trampoline_addr(SB)/8, $libc_mknodat_trampoline<>(SB)

TEXT libc_nanosleep_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_nanosleep(SB)
	RET
GLOBL	·libc_nanosleep_trampoline_addr(SB), RODATA, $8
DATA	·libc_nanosleep_trampoline_addr(SB)/8, $libc_nanosleep_trampoline<>(SB)

TEXT libc_open_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_open(SB)
	RET
GLOBL	·libc_open_trampoline_addr(SB), RODATA, $8
DATA	·libc_open_trampoline_addr(SB)/8, $libc_open_trampoline<>(SB)

TEXT libc_openat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_openat(SB)
	RET
GLOBL	·libc_openat_trampoline_addr(SB), RODATA, $8
DATA	·libc_openat_trampoline_addr(SB)/8, $libc_openat_trampoline<>(SB)

TEXT libc_pathconf_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_pathconf(SB)
	RET
GLOBL	·libc_pathconf_trampoline_addr(SB), RODATA, $8
DATA	·libc_pathconf_trampoline_addr(SB)/8, $libc_pathconf_trampoline<>(SB)

TEXT libc_pread_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_pread(SB)
	RET
GLOBL	·libc_pread_trampoline_addr(SB), RODATA, $8
DATA	·libc_pread_trampoline_addr(SB)/8, $libc_pread_trampoline<>(SB)

TEXT libc_pwrite_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_pwrite(SB)
	RET
GLOBL	·libc_pwrite_trampoline_addr(SB), RODATA, $8
DATA	·libc_pwrite_trampoline_addr(SB)/8, $libc_pwrite_trampoline<>(SB)

TEXT libc_read_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_read(SB)
	RET
GLOBL	·libc_read_trampoline_addr(SB), RODATA, $8
DATA	·libc_read_trampoline_addr(SB)/8, $libc_read_trampoline<>(SB)

TEXT libc_readlink_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_readlink(SB)
	RET
GLOBL	·libc_readlink_trampoline_addr(SB), RODATA, $8
DATA	·libc_readlink_trampoline_addr(SB)/8, $libc_readlink_trampoline<>(SB)

TEXT libc_readlinkat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_readlinkat(SB)
	RET
GLOBL	·libc_readlinkat_trampoline_addr(SB), RODATA, $8
DATA	·libc_readlinkat_trampoline_addr(SB)/8, $libc_readlinkat_trampoline<>(SB)

TEXT libc_rename_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_rename(SB)
	RET
GLOBL	·libc_rename_trampoline_addr(SB), RODATA, $8
DATA	·libc_rename_trampoline_addr(SB)/8, $libc_rename_trampoline<>(SB)

TEXT libc_renameat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_renameat(SB)
	RET
GLOBL	·libc_renameat_trampoline_addr(SB), RODATA, $8
DATA	·libc_renameat_trampoline_addr(SB)/8, $libc_renameat_trampoline<>(SB)

TEXT libc_revoke_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_revoke(SB)
	RET
GLOBL	·libc_revoke_trampoline_addr(SB), RODATA, $8
DATA	·libc_revoke_trampoline_addr(SB)/8, $libc_revoke_trampoline<>(SB)

TEXT libc_rmdir_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_rmdir(SB)
	RET
GLOBL	·libc_rmdir_trampoline_addr(SB), RODATA, $8
DATA	·libc_rmdir_trampoline_addr(SB)/8, $libc_rmdir_trampoline<>(SB)

TEXT libc_lseek_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_lseek(SB)
	RET
GLOBL	·libc_lseek_trampoline_addr(SB), RODATA, $8
DATA	·libc_lseek_trampoline_addr(SB)/8, $libc_lseek_trampoline<>(SB)

TEXT libc_select_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_select(SB)
	RET
GLOBL	·libc_select_trampoline_addr(SB), RODATA, $8
DATA	·libc_select_trampoline_addr(SB)/8, $libc_select_trampoline<>(SB)

TEXT libc_setegid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setegid(SB)
	RET
GLOBL	·libc_setegid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setegid_trampoline_addr(SB)/8, $libc_setegid_trampoline<>(SB)

TEXT libc_seteuid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_seteuid(SB)
	RET
GLOBL	·libc_seteuid_trampoline_addr(SB), RODATA, $8
DATA	·libc_seteuid_trampoline_addr(SB)/8, $libc_seteuid_trampoline<>(SB)

TEXT libc_setgid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setgid(SB)
	RET
GLOBL	·libc_setgid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setgid_trampoline_addr(SB)/8, $libc_setgid_trampoline<>(SB)

TEXT libc_setlogin_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setlogin(SB)
	RET
GLOBL	·libc_setlogin_trampoline_addr(SB), RODATA, $8
DATA	·libc_setlogin_trampoline_addr(SB)/8, $libc_setlogin_trampoline<>(SB)

TEXT libc_setpgid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setpgid(SB)
	RET
GLOBL	·libc_setpgid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setpgid_trampoline_addr(SB)/8, $libc_setpgid_trampoline<>(SB)

TEXT libc_setpriority_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setpriority(SB)
	RET
GLOBL	·libc_setpriority_trampoline_addr(SB), RODATA, $8
DATA	·libc_setpriority_trampoline_addr(SB)/8, $libc_setpriority_trampoline<>(SB)

TEXT libc_setregid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setregid(SB)
	RET
GLOBL	·libc_setregid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setregid_trampoline_addr(SB)/8, $libc_setregid_trampoline<>(SB)

TEXT libc_setreuid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setreuid(SB)
	RET
GLOBL	·libc_setreuid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setreuid_trampoline_addr(SB)/8, $libc_setreuid_trampoline<>(SB)

TEXT libc_setresgid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setresgid(SB)
	RET
GLOBL	·libc_setresgid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setresgid_trampoline_addr(SB)/8, $libc_setresgid_trampoline<>(SB)

TEXT libc_setresuid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setresuid(SB)
	RET
GLOBL	·libc_setresuid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setresuid_trampoline_addr(SB)/8, $libc_setresuid_trampoline<>(SB)

TEXT libc_setrlimit_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setrlimit(SB)
	RET
GLOBL	·libc_setrlimit_trampoline_addr(SB), RODATA, $8
DATA	·libc_setrlimit_trampoline_addr(SB)/8, $libc_setrlimit_trampoline<>(SB)

TEXT libc_setrtable_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setrtable(SB)
	RET
GLOBL	·libc_setrtable_trampoline_addr(SB), RODATA, $8
DATA	·libc_setrtable_trampoline_addr(SB)/8, $libc_setrtable_trampoline<>(SB)

TEXT libc_setsid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setsid(SB)
	RET
GLOBL	·libc_setsid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setsid_trampoline_addr(SB)/8, $libc_setsid_trampoline<>(SB)

TEXT libc_settimeofday_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_settimeofday(SB)
	RET
GLOBL	·libc_settimeofday_trampoline_addr(SB), RODATA, $8
DATA	·libc_settimeofday_trampoline_addr(SB)/8, $libc_settimeofday_trampoline<>(SB)

TEXT libc_setuid_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_setuid(SB)
	RET
GLOBL	·libc_setuid_trampoline_addr(SB), RODATA, $8
DATA	·libc_setuid_trampoline_addr(SB)/8, $libc_setuid_trampoline<>(SB)

TEXT libc_stat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_stat(SB)
	RET
GLOBL	·libc_stat_trampoline_addr(SB), RODATA, $8
DATA	·libc_stat_trampoline_addr(SB)/8, $libc_stat_trampoline<>(SB)

TEXT libc_statfs_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_statfs(SB)
	RET
GLOBL	·libc_statfs_trampoline_addr(SB), RODATA, $8
DATA	·libc_statfs_trampoline_addr(SB)/8, $libc_statfs_trampoline<>(SB)

TEXT libc_symlink_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_symlink(SB)
	RET
GLOBL	·libc_symlink_trampoline_addr(SB), RODATA, $8
DATA	·libc_symlink_trampoline_addr(SB)/8, $libc_symlink_trampoline<>(SB)

TEXT libc_symlinkat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_symlinkat(SB)
	RET
GLOBL	·libc_symlinkat_trampoline_addr(SB), RODATA, $8
DATA	·libc_symlinkat_trampoline_addr(SB)/8, $libc_symlinkat_trampoline<>(SB)

TEXT libc_sync_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_sync(SB)
	RET
GLOBL	·libc_sync_trampoline_addr(SB), RODATA, $8
DATA	·libc_sync_trampoline_addr(SB)/8, $libc_sync_trampoline<>(SB)

TEXT libc_truncate_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_truncate(SB)
	RET
GLOBL	·libc_truncate_trampoline_addr(SB), RODATA, $8
DATA	·libc_truncate_trampoline_addr(SB)/8, $libc_truncate_trampoline<>(SB)

TEXT libc_umask_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_umask(SB)
	RET
GLOBL	·libc_umask_trampoline_addr(SB), RODATA, $8
DATA	·libc_umask_trampoline_addr(SB)/8, $libc_umask_trampoline<>(SB)

TEXT libc_unlink_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_unlink(SB)
	RET
GLOBL	·libc_unlink_trampoline_addr(SB), RODATA, $8
DATA	·libc_unlink_trampoline_addr(SB)/8, $libc_unlink_trampoline<>(SB)

TEXT libc_unlinkat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_unlinkat(SB)
	RET
GLOBL	·libc_unlinkat_trampoline_addr(SB), RODATA, $8
DATA	·libc_unlinkat_trampoline_addr(SB)/8, $libc_unlinkat_trampoline<>(SB)

TEXT libc_unmount_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_unmount(SB)
	RET
GLOBL	·libc_unmount_trampoline_addr(SB), RODATA, $8
DATA	·libc_unmount_trampoline_addr(SB)/8, $libc_unmount_trampoline<>(SB)

TEXT libc_write_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_write(SB)
	RET
GLOBL	·libc_write_trampoline_addr(SB), RODATA, $8
DATA	·libc_write_trampoline_addr(SB)/8, $libc_write_trampoline<>(SB)

TEXT libc_mmap_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_mmap(SB)
	RET
GLOBL	·libc_mmap_trampoline_addr(SB), RODATA, $8
DATA	·libc_mmap_trampoline_addr(SB)/8, $libc_mmap_trampoline<>(SB)

TEXT libc_munmap_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_munmap(SB)
	RET
GLOBL	·libc_munmap_trampoline_addr(SB), RODATA, $8
DATA	·libc_munmap_trampoline_addr(SB)/8, $libc_munmap_trampoline<>(SB)

TEXT libc_utimensat_trampoline<>(SB),NOSPLIT,$0-0
	CALL	libc_utimensat(SB)
	RET
GLOBL	·libc_utimensat_trampoline_addr(SB), RODATA, $8
DATA	·libc_utimensat_trampoline_addr(SB)/8, $libc_utimensat_trampoline<>(SB)
