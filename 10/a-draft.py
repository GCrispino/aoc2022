with open("./10/input/test.txt") as f:
    lines = f.readlines()
    cycle_num = 1
    reg_x_val = 1
    for line in lines:
        print("cycle:", cycle_num)
        if cycle_num == 20 or (cycle_num - 20) % 40 == 0:
          print(" ", reg_x_val)

        spl = line.split(" ")

        cmd = spl[0]

        if cmd == "addx":
          reg_x_val += int(spl[1])
          cycle_num += 2
        else:
          cycle_num += 1
