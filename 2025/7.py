filename = "7.txt"


grid = []
with open(filename, "r") as f:
    for line in f:
        grid.append(list(line.strip()))

visited = ()
count = 0

s_index = len(grid[0]) // 2
grid[1][s_index] = "|"


def check_bounds(r, c):
    if r < 0 or c < 0:
        return False
    if r >= len(grid) or c >= len(grid[r]):
        return False
    return True


for r in range(1, len(grid)):
    for c in range(len(grid[r])):
        should_count = False
        if grid[r][c] == "^":
            if grid[r - 1][c] in ["|", "S"]:
                if check_bounds(r + 1, c - 1):
                    grid[r + 1][c - 1] = "|"
                    should_count = True
                if check_bounds(r - 1, c + 1):
                    grid[r + 1][c + 1] = "|"
                    should_count = True
                grid[r][c] = str(count % 10)
        elif grid[r - 1][c] == "|":
            grid[r][c] = "|"
        if should_count:
            count += 1
    # for row in grid:
    #     print(row)
    # print()
print(f"Part 1 answer: {count}")


def helper(r, c):
    if not check_bounds(r,c):
        return 0
    
    if grid[r][c] == "^" and grid[r - 1][c] == "|":
        if check_bounds(r + 1, c - 1):
            return 1 + helper(r + 1, c - 1)
    elif grid[r-1][c] == "|"
        if check_bounds(r + 1, c - 1):
            return helper(r + 1, c - 1)
    if 