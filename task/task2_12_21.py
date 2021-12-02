from os import get_terminal_size


def get_input(file_name):
  directions = [(y[0], int(y[1])) for y in [x.split() for x in open(file_name, 'r').readlines()]]

  return directions


def get_final_direction(directions):
  total_forward = sum([x[1] for x in directions if x[0] == 'forward'])
  total_depth = sum([-y[1] if y[0] == 'up' else y[1] for y in [x for x in directions if x[0] in ('down', 'up')]])
  return total_depth * total_forward


def get_final_direction_with_aim(directions):
  aim = 0
  forward = 0
  depth = 0
  for direction in directions:
    if direction[0] == 'forward':
      forward += direction[1]
      depth += aim * direction[1]
    elif direction[0] == 'up':
      aim -= direction[1]
    elif direction[0] == 'down':
      aim += direction[1]

  return forward * depth


if __name__ == "__main__":
  directions = get_input('./input/input2.txt')
  print(get_final_direction(directions))
  print(get_final_direction_with_aim(directions))