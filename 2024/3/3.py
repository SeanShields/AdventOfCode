import re

def solvePart1(fileName):
  input = readFile(fileName)
  matches = re.findall("mul\(\d*,\d*\)", input)
  return multiplyMatches(matches)

def solvePart2(fileName):
  input = readFile(fileName)
  matches = re.findall("mul\(\d*,\d*\)|do\(\)|don't\(\)", input)
  return multiplyMatches(matches)

def multiplyMatches(matches):
  total = 0
  do = True
  for match in matches:
    isDo = match == "do()"
    isDont = match == "don't()"
    if isDo or isDont:
      do = True if isDo else False
      continue
  
    if not do:
      continue

    split = match.split(',')
    x = int(split[0].lstrip("mul("))
    y = int(split[1].rstrip(")"))
    total += x * y
  return total

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('input.txt'))