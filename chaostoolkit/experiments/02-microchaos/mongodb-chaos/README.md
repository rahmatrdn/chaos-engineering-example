To run the experiment and generate a PDF report, use the following commands:

### Toxic Latency
The "Toxic Latency" experiment is designed to test the resilience of the MongoDB system by introducing artificial latency to its connections. By running the commands below, you can observe how the application responds to network delays and generate a PDF report of the experiment results.

```bash
-- Run and Generate Report
chaos run add-toxic-latency.json --journal-path=reports/report-add-toxic-latency.json

-- Convert report to PDF
chaos report --export-format=pdf reports/report-add-toxic-latency.json reports/report-add-toxic-latency.pdf
```

### Toxic Bandwidth Degradation
The "Toxic Bandwidth Degradation" experiment is designed to test the resilience of the MongoDB system by introducing artificial bandwidth degradation to its connections. By running the commands below, you can observe how the application responds to reduced network bandwidth and generate a PDF report of the experiment results.

```bash
-- Run and Generate Report
chaos run add-toxic-bandwidth-degradation.json --journal-path=reports/report-add-toxic-bandwidth-degradation.json

-- Convert report to PDF
chaos report --export-format=pdf reports/report-add-toxic-bandwidth-degradation.json reports/report-add-toxic-bandwidth-degradation.pdf
```

### Toxic Timeout
The "Toxic Timeout" experiment is designed to test the resilience of the MongoDB system by introducing artificial timeouts to its connections. By running the commands below, you can observe how the application responds to forced timeouts and generate a PDF report of the experiment results.

```bash
-- Run and Generate Report
chaos run add-toxic-timeout.json --journal-path=reports/report-add-toxic-timeout.json

-- Convert report to PDF
chaos report --export-format=pdf reports/report-add-toxic-timeout.json reports/report-add-toxic-timeout.pdf
```

### Toxic Slow Close
The "Toxic Slow Close" experiment is designed to test the resilience of the MongoDB system by introducing artificial delays when closing connections to the database. This simulates scenarios where network connections are not closed promptly, which can impact application performance and resource usage. By running the commands below, you can observe how the application responds to slow connection closes and generate a PDF report of the experiment results.

```bash
-- Run and Generate Report
chaos run add-toxic-slow-close.json --journal-path=reports/report-add-toxic-slow-close.json

-- Convert report to PDF
chaos report --export-format=pdf reports/report-add-toxic-slow-close.json reports/report-add-toxic-slow-close.pdf
```

### Toxic Slicer
The "Toxic Slicer" experiment is designed to test the resilience of the MongoDB system by introducing artificial packet fragmentation (slicer toxic) to its connections. This simulates scenarios where network packets are split into smaller fragments, potentially affecting the application's ability to process database responses efficiently. By running the commands below, you can observe how the application responds to packet fragmentation and generate a PDF report of the experiment results.

```bash
-- Run and Generate Report
chaos run add-toxic-slicer.json --journal-path=reports/report-add-toxic-slicer.json

-- Convert report to PDF
chaos report --export-format=pdf reports/report-add-toxic-slicer.json reports/report-add-toxic-slicer.pdf
```

### Toxic Limiter
The "Toxic Limiter" experiment is designed to test the resilience of the MongoDB system by introducing artificial connection limiting (limiter toxic) to its connections. This simulates scenarios where the bandwidth or number of bytes allowed through the connection is restricted, potentially impacting the application's ability to communicate with the database efficiently. By running the commands below, you can observe how the application responds to connection limiting and generate a PDF report of the experiment results.

```bash
-- Run and Generate Report
chaos run add-toxic-limiter.json --journal-path=reports/report-add-toxic-limiter.json

-- Convert report to PDF
chaos report --export-format=pdf reports/report-add-toxic-limiter.json reports/report-add-toxic-limiter.pdf
```