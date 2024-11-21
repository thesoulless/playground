# if defined (_WIN32)
# define OS_WIN 1
# elif defined (__gnu_linux__)
# define OS_LINUX 1
# elif defined (__APPLE__) && defined (__MACH__)
# define OS_MAC 1
# else
# error "Unknown OS"
# endif

# if defined (__amd64__)
# define ARCH64 1
# define ARCH_X64 1
# elif defined (__i386__)
# define ARCH32 1
# define ARCH_X86 1
# elif defined (__arm__)
# define ARCH_ARM 1
# elif defined (__aarch64__)
# define ARCH64 1
# define ARCH_ARM64 1
# else
# error "Unknown ARCH"
# endif


# if !defined(ARCH_ARM64)
# define ARCH_ARM64 0
# endif
# if !defined(ARCH_ARM)
# define ARCH_ARM 0
# endif
# if !defined(ARCH_X86)
# define ARCH_X86 0
# endif
# if !defined(ARCH_X64)
# define ARCH_X64 0
# endif
# if !defined(ARCH32)
# define ARCH32 0
# endif
# if !defined(ARCH64)
# define ARCH64 0
# endif
# if !defined(OS_WIN)
# define OS_WIN 0
# endif
# if !defined(OS_LINUX)
# define OS_LINUX 0
# endif
# if !defined(OS_MAC)
# define OS_MAC 0
# endif

# if !defined(ENABLE_ASSERT)
# define ENABLE_ASSERT
# endif

# define Stmnt(x) do { x } while (0)

# if !defined(AssertBreak)
# define AssertBreak() (*(int *)0 = 0)
# endif

# if defined(ENABLE_ASSERT)
# define Assert(x) Stmnt( if (!(x)) { AssertBreak(); } )
# else
# define Assert(x)
# endif

#define Stringify_(s) #s
#define Stringify(s) Stringify_(s)
#define Concat_(a, b) a##b
#define Concat(a, b) Concat_(a, b)

#define ArrayCount(a) (sizeof(a) / sizeof(*(a)))

#define IntFromPtr(p) (unsigned long long)((char*)p - (char*)0)
#define PtrFromInt(i) (void*)((char*)0 + i)

#define Member(T,m) (((T*)0)->m)
#define OffsetOfMember(T,m) IntFromPtr(&Member(T,m))

#define global static
#define local static
#define function static

#define c_linkage_begin extern "C" {
#define c_linkage_end }
#define c_linkage extern "C"

#define EvalPrint(x) printf(#x " = %d\n", x)
#define EvalPrintF(x) printf(#x " = %f\n", x)

////////////////////////////////////////
// NOTE(hamed): Basic types

#include <stdint.h>
typedef uint64_t U64;
typedef uint8_t U8;
typedef int8_t S8;

typedef void VoidFunc(void);
typedef float F32;
typedef double F64;

global F32 machine_epsilon_F32 = 1.19209290e-07f;

union V3F32 {
    struct {
        F32 x, y, z;
    };
    F32 v[3];
};

function F32 vec_dot(union V3F32 a, union V3F32 b);

////////////////////////////////////////
// NOTE(hamed): Consts

enum Month {
    Month_Jan,
    Month_Feb,
    Month_Mar,
    Month_Apr,
    Month_May,
    Month_Jun,
    Month_Jul,
    Month_Aug,
    Month_Sep,
    Month_Oct,
    Month_Nov,
    Month_Dec,
};

enum DayOfWeek {
    DayOfWeek_Sun,
    DAYOfWeek_Mon,
    DAYOfWeek_Tue,
    DAYOfWeek_Wed,
    DAYOfWeek_Thu,
    DAYOfWeek_Fri,
    DAYOfWeek_Sat,
};

function char* string_from_month(enum Month m);
function char* string_from_day_of_week(enum DayOfWeek d);


////////////////////////////////////////
// NOTE(hamed): Linked List
// 
// Doubly linked list
//   PushBack
//   PushFront
//   Remove
// Singly linked list queue
//   PushBack
//   PushFront
//   Pop
// Singly linked list stack
//   Push
//   Pop
#define DLLPushBack_NP(f,l,n,next,prev) ((f)==0?\
             ((f)=(l)=(n),(n)->next=(n)->prev=0):\
             ((n)->prev=(l),(l)->next=(n),(l)=(n),(n)->next=0))
#define DLLPushBack(f,l,n) DLLPushBack_NP(f,l,n,next,prev)
#define DLLPushFront(f,l,n) DLLPushBack_NP(l,f,n,prev,next)
#define DLLRemove_NP(f,l,n,next,prev) (((f)==(n)?\
            ((f)=(f)->next,(f)->prev=0):\
            (l)==(n)?\
            ((l)=(l)->prev,(l)->next=0):\
            ((n)->next->prev=(n)->prev,\
             (n)->prev->next=(n)->next)))
#define DLLRemove(f,l,n) DLLRemove_NP(f,l,n,next,prev)


#define SLLQueuePush_N(f,l,n,next) ((f)==0?\
        (f)=(l)=(n):\
        ((l)->next=(n),(l)=(n)),\
        (n)->next=0)
#define SLLQueuePush(f,l,n) SLLQueuePush_N(f,l,n,next)
#define SLLQueuePushFront_N(f,l,n,next) ((f)==0?\
        ((f)=(l)=(n),(n)->next=0):\
        ((n)->next=(f),(f)=(n)))
#define SLLQueuePushFront(f,l,n) SLLQueuePushFront_N(f,l,n,next)
#define SLLQueuePop_N(f,l,next) ((f)==(l)?\
        ((f)=(l)=0):\
        ((f)=(f)->next))
#define SLLQueuePop(f,l) SLLQueuePop_N(f,l,next)


#define SLLStackPush_N(f,n,next) ((n)->next=(f),(f)=(n))
#define SLLStackPush(f,n) SLLStackPush_N(f,n,next)
#define SLLStackPop_N(f,next) ((f)==0?0:\
        (f)=(f)->next)
#define SLLStackPop(f) SLLStackPop_N(f,next)
