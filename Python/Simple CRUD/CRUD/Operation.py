import os
import time

from . import Database
from .Util import generate_id


def add():
    title = input("Title: ")
    author = input("Author: ")
    year = input("Year: ")

    data = Database.TEMPLATE.copy()

    data["id"] = generate_id(6)
    data["date_add"] = time.strftime("%Y-%m-%d-%H-%M-%S%z", time.gmtime())
    data["title"] = title + Database.TEMPLATE["title"][len(title):]
    data["author"] = author + Database.TEMPLATE["author"][len(author):]
    data["year"] = year

    data_str = f'{data["id"]},{data["date_add"]},{data["title"]},{data["author"]},{data["year"]}\n'

    try:
        with open(Database.DB_NAME, 'w', encoding='utf-8') as file:
            file.write(data_str)
            file.close()
    except IOError:
        print("Database could not be saved")


def read(**kwargs):
    try:
        with open(Database.DB_NAME, 'r') as file:
            content = file.readlines()
            book_count = len(content)
            if "index" in kwargs:
                book_index = kwargs["index"]-1
                if book_index < 0 or book_index > book_count:
                    return False
                else:
                    return content[book_index]
            else:
                return content
    except FileNotFoundError:
        print("Can't read data from database")
        return False


def create(title, author, year):
    data = Database.TEMPLATE.copy()

    data["id"] = generate_id(6)
    data["date_add"] = time.strftime("%Y-%m-%d-%H-%M-%S%z", time.gmtime())
    data["title"] = title + Database.TEMPLATE["title"][len(title):]
    data["author"] = author + Database.TEMPLATE["author"][len(author):]
    data["year"] = year

    data_str = f'{data["id"]},{data["date_add"]},{data["title"]},{data["author"]},{data["year"]}\n'

    try:
        with open(Database.DB_NAME, 'a', encoding='utf-8') as file:
            file.write(data_str)
            file.close()
    except IOError:
        print("Data could not be saved")


def update(book_number, id, date_add, title, author, year):
    data = Database.TEMPLATE.copy()

    data["id"] = id
    data["date_add"] = date_add
    data["title"] = title + Database.TEMPLATE["title"][len(title):]
    data["author"] = author + Database.TEMPLATE["author"][len(author):]
    data["year"] = str(year)

    data_str = f'{data["id"]},{data["date_add"]},{data["title"]},{data["author"]},{data["year"]}\n'

    length = len(data_str)

    try:
        with(open(Database.DB_NAME, 'r+', encoding="utf-8")) as file:
            file.seek(length * (book_number - 1))
            file.write(data_str)
    except FileNotFoundError:
        print("error updating data")


def delete(book_number):
    try:
        with open(Database.DB_NAME, 'r') as file:
            counter = 0

            while (True):
                content = file.readline()
                if len(content) == 0:
                    break
                elif counter == book_number - 1:
                    pass
                else:
                    with open("data_temp.txt", 'a', encoding="utf-8") as temp_file:
                        temp_file.write(content)
                counter += 1
    except:
        print("database error")

    os.rename("data_temp.txt", Database.DB_NAME)