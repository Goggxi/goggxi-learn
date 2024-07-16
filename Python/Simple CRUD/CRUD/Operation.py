import time

from . import Database
from .Util import generate_id

def add():
    title = input("Title: ")
    author = input("Author: ")
    year = input("Year: ")

    data = Database.TEMPLATE.copy()

    data["id"] = generate_id(6)
    data["date_add"] = time.strftime("%Y-%m-%d-%H-%M-%S%z",time.gmtime())
    data["title"] = title + Database.TEMPLATE["title"][len(title):]
    data["author"] = author + Database.TEMPLATE["author"][len(author):]
    data["year"] = year

    data_str = f'{data["id"]},{data["date_add"]},{data["title"]},{data["author"]},{data["year"]}\n'

    try:
        with open(Database.DB_NAME, 'w', encoding='utf-8') as file:
            file.write(data_str)
            file.close()
            print("Database saved")
    except IOError:
        print("Database could not be saved")

def read():
    try:
        with open(Database.DB_NAME, 'r', encoding='utf-8') as file:
            content = file.readlines()
            return content
    except IOError:
        print("Database could not be read")
        return False
