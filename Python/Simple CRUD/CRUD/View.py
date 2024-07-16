from . import Operation


def read_console():
    data_file = Operation.read()
    index = "No"
    title = "Title"
    author = "Author"
    year = "Year"

    # Header
    print("\n" + "=" * 100)
    print(f"{index:4} | {title:40} | {author:40} | {year:5}")
    print("-" * 100)

    # Data
    for index, data in enumerate(data_file):
        data_break = data.split(",")
        id = data_break[0]
        date_add = data_break[1]
        title = data_break[2]
        author = data_break[3]
        year = data_break[4]
        print(f"{index + 1:4} | {title:.40} | {author:.40} | {year:4}", end="")

    # Footer
    print("=" * 100 + "\n")


def create_console():
    print("\n\n" + "=" * 100)
    print("Input new book")
    title = input("Title\t: ")
    author = input("Author\t: ")
    while (True):
        try:
            year = int(input("Year\t: "))
            if len(str(year)) == 4:
                break
            else:
                print("Year length must be 4")
        except ValueError:
            print("Year must be an integer")

    Operation.create(title, author, year)
    print("\nnew book created")
    read_console()


def update_console():
    read_console()
    while True:
        print("choice book number on update")
        book_number = int(input("Book Number: "))
        book_data = Operation.read(index=book_number)

        if book_data:
            break
        else:
            print("number not found")

    data_break = book_data.split(',')
    id = data_break[0]
    date_add = data_break[1]
    title = data_break[2]
    author = data_break[3]
    year = data_break[4][:-1]

    while True:
        print("\n" + "=" * 100)
        print("Please input number data you want to update")
        print(f"1. Title\t: {title:.40}")
        print(f"2. Author\t: {author:.40}")
        print(f"3. Year\t: {year:4}")

        user_option = input("Use data [1,2,3]: ")
        print("\n" + "=" * 100)
        match user_option:
            case "1":
                title = input("title\t: ")
            case "2":
                author = input("author\t: ")
            case "3":
                while True:
                    try:
                        year = int(input("Year\t: "))
                        if len(str(year)) == 4:
                            break
                        else:
                            print("year must 4 characters")
                    except ValueError:
                        print("year must be an integer")
            case _:
                print("index not found")

        print("New Data")
        print(f"1. Title\t: {title:.40}")
        print(f"2. Author\t: {author:.40}")
        print(f"3. Year\t: {year:4}")
        is_done = input("Data is done? [y/n]: ")
        if is_done == "y" or is_done == "Y":
            break

    Operation.update(book_number, id, date_add, title, author, year)
