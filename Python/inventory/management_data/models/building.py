from django.db import models
from django.contrib.auth.models import User
from django.contrib.contenttypes.fields import GenericRelation
from management_data.models import Room, IntangibleItem, MovableItem, ImmovableItem, Attachment


class Building(models.Model):
    code = models.CharField(max_length=50, unique=True, help_text="building code")
    name = models.CharField(max_length=255, help_text="building name")
    width = models.FloatField(help_text="width in meter")
    length = models.FloatField(help_text="length in meter")
    asset_change_value = models.IntegerField(help_text="asset change value in IDR")
    maintenance_estimate = models.PositiveIntegerField(help_text="maintenance estimate in month")
    last_renovation = models.PositiveIntegerField(help_text="last renovation year")
    build_year = models.PositiveIntegerField(help_text="build year")
    description = models.TextField(blank=True, help_text="building description")
    latitude = models.FloatField(help_text="latitude")
    longitude = models.FloatField(help_text="longitude")
    total_area = models.FloatField(help_text="area in square meter")
    floors = models.PositiveIntegerField(help_text="number of floors")
    occupied_by = models.CharField(max_length=255, help_text="building occupied by")
    responsible_person = models.ForeignKey(User, on_delete=models.SET_NULL, null=True,
                                           help_text="responsible person for this building")

    rooms = GenericRelation(Room, related_query_name='building')
    movable_items = GenericRelation(MovableItem, related_query_name='building')
    immovable_items = GenericRelation(ImmovableItem, related_query_name='building')
    intangible_items = GenericRelation(IntangibleItem, related_query_name='building')
    attachments = GenericRelation(Attachment, related_query_name='building')

    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.name
