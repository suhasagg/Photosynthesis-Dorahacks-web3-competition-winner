<!-- markdown-link-check-disable -->

# Gaia는 무엇인가요?

가이아(`gaia`)는 코스모스 허브의 코스모스 SDK 애플리케이션의 이름입니다. 가이아
는 두개의 엔트리 포인트로 구성돼있습니다:

*   `gaiad`: 가이아 데몬, `gaia` 애플리케이션의 풀노드를 운영합니다.
*   `gaiad`: 가이아 커맨드 라인 인터페이스는 유저가 가이아 풀노드와 소통할 수 있게
    합니다.

`gaia`는 코스모스 SDK의 다음 모듈들을 이용해 제작되었습니다:

*   `x/auth`: 계정 및 서명
*   `x/bank`: 토큰 전송
*   `x/staking`: 스테이킹 로직
*   `x/mint`: 인플레이션 로직
*   `x/distribution`: 수수료(보상) 분배 로직(fee distribution logic)
*   `x/slashing`: 슬래싱 로직
*   `x/gov`: 거버넌스 로직
*   `ibc-go/modules`: 인터블록체인 전송
*   `x/params`: 앱레벨 파라미터 관리

> 코스모스 허브에 대해서: 코스모스 허브는 코스모스 네트워크의 최초 허브입니다.
> 허브는 블록체인 간 전송을 수행하는 역할을 합니다. IBC를 통해 특정 허브에 연결
> 된 블록체인은 해당 허브에 연결되어있는 모든 블록체인과 함께 연결됩니다. 코스모
> 스 허브는 지분증명 기반 퍼블릭 블록체인이며, 고유 토큰은 아톰(Atom)입니다. 다
> 음은 Gaia를 [설치하는 방법](./installation.md)을 알아보겠습니다.

<!-- markdown-link-check-enable -->
