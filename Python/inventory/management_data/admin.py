from django.contrib import admin
from .models import Building, Room, ImmovableItem, MovableItem, IntangibleItem, Attachment


# @admin.register(Building)
# class BuildingAdmin(admin.ModelAdmin):
#     list_display = ('code', 'name', 'floors', 'total_area', 'responsible_person')
#     search_fields = ('code', 'name', 'responsible_person__username')
#
#
# @admin.register(Room)
# class RoomAdmin(admin.ModelAdmin):
#     list_display = ('code', 'name', 'floor', 'area', 'building', 'responsible_person')
#     search_fields = ('code', 'name', 'building__name', 'responsible_person__username')
#     list_filter = ('building', 'floor')
#
#
# @admin.register(ImmovableItem)
# class ImmovableItemAdmin(admin.ModelAdmin):
#     list_display = ('code', 'name', 'quantity', 'year_acquired', 'category', 'building', 'room', 'responsible_person')
#     search_fields = ('code', 'name', 'building__name', 'room__name', 'responsible_person__username')
#     list_filter = ('state', 'category', 'building', 'room')
#
#
# @admin.register(MovableItem)
# class MovableItemAdmin(admin.ModelAdmin):
#     list_display = ('code', 'name', 'quantity', 'year_acquired', 'brand', 'model', 'condition', 'building', 'room',
#                     'responsible_person')
#     search_fields = ('code', 'name', 'brand', 'model', 'building__name', 'room__name', 'responsible_person__username')
#     list_filter = ('state', 'condition', 'building', 'room')
#
#
# @admin.register(IntangibleItem)
# class IntangibleItemAdmin(admin.ModelAdmin):
#     list_display = ('code', 'name', 'quantity', 'year_acquired', 'version', 'serial_number', 'responsible_person')
#     search_fields = ('code', 'name', 'version', 'serial_number', 'responsible_person__username')
#     list_filter = ('state',)
#
#
# @admin.register(Attachment)
# class AttachmentAdmin(admin.ModelAdmin):
#     list_display = ('file', 'content_type', 'object_id', 'created_at')
#     search_fields = ('file', 'content_type__model')
#     list_filter = ('content_type',)
