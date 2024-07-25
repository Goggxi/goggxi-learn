from rest_framework import generics
from management_data.models import Building
from management_data.serializers import BuildingSerializer
from inventory.pagination import StandardResultsSetPagination


class BuildingListCreateView(generics.ListCreateAPIView):
    queryset = Building.objects.all()
    serializer_class = BuildingSerializer
    pagination_class = StandardResultsSetPagination


class BuildingRetrieveUpdateDestroyView(generics.RetrieveUpdateDestroyAPIView):
    queryset = Building.objects.all()
    serializer_class = BuildingSerializer
