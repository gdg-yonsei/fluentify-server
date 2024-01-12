# Contributing Guide

## Branch Guide

- 브랜치 이름은 `이름/설명` 형식으로 작성한다. (Ex. `jason/add-api-logging`)

## Commit Message

### 포맷

`type: subject`

```
- fix: 버스 픽스
- feat: 새로운 기능 추가
- refactor: 리팩토링 (버그픽스나 기능추가없는 코드변화)
- docs: 문서만 변경
- style: 코드의 의미가 변경 안 되는 경우 (띄어쓰기, 포맷팅, 줄바꿈 등)
- test: 테스트코드 추가/수정
- chore: 빌드 테스트 업데이트, 패키지 매니저를 설정하는 경우 (프로덕션 코드 변경 X)
```

## Pull Request

- 모든 코드의 변경은 PR을 통해서만 이루어진다.
- 필수적으로 PR Approve를 받아야 한다. PR을 올린 이후엔 리뷰어를 지정한다.
- Review 를 받은 내용을 반영했다면 해당 커밋 해시을 답글로 달고 리뷰를 재요청한다.
