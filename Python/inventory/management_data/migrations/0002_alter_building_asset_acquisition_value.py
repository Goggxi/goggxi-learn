# Generated by Django 5.0.7 on 2024-07-24 04:28

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('management_data', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='building',
            name='asset_acquisition_value',
            field=models.IntegerField(default=0, help_text='in IDR (indonesia)'),
        ),
    ]
