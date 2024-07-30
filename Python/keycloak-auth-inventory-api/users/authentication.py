from rest_framework.authentication import BaseAuthentication
from rest_framework.exceptions import AuthenticationFailed
from keycloak import KeycloakOpenID
from django.conf import settings
from .models import Employee


class KeycloakAuthentication(BaseAuthentication):
    def authenticate(self, request):
        auth_header = request.META.get('HTTP_AUTHORIZATION')
        if not auth_header:
            return None

        try:
            token = auth_header.split()[1]
            keycloak_openid = KeycloakOpenID(
                server_url=settings.KEYCLOAK_SERVER_URL,
                client_id=settings.KEYCLOAK_CLIENT_ID,
                realm_name=settings.KEYCLOAK_REALM,
                client_secret_key=settings.KEYCLOAK_CLIENT_SECRET
            )
            token_info = keycloak_openid.introspect(token)

            if not token_info['active']:
                raise AuthenticationFailed('Invalid token')

            employee = Employee.objects.get(keycloak_id=token_info['sub'])
            return (employee.user, None)
        except Exception as e:
            raise AuthenticationFailed('Invalid token')
