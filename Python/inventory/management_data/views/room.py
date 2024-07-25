from rest_framework import generics
from management_data.models import Room
from management_data.serializers import RoomSerializer
from inventory.pagination import StandardResultsSetPagination


class RoomListCreateView(generics.ListCreateAPIView):
    queryset = Room.objects.all()
    serializer_class = RoomSerializer
    pagination_class = StandardResultsSetPagination


class RoomRetrieveUpdateDestroyView(generics.RetrieveUpdateDestroyAPIView):
    queryset = Room.objects.all()
    serializer_class = RoomSerializer
