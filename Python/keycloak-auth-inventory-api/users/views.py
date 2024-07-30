from rest_framework import viewsets, status
from rest_framework.decorators import action
from rest_framework.response import Response
from rest_framework.permissions import AllowAny, IsAdminUser
from django.contrib.auth.models import User
from keycloak import KeycloakAdmin, KeycloakOpenID
from .models import Employee
from .serializers import EmployeeSerializer
from django.conf import settings


class EmployeeViewSet(viewsets.ModelViewSet):
    queryset = Employee.objects.all()
    serializer_class = EmployeeSerializer

    def get_permissions(self):
        if self.action in ['create', 'destroy']:
            permission_classes = [IsAdminUser]
        elif self.action == 'list':
            permission_classes = [IsAdminUser]
        else:
            permission_classes = [AllowAny]
        return [permission() for permission in permission_classes]

    @action(detail=False, methods=['post'])
    def register(self, request):
        keycloak_admin = KeycloakAdmin(
            server_url=settings.KEYCLOAK_SERVER_URL,
            username='admin',
            password='admin',
            realm_name=settings.KEYCLOAK_REALM,
            client_id=settings.KEYCLOAK_CLIENT_ID,
            client_secret_key=settings.KEYCLOAK_CLIENT_SECRET,
            verify=True)

        user_data = {
            'username': request.data['username'],
            'email': request.data['email'],
            'firstName': request.data['first_name'],
            'lastName': request.data['last_name'],
            'enabled': True,
            'credentials': [{'type': 'password', 'value': request.data['password'], 'temporary': False}]
        }

        keycloak_user = keycloak_admin.create_user(user_data)

        django_user = User.objects.create_user(
            username=request.data['username'],
            email=request.data['email'],
            first_name=request.data['first_name'],
            last_name=request.data['last_name'],
            password=request.data['password']
        )

        employee = Employee.objects.create(
            user=django_user,
            keycloak_id=keycloak_user,
            role='staff'
        )

        serializer = self.get_serializer(employee)
        return Response(serializer.data, status=status.HTTP_201_CREATED)

    @action(detail=False, methods=['post'])
    def login(self, request):
        keycloak_openid = KeycloakOpenID(
            server_url=settings.KEYCLOAK_SERVER_URL,
            client_id=settings.KEYCLOAK_CLIENT_ID,
            realm_name=settings.KEYCLOAK_REALM,
            client_secret_key=settings.KEYCLOAK_CLIENT_SECRET
        )

        token = keycloak_openid.token(
            username=request.data['username'],
            password=request.data['password']
        )

        return Response(token)

    @action(detail=False, methods=['post'])
    def logout(self, request):
        keycloak_openid = KeycloakOpenID(
            server_url=settings.KEYCLOAK_SERVER_URL,
            client_id=settings.KEYCLOAK_CLIENT_ID,
            realm_name=settings.KEYCLOAK_REALM,
            client_secret_key=settings.KEYCLOAK_CLIENT_SECRET
        )

        keycloak_openid.logout(request.data['refresh_token'])
        return Response(status=status.HTTP_204_NO_CONTENT)
