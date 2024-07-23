from django.urls import path

from api.views.artist import ArtistListCreateAPIView, ArtistRetrieveUpdateDestroyAPIView

urlpatterns = [
    path('artists/', ArtistListCreateAPIView.as_view(), name='artist-list-create'),
    path('artists/<int:pk>/', ArtistRetrieveUpdateDestroyAPIView.as_view(), name='artist-retrieve-update-destroy'),
]