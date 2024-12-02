def solvePart1(fileName):
  lines = readFile(fileName).split('\n')
  safe = 0
  for l in lines:
     line = l.split(' ')
     if lineSafe(line):
        safe += 1
  return safe

def solvePart2(fileName):
  lines = readFile(fileName).split('\n')
  safe = 0
  for l in lines:
     line = l.split(' ')
     if lineSafe(line):
        safe += 1
  return safe


def increasing(current, next):
   return next > current

def lineSafe(line):
  minDiff = 1
  maxDiff = 3
  increasing = int(line[0]) < int(line[1])

  for i, c in enumerate(line):
    if i == len(line) - 1:
      break

    current = int(c)
    next = int(line[i + 1])
    
    # check if increase or decrease
    if (increasing and current >= next) or (not increasing and current <= next):
      return False

    # check the diff
    diff = abs(current - next)
    if not minDiff <= diff <= maxDiff:
       return False

  return True

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('test.txt'))
    print('Part 2:', solvePart2('test.txt'))