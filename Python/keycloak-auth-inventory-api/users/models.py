from django.contrib.auth.models import User
from django.db import models


class Employee(models.Model):
    user = models.OneToOneField(User, on_delete=models.CASCADE)
    keycloak_id = models.CharField(max_length=255, unique=True)
    role = models.CharField(max_length=20, choices=[('staff', 'Staff'), ('admin', 'Admin')])
