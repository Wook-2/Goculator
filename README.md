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
	3. Now Goculator is On.<br>
	<img src ="https://github.com/Wook-2/CS_Summary/blob/main/image/goculator.PNG?raw=true" width = "250px"></img>
		- First, you have to type full numerical expression like "10 + 20"<br>
		<img src ="https://github.com/Wook-2/CS_Summary/blob/main/image/1.gif?raw=true" width="250px"></img>
		- Next, 2 ways to type
			1. number + operator + number ex) "20 * 5"<br> It will save recent result in storage<br>
			<img src="https://github.com/Wook-2/CS_Summary/blob/main/image/2.gif?raw=true" width="250px"></img>
			2. operator + number ex) "- 3"<br>
			<img src="https://github.com/Wook-2/CS_Summary/blob/main/image/3.gif?raw=true" width="250px"></img>
		- You can load saved number by typing like "s[0] + 20" or "- s[1]"<br>
		<img src="https://github.com/Wook-2/CS_Summary/blob/main/image/4.gif?raw=true" width="250px"></img>

### Will be update
- [ ] Log, 삼각함수, 나머지연산 기능 추가
- [x] 새로운 계산을 시작할 때 기존 결과값을 저장하는 기능
- [x] 저장된 결과값을 불러올 수 있는 기능
- [ ] 다양한 상황에서의 에러, 예외처리
- [x] UX/UI 개선을 위한 Display함수
- [ ] webapp으로 확장
