# sap-api-time-value-converter
sap-api-time-value-converter は、主にエッジコンピューティング環境において、[sap-sandbox](https://github.com/latonaio/sap-sandbox) が対象とする SAP APIs ならびに 各 SAP API Integrations の Runtimes について、当該 Runtimes において Get または
Post する Json 内の時間のフォーマットを変換するマイクロサービスです。

## 動作環境  
sap-api-time-value-converter は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  

* エッジ Kubernetes （推奨）  
* AION のリソース （推奨)  
* OS: LinuxOS （必須）  
* CPU: ARM/AMD/Intel（いずれか必須）  
* Golang Runtime 

## クラウド環境での利用
sap-api-time-value-converter は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## sap-api-time-value-converter による変換例
[sap-sandbox](https://github.com/latonaio/sap-sandbox) が対象とする SAP APIs ならびに 各 SAP API Integrations を callした際の時間のフォーマットは、time_converter_test.goの以下の部分にある \/Date(1642757478000)\/ のフォーマットで返されます。  
本マイクロサービスではこのフォーマットを、他のマイクロサービスや APIs で活用するための適切なフォーマットに変換できます。
この場合、時間のフォーマットである　\/Date(1642757478000)\/　が time.Date(2022, 1, 21, 9, 31, 18, 0, time.UTC)　に変換されています。

```
		func() testStr {
			return testStr{
				name: "OK now time",
				args: args{
					sapTime: `\/Date(1642757478000)\/`,
				},
				want: time.Date(2022, 1, 21, 9, 31, 18, 0, time.UTC),
			}
```

## go.mod / go.sum
sap-api-time-value-converter は、ライブラリであり、go.mod / go.sum に設定することで、他のレポジトリやランタイムで実行できます。  
sap-api-time-value-converter は、[sap-sandbox](https://github.com/latonaio/sap-sandbox)における SAP APIs ならびに 各 SAP API Integrations の Runtimes を対象としています。  
sap-api-time-value-converter は、マイクロサービスとして利用されることができます。  

