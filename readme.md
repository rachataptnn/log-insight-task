# Message From Candidate

## Testing
1. Update the log file name in `config.yaml`.
2. Execute the command `go run .` to run the program.

## Example File Issue
- Regarding the provided file ([logfile.txt](https://github.com/rachataptnn/log-insight-task/blob/feature/log-analyzer/logfile.txt)), I didn't find records for other timezones or hosts. Thus, I added some based on my understanding and created a new file called  [logfileAdded.txt](https://github.com/rachataptnn/log-insight-task/blob/feature/log-analyzer/logfileAdded.txt).


## Result File Location
- The result file is expected to be located in the folder `/result-reports`. The name is changed based on the date and time. Additionally, I have already generated the [example file](https://github.com/rachataptnn/log-insight-task/blob/feature/log-analyzer/result-reports/log-report24-03-10T02-08.md)
<br><br><br><br>

# Origin Message
* Build a system that can efficiently analyze and process large amounts of log data
* Process each line in the log file simultaneously to perform log analysis (e.g. error log count, log detection of specific pattern, etc.).
* Analysis results must be output to the console or saved as a new file.
* Design the system considering optimized processing performance and memory usage
1. Error log number and ratio analysis:
Calculate the number of logs for error levels (“ERROR”, “WARNING”, “DEBUG”) in the log file, calculate the ratio for each, and output.
2. Log analysis by HTTP status code:
Calculate and output the number of logs for each HTTP status code in the log file.
3. Response time analysis:
Extract the response time from the log file and calculate and output the average response time, minimum response time, and maximum response time.
Calculate and output the percentage of requests with response times exceeding the specified threshold (500ms).
4. Parse the request URI:
Find the most requested domains in the log file and print them, up to 5th place in ranking.
5. Analysis by time period
Analyze and output request frequency by time zone.
6. (Additional items to do if you have time) Build with cli
7. Branch Rule: Use the format feature/name for branch names.
