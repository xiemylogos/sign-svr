# sign-svr

offline signature of the original transaction

### 1、use ./sign-svr signtx cli sign unsigned_transaction
example:
```
./sign-svr signtx --address AGc9NrdF5MuMJpkFfZ3MWKa67ds6H2fzud --payeraddr Af6xrG7WB9wUKQ3aRDXnfba2G5DXjqejMS --rawtx  00d10585b4f4c409000000000000204e000000000000ffe723aefd01bac311d8b16ff8bfd594d77f31ee7100c66b14092118e0112274581b60dfb6fedcbfdcfc044be76a7cc814ffe723aefd01bac311d8b16ff8bfd594d77f31ee6a7cc8516a7cc86c51c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b65000051c1087472616e736665721400000000000000000000000000000000000000010068164f6e746f6c6f67792e4e61746976652e496e766f6b650000

Password:
Password:

publickey := 03944e3ff777b14add03a76fd6767aaf4a65c227ec201375d9118d4e6b272494c7

sigdata := 190d5bb879964e83ab71d5e85b2c1387c7fe7ac51f15833b38d40f44142f6e2fa529cfe28593487a93b9396d8e86ceee6aaf2171c901390067aaa3b2e4dd179a

publickey := 02263e2e1eecf7a45f21e9e0f865510966d4e93551d95876ecb3c42acf2b68aaae

sigdata := 0ae2b459177b40851c3d8fa529d88042e30c66500de551b368c4a9f60a199aa81b8fc6412305f8aef1a4c6700d4711642e56c6e87bd9bdd8c446ba1c289dfee6

```

### 2、use ./sign-svr signtxhash cli sign txhash 

```
./sign-svr signtxhash --address AGc9NrdF5MuMJpkFfZ3MWKa67ds6H2fzud --txhash f7079fd564d84e68601b2a2802924aa35a089aef54d3df364d9058ff2a142f12

Password:
signed txHash:1bf1bd2421cd010eeb402101cf0a15ca3fca8e6c91792086f5fb084a91936c870c6747ed8c2d3a8b08a9735594cc182571d27800984006968e2a52dfd25b175d
```




