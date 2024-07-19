# django-songs-api

# Create Project CLI
```bash
# Make a new directory and change into it
mkdir django-songs-api
cd django-songs-api

# Create a new virtual environment and activate it for manage versions of Python (optional)
pyenv virtualenv 3.x.x django-songs-api-env  # Create a new virtual environment
pyenv local django-songs-api-env  # Activate the virtual environment
# Check the list of virtual environments
pyenv virtualenvs
# Deactivate the virtual environment
pyenv uninstall django-songs-api-env

# Create a new folder .venv and activate the virtual environment
python -m venv .venv
source .venv/bin/activate # Linux or MacOS
# .venv\Scripts\activate  # Windows

# Install Django
pip install django

# django-admin startproject <project_name> .
django-admin startproject config .

# Create a new file requirements.txt
pip freeze > requirements.txt

# Run the server
python manage.py runserver

# Create a new app inside the project
python manage.py startapp hello
```