func strStr(haystack string, needle string) int {
    if !strings.Contains(haystack, needle) {
        return -1
    }
    return strings.Index(haystack, needle)
}