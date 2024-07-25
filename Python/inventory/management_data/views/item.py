from rest_framework import generics
from management_data.models import ImmovableItem, MovableItem, IntangibleItem
from management_data.serializers import ImmovableItemSerializer, MovableItemSerializer, IntangibleItemSerializer
from inventory.pagination import StandardResultsSetPagination


class ImmovableItemListCreateView(generics.ListCreateAPIView):
    queryset = ImmovableItem.objects.all()
    serializer_class = ImmovableItemSerializer
    pagination_class = StandardResultsSetPagination


class ImmovableItemRetrieveUpdateDestroyView(generics.RetrieveUpdateDestroyAPIView):
    queryset = ImmovableItem.objects.all()
    serializer_class = ImmovableItemSerializer


class MovableItemListCreateView(generics.ListCreateAPIView):
    queryset = MovableItem.objects.all()
    serializer_class = MovableItemSerializer
    pagination_class = StandardResultsSetPagination


class MovableItemRetrieveUpdateDestroyView(generics.RetrieveUpdateDestroyAPIView):
    queryset = MovableItem.objects.all()
    serializer_class = MovableItemSerializer


class IntangibleItemListCreateView(generics.ListCreateAPIView):
    queryset = IntangibleItem.objects.all()
    serializer_class = IntangibleItemSerializer
    pagination_class = StandardResultsSetPagination


class IntangibleItemRetrieveUpdateDestroyView(generics.RetrieveUpdateDestroyAPIView):
    queryset = IntangibleItem.objects.all()
    serializer_class = IntangibleItemSerializer
