package config

const defaultYAML string = `
service:
  name: omo.api.msa.account
  address: :9601
  ttl: 15
  interval: 10
logger:
  level: trace
  dir: /var/log/msa/
database:
  lite: true
  mysql:
    address: 127.0.0.1:3306
    user: root
    password: mysql@OMO
    db: msa_account
  sqlite:
    path: /tmp/msa-account.db
encrypt:
  secret: 964E50CA8F603714BF373A4C03E07739
token:
  jwt:
    expiry: 1
publisher:
- /auth/signup
- /auth/signin
- /auth/signout
- /auth/resetpassword
- /profile/update
`
