import heapq
import math

points = []
with open("8-test.txt", "r") as f:
    for line in f:
        vals = line.strip().split(",")
        points.append(tuple([int(val) for val in vals]))


def get_distance(p1, p2):
    return math.sqrt((p1[0] - p2[0]) ** 2 + (p1[1] - p2[1]) ** 2 + (p1[2] - p2[2]) ** 2)


print(points)
distances = []
for i in range(len(points)):
    for j in range(i + 1, len(points)):
        dist = get_distance(points[i], points[j])
        distances.append([dist, points[i], points[j]])

heapq.heapify(distances)
for i in range(10):
    print(heapq.heappop(distances))
print()

parents = {p: p for p in points}
set_size = {p: 1 for p in points}


def find(point: tuple) -> tuple:
    # find parent
    while parents[point] != point:
        parents[point] = parents[parents[point]]
        point = parents[point]
    return point


def union(point1: tuple, point2: tuple) -> bool:
    # join sets
    p1_parent = find(point1)
    p2_parent = find(point2)
    if p1_parent == p2_parent:
        return False

    elif p1_parent != p2_parent:
        # set p2 to p1 (could do union by rank though)
        if set_size[p1_parent] >= set_size[p2_parent]:
            parents[p2_parent] = parents[p1_parent]
            set_size[p1_parent] += set_size[p2_parent]
            set_size[p2_parent] = 0
        parents[p2_parent] = p1_parent
        return True  # return parent of set?


n_merged = 0
while n_merged < 10:
    shortest_dist, point1, point2 = heapq.heappop(distances)
    print(f"on {shortest_dist=} {point1=} {point2=}")
    # if union(point1, point2):
    union(point1, point2)
    print(f"Merged {point1} and {point2}")
    n_merged += 1

sizes = []
for k, v in set_size.items():
    sizes.append(v)

s = sorted(sizes, reverse=True)
print(f"Ans: {s[0] * s[1] * s[2]}")
