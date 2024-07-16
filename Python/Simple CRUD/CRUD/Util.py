import random
import string

def generate_id(length:int) -> str:
    return "".join(random.choices(string.ascii_uppercase + string.digits, k=length))