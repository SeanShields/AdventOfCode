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
    solvePart1('input.txt')
    solvePart2('input.txt')