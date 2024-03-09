แปล
- สร้าง system ที่สามารถวิเคราะห์และ process log data ขนาดใหญ่ **ได้อย่างมีประสิทธิภาพ**
- process แต่ละ line ของ log file พร้อมๆกันกับ วิเคราะห์ log ไปด้วย เช่น
  - นับ log (นับว่ากี่บรรทัดใช่ป่าววะ)
  - ตรวจสอบ log รูปแบบเฉพาะ
- ผลการวิเคราะห์จะต้องถูกพ่นออกทาง console หรือ ถูกเซฟเป็นไฟล์
- สุดท้าย. ระบบที่จะออกแบบเนี่ย ให้คำนึงถึงประสิทธิภาพการประมวลผล และ การ memory usage อย่างคุ้มๆด้วย
1. Error log number and ratio analysis
   - ข้อนี้อยากให้คำนวนจำนวนของ logs แต่ละ level เช่น
     - ERROR : 7 ea
     - WARNING : 2024 ea
     - DEBUG : 100 ea
2. Log analysis by HTTP status code:
   - แล้วก็คำนวนด้วยว่าจำนวนของ HTTP status code เนี่ย แต่ละ code มีเท่าไหร่ เช่น
      - 200: 19000 ea
      - 500: 10 ea
      - 404: 180 ea
3. คำนวน responnse time 
   - เค้าบอกว่าให้ extract response time ออกมาจาก log file แล้วก็คำนวนและแสดง output ออกมาว่า 
     - avg response time เท่าไหร่ 
     - min response time เท่าไหร่
     - max response time เท่าไหร่
     - stat สุดท้าย -> คำนวนมาว่า มี request กี่ % ที่ผ่าน threshold (ถ้า request ใช้เวลา >= 500ms ถือว่าช้า)
4. parse request url ออกมาด้วย
   - หามาว่า domain ไหนที่ถูกยิงไปมากที่สุด จาก log file ที่ให้มานิ, print ออกมาด้วย
   - เอา 5 อันดับพอ
5. Analysis by time period
   - คำนวนและ output ออกมาให้ดูหน่อยว่า ความถี่ของ request ที่แต่ละ time zone อะ เป็นไง(กุต้องทำ graph, diagram รึวาดอะไรออกมาโชว์มั้ยวะ?) 

* Build a system that can efficiently analyze and process large amounts of log data
* Process each line in the log file simultaneously to perform log analysis (e.g. error log count, log detection of specific pattern, etc.).
* Analysis results must be output to the console or saved as a new file.
* Design the system considering optimized processing performance and memory usage

1. Error log number and ratio analysis:
  Calculate the number of logs for error levels (“ERROR”, “WARNING”, “DEBUG”) in the log file, calculate the ratio for each, and output.

1. Log analysis by HTTP status code:
  Calculate and output the number of logs for each HTTP status code in the log file.

1. Response time analysis:
  Extract the response time from the log file and calculate and output the average response time, minimum response time, and maximum response time. Calculate and output the percentage of requests with response times exceeding the specified threshold (500ms).
  
1. Parse the request URI:
  Find the most requested domains in the log file and print them, up to 5th place in ranking.

1. Analysis by time period 
  Analyze and output request frequency by time zone.

1. (Additional items to do if you have time) Build with cli

2. Branch Rule: Use the format feature/name for branch names.

objective
