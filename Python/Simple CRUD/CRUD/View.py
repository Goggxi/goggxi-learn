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
        print(f"{index+1:4} | {title:.40} | {author:.40} | {year:4}", end="")

    # Footer
    print("=" * 100 + "\n")


def create_console():
    print("\n\n"+"="*100)
    print("Input new book")
    title = input("Title\t: ")
    author = input("Author\t: ")
    while(True):
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

