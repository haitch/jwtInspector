'use strict'
const core = require('@actions/core');

const token = core.getInput('token')
const tokens = token.split(".");

tokens.forEach((token, index) => {
    if (index >= 2) {
        return;
    }
    core.info(JSON.stringify(JSON.parse(atob(token)), null, 4));
});