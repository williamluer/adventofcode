import heapq

def largest_two_digits(vals: list[int]) -> int:
    max_to_n = vals[-1]
    maxes = [vals[-1]]
    for n in reversed(vals[:-1]):
        maxes.append(max(maxes[-1], n))
    maxes.reverse()

    two_largest = heapq.nlargest(2, vals)

    max_val = 0
    for option in two_largest:
        l = option
        i = vals.index(l)
        if i < len(vals) - 1:
            r = maxes[i+1]
            max_val = max(max_val, l * 10 + r)

    return max_val


def largest_n_digits(vals: list[int], n: int) -> int:

    # find largest digit in first [:len-n]
    # for i in range(n):
    # then find largest digit in first :len-n+1

    ans = []
    last_index = -1
    print()
    print()
    print("\n\n\n")
    print(vals)
    for i in range(n):
        print(f"Looking at: {vals[last_index+1:len(vals)-n+i+1]}")
        highest_val = max(vals[last_index+1:len(vals)-n+i+1])
        print(f"found highest val: {highest_val}")
        index_of_highest = vals.index(highest_val, last_index+1, len(vals)-n+i+1)
        print(f"index is: {index_of_highest}")
        last_index = index_of_highest
        ans.append(highest_val)
        print(ans)
        print()
    print(ans)

    result = 0
    for i, val in enumerate(ans):
       result += val * 10 ** (n-i-1)

    return result



with open("3.txt", "r") as f:
    ans = 0
    for line in f:
        vals = [int(n) for n in line.strip()]
        joltage = largest_two_digits(vals)
        ans += joltage
    print(f"part 1: {ans}")

with open("3.txt", "r") as f:
    ans = 0
    for line in f:
        vals = [int(n) for n in line.strip()]
        joltage = largest_n_digits(vals, 12)
        ans += joltage
    print(f"part 1: {ans}")


