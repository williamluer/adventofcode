

position = 50
count = 0
with open("1.txt","r") as f:
    for line in f:
        direction = line[0]
        value = int(line[1:])
        if direction == "R":
            position += value
        else:
            position -= value
        print(f"{direction=} {value=} {position=}")
        if position % 100 == 0:
            print("COUNTING")
            count += 1
print(count)

