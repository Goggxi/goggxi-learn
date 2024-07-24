from django.db import models


class Employee(models.Model):
    name = models.CharField(max_length=150)
    email = models.CharField(max_length=150)
    uuid = models.UUIDField(blank=True, null=True)
    username = models.CharField(max_length=150)

    def __str__(self):
        return self.name

    def __unicode__(self):
        return
