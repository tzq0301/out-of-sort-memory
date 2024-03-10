import json

# Adjusting the previous JSON object to increase its size to over 500KB

# Increase the number of users to generate a larger JSON file
data_large = {
    "users": [
        {
            "id": i,
            "username": f"user_{i}",
            "email": f"user_{i}@example.com",
            "profile": {
                "age": i % 100,
                "country": "Country_" + str(i % 10),
                "preferences": {
                    "likes": ["Item_" + str(i % 10), "Item_" + str((i + 1) % 10), "Item_" + str((i + 2) % 10)],
                    "dislikes": ["Item_" + str((i + 3) % 10), "Item_" + str((i + 4) % 10), "Item_" + str((i + 5) % 10)],
                }
            }
        }
        for i in range(5000)  # Increase the range to generate more user entries
    ]
}

# Write the larger JSON data to the same file 'data.json'
file_path_large = 'data.json'
with open(file_path_large, 'w') as json_file_large:
    json.dump(data_large, json_file_large)
