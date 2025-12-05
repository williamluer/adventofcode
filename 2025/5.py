
filename = "5-test.txt"
intervals = []
ids = []
with open(filename, "r") as f:
    on_intervals = True
    for line in f:
        if line.strip() == "":
            on_intervals = False
            continue
        if on_intervals:
            n_strs = line.strip().split("-")
            intervals.append([int(n_strs[0]), int(n_strs[1])])
        else:
            ids.append(int(line.strip()))

print(intervals)
print(ids)

"""
sort intervals by starting ID
merge by:
l = min(l1,l1)
if l2 <= l1:
    l = min(l1,l2)
    r = max(r1,r2)

if l2 > l1 and l2 < r1:
    l = l1
    r = max(r1,r2)

JUST CHECK IF N IS >= L and <= R


"""


