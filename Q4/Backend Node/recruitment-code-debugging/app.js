const express = require("express");
const authService = require("./src/services/authService");
const AuthCallbackService = require("./src/services/authCallbackService");

const config = require("./src/config");

const app = express();

app.get("/", (req, res) => {
  const auth = authService.redirectUri();
  res.redirect(auth);
});

app.get("/oauth-github-callback", (req, res) => {
  return AuthCallbackService.callback(req, res);
});

app.listen(config.port);
console.log(`App listening on http://localhost:${config.port}`);
