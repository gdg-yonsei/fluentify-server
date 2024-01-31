## Firebase Emulator

### Run Firebase Emulator using Docker
```
docker build -t 'name:tag' .
docker run -p 9199:9199 -p 9099:9099 -p 9000:9000 -p 8080:8080 -p 4000:4000 'name:tag' firebase emulators:start --token 'FIREBASE_TOKEN'
```
- FIREBASE_TOKEN will be generated at the first time when you login with google email used for creating firebase project
- project.default in .firebaserc need to be changed into each firebase project/GCP id