filename = "7-test.txt"


grid = []
with open(filename, "r") as f:
    for line in f:
        grid.append(list(line.strip()))

visited = {}  # r,c -> count

s_index = len(grid[0]) // 2


def check_bounds(r, c):
    if r < 0 or c < 0:
        return False
    if r >= len(grid) or c >= len(grid[r]):
        return False
    return True


def helper(r, c):
    # print(f"r: {r} c: {c}")
    # for row in grid:
    #     print(row)
    # print()
    if (r, c) in visited:
        return visited[(r, c)]
    # always 1 path at the end
    if r == len(grid):
        visited[(r, c)] = 1
        return 1

    # OOB
    if not check_bounds(r, c):
        return 0

    # if splitter, go left and right with backtracking
    if grid[r][c] == "^" and grid[r - 1][c] == "|":

        # go left
        # if check_bounds(r + 1, c - 1):
        grid[r + 1][c - 1] = "|"
        l_count = helper(r + 1, c - 1)
        grid[r + 1][c - 1] = "."

        # go right
        grid[r + 1][c + 1] = "|"
        r_count = helper(r + 1, c + 1)
        grid[r + 1][c + 1] = "."
        visited[(r, c)] = l_count + r_count
        return l_count + r_count

    # continue a path
    elif grid[r][c] == "." and grid[r - 1][c] in ["|", "S"]:
        grid[r][c] = "|"
        count = helper(r + 1, c)
        grid[r][c] = "."
        return count

    elif grid[r][c] == "|":
        return helper(r + 1, c)

    else:
        print(f"r: {r} c: {c} val: {grid[r][c]}")
        print(f"HERE ?")
        return 0


ans = helper(1, s_index)
print(f"part 2 ans: {ans}")
