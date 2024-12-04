import re

nums = {
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
  lines = readFile(fileName).split('\n')
  calibration = 0
  for line in lines:
    digits = re.findall("\d", line)
    calibration += int(digits[0] + digits[len(digits) - 1])
  return calibration

def solvePart2(fileName):
  lines = readFile(fileName).split('\n')
  calibration = 0
  for line in lines:
      print(parseLine(line))
    # calibration += int(digits[0] + digits[len(digits) - 1])
  return calibration

def parseLine(line):
  current = ""
  for char in reversed(line):
    for num in nums:
       if char + current == num:
          current = 
  return current

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    # print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('test.txt'))