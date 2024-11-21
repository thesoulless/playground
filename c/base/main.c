#include "base_inc.h"
#include <stdio.h>
#include <stdlib.h>
#include "base_inc.c"
#include "base_memory.h"

struct Node {
    struct Node *next;
    struct Node *prev;

    int x;
};

int main() {


  printf("arm64 = %d\n", ARCH_ARM64);
  printf("arm = %d\n", ARCH_ARM);
  printf("x86 = %d\n", ARCH_X86);
  printf("x64 = %d\n", ARCH_X64);
  printf("32 = %d\n", ARCH32);
  printf("64 = %d\n", ARCH64);

  // os
  printf("win = %d\n", OS_WIN);
  printf("linux = %d\n", OS_LINUX);
  printf("mac = %d\n", OS_MAC);

  Assert(10 < 100);

  /*int i = 10;

  printf("i = %d\n", i);
  printf("i PTR = %d\n", PtrFromInt(i));
  printf("i Int = %d\n", IntFromPtr(&i));*/

  union V3F32 v = {1, 2, 3};
  for (int i = 0; i < 3; i++) {
    EvalPrintF(v.v[i]);
  }

  printf("month = %s\n", string_from_month(Month_Jan));
  printf("day = %s\n", string_from_day_of_week(DayOfWeek_Sun));

  struct M_BaseMemory *base_memory = m_malloc_base_memory();
  // struct Node nodes[10];

  U64 node_count = 10;
  struct Node *nodes = (struct Node*)base_memory->reserve(base_memory->ctx, sizeof(struct Node) * node_count);
  for (int i = 0; i < node_count; i+=1) {
    nodes[i].x = i;
  }

  {
      printf("DLL\n");
      struct Node *first = 0;
      struct Node *last = 0;
      for (int i = 0; i < 5; i+=1) {
          DLLPushBack(first, last, &nodes[i]);
      }
      for (int i = 5; i < 10; i+=1) {
          DLLPushFront(first, last, &nodes[i]);
      }

      for (struct Node *node = first; node != 0; node = node->next) {
          EvalPrint(node->x);
      }
      printf("\n");
  }

  {
      printf("Queue\n");
      struct Node *first = 0;
      struct Node *last = 0;
      for (int i = 0; i < 5; i+=1) {
          SLLQueuePush(first, last, &nodes[i]);
      }

      for (int i = 5; i < 10; i+=1) {
          SLLQueuePushFront(first, last, &nodes[i]);
      }

      SLLQueuePop(first, last);
      SLLQueuePop(first, last);
      for (struct Node *node = first; node != 0; node = node->next) {
          EvalPrint(node->x);
      }
      printf("\n");
  }

  {
      printf("Stack\n");

      struct Node *first = 0;
      struct Node *last = 0;
      for (int i = 0; i < 5; i+=1) {
          SLLStackPush(first, &nodes[i]);
      }

      for (struct Node *node = first; node != 0; node = node->next) {
          EvalPrint(node->x);
      }

      printf("\n");
  }

  /* struct String8 str = str8_lit("Hello World!"); */
  struct String8 str = str8_lit("سلام دنیا!");
  /* printf("str = %.*s, len = %d\n", (int)str.size, str.str, (int)str.size); */
  printf("str = %.*s\n", (int)str.size, str.str);

  return (0);
}
