/**
 * @type {Promise<ServerConfig> | null}
 * @private
 */
let _serverConfig = null;

export async function serverConfig() {
  if (_serverConfig === null) {
    /** @type {string} */
    const url = import.meta.env.VITE_CONFIG_URL || '/__/scos/config.json';
    console.debug('server config path:', url);
    _serverConfig = fetch(url)
        .then(res => /** @type {Promise<ServerConfig>} */ res.json())
        .catch(() => ({}));
  }
  // todo: retry on network failure
  return _serverConfig;
}

/**
 * @returns {Promise<string>}
 */
export async function grpcWebEndpoint() {
  const envEndpoint = import.meta.env.VITE_GRPC_ENDPOINT || '';
  if (envEndpoint) {
    return envEndpoint;
  }

  // Attempt to read the config from a well known path on the server that hosts this application.
  const config = await serverConfig();
  const address = configAddress(config);
  if (!address) {
    // If there's no configured address we assume we're connecting to the same server we are hosting on,
    // using the same scheme that we are served under.
    return '//' + location.host
  }
  const protocol = (config.insecure || !config.httpsAddress) ? 'http://' : 'https://';
  return protocol + address;
}

function configAddress(config) {
  const address = (config.insecure ? config.httpAddress : (config.httpsAddress || config.httpAddress));
  if (address?.[0] === ':') {
    // no host
    return location.hostname + address;
  }
  return address;
}
