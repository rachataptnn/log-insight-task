## Log Report

#### 1. Error log number and ratio analysis
    Total: %s
    Error: %s
    Error Ratio: %s
    
    Overall:
        <%s> (loop every element in map)
        ex. INFO: <%s>
            ERROR: <%s>

#### 2. Log analysis by HTTP status code
    Overall:
        <%s> (loop every element in map)
        ex. 200: 14
            404: 6
            500: 1

#### 3. Response time analysis
    All Routes:
        Total Req: <%s> req 
        Slow Req:  <%s> req
        SlowRate:  <%s>%
        Min: <%s> ms
        Avg: <%s> ms
        Max: <%s> ms
        

    Each Routes
        <%s> (loop every element in map)
        ex. "Key1":
            Total Req: <%s> req 
            Slow Req:  <%s> req
            SlowRate:  <%s>%
            Min: <%s> ms
            Avg: <%s> ms
            Max: <%s> ms

#### 4. Parse the request URI
        <%s> (loop every element in map)
        ex. API.com : 15
            API.another1.com : 7
            API.another2.com : 3
            API.another4.com : 3
            API.another5.com : 1

#### 5. Analysis by time period 
        <%s> (loop every element in map)
        ex. Z : 204
            +07:00 : 15
            -05:00 : 2
        