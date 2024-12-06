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

def checkOrder(update, rules):
  for rule in rules:
    args = rule.split('|')
    if (args[0] not in update and args[1] not in update) \
      or (args[0] in update and args[1] not in update) \
      or (args[1] in update and args[0] not in update):
      continue
    if update.index(args[0]) > update.index(args[1]):
      return False
  return True

def solvePart2(fileName):
  file = readFile(fileName).split('\n')
  return file

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    # print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('test.txt'))