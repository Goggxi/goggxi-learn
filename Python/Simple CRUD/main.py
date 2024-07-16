import os
import CRUD as CRUD

if __name__ == '__main__':
    operating_system = os.name

    match operating_system:
        case 'posix':
            os.system('clear')
        case 'nt':
            os.system('cls')

    print("Welcome to program")
    print("Library Databases")
    print("===================")

    CRUD.init_console()

while(True):
    match operating_system:
        case 'posix':
            os.system('clear')
        case 'nt':
            os.system('cls')

    print("Welcome to program")
    print("Library Databases")
    print("===================")

    print("1. Read Data")
    print("2. Create Data")
    print("3. Update Data")
    print("4. Delete Data")
    print("5. Exit\n")

    user_input = input("Input Option: ")

    match user_input:
        case '1': CRUD.read_console()
        case '2': CRUD.create_console()
        case '3': print("Update Data")
        case '4': print("Delete Data")
        case '5': print("Exit")

    is_done = input("is done (y/n)?")
    if is_done == 'y' or is_done == 'Y':
        break

    print("Thank you for using program")

