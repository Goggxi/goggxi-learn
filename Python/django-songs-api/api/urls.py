from django.urls import path

from api.views.artist import ArtistListCreateAPIView, ArtistRetrieveUpdateDestroyAPIView
from api.views.album import AlbumListCreateAPIView, AlbumRetrieveUpdateDestroyAPIView

urlpatterns = [
    path('artists/', ArtistListCreateAPIView.as_view(), name='artist-list-create'),
    path('artists/<int:pk>/', ArtistRetrieveUpdateDestroyAPIView.as_view(), name='artist-retrieve-update-destroy'),
    path('albums/', AlbumListCreateAPIView.as_view(), name='album-list-create'),
    path('albums/<int:pk>/', AlbumRetrieveUpdateDestroyAPIView.as_view(), name='album-retrieve-update-destroy'),
]