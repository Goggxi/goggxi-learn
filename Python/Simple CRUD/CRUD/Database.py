from . import Operation

DB_NAME = "data.txt"
TEMPLATE = {
    "id": "XXXXXX",
    "date_add": "yyyy-mm-dd",
    "title": 255 * " ",
    "author": 255 * " ",
    "year": "yyyy",
}


def init_console():
    try:
        with open(DB_NAME, "r") as file:
            print("Database initialized successfully")
    except FileNotFoundError:
        print("Database file not found & Create new database")
        with open(DB_NAME, "w", encoding="utf-8") as file:
            Operation.add()
