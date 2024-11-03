def binary_search(haystack: list[int], needle: int) -> int:
    """
    Performs binary search on a sorted array
    Returns index of needle if found, -1 otherwise
    """
    low, high = 0, len(haystack) - 1

    while low <= high:
        middle = (low + high) // 2
        value = haystack[middle]

        if value == needle:
            return True
        elif value < needle:
            low = middle + 1
        else:
            high = middle - 1

    return False
