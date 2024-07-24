from rest_framework import serializers
from management_data.models.building import Building
from management_data.serializers.attachment import AttachmentSerializer


class BuildingSerializer(serializers.ModelSerializer):
    attachments = AttachmentSerializer(many=True, read_only=True)

    class Meta:
        model = Building
        fields = '__all__'
