

position = 50
count = 0
with open("1.txt","r") as f:
    for line in f:
        direction = line[0]
        value = int(line[1:])
        print(f"{direction=} {value=} {position=}")
        while value > 0:
            if direction == "R":
                position += 1
            else:
                position -= 1
            value -= 1
            if position % 100 == 0:
                count += 1
print(count)

