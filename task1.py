def get_input(file_name):
  numbers = []
  with open(file_name, 'r') as f:
    numbers = [int(x) for x in f.readlines()]

  return numbers

def get_increasing_depth_sum_with_time_window(numbers, window_size):
  previous = sum(numbers[:window_size])
  increasing_depth = 0
  
  for i in range(window_size, len(numbers)+1):
    current = sum(numbers[i-window_size:i])
    if current > previous:
      increasing_depth += 1

    previous = current

  return increasing_depth

if __name__ == "__main__":
  numbers = get_input('./input/input1_1.txt')
  print(get_increasing_depth_sum_with_time_window(numbers, 1))
