from rest_framework import serializers
from management_data.models import Room, Building
from management_data.serializers.attachment import AttachmentSerializer


class RoomSerializer(serializers.ModelSerializer):
    building = serializers.PrimaryKeyRelatedField(queryset=Building.objects.all())
    items = serializers.PrimaryKeyRelatedField(many=True, read_only=True)
    attachments = AttachmentSerializer(many=True, read_only=True)

    class Meta:
        model = Room
        fields = '__all__'
        read_only_fields = ['created_at', 'updated_at']
