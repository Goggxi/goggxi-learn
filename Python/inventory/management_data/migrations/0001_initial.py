# Generated by Django 5.0.7 on 2024-07-24 04:21

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Attachment',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('file', models.FileField(upload_to='attachments/')),
                ('created_at', models.DateTimeField(auto_now_add=True, null=True)),
                ('updated_at', models.DateTimeField(auto_now=True, null=True)),
            ],
        ),
        migrations.CreateModel(
            name='Building',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('code', models.CharField(max_length=50, unique=True)),
                ('name', models.CharField(max_length=150)),
                ('length_in_meters', models.IntegerField(default=0, help_text='in meters')),
                ('width_in_meters', models.IntegerField(default=0, help_text='in meters')),
                ('building_area_sqm', models.IntegerField(default=0, help_text='in square meters')),
                ('number_of_floors', models.IntegerField(default=1)),
                ('asset_acquisition_value', models.DecimalField(decimal_places=2, default=0.0, max_digits=10)),
                ('maintenance_estimate_in_month', models.IntegerField(default=0, help_text='in months')),
                ('last_renovation_time', models.DateTimeField(blank=True, null=True)),
                ('description', models.TextField(blank=True, null=True)),
                ('latitude', models.FloatField(blank=True, null=True)),
                ('longitude', models.FloatField(blank=True, null=True)),
                ('occupied_by', models.CharField(blank=True, max_length=150, null=True)),
                ('responsible_person', models.CharField(blank=True, max_length=150, null=True)),
                ('created_at', models.DateTimeField(auto_now_add=True, null=True)),
                ('updated_at', models.DateTimeField(auto_now=True, null=True)),
                ('attachments', models.ManyToManyField(blank=True, related_name='buildings', to='management_data.attachment')),
            ],
        ),
    ]