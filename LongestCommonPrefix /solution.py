def longestCommonPrefix(strs: list):
    prefix = []
    for x in zip(*strs):
        print(x)
        if len(set(x)) == 1:
            prefix.append(x[0])
        else:
            break
    return "".join(prefix)

