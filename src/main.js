'use strict'
const core = require('@actions/core');

const token = core.getInput('username')
const tokens = token.split(".");

tokens.forEach((token, index) => {
    core.info(JSON.parse(atob(token)));
});