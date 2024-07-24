from django.db import models
from .attachment import Attachment
from django.contrib.contenttypes.fields import GenericRelation


class Building(models.Model):
    code = models.CharField(max_length=50, unique=True)
    name = models.CharField(max_length=150)
    length_in_meters = models.IntegerField(help_text='in meters', default=0)
    width_in_meters = models.IntegerField(help_text='in meters', default=0)
    building_area_sqm = models.IntegerField(help_text='in square meters', default=0)
    number_of_floors = models.IntegerField(default=1)
    asset_acquisition_value = models.IntegerField(help_text='in IDR (indonesia)', default=0)
    maintenance_estimate_in_month = models.IntegerField(help_text='in months', default=0)
    last_renovation_time = models.DateTimeField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    latitude = models.FloatField(blank=True, null=True)
    longitude = models.FloatField(blank=True, null=True)
    occupied_by = models.CharField(max_length=150, blank=True, null=True)
    responsible_person = models.CharField(max_length=150, blank=True, null=True)
    attachments = GenericRelation(Attachment, related_query_name='buildings')
    created_at = models.DateTimeField(auto_now_add=True, null=True, editable=False)
    updated_at = models.DateTimeField(auto_now=True, null=True, editable=False)

    def __str__(self):
        return self.name
