### Dell iDRAC Exporter in Go

### How to code
1. create new metric in config/variables.go
2. if metric in chassis -> create file with metric_name in chassis/ folder
3. define new struct `type PowerControl struct{}`
4. define 2 interface `Describe` and `Collect`
   1. Describe interface -> use for description of metric
   2. Collect interface -> use for added value/ label value
5. declare and register metric in main.go

