'use strict'
const core = require('@actions/core');

const token = core.getInput('username')
const tokens = token.split(".");

tokens.forEach((token, index) => {
    core.info(`token part ${index + 1}: ${token}`);
    core.info(`token part ${index + 1}: ${token}`);
    core.info(JSON.parse(atob(token)));
});