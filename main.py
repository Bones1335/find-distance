import csv
import sys

file = sys.argv[1]
default_start_city = ('25000', 'BESANCON')

with open(file, newline="") as csvfile:
    reader = csv.DictReader(csvfile)
    for row in reader:
        city_1 = (row["zip_code_1"], row["city_1"])
        city_2 = (row["zip_code_2"], row["city_2"])
        destination_city = (row["destination_zip_code"], row["destination_city"])
        print(default_start_city)
        print(city_1)
        print(city_2)
        print(destination_city)

