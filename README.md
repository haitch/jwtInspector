# jwtInspector
A GitHub Action to decode and dump JWT tokens locally without external communication. It does not verify the token signature or expiration.

## Spec
#### Input
- `token`, the jwt token to inspect

#### Output
- `issuer`, from jwt `iss` claim
- `subject`, from jwt `sub` claim
- `audience`, from jwt `aud` claim
- `claim_{key}`, this is output for each other claim in the token, i.e. `ref_type` claim in token would be output as `claim_ref_type`.

## Usage
```yaml
    - name: Get Id Token
      uses: actions/github-script@v6
      id: idtoken
      with:
        script: |
          let id_token = await core.getIDToken('test')
          core.setOutput('id_token', id_token)
          core.info(id_token.length)

    - name: Token Inspect
      id: tokenInspect
      uses: haitch/jwtInspector
      with:
        token: ${{ steps.idtoken.outputs.id_token }}

    - name: output consumer example
      run : |
        echo "Issuer: ${{ steps.tokenInspect.outputs.issuer }}"
        echo "Subject: ${{ steps.tokenInspect.outputs.subject }}"
        echo "Audience: ${{ steps.tokenInspect.outputs.audience }}"
        echo "ref_type: ${{ steps.tokenInspect.outputs.claim_ref_type }}"
        echo "event_name: ${{ steps.tokenInspect.outputs.claim_event_name }}"
```

#### Output sample
```
Header:
{
    "typ": "JWT",
    "alg": "RS256",
    "x5t": "Hyq4NATAjsnqC7mdrtAhhrCR2_Q",
    "kid": "1F2AB83404C08EC9EA0BB99DAED02186B091DBF4"
}
Payload:
{
    "jti": "381ca463-fd09-4f7e-b3ee-77f6fc06b7f7",
    "sub": "repo:haitch/jwtInspector:environment:Production",
    "aud": "test",
    "ref": "refs/heads/main",
    "sha": "8b42ec0b6b973f5465ddc4e667cee9ba26aeabb5",
    "repository": "haitch/jwtInspector",
    "repository_owner": "haitch",
    "repository_owner_id": "1167635",
    "run_id": "10988497055",
    "run_number": "19",
    "run_attempt": "1",
    "repository_visibility": "public",
    "repository_id": "860122172",
    "actor_id": "1167635",
    "actor": "haitch",
    "workflow": "TokenTest",
    "head_ref": "",
    "base_ref": "",
    "event_name": "push",
    "ref_protected": "false",
    "ref_type": "branch",
    "workflow_ref": "haitch/jwtInspector/.github/workflows/tokenTest.yml@refs/heads/main",
    "workflow_sha": "8b42ec0b6b973f5465ddc4e667cee9ba26aeabb5",
    "environment": "Production",
    "environment_node_id": "EN_kwDOM0RsPM7xAwGo",
    "job_workflow_ref": "haitch/jwtInspector/.github/workflows/tokenTest.yml@refs/heads/main",
    "job_workflow_sha": "8b42ec0b6b973f5465ddc4e667cee9ba26aeabb5",
    "runner_environment": "github-hosted",
    "iss": "https://token.actions.githubusercontent.com",
    "nbf": 1727070587,
    "exp": 1727071487,
    "iat": 1727071187
}
```

```
Issuer: https://token.actions.githubusercontent.com
Subject: repo:haitch/jwtInspector:environment:Production
Audience: test
ref_type: branch
event_name: push
```
