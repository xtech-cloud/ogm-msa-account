package config

const defaultYAML string = `
service:
  name: xtc.api.ogm.account
  address: :9601
  ttl: 15
  interval: 10
logger:
  level: trace
  dir: /var/log/ogm/
database:
  lite: true
  mysql:
    address: localhost:3306
    user: root
    password: mysql@OMO
    db: ogm_account
  sqlite:
    path: /tmp/ogm-account.db
encrypt:
  secret: 964E50CA8F603714BF373A4C03E07739
token:
  jwt:
    expiry: 1
`
