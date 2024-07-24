from rest_framework import generics
from management_data.models import Building
from management_data.serializers import BuildingSerializer
from inventory.pagination import StandardResultsSetPagination


class BuildingListCreateAPIView(generics.ListCreateAPIView):
    serializer_class = BuildingSerializer
    pagination_class = StandardResultsSetPagination

    def get_queryset(self):
        return Building.objects.all().order_by('created_at')

class BuildingRetrieveUpdateDestroyAPIView(generics.RetrieveUpdateDestroyAPIView):
    queryset = Building.objects.all()
    serializer_class = BuildingSerializer
