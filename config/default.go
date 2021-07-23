package config

const defaultYAML string = `
service:
  name: xtc.ogm.account
  address: :18801
  ttl: 15
  interval: 10
logger:
  level: trace
  dir: /var/log/ogm/
database:
  # 驱动类型，可选值为 [sqlite,mysql]
  driver: sqlite
  mysql:
    address: localhost:3306
    user: root
    password: mysql@XTC
    db: ogm
  sqlite:
    path: /tmp/ogm-account.db
token:
  jwt:
    # 过期时间（小时）
    expiry: 12
    # jwt密钥
    secret: c56de585baa85b8d689116a391371035
`
