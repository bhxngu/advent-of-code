file = open("day-2/input.txt")


def isSafeReport(numbers, recursive):
  seriesPattern = None

  for i in range(len(numbers) - 1):

    first, second = numbers[i], numbers[i + 1]

    pairPattern = second > first

    difference = abs(second - first)

    if difference >= 1 and difference <= 3:
      if seriesPattern is None:
        seriesPattern = pairPattern
      elif seriesPattern != pairPattern:
        if recursive:
          return False
        firstAns = isSafeReport(numbers[:i] + numbers[i + 1:],
                                True) or isSafeReport(
                                    numbers[:i + 1] + numbers[i + 2:], True)
        if i - 1 == 0:
          secondAns = isSafeReport(numbers[1:], True)
          return firstAns or secondAns
        return firstAns
      elif i == len(numbers) - 2:
        return True
    else:
      if recursive:
        return False

      return isSafeReport(numbers[:i] + numbers[i + 1:], True) or isSafeReport(
          numbers[:i + 1] + numbers[i + 2:], True)


count = 0
for line in file.readlines():
  numbers = line.split()
  numbers = list(map(int, numbers))

  if isSafeReport(numbers, False):
    count += 1

print(count)
