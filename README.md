# jwtInspector
A github action to decode/dump JWT token, locally only, no external communication. it doesn't verify the token signature, or expiration.

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