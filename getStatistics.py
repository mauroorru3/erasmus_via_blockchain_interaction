import json

file_it_chain = open("university_chain_it/log/val-unipi/stats.txt", "r")
file_de_chain = open("university_chain_de/log/val-tum/stats.txt", "r")
file_hub_chain = open("hub/log/val-hub-instance-1/stats.txt", "r")

data_it_chain = []

for line in file_it_chain:
    data.append(json.loads(line))

print(len(data))
