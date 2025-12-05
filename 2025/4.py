import copy

grid = []
with open("4.txt") as f:
    for line in f:
        grid.append(list(line.strip()))

grid_copy = copy.deepcopy(grid)
dirs = [(0,1),(1,0),(0,-1),(-1,0),(1,1),(1,-1),(-1,1),(-1,-1)]

n_rolls = 0
diff = -1
while diff != 0:
    diff = 0
    for r in range(len(grid)):
        for c in range(len(grid[r])):
            if grid[r][c] != "@":
                continue
            count = 0
            for dr,dc in dirs:
                check_r = dr + r
                check_c = dc + c
                if 0 <= check_r < len(grid) and 0 <= check_c < len(grid[r]):
                    if grid[check_r][check_c] == "@":
                        count += 1
            if count < 4:
                diff += 1
                grid[r][c] = "."
                # for row in grid:
                #     print(row)
                # print()
    n_rolls += diff


print(f"part 1 answer: {n_rolls}")

