from os import get_terminal_size


def get_input(file_name):
  diagnostics = [x.strip() for x in open(file_name, 'r').readlines()]

  return diagnostics

def invert_binary(binary_str):
  return binary_str.replace("1", "2").replace("0", "1").replace("2", "0")

def multiply_binary_string_list(a, b):
  return int("".join(a), 2) * int("".join(b), 2)


def get_gamma_epsylon(diagnostics):
  zero_counter = [0 for x in range(len(diagnostics[0]))]
  
  for i in range(len(zero_counter)):
    zero_counter[i] = sum([True if x == "0" else False for x in [y[i] for y in diagnostics]])
  
  gamma_str = "".join(["0" if zero_counter[i] > 0.5 * len(diagnostics) else "1" for i in range(len(zero_counter))])
  eps_str = invert_binary(gamma_str)

  return multiply_binary_string_list(gamma_str, eps_str)

def get_oxygen_with_co2(diagnostics):
  leftover_most_common_list = list(diagnostics)
  leftover_least_common_list = list(diagnostics)

  for index in range(len(diagnostics[0])):
    zero_most_counter = sum([True if x == "0" else False for x in [y[index] for y in leftover_most_common_list]])
    zero_least_counter = sum([True if x == "0" else False for x in [y[index] for y in leftover_least_common_list]])

    most_common_value = "0" if zero_most_counter > 0.5 * len(leftover_most_common_list) else "1"
    least_common_value = "0" if zero_least_counter <= 0.5 * len(leftover_least_common_list) else "1"
    
    if len(leftover_most_common_list) != 1:
      leftover_most_common_list = list(filter(lambda x: x[index] == most_common_value, leftover_most_common_list))
    if len(leftover_least_common_list) != 1:
      leftover_least_common_list = list(filter(lambda x: x[index] == least_common_value, leftover_least_common_list))

  return multiply_binary_string_list(leftover_most_common_list, leftover_least_common_list)



if __name__ == "__main__":
  diagnostics = get_input('./input/input3.txt')
  print(get_gamma_epsylon(diagnostics))
  print(get_oxygen_with_co2(diagnostics))