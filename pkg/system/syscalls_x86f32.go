package system

const (
	SyscallX86MaxNum32   = 358
	SyscallX86LastName32 = "execveat"
)

// line numbers are aligned with the syscall number (-10)
var syscallNumTableX86Family32 = [...]string{
	"restart_syscall",
	"_exit",
	"fork",
	"read",
	"write",
	"open",
	"close",
	"waitpid",
	"creat",
	"link",
	"unlink",
	"execve",
	"chdir",
	"time",
	"mknod",
	"chmod",
	"lchown",
	"break",
	"oldstat",
	"lseek",
	"getpid",
	"mount",
	"umount",
	"setuid",
	"getuid",
	"stime",
	"ptrace",
	"alarm",
	"oldfstat",
	"pause",
	"utime",
	"stty",
	"gtty",
	"access",
	"nice",
	"ftime",
	"sync",
	"kill",
	"rename",
	"mkdir",
	"rmdir",
	"dup",
	"pipe",
	"times",
	"prof",
	"brk",
	"setgid",
	"getgid",
	"signal",
	"geteuid",
	"getegid",
	"acct",
	"umount2",
	"lock",
	"ioctl",
	"fcntl",
	"mpx",
	"setpgid",
	"ulimit",
	"oldolduname",
	"umask",
	"chroot",
	"ustat",
	"dup2",
	"getppid",
	"getpgrp",
	"setsid",
	"sigaction",
	"sgetmask",
	"ssetmask",
	"setreuid",
	"setregid",
	"sigsuspend",
	"sigpending",
	"sethostname",
	"setrlimit",
	"getrlimit",
	"getrusage",
	"gettimeofday",
	"settimeofday",
	"getgroups",
	"setgroups",
	"oldselect",
	"symlink",
	"oldlstat",
	"readlink",
	"uselib",
	"swapon",
	"reboot",
	"readdir",
	"old_mmap",
	"munmap",
	"truncate",
	"ftruncate",
	"fchmod",
	"fchown",
	"getpriority",
	"setpriority",
	"profil",
	"statfs",
	"fstatfs",
	"ioperm",
	"socketcall",
	"syslog",
	"setitimer",
	"getitimer",
	"stat",
	"lstat",
	"fstat",
	"olduname",
	"iopl",
	"vhangup",
	"idle",
	"vm86old",
	"wait4",
	"swapoff",
	"sysinfo",
	"ipc",
	"fsync",
	"sigreturn",
	"clone",
	"setdomainname",
	"uname",
	"modify_ldt",
	"adjtimex",
	"mprotect",
	"sigprocmask",
	"create_module",
	"init_module",
	"delete_module",
	"get_kernel_syms",
	"quotactl",
	"getpgid",
	"fchdir",
	"bdflush",
	"sysfs",
	"personality",
	"afs_syscall",
	"setfsuid",
	"setfsgid",
	"_llseek",
	"getdents",
	"select",
	"flock",
	"msync",
	"readv",
	"writev",
	"getsid",
	"fdatasync",
	"_sysctl",
	"mlock",
	"munlock",
	"mlockall",
	"munlockall",
	"sched_setparam",
	"sched_getparam",
	"sched_setscheduler",
	"sched_getscheduler",
	"sched_yield",
	"sched_get_priority_max",
	"sched_get_priority_min",
	"sched_rr_get_interval",
	"nanosleep",
	"mremap",
	"setresuid",
	"getresuid",
	"vm86",
	"query_module",
	"poll",
	"nfsservctl",
	"setresgid",
	"getresgid",
	"prctl",
	"rt_sigreturn",
	"rt_sigaction",
	"rt_sigprocmask",
	"rt_sigpending",
	"rt_sigtimedwait",
	"rt_sigqueueinfo",
	"rt_sigsuspend",
	"pread64",
	"pwrite64",
	"chown",
	"getcwd",
	"capget",
	"capset",
	"sigaltstack",
	"sendfile",
	"getpmsg",
	"putpmsg",
	"vfork",
	"ugetrlimit",
	"mmap2",
	"truncate64",
	"ftruncate64",
	"stat64",
	"lstat64",
	"fstat64",
	"lchown32",
	"getuid32",
	"getgid32",
	"geteuid32",
	"getegid32",
	"setreuid32",
	"setregid32",
	"getgroups32",
	"setgroups32",
	"fchown32",
	"setresuid32",
	"getresuid32",
	"setresgid32",
	"getresgid32",
	"chown32",
	"setuid32",
	"setgid32",
	"setfsuid32",
	"setfsgid32",
	"pivot_root",
	"mincore",
	"madvise",
	"getdents64",
	"fcntl64",
	"unused.222", //222 - unused
	"security",   //223 - supposed to be unused too
	"gettid",
	"readahead",
	"setxattr",
	"lsetxattr",
	"fsetxattr",
	"getxattr",
	"lgetxattr",
	"fgetxattr",
	"listxattr",
	"llistxattr",
	"flistxattr",
	"removexattr",
	"lremovexattr",
	"fremovexattr",
	"tkill",
	"sendfile64",
	"futex",
	"sched_setaffinity",
	"sched_getaffinity",
	"set_thread_area",
	"get_thread_area",
	"io_setup",
	"io_destroy",
	"io_getevents",
	"io_submit",
	"io_cancel",
	"fadvise64",
	"unused.251", //251 - unused (available for reuse) was used by sys_set_zone_reclaim
	"exit_group",
	"lookup_dcookie",
	"epoll_create",
	"epoll_ctl",
	"epoll_wait",
	"remap_file_pages",
	"set_tid_address",
	"timer_create",
	"timer_settime",
	"timer_gettime",
	"timer_getoverrun",
	"timer_delete",
	"clock_settime",
	"clock_gettime",
	"clock_getres",
	"clock_nanosleep",
	"statfs64",
	"fstatfs64",
	"tgkill",
	"utimes",
	"fadvise64_64",
	"vserver",
	"mbind",
	"get_mempolicy",
	"set_mempolicy",
	"mq_open",
	"mq_unlink",
	"mq_timedsend",
	"mq_timedreceive",
	"mq_notify",
	"mq_getsetattr",
	"kexec_load",
	"waitid",
	"setaltroot", //285 - sys_setaltroot (sort of used :))
	"add_key",
	"request_key",
	"keyctl",
	"ioprio_set",
	"ioprio_get",
	"inotify_init",
	"inotify_add_watch",
	"inotify_rm_watch",
	"migrate_pages",
	"openat",
	"mkdirat",
	"mknodat",
	"fchownat",
	"futimesat",
	"fstatat64",
	"unlinkat",
	"renameat",
	"linkat",
	"symlinkat",
	"readlinkat",
	"fchmodat",
	"faccessat",
	"pselect6",
	"ppoll",
	"unshare",
	"set_robust_list",
	"get_robust_list",
	"splice",
	"sync_file_range",
	"tee",
	"vmsplice",
	"move_pages",
	"getcpu",
	"epoll_pwait",
	"utimensat",
	"signalfd",
	"timerfd_create",
	"eventfd",
	"fallocate",
	"timerfd_settime",
	"timerfd_gettime",
	"signalfd4",
	"eventfd2",
	"epoll_create1",
	"dup3",
	"pipe2",
	"inotify_init1",
	"preadv",
	"pwritev",
	"rt_tgsigqueueinfo",
	"perf_event_open",
	"recvmmsg",
	"fanotify_init",
	"fanotify_mark",
	"prlimit64",
	"name_to_handle_at",
	"open_by_handle_at",
	"clock_adjtime",
	"syncfs",
	"sendmmsg",
	"setns",
	"process_vm_readv",
	"process_vm_writev",
	"kcmp",
	"finit_module",
	"sched_setattr",
	"sched_getattr",
	"renameat2",
	"seccomp",
	"getrandom",
	"memfd_create",
	"bpf",
	"execveat",
}

func callNameX86Family32(num uint32) string {
	if num > SyscallX86MaxNum32 {
		return SyscallX86UnknownName
	}

	return syscallNumTableX86Family32[num]
}

func callNumTableIsOkX86Family32() bool {
	if (len(syscallNumTableX86Family32) == SyscallX86MaxNum32+1) &&
		syscallNumTableX86Family32[SyscallX86MaxNum32] == SyscallX86LastName32 {
		return true
	}

	return false
}

func callNumberX86Family32(name string) (uint32, bool) {
	num, ok := syscallNameTableX86Family32[name]
	return num, ok
}

var syscallNameTableX86Family32 map[string]uint32

func init() {
	syscallNameTableX86Family32 = make(map[string]uint32, len(syscallNumTableX86Family32))

	for callNum, callName := range syscallNumTableX86Family32 {
		syscallNameTableX86Family32[callName] = uint32(callNum)
	}
}
