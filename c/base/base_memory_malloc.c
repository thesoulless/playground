// #include <stdlib.h>

function void*
m_malloc_reserve(void *ctx, U64 size){
    return(malloc(size));
}

function void
m_malloc_commit(void *ctx, void *ptr, U64 size){
    // Do nothing
}

function void
m_malloc_decommit(void *ctx, void *ptr, U64 size){
    // Do nothing
}

function void
m_malloc_release(void *ctx, void *ptr, U64 size){
    free(ptr);
}

function struct M_BaseMemory*
m_malloc_base_memory(void){
    local struct M_BaseMemory memory = {};
    if (memory.reserve == 0){
        memory.reserve = m_malloc_reserve;
        memory.commit = m_malloc_commit;
        memory.decommit = m_malloc_decommit;
        memory.release = m_malloc_release;
    }
    return(&memory);
}
