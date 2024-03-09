# Log Report

### 1. Error log number and ratio analysis
	Total: 20000 Error: 638

	Error Ratio: 3.19%

| Level | Amount |
|-----------|-------|
| ERROR | 638 |
| DEBUG | 668 |
| INFO | 17948 |
| WARNING | 746 |

---


<br>

### 2. Log analysis by HTTP status code
| Code | Amount |
|-----------|-------|
| 200 | 9362 |
| 401 | 92 |
| 503 | 78 |
| 500 | 97 |
| 502 | 91 |
| 404 | 84 |
| log doesn't have status code | 10000 |
| 400 | 86 |
| 403 | 110 |

---


<br>

### 3. Response time analysis
	All Routes:
		Min: 1 ms
		Avg: 300.63 ms
		Max: 999 ms
		
		Total Req: 9982 req 
		Slow Req:  990 req
		SlowRate:  9.92
	
	Each Routes
		/payment:
				Min: 1 ms
				Avg: 306.58 ms
				Max: 999 ms

				Total Req: 1132 req 
				Slow Req:  120 req
				SlowRate:  10.60%

		/user:
				Min: 1 ms
				Avg: 301.17 ms
				Max: 998 ms

				Total Req: 1125 req 
				Slow Req:  107 req
				SlowRate:  9.51%

		/notification:
				Min: 1 ms
				Avg: 300.05 ms
				Max: 982 ms

				Total Req: 1119 req 
				Slow Req:  110 req
				SlowRate:  9.83%

		/order:
				Min: 1 ms
				Avg: 302.41 ms
				Max: 998 ms

				Total Req: 1088 req 
				Slow Req:  103 req
				SlowRate:  9.47%

		/chat:
				Min: 1 ms
				Avg: 297.57 ms
				Max: 997 ms

				Total Req: 1097 req 
				Slow Req:  99 req
				SlowRate:  9.02%

		/cart:
				Min: 1 ms
				Avg: 298.57 ms
				Max: 991 ms

				Total Req: 1110 req 
				Slow Req:  111 req
				SlowRate:  10.00%

		/product:
				Min: 1 ms
				Avg: 295.25 ms
				Max: 993 ms

				Total Req: 1125 req 
				Slow Req:  110 req
				SlowRate:  9.78%

		/trade:
				Min: 1 ms
				Avg: 304.03 ms
				Max: 998 ms

				Total Req: 1097 req 
				Slow Req:  116 req
				SlowRate:  10.57%

		/search:
				Min: 1 ms
				Avg: 300.04 ms
				Max: 998 ms

				Total Req: 1089 req 
				Slow Req:  114 req
				SlowRate:  10.47%





<br>

### 4. Parse the request URI
| Host | Amount |
|-----------|-------|
| https://api.example.com | 9982 |

---


<br>

### 5. Analysis by time period 
| Timezone | Amount |
|-----------|-------|
| Z | 9982 |

---
