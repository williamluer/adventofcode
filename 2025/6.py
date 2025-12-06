import operator

filename = "6.txt"

OPS = {"*": operator.mul, "/": operator.truediv, "+": operator.add, "-": operator.sub}
problems = []
with open(filename, "r") as f:
    for i, line in enumerate(f):
        line = line.strip().split()
        for j, val in enumerate(line):
            if i == 0:
                problems.append([val])
            else:
                problems[j].append(val)

# print(problems)

tot_sum = 0
for problem in problems:
    op = OPS[problem[-1]]
    ans = int(problem[0])
    for val in problem[1:-1]:
        ans = op(ans, int(val))
    tot_sum += ans

print(f"part 1:  {tot_sum}")


# ------------------------------------------------------------

OPS = {"*": operator.mul, "/": operator.truediv, "+": operator.add, "-": operator.sub}
grid = []
with open(filename, "r") as f:
    for i, line in enumerate(f):
        arr = []
        for val in line.replace("\n", ""):
            arr.append(val)
        grid.append(arr)

stack = []
total_sum = 0
for c in range(len(grid[0]) - 1, -1, -1):
    # print(stack)
    working_val = ""
    for r in range(len(grid)):
        # print(f"r: {r} c: {c} val: {grid[r][c]}")
        if grid[r][c] in OPS:
            # print(f"Working val here: {working_val}")
            op = OPS[grid[r][c]]
            stack.append(int(working_val))
            ans = int(stack[0])
            for val in stack[1:]:
                ans = op(ans, int(val))
            total_sum += ans
            # do math
            stack = []
            working_val = ""
        elif grid[r][c].isdigit():
            working_val += grid[r][c]
    if working_val.isdigit():
        # print(f"Appending {working_val}")
        stack.append(int(working_val))
    # else:
    #     print(f"not appending: {working_val}")
    # print()
print(f"part 2 ans: {total_sum}")
