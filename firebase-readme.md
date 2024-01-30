# firebase emulator

required node.js >= 16.0 or java JDK >= 11
1. install firebase CLI
'''
curl -sL https://firebase.tools or npm install -g firebase-tools
'''
2. login at the first time
'''
firebase login:ci (--no-localhost)
'''
Token will be generated, save it to FIREBASE_TOKEN env var, system will automatically use the token. 
Or run all firebase commands with the --token TOKEN flag.
3. check your project id in firebase.json is correct

# run emulator
```
firebase emulators:start
```