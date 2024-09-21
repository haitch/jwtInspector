'use strict';

import { getInput, info, setOutput } from '@actions/core';

// Getting the 'token' input from the githubAction
const token = getInput('token');

// jwt token is constructed as 'header.payload.signature'
const tokens = token.split(".");

// decode each section
tokens.forEach((token, index) => {
    // signature part, we are not verifying it for now.
    if (index >= 2) {
        return;
    }

    // header part
    if (index === 0) {
        info("Header:");
        // Base64 decode -> Parse -> Pretty print
        info(JSON.stringify(JSON.parse(atob(token)), null, 4));
        return;
    }

    info("Payload:");
    // Base64 decode -> Parse
    const claims = JSON.parse(atob(token));
    info(JSON.stringify(claims, null, 4));

    // output the important claims
    if (claims.iss) {
        setOutput('issuer', claims.iss);
    }
    if (claims.sub) {
        setOutput('subject', claims.sub);
    }
    if (claims.aud) {
        setOutput('audience', claims.aud);
    }

    // output the rest with a prefix
    for (const [key, value] of Object.entries(claims)) {
        setOutput(`claim_${key}`, JSON.stringify(value));
    }
});