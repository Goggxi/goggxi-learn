from django.urls import path
from management_data.views import (
    BuildingListCreateView,
    BuildingRetrieveUpdateDestroyView,
    RoomListCreateView,
    RoomRetrieveUpdateDestroyView,
    ImmovableItemListCreateView,
    ImmovableItemRetrieveUpdateDestroyView,
    MovableItemListCreateView,
    MovableItemRetrieveUpdateDestroyView,
    IntangibleItemListCreateView,
    IntangibleItemRetrieveUpdateDestroyView,

)

urlpatterns = [
    path('buildings/', BuildingListCreateView.as_view(), name='building-list-create'),
    path('buildings/<int:pk>/', BuildingRetrieveUpdateDestroyView.as_view(), name='building-detail'),
    path('rooms/', RoomListCreateView.as_view(), name='room-list-create'),
    path('rooms/<int:pk>/', RoomRetrieveUpdateDestroyView.as_view(), name='room-detail'),
    path('immovable-items/', ImmovableItemListCreateView.as_view(), name='immovable-item-list-create'),
    path('immovable-items/<int:pk>/', ImmovableItemRetrieveUpdateDestroyView.as_view(),
         name='immovable-item-detail'),
    path('movable-items/', MovableItemListCreateView.as_view(), name='movable-item-list-create'),
    path('movable-items/<int:pk>/', MovableItemRetrieveUpdateDestroyView.as_view(), name='movable-item-detail'),
    path('intangible-items/', IntangibleItemListCreateView.as_view(), name='intangible-item-list-create'),
    path('intangible-items/<int:pk>/', IntangibleItemRetrieveUpdateDestroyView.as_view(),
         name='intangible-item-detail')
]
