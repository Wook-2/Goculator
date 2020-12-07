# Goculator
##### 고언어로 만든 계산기
### Feature
- 사칙연산(+, -, *, / ), 제곱승 계산(^)
- Continuous calculate .

### Requirements
- golang 1.10.4 linux/amd64

### Usage
- #### Command Line Interface
	1.  Clone this repo<br>
	`git clone https://github.com/Wook-2/Goculator.git`
	2. Run main.go<br>
	`go run main.go`
	3. you can see like below<br>
	<img src ="https://github.com/Wook-2/CS_Summary/blob/main/image/goculator.PNG?raw=true" width = "300px"></img>
		- First, you have to type full numerical expression like "10 + 20"
		- Next, 2 ways to type
			1. number + operator + number ex) "20 * 5"<br> It will save recent result in storage
			2. operator + number ex) "- 3"
		- You can load saved number by typing like "s[0] + 20" or "- s[1]"

### Will be update
- [ ] Log, 삼각함수, 나머지연산 기능 추가
- [x] 새로운 계산을 시작할 때 기존 결과값을 저장하는 기능
- [x] 저장된 결과값을 불러올 수 있는 기능
- [ ] 다양한 상황에서의 에러, 예외처리
- [x] UX/UI 개선을 위한 Display함수
- [ ] webapp으로 확장
