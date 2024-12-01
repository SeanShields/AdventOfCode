class NumberLocation:
  def __init__(self, number, start, end):
    self.number = number
    self.start = start
    self.end = end

def solvePart1(fileName):
  file = readFile(fileName)
  grid = file.split('\n')
  total = 0
  for rowNum, row in enumerate(grid):
    numberLocations = numbersInRow(row)
    for numberLocation in numberLocations:
      for colNum in range(numberLocation.start, numberLocation.end + 1):
        hasAdjacentSymbols = hasAdjacentSymbol(getAdjacents(rowNum, colNum, grid))
        if hasAdjacentSymbols:
          total += numberLocation.number
          break
  print('total:', total)

def solvePart2(fileName):
  file = readFile(fileName)
  grid = file.split('\n')
  total = 0
  numbers = []
  for rowNum, row in enumerate(grid):
    numbers.append(numbersInRow(row))
    # find each start
    # check if it has exactly 2 number adjacents
    # multiply those numbers together
    # add to total
  for num in enumerate(numbers):
    print(num.number, num.start, num.end)
  print('total:', total)

def numbersInRow(row):
  numbers = []
  current = ''
  for i in range(0, len(row) + 1):
    num = False
    if i < len(row):
      _, num = intTryParse(row[i])
    if not num and current != '':
      numbers.append(NumberLocation(int(current), i - len(current), i - 1))
      current = ''
    elif num:
      current += row[i]
      num = True
  return numbers

def getAdjacents(row, col, grid):
  top = None
  topRight = None
  right = None
  bottomRight = None
  bottom = None
  bottomLeft = None
  left = None
  topLeft = None

  topEdge = row == 0
  rightEdge = col == len(grid[0]) - 1
  bottomEdge = row == len(grid) - 1
  leftEdge = col == 0

  # top
  if not topEdge:
    top = grid[row - 1][col]

  # top-right
  if not topEdge and not rightEdge:
    topRight = grid[row - 1][col + 1]
  
  # right
  if not rightEdge:
    right = grid[row][col + 1]

  # bottom-right
  if not bottomEdge and not rightEdge:
    bottomRight = grid[row + 1][col + 1]

  # bottom
  if not bottomEdge:
    bottom = grid[row + 1][col]

  # bottom-left
  if not leftEdge and not bottomEdge:
    bottomLeft = grid[row + 1][col - 1]

  # left
  if not leftEdge:
    left = grid[row][col - 1]

  # top-left
  if not leftEdge and not topEdge:
    topLeft = grid[row - 1][col - 1]

  return top, topRight, right, bottomRight, bottom, bottomLeft, left, topLeft

def hasAdjacentSymbol(chars):
  for char in chars:
    if char is None:
       continue
    _, isInt = intTryParse(char)
    if not isInt and char != '.':
        return True
  return False

def intTryParse(value):
  try:
      return int(value), True
  except ValueError:
      return value, False

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    # solvePart1('input.txt')
    solvePart2('test.txt')