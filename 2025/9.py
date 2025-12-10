points = []
distances = []
with open("9-test.txt", "r") as f:
    for line in f:
        l = line.strip().split(",")
        points.append([int(x) for x in l])

# print(points)


def compute_area(p1, p2):
    return abs(p1[0] - p2[0] + 1) * abs(p1[1] - p2[1] + 1)


# PART 1
max_x = 0
max_y = 0
min_x = float("inf")
min_y = float("inf")

for i in range(len(points)):
    max_x = max(max_x, points[i][0])
    max_y = max(max_y, points[i][1])
    min_x = min(min_x, points[i][0])
    min_y = min(min_y, points[i][1])

    for j in range(i + 1, len(points)):
        distances.append([compute_area(points[i], points[j]), points[i], points[j]])

distances = sorted(distances, key=lambda x: x[0], reverse=True)
print()
print(distances[:3])

print(f"x: {min_x, max_x} and y: {min_y, max_y}")

tiles = set()
tiles_l = []

lx, ly = points[0]
tiles.add((lx, ly))
points.append(points[0])
for i in range(1, len(points)):
    cx, cy = points[i]
    tiles.add((cx, cy))
    for x in range(cx, lx):
        tiles.add((x, cy))
    for x in range(cx, lx, -1):
        tiles.add((x, cy))
    for y in range(cy, ly):
        tiles.add((cx, y))
    for y in range(cy, ly, -1):
        tiles.add((cx, y))
    lx, ly = cx, cy
# print(tiles_l)

# for row in range(min_y, 7 + 1):
#     for col in range(min_x, 11 + 1):
#         if [col, row] in points:
#             print("#", end="")
#         elif (col, row) in tiles:
#             print("X", end="")
#         else:
#             print(".", end="")
#     print()

from collections import deque

visited_overall = set()
not_tiles = set()

for row in range(min_y, max_y + 1):
    for col in range(min_x, max_x + 1):
        if (
            (col, row) in tiles
            or (col, row) in visited_overall
            or (col, row) in not_tiles
        ):
            continue
        q = deque([(col, row)])
        region = set()
        touches_wall = False
        while q:
            x, y = q.popleft()
            if (
                (x, y) in tiles
                or (x, y) in visited_overall
                or (x, y) in not_tiles
                or (x, y) in region
            ):
                continue
            region.add((x, y))
            if x == min_x or x == max_x or y == min_y or y == max_y:
                touches_wall = True
            for dx, dy in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
                nx, ny = x + dx, y + dy
                if min_x <= nx <= max_x and min_y <= ny <= max_y:
                    if (
                        (nx, ny) not in tiles
                        and (nx, ny) not in visited_overall
                        and (nx, ny) not in not_tiles
                        and (nx, ny) not in region
                    ):
                        q.append((nx, ny))
        if touches_wall:
            not_tiles.update(region)
        else:
            tiles.update(region)
        visited_overall.update(region)

# Find first rectangle with all internal points in tiles
for dist in distances:
    area, p1, p2 = dist
    x1, y1 = p1
    x2, y2 = p2
    min_rx, max_rx = min(x1, x2), max(x1, x2)
    min_ry, max_ry = min(y1, y2), max(y1, y2)

    all_in_tiles = True
    for x in range(min_rx, max_rx + 1):
        if not all_in_tiles:
            break
        for y in range(min_ry, max_ry + 1):
            if (x, y) not in tiles:
                all_in_tiles = False
                break

    if all_in_tiles:
        print(f"First valid rectangle: area={area}, p1={p1}, p2={p2}")
        break
