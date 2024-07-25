from rest_framework import serializers
from management_data.models import Building, Room, ImmovableItem, MovableItem, IntangibleItem
from management_data.serializers.attachment import AttachmentSerializer


class ImmovableItemSerializer(serializers.ModelSerializer):
    building = serializers.PrimaryKeyRelatedField(queryset=Building.objects.all(), allow_null=True, required=False)
    room = serializers.PrimaryKeyRelatedField(queryset=Room.objects.all(), allow_null=True, required=False)
    attachments = AttachmentSerializer(many=True, read_only=True)

    class Meta:
        model = ImmovableItem
        fields = '__all__'
        read_only_fields = ['created_at', 'updated_at']

    def validate(self, data):
        building = data.get('building')
        room = data.get('room')
        if room and not building:
            raise serializers.ValidationError("Jika ruangan diisi, bangunan juga harus diisi.")
        if room and building and room.building != building:
            raise serializers.ValidationError("Ruangan harus berada dalam bangunan yang dipilih.")
        return data


class MovableItemSerializer(serializers.ModelSerializer):
    attachments = AttachmentSerializer(many=True, read_only=True)

    class Meta:
        model = MovableItem
        fields = '__all__'
        read_only_fields = ['created_at', 'updated_at']


class IntangibleItemSerializer(serializers.ModelSerializer):
    attachments = AttachmentSerializer(many=True, read_only=True)

    class Meta:
        model = IntangibleItem
        fields = '__all__'
        read_only_fields = ['created_at', 'updated_at']
