# 미션 - 자동차 경주

## 🔍 진행 방식

- 미션은 **기능 요구 사항, 프로그래밍 요구 사항, 과제 진행 요구 사항** 세 가지로 구성되어 있다.
- 세 개의 요구 사항을 만족하기 위해 노력한다. 특히 기능을 구현하기 전에 기능 목록을 만들고, 기능 단위로 커밋 하는 방식으로 진행한다.
- 기능 요구 사항에 기재되지 않은 내용은 스스로 판단하여 구현한다.

## 📮 미션 제출 방법

- 미션 구현을 완료한 후 GitHub을 통해 제출해야 한다.
- go-racing-car 저장소를 개인 저장소로 Fork 한다.
- 개인 저장소에서 본인의 Github ID로 브랜치를 분기하여 미션을 구현한다.
- 미션 구현이 완료되면, `poi1649/go-racing-car` [저장소](https://github.com/poi1649/go-racing-car)에 자신의 Github ID에 해당하는 브랜치에 Pull Request를 생성한다.

### 테스트 실행 가이드

- 터미널에서 `go version`을 실행하여 go 버전이 1.21.5인지 확인한다.
- 터미널에서 `go test` 명령을 실행하고, 명령을 실행할 때 모든 테스트가 아래와 같이 통과하는지 확인한다.

```
=== RUN   TestHello
--- PASS: TestHello (0.00s)
PASS
```

---

## 🚀 기능 요구 사항

초간단 자동차 경주 게임을 구현한다.

- 주어진 횟수 동안 n대의 자동차는 전진 또는 멈출 수 있다.
- 각 자동차에 이름을 부여할 수 있다. 전진하는 자동차를 출력할 때 자동차 이름을 같이 출력한다.
- 자동차 이름은 쉼표(,)를 기준으로 구분하며 이름은 5자 이하만 가능하다.
- 사용자는 몇 번의 이동을 할 것인지를 입력할 수 있어야 한다.
- 전진하는 조건은 0에서 9 사이에서 무작위 값을 구한 후 무작위 값이 4 이상일 경우이다.
- 자동차 경주 게임을 완료한 후 누가 우승했는지를 알려준다. 우승자는 한 명 이상일 수 있다.
- 우승자가 여러 명일 경우 쉼표(,)를 이용하여 구분한다.
- 사용자가 잘못된 값을 입력할 경우 `panic`을 발생시킨 후 애플리케이션은 종료되어야 한다. (추후 수정 예정)

### 입출력 요구 사항

#### 입력

- 경주 할 자동차 이름(이름은 쉼표(,) 기준으로 구분)

```
poi, bebe, phili
```

- 시도할 회수

```
5
```

#### 출력

- 각 차수별 실행 결과

```
poi : --
bebe : ----
phili : ---
```

- 단독 우승자 안내 문구

```
최종 우승자 : poi
```

- 공동 우승자 안내 문구

```
최종 우승자 : poi, bebe
```

#### 실행 결과 예시

```
경주할 자동차 이름을 입력하세요.(이름은 쉼표(,) 기준으로 구분)
poi,bebe,phili
시도할 회수는 몇회인가요?
5

실행 결과
poi : -
bebe : 
phili : -

poi : --
bebe : -
phili : --

poi : ---
bebe : --
phili : ---

poi : ----
bebe : ---
phili : ----

poi : -----
bebe : ----
phili : -----

최종 우승자 : poi, phili
```

---

## 🎯 프로그래밍 요구 사항

- Go 1.21.5 버전에서 실행 가능해야 한다. 
- 프로그램 실행의 시작점은 `application.go`의 `main()`이다.
- [Go 공식 코딩 가이드](https://go.dev/doc/effective_go)와 [뱅크 샐러드 Go 코딩 컨벤션](https://blog.banksalad.com/tech/go-best-practice-in-banksalad/)을 준수하며 프로그래밍한다.
  - 커밋을 하기 전 반드시 **gofmt** 기능을 이용해 코드를 포맷팅한다.
- 프로그램 구현이 완료되면 모든 테스트가 성공해야 한다. 테스트는 `go test` 명령으로 실행할 수 있다.
- 프로그래밍 요구 사항에서 달리 명시하지 않는 한 파일, 패키지 이름을 수정하거나 이동하지 않는다.


---

## ✏️ 과제 진행 요구 사항

- 미션은 [go-racing-car](https://github.com/poi1649/go-racing-car) 저장소를 Fork & Clone해 시작한다.
- **기능을 구현하기 전 `docs/README.md`에 구현할 기능 목록을 정리**해 추가한다.
- **Git의 커밋 단위는 앞 단계에서 `docs/README.md`에 정리한 기능 목록 단위**로 추가한다.
    - [커밋 메시지 컨벤션](https://gist.github.com/stephenparish/9941e89d80e2bc58a153) 가이드를 참고해 커밋 메시지를 작성한다.