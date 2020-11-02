
with open("map.txt", "r") as f:
    map = f.readlines()

updated_map = ""

for line in map:
    updated_map += line.strip()

with open("map_sim.txt", "w") as f:
    f.writelines(updated_map)
