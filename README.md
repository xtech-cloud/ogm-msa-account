# StartKit

See [omo-msa-startkit](https://github.com/xtech-cloud/omo-msa-startkit)

# 消息订阅

- 地址
  omo.msa.account.notification

- 消息
  | Action | Head | Body|
  |:--|:--|:--|
  |/signup||uuid|
  |/signin|accessToken|uuid|
  |/signout|accessToken||
  |/reset/password|accessToken||
  |/profile/update|accessToken||
  |/profile/query|accessToken||
