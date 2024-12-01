import re

class Counts:
  def __init__(self, red, green, blue):
    self.red = red
    self.green = green
    self.blue = blue

def solvePart1(fileName):
  games = readFile(fileName).split('\n')
  maximums = Counts(12, 13, 14)
  result = 0
  for game in games:
    id, sets = parseGame(game)
    if possible(sets, maximums):
      result += id
  print(result)

def solvePart2(fileName):
  games = readFile(fileName).split('\n')
  result = 0
  for game in games:
    _, sets = parseGame(game)
    mins = minumums(sets)
    result += power(mins)
  print(result)

def parseGame(game):
  idGroup = re.search(r'Game\s[0-9]+', game).group()
  id = int(idGroup.split(' ')[1])
  sets = re.sub(idGroup + ': ', '', game)
  return id, sets

def minumums(sets):
  minumums = Counts(0, 0, 0)
  for set in sets.split('; '):
    for pair in set.split(', '):
      count, color = parseCountColor(pair)
      if color == 'red' and count > minumums.red:
        minumums.red = count
      if color == 'green' and count > minumums.green:
        minumums.green = count
      if color == 'blue' and count > minumums.blue:
        minumums.blue = count
  return minumums

def possible(sets, maximums):
  for set in sets.split('; '):
    for pair in set.split(', '):
      count, color = parseCountColor(pair)
      if (color == 'red' and count > maximums.red) \
        or (color == 'green' and count > maximums.green) \
        or (color == 'blue' and count > maximums.blue):
          return False
  return True

def parseCountColor(pair):
  group = pair.split(' ')
  count = int(group[0])
  color = group[1]
  return count, color

def power(counts):
  return counts.red * counts.green * counts.blue

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    solvePart1('input.txt')
    solvePart2('input.txt')