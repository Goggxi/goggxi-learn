# Generated by Django 5.0.7 on 2024-07-23 12:58

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('api', '0001_initial'),
    ]

    operations = [
        migrations.RenameField(
            model_name='artist',
            old_name='description',
            new_name='bio',
        ),
    ]
