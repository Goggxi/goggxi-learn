import os

if __name__ == '__main__':
    operating_system = os.name

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

    print("\n===================\n")

    match user_input:
        case '1': print("Read Data")
        case '2': print("Create Data")
        case '3': print("Update Data")
        case '4': print("Delete Data")
        case '5': print("Exit")

    print("\n===================\n")
    is_done = input("is done (y/n)?")
    if is_done == 'y' or is_done == 'Y':
        break

    print("Thank you for using program")

