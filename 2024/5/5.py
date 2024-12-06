def solvePart1(fileName):
  file = readFile(fileName)
  parts = file.split('\n\n')
  rules = parts[0].split('\n')
  updates = parts[1].split('\n')
  total = 0
  for update in updates:
    nums = update.split(',')
    if checkOrder(nums, rules):
      total += int(nums[round(len(nums) // 2)])
  return total

def solvePart2(fileName):
  file = readFile(fileName)
  parts = file.split('\n\n')
  rules = parts[0].split('\n')
  updates = parts[1].split('\n')
  total = 0
  for update in updates:
    nums = update.split(',')
    if not checkOrder(nums, rules):
      corrected = correctOrder(nums, rules)
      total += int(corrected[round(len(corrected) // 2)])
  return total

def correctOrder(nums, rules):
  for _ in range(len(nums)):
    for rule in rules:
      args = rule.split('|')
      leftIndex = None
      rightIndex = None
      try:
        leftIndex = nums.index(args[0])
      except:
        pass
      try:
        rightIndex = nums.index(args[1])
      except:
          pass
      if leftIndex is None or rightIndex is None:
        continue
      if leftIndex > rightIndex:
        nums[leftIndex], nums[rightIndex] = nums[rightIndex], nums[leftIndex]
  return nums

def checkOrder(update, rules):
  for rule in rules:
    args = rule.split('|')
    left = args[0]
    right = args[1]
    if (left not in update and right not in update) \
      or (left in update and right not in update) \
      or (right in update and left not in update):
      continue
    if update.index(left) > update.index(right):
      return False
  return True

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('input.txt'))