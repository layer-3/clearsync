# 0005 - Sub-key Management

## Status

Work in progress

## Context

End user must use external wallet key management, such as ledger, Metamask or WalletConnect but it must be able to delegate the signing of channels states using another key with a background process which will auto-sign from the cloud.

End User MUST be able to set expiration time of the subkey, revoke issued subkey, counterparties should be able to verify that this subkey was certified by the correct end-user.

## Decision

Certainly! A lightweight certification protocol for your scenario could utilize JSON Web Tokens (JWT), which are compact, URL-safe means of representing claims to be transferred between two parties. JWTs can include metadata and are digitally signed using a key pair. Here's a simplified example of how Alice can use JWTs to certify the `clearport` subkey:

### Certification Protocol Overview

1. **Key Generation:** Both Alice and `clearport` generate their key pairs. Alice's keypair is her wallet keypair, and `clearport` generates a subkey using KMS.
2. **Certificate Creation:** Alice creates a JWT to serve as the certificate for the `clearport` subkey. The JWT contains claims about the `clearport` subkey, such as the public key, ENS, expiration, and other metadata.
3. **Signing Certificate:** Alice signs the JWT using her private key.
4. **Verification:** Any party can verify the JWT using Alice's public key.
5. **Revocation:** To revoke the certificate, Alice can maintain a revocation list or use a smart contract to flag the revoked subkeys.

### Example Code

Here's a simple example using Ruby with the `ruby-jwt` gem for JWT handling. This is just to illustrate the process; you'll need to adapt it to your exact requirements and ensure security best practices.

```ruby
require 'jwt'
require 'openssl'

# Alice's wallet keypair
alice_private_key = OpenSSL::PKey::RSA.generate(2048)
alice_public_key = alice_private_key.public_key

# Clearport's subkey from KMS
clearport_subkey = OpenSSL::PKey::RSA.generate(2048)
clearport_subkey_public = clearport_subkey.public_key.to_pem

# Current time and expiration time for the JWT
current_time = Time.now.to_i
expiration_time = current_time + 365 * 24 * 60 * 60 # 1 year from now

# Certificate (JWT) payload with standard claims
claims = {
  iss: "0xAliceWalletAddress",          # Issuer: Alice wallet
  ens: "alice.eth"			  			# Name: Alice
  sub: "cert:0xClearport_subkey",       # Subject: Purpose of the JWT
  aud: "canarynet",                     # Audience: Intended recipients
  exp: expiration_time,                 # Expiration Time
  nbf: current_time,                    # Not Before
  iat: current_time,                    # Issued At
  jti: SecureRandom.uuid                # JWT ID: Unique identifier for the JWT
}

# Alice signs the JWT using her private key
cert_token = JWT.encode(claims, alice_private_key, 'RS256')

# Verification (any party would use Alice's public key)
begin
  decoded_token = JWT.decode(cert_token, alice_public_key, true, { algorithm: 'RS256' })
  puts "Certificate is valid: #{decoded_token}"
rescue JWT::DecodeError => e
  puts "Invalid certificate: #{e.message}"
end

# Revocation method
def revoke_subkey(subkey_id)
  # Logic to revoke the key, e.g., update a list or a smart contract
  puts "Subkey #{subkey_id} revoked."
end

```

## Consequences

- Use JWTs for creating a certificate for the `clearport` subkey.
- Include metadata like the public key, ENS, wallet address, and expiration in the JWT claims.
- Alice signs the JWT with her private key to issue the certificate.
- Verification is done using Alice's public key.
- Revocation can be handled with a list or smart contract.

Remember to handle key management securely, use appropriate cryptographic algorithms, and implement proper error handling in your actual code.

