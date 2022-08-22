import {keycloak} from './keycloak.js';

/**
 * @returns {Promise<string | null>}
 */
export async function apiToken() {
  return keycloak()
      .then(kc => kc.token)
      .catch(() => null);
}