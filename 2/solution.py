print(sum([1 for line in map(lambda x: x.strip('\n'), open('input')) if __import__('re').match(__import__('re').sub('(\d+)-(\d+) (\w): (\w+)', '^\g<3>{\g<1>,\g<2>}$', line), __import__('re').sub(__import__('re').sub('.* (\w):.*', '[^\g<1>]', line), '', line.rpartition(' ')[2]))]))

# construct two regexes, one for each position
# the regexes should be .{position_one - 1}search_char.{position_two-1}search_char
import re
with  open('input') as file:
  count = 0
  for line in map(lambda x: x.strip('\n'), file):
    print(line)
    first_pos = int(re.match('(\d+)', line)[1]) - 1
    print(first_pos)
    second_pos = int(re.match('\d+-(\d+)', line)[1]) - 1
    print(second_pos)
    search_char = re.match('\d+-\d+ (\w)', line)[1]
    matcher = ('.' * first_pos) + search_char + ('.' * second_pos) + search_char
    if re.match(matcher, line.rpartition(' ')[2]):
      count += 1
  print(count)