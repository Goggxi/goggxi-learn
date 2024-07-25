from django.db import models
from django.contrib.auth.models import User
from django.contrib.contenttypes.fields import GenericRelation, GenericForeignKey
from django.contrib.contenttypes.models import ContentType
from management_data.models import IntangibleItem, MovableItem, ImmovableItem, Attachment


class Room(models.Model):
    code = models.CharField(max_length=50, unique=True, help_text="room code")
    name = models.CharField(max_length=255, help_text="room name")
    floor = models.PositiveIntegerField(help_text="number of floor")
    length = models.FloatField(help_text="length in meter")
    width = models.FloatField(help_text="width in meter")
    area = models.FloatField(help_text="area in square meter")
    responsible_person = models.ForeignKey(User, on_delete=models.SET_NULL, null=True,
                                           help_text="responsible person for this room")
    description = models.TextField(blank=True, help_text="room description")

    content_type = models.ForeignKey(ContentType, null=True, on_delete=models.SET_NULL)
    object_id = models.PositiveIntegerField()
    content_object = GenericForeignKey('content_type', 'object_id')

    movable_items = GenericRelation(MovableItem, related_query_name='building')
    immovable_items = GenericRelation(ImmovableItem, related_query_name='building')
    intangible_items = GenericRelation(IntangibleItem, related_query_name='building')
    attachments = GenericRelation(Attachment, related_query_name='building')

    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self):
        return f"{self.name} - {self.content_object.name}"
