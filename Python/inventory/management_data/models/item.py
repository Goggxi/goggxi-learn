from django.db import models
from django.contrib.auth.models import User
from django.contrib.contenttypes.fields import GenericRelation, GenericForeignKey
from django.contrib.contenttypes.models import ContentType
from management_data.models import Building, Room, Attachment


class Item(models.Model):
    STATE_CHOICES = [
        (1, 'Internal'),
        (2, 'BMN'),
    ]

    code = models.CharField(max_length=50, unique=True, help_text="item code")
    name = models.CharField(max_length=255, help_text="item name")
    state = models.CharField(max_length=10, choices=STATE_CHOICES, help_text="item state (BMN or Internal)", null=True)
    quantity = models.PositiveIntegerField(default=1, help_text="item quantity")
    responsible_person = models.ForeignKey(User, on_delete=models.SET_NULL, null=True,
                                           help_text="responsible person for this item")
    description = models.TextField(blank=True, help_text="item description")
    qr_code = models.SlugField(max_length=255, unique=True, help_text="QR code")

    building = models.ForeignKey(Building, on_delete=models.SET_NULL, null=True, blank=True, related_name='items')
    room = models.ForeignKey(Room, on_delete=models.SET_NULL, null=True, blank=True, related_name='items')

    content_type = models.ForeignKey(ContentType, null=True, on_delete=models.SET_NULL)
    object_id = models.PositiveIntegerField()
    content_object = GenericForeignKey('content_type', 'object_id')

    attachments = GenericRelation(Attachment, related_query_name='item')

    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    class Meta:
        abstract = True


class ImmovableItem(Item):
    building = models.ForeignKey(
        Building,
        related_name='immovable_items',
        on_delete=models.CASCADE
    )
    room = models.ForeignKey(
        Room,
        related_name='immovable_items',
        on_delete=models.CASCADE
    )

    CATEGORY_CHOICES = [
        ('FURNITURE', 'Furniture'),
        ('ELECTRONIC', 'Elektronik'),
        ('OTHER', 'Lainnya'),
        # Add more category here
    ]

    is_storage = models.BooleanField(default=False, help_text="is this item a storage?")
    category = models.CharField(max_length=20, choices=CATEGORY_CHOICES, help_text="category of the item")
    item_type = models.CharField(max_length=255, help_text="type of the item")

    def __str__(self):
        return f"{self.name} ({self.get_category_display()})"


class MovableItem(Item):
    building = models.ForeignKey(
        Building,
        related_name='movable_items',
        on_delete=models.CASCADE
    )
    room = models.ForeignKey(
        Room,
        related_name='movable_items',
        on_delete=models.CASCADE
    )


    CONDITION_CHOICES = [
        ('GOOD', 'Baik'),
        ('DAMAGED', 'Rusak'),
        ('REPAIR', 'Dalam Perbaikan'),
        # Add more condition here
    ]

    brand = models.CharField(max_length=255, help_text="item brand")
    model = models.CharField(max_length=255, help_text="item model")
    acquisition_date = models.DateField(help_text="date of acquisition")
    acquisition_value = models.IntegerField(help_text="acquisition value in Rupiah")
    annual_tax = models.IntegerField(help_text="annual tax in Rupiah")
    chassis_number = models.CharField(max_length=255, blank=True, help_text="chassis number")
    engine_number = models.CharField(max_length=255, blank=True, help_text="engine number")
    sale_value = models.IntegerField(help_text="sale value in Rupiah")
    service_interval = models.PositiveIntegerField(help_text="service interval in months")
    condition = models.CharField(max_length=10, choices=CONDITION_CHOICES, help_text="item condition")
    license_plate = models.CharField(max_length=20, blank=True, help_text="license plate")
    annual_depreciation = models.IntegerField(help_text="annual depreciation in Rupiah")

    def __str__(self):
        return f"{self.brand} {self.model} - {self.name}"


class IntangibleItem(Item):
    building = models.ForeignKey(
        Building,
        related_name='intangible_items',
        on_delete=models.CASCADE
    )
    room = models.ForeignKey(
        Room,
        related_name='intangible_items',
        on_delete=models.CASCADE
    )

    version = models.CharField(max_length=50, help_text="version")
    serial_number = models.CharField(max_length=255, unique=True, help_text="serial number")
    acquisition_date = models.DateTimeField(help_text="acquisition date")
    expiry_date = models.DateTimeField(help_text="expiry date")

    def __str__(self):
        return f"{self.name} v{self.version}"
