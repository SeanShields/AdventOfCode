def solvePart1(fileName):
  lines = readFile(fileName).split('\n')
  a, b = getLists(lines)
  a.sort()
  b.sort()

  # assumes both lists are the same length
  total = 0
  for i, a in enumerate(a):
    total += abs(a - b[i])
  return total

def solvePart2(fileName):
  lines = readFile(fileName).split('\n')
  a, b = getLists(lines)

  similarity = 0
  for num in a:
    similarity += getSimilarity(num, timesInList(num, b))
  return similarity

def timesInList(num, list):
  total = 0
  for x in list:
    if x == num:
      total += 1
  return total

def getSimilarity(num, count):
   return num * count

def getLists(lines):
  a = []
  b = []
  for line in lines:
     pair = line.split('   ')
     a.append(int(pair[0]))
     b.append(int(pair[1]))
  return a, b

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('input.txt'))