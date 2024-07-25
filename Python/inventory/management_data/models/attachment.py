from django.db import models
from django.contrib.contenttypes.fields import GenericForeignKey
from django.contrib.contenttypes.models import ContentType


def attachment_upload_path(instance, filename):
    return f'attachments/{instance.content_type.model}/{instance.object_id}/{filename}'


class Attachment(models.Model):
    file = models.FileField(upload_to=attachment_upload_path, help_text="File attachment")
    content_type = models.ForeignKey(ContentType, null=True, on_delete=models.SET_NULL)
    object_id = models.PositiveIntegerField()
    content_object = GenericForeignKey('content_type', 'object_id')

    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self):
        return f"a file attachment for {self.content_object}"
