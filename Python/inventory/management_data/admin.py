from django.contrib import admin
from .models import Building

@admin.register(Building)
class AdminBuilding(admin.ModelAdmin):
    list_filter = ('name',)
    ordering = ('id',"name",)
    list_display = [field.name for field in Building._meta.fields]

