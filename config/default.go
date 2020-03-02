package config

const defaultYAML string = `
service:
    ttl: 15
    interval: 10
logger:
    level: 5
database:
    lite: true
    mysql:
        address: 127.0.0.1:3306
        user: root
        password: mysql@OMO
        db: account
    sqlite:
        path: /tmp/msa-account.db
encrypt:
    secret: 964E50CA8F603714BF373A4C03E07739
`
