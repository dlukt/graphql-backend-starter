import type {UserManagerSettings} from 'oidc-client-ts';

export const oidcConfig: UserManagerSettings = {
    authority: import.meta.env.VITE_OIDC_AUTHORITY,
    client_id: import.meta.env.VITE_OIDC_CLIENT_ID,
    redirect_uri: import.meta.env.VITE_OIDC_REDIRECT_URI,
    scope: import.meta.env.VITE_OIDC_SCOPE,
}

