const config = require("../config");

function redirectUri() {
    return `${config.oauthUrl}/authorize?client_id=${config.clientId}`;
  }

module.export = {
    redirectUri: redirectUri
}
