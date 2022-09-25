import json

import matplotlib.pyplot as plt

DATA_PATH = "anajson"
BAR_WIDTH = 0.15


with open(DATA_PATH) as f:
    results = json.load(f)["results"]

commands = {
    batch["command"]: [sum(batch["times"]) / len(batch["times"])] for batch in results
}

sorted_results = {k: v for k, v in sorted(commands.items(), key=lambda item: item[1])}
sorted_key_list = list(sorted_results)

# X axis position
rust_position = sorted_key_list.index("Rust")
go_position = sorted_key_list.index("Go")
nim_position = sorted_key_list.index("Nim")
python_position = sorted_key_list.index("Python")
v_position = sorted_key_list.index("V")

# Medians
rust_median = commands["Rust"]
go_median = commands["Go"]
python_median = commands["Python"]
nim_median = commands["Nim"]
v_median = commands["V"]

# Plot size
plt.subplots(figsize=(15, 10))

# Make the plot
plt.bar(rust_position, rust_median, color="r", width=BAR_WIDTH, edgecolor="grey", label="Rust", align="center")
plt.bar(go_position, go_median, color="c", width=BAR_WIDTH, edgecolor="grey", label="Go", align="center")
plt.bar(python_position, python_median, color="b", width=BAR_WIDTH, edgecolor="grey", label="Python", align="center")
plt.bar(nim_position, nim_median, color="y", width=BAR_WIDTH, edgecolor="grey", label="Nim", align="center")
plt.bar(v_position, v_median, color="k", width=BAR_WIDTH, edgecolor="grey", label="Nim", align="center")

plt.xlabel("Programming Language", fontweight="bold", fontsize=15)
plt.ylabel("Time (s)", fontweight="bold", fontsize=15)

plt.xticks(
    [rust_position, go_position, python_position, nim_position, v_position],
    ["Rust", "Go", "Python", "Nim", "V"],
)

plt.legend()
plt.savefig("testplot.png")
plt.show()
