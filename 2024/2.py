numbers = []
with open("2-sample.txt", "r") as f:
    line = f.readline()
    while line:
        nums = line.split()
        numbers.append([int(n) for n in nums])
        line = f.readline()

count = 0
for level in numbers:
    # print()
    print(level)
    n_incorrect = 0
    asc = None
    last_val = None
    for i, n in enumerate(level):
        if last_val == n:
            # print("equal so removing")
            # break
            if n_incorrect >= 1:
                print("too many bads, equal so removing")
                break
            n_incorrect += 1
        print(f"i: {i}, last val: {last_val}, current: {n}, asc: {asc}")
        if i == 0:
            last_val = n
        elif i == 1:
            if n > last_val and n - last_val >= 1 and n - last_val <= 3:
                print("setting to ascending")
                asc = True
                last_val = n
            elif n < last_val and last_val - n >= 1 and last_val - n <= 3:
                print("setting to descending")
                asc = False
                last_val = n
            else:
                if n_incorrect >= 1:
                    print("too many bads")
                    break
                n_incorrect += 1
                # last_val = n
                # break

        elif i > 1:
            if asc is True:
                # print("asc is true")
                if last_val < n:
                    if n - last_val >= 1 and n - last_val <= 3:
                        last_val = n
                        if i == len(level) - 1:
                            # print(
                            #     f"ASC adding for {level} while n: {n} and last_val: {last_val}"
                            # )
                            count += 1
                            break
                        continue
                    else:
                        # print(
                        #     f"ASC Not adding for {level} b/c out of range for {last_val} and {n}"
                        # )
                        # last_val = n
                        # break
                        if n_incorrect >= 1:
                            break
                        n_incorrect += 1
                else:
                    # print(f"Not all ascending for {level}")
                    # break
                    if n_incorrect >= 1:
                        break
                    n_incorrect += 1

            elif asc is False:
                # print("asc is false")
                if last_val > n:
                    if last_val - n >= 1 and last_val - n <= 3:
                        last_val = n
                        if i == len(level) - 1:
                            # print(f"DESC adding for {level}")
                            count += 1
                            break
                        continue
                    else:
                        # print(
                        # f"DESC not adding for {level} b/c out of range for {last_val} and {n}"
                        # )
                        # last_val = n
                        # break
                        if n_incorrect >= 1:
                            break
                        n_incorrect += 1
                else:
                    # print(f"Not all descending for {level}")
                    if n_incorrect >= 1:
                        break
                    n_incorrect += 1
                    # break

print(count)
