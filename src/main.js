'use strict'
const core = require('@actions/core');

const token = core.getInput('token')
const tokens = token.split(".");

tokens.forEach((token, index) => {
    core.info(`token part ${index + 1}: ${token}`);
    core.info(atob(token));
});