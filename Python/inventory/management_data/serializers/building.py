from rest_framework import serializers
from management_data.models.building import Building
from management_data.serializers.attachment import AttachmentSerializer


class BuildingSerializer(serializers.ModelSerializer):
    items = serializers.PrimaryKeyRelatedField(many=True, read_only=True)
    attachments = AttachmentSerializer(many=True, read_only=True)

    class Meta:
        model = Building
        fields = '__all__'
        read_only_fields = ['created_at', 'updated_at']
