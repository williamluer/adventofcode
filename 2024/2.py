def is_valid_sequence(nums):
    if len(nums) < 2:
        return True

    # Check if ascending
    is_ascending = True
    for i in range(1, len(nums)):
        if nums[i] <= nums[i - 1] or nums[i] - nums[i - 1] > 3:
            is_ascending = False
            break

    if is_ascending:
        return True

    # Check if descending
    is_descending = True
    for i in range(1, len(nums)):
        if nums[i] >= nums[i - 1] or nums[i - 1] - nums[i] > 3:
            is_descending = False
            break

    return is_descending


# Read input
numbers = []
with open("2.txt", "r") as f:
    for line in f:
        nums = [int(n) for n in line.split()]
        numbers.append(nums)

count = 0
for sequence in numbers:
    # First check if sequence is valid without removing anything
    if is_valid_sequence(sequence):
        count += 1
        continue

    # Try removing one number at a time
    for i in range(len(sequence)):
        test_sequence = sequence[:i] + sequence[i + 1 :]
        if is_valid_sequence(test_sequence):
            count += 1
            break

print(count)
