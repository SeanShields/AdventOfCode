def solvePart1(fileName):
  file = readFile(fileName)
  print(file)

def solvePart2(fileName):
  file = readFile(fileName)
  print(file)

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('input.txt'))