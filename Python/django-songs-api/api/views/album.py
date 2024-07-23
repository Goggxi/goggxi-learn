from rest_framework import generics
from rest_framework.response import Response
from api.models import Album
from api.serializers import AlbumSerializer


class AlbumListCreateAPIView(generics.ListCreateAPIView):
    queryset = Album.objects.all()
    serializer_class = AlbumSerializer

    def list(self, request, *args, **kwargs):
        response = super().list(request, *args, **kwargs)
        return Response({
            'message': 'Albums retrieved successfully',
            'data': response.data
        })

    def create(self, request, *args, **kwargs):
        response = super().create(request, *args, **kwargs)
        return Response({
            'message': 'Album created successfully',
            'data': response.data
        })


class AlbumRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    queryset = Album.objects.all()
    serializer_class = AlbumSerializer
