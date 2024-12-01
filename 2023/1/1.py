import re
import numbers

namesAndNums = {
   'one': 1,
   'two': 2,
   'three': 3,
   'four': 4,
   'five': 5,
   'six': 6,
   'seven': 7,
   'eight': 8,
   'nine': 9
}

def solvePart1(fileName):
  total = 0
  lines = readFile(fileName).split('\n')
  for line in lines:
    nums = re.findall('[0-9]', line)
    total += int(str(nums[0]) + nums[len(nums) - 1])
  print(total)

def solvePart2(fileName):
  total = 0
  lines = readFile(fileName).split('\n')
  for line in lines:
     #nums = namesToNums(line)
     nums = parseLine(line)
     print(nums)
    # total += nums[0] + nums[len(nums) - 1]
    #print(nums)
  #print(total)

def parseLine(line):
  nums = []
  current = ''
  i = 0
  while i < len(line):
    print(current)
    current += line[i]
    num, isInt = intTryParse(line[i])
    if current in namesAndNums:
      nums.append(namesAndNums[current])
      current = ''
    else:
      i+=1
    if isInt:
      nums.append(num)

  return nums

def intTryParse(value):
    try:
        return int(value), True
    except ValueError:
        return value, False

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    # solvePart1('test.txt')
    solvePart2('test2.txt')