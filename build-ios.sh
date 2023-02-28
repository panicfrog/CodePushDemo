#/usr/bin/sh

set -e 

npx react-native bundle \
  --platform ios \
  --dev false \
  --entry-file index.js \
  --bundle-output out/js/ios/main.jsbundle \
  --sourcemap-output out/js/ios/main.jsbundle.map \
  --assets-dest out/assets/ios/

node_modules/react-native/sdks/hermesc/osx-bin/hermesc \
  -O \
  -emit-binary \
  -output-source-map \
  -out=out/hbc/ios/main.bundle.hbc \
  out/js/ios/main.jsbundle
