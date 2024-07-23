from rest_framework import serializers
from api.models import Album
from api.serializers.artist import ArtistSerializer


class AlbumSerializer(serializers.ModelSerializer):
    artist = ArtistSerializer(read_only=True)

    class Meta:
        model = Album
        fields = [
            'id',
            'title',
            'release_date',
            'artist',
            'created_at',
            'updated_at',
        ]
        read_only_fields = ['created_at', 'updated_at']