import {
  Environment,
  Network,
  RecordSource,
  Store,
  type FetchFunction,
} from "relay-runtime";
import {oidcConfig} from './oidc-config.ts';
import {User} from 'oidc-client-ts';

const HTTP_ENDPOINT = import.meta.env.VITE_HTTP_ENDPOINT;
const file = "file"
const files = "files"

function getUser() {
  let oidcStorage = localStorage.getItem(`oidc.user:${oidcConfig.authority}:${oidcConfig.client_id}`)
  if (!oidcStorage) {
    oidcStorage = sessionStorage.getItem(`oidc.user:${oidcConfig.authority}:${oidcConfig.client_id}`)
    if (!oidcStorage) {
      return null;
    }
  }

  return User.fromStorageString(oidcStorage);
}

const fetchFn: FetchFunction = async (request, variables, _cacheConfig, uploadables) => {
  const reqLabel = `[Relay] ${request.name}`
  if (import.meta.env.DEV) {
    try {
      // Simple client-side trace for debugging stuck loads
      console.debug(`${reqLabel} ->`, variables)
    } catch {
      // ignore
    }
  }
  let payload
  const user = getUser()
  const access_token = user?.access_token

  const headers = new Headers()
  headers.set("Accept", "application/graphql-response+json; charset=utf-8, application/json; charset=utf-8")
  // headers.set("Content-Type", "application/json; charset=utf-8")
  if (access_token) {
    headers.set("Authorization", `Bearer ${access_token}`)
  }

  if (uploadables) {
    if (!window.FormData) {
      throw new Error("Uploading files without `FormData` not supported.");
    }
    const formData = new FormData();
    formData.append(
        'operations',
        JSON.stringify({
          query: request.text,
          variables: variables,
        })
    )
    if (file in variables) {
      formData.append('map', JSON.stringify({'0': ['variables.file']}))
    } else if (files in variables) {
      const fo: Map<string, Array<string>> = new Map()
      for (const uploadable in uploadables) {
        if (Object.prototype.hasOwnProperty.call(uploadables, uploadable)) {
          fo.set(uploadable, [`variables.files.${uploadable}`])
        }
      }
      formData.append('map', JSON.stringify(fo));
    } else {
      console.error("uploadables provided, but no file or files in variables.")
      return
    }

    for (const uploadable in uploadables) {
      if (Object.prototype.hasOwnProperty.call(uploadables, uploadable)) {
        formData.append(uploadable, uploadables[uploadable])
      }
    }
    payload = {
      method: "POST",
      headers: headers,
      body: formData
    }
  } else {
    headers.set("Content-Type", "application/json; charset=utf-8")
    payload = {
      method: "POST",
      headers: headers,
      body: JSON.stringify({
        query: request.text, // <-- The GraphQL document composed by Relay
        variables,
      }),
    }
  }

  // Add a safety timeout so we can surface stuck requests as errors in dev
  const controller = new AbortController()
  const timeoutMs = 30000
  const timeout = setTimeout(() => controller.abort(), timeoutMs)
  try {
    const resp = await fetch(HTTP_ENDPOINT, { ...payload, signal: controller.signal })
    const isJson = resp.headers.get('content-type')?.includes('json')
    const data = isJson ? await resp.json() : undefined
    if (!resp.ok) {
      if (import.meta.env.DEV) {
        console.error(`${reqLabel} <- HTTP ${resp.status}`, data ?? await resp.text())
      }
      throw new Error(`GraphQL HTTP ${resp.status}`)
    }
    if (import.meta.env.DEV) {
      try { 
        console.debug(`${reqLabel} <- ok`)
      } catch {
        // ignore
      }
    }
    return data
  } catch (err) {
    if (import.meta.env.DEV) {
      console.error(`${reqLabel} <- error`, err)
    }
    throw err
  } finally {
    clearTimeout(timeout)
  }
};


function createRelayEnvironment() {
  return new Environment({
    network: Network.create(fetchFn),
    store: new Store(new RecordSource()),
  });
}

export const RelayEnvironment = createRelayEnvironment();

