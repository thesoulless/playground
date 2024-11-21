function struct String8
str8_substr(struct String8 str, U64 first, U64 opl) {
    struct String8 result = {0};
    result.str = str.str + first;
    result.size = opl;
    return result;
}

function struct String8
str8(U8 *str, U64 size) {
    struct String8 result = {0};
    result.str = str;
    result.size = size;
    return result;
}
