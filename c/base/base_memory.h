
#ifndef BASE_MEMORY_H
#define BASE_MEMORY_H

///////////////////////////////
/// NOTE(hamed): Memory V-Table

typedef void* M_ReserveFunc(void *ctx, U64 size);
typedef void  M_ChangeMemoryFunc(void *ctx, void *ptr, U64 size);

struct M_BaseMemory{
    M_ReserveFunc *reserve;
    M_ChangeMemoryFunc *commit;
    M_ChangeMemoryFunc *decommit;
    M_ChangeMemoryFunc *release;
    void *ctx;
};

function void m_change_memory_noop(void *ctx, void *ptr, U64 size);

#endif // BASE_MEMORY_H
