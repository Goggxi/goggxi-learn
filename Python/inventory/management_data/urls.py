from django.contrib import admin
from django.urls import path
from management_data.views import BuildingListCreateAPIView, BuildingRetrieveUpdateDestroyAPIView


urlpatterns = [
    path('buildings/', BuildingListCreateAPIView.as_view(), name='building-list-create'),
    path('buildings/<int:pk>/', BuildingRetrieveUpdateDestroyAPIView.as_view(), name='building-retrieve-update-destroy'),
]