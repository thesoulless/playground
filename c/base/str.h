
#ifndef STR_H
#define STR_H

struct String8 {
    U8 *str;
    U64 size;
};

struct String8Node {
    struct String8Node *next;
    struct String8 str;
};

struct String8List {
    struct String8Node *first;
    struct String8Node *last;
    U64 node_count;
    U64 total_size;
};

function struct String8 str8(U8 *str, U64 size);

#define str8_lit(s) str8((U8*)(s), sizeof(s) - 1)

function struct String8 str8_substr(struct String8 str, U64 first, U64 opl);

#endif // STR_H
