import csv
import sys

file = sys.argv[1]

with open(file, newline="") as csvfile:
    reader = csv.DictReader(csvfile)
    for row in reader:
        print(row["zip_code_1"], row["city_1"])
        print(row["zip_code_2"], row["city_2"])
        print(row["destination_zip_code"], row["destination_city"])

