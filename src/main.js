'use strict'
const core = require('@actions/core');

const token = core.getInput('token')
const tokens = token.split(".");

tokens.forEach((token, index) => {
    if (index >= 2) {
        return;
    }
    if (index === 0) {
        core.info("Header:");
        core.info(JSON.stringify(JSON.parse(atob(token)), null, 4));
        return;
    }

    core.info("Payload:");
    const claims = JSON.parse(atob(token))
    core.info(JSON.stringify(claims, null, 4));
    if (claims.iss) {
        core.setOutput('issuer', claims.iss);
    }
    if (claims.sub) {
        core.setOutput('subject', claims.sub);
    }
    if (claims.aud) {
        core.setOutput('audience', claims.aud);
    }

    for (const [key, value] of Object.entries(object1)) {
        core.setOutput(`claim_${key}`, JSON.stringify(value));
        core.debug(`adding output claim_${key}: ${JSON.stringify(value)}`);
    }
});