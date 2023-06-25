{
    "name": "compile-proto",
    "version": "1.0.0",
    "description": "",
    "scripts": {
      "build-proto": "npm run build-proto:pbjs && npm run build-proto:pbts",
      "build-proto:pbjs": "pbjs --target static-module --wrap ./wrap-pbjs.js --out ../../assets/scripts/models/msg.js ../../assets/model.proto",
      "build-proto:pbts": "pbts --main --out ../../assets/scripts/models/msg.d.ts ../../assets/scripts/models/msg.js && node ./wrap-pbts"
    },
    "author": "",
    "license": "ISC",
    "devDependencies": {
      "fs-extra": "^9.0.0",
      "protobufjs": "^6.8.9"
    }
  }