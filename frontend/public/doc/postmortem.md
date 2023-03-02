# Postmortem: Storage Service went offline due to Oracle DB failure


|     |     |
| --- | --- |
| Owner 		| Long (long@) |
| Collaborators | Kevin (kevin@), Ryan (ryan@) |
| Status     	| draft |
| Created  		| 2023-03-02 |
| Last Updated	| 2023-03-02 |


## Executive Summary

|     |     |
| --- | --- |
| Impact | On 2023-03-01, Storage Service was down from 10:23 AM to 7:00 PM, caused data processing jobs hang or fail, client meetings cancelled. |
| Root Cause | Oracle cluster had a hardware failure. |


## Incident Summary

|     |     |
| --- | --- |
| Duration 			| 8 hours 37 minutes |
| Severity  		| High |
| Products affected | data processing, client meetings |
| User impact  		| 100% of data processing and QC |
| Revenue impact 	| 1% of annual revenue |
| Detection 		| alert, user report |
| Resolution 		| database restore from backup |


## Background

![Storage Service](storage/storage_service.svg)


## Impact

### User Impact

* All CPU and GPU processing jobs hang on new file creation.
* Data Manager web UI and command line tool failed to work.
* Data QC app was not able to load datasets from Storage Service.
* Client meetings were cancelled for 3/1.

### Revenue Impact

* ~1% of annual revenue.


## Root Causes and Trigger

### Trigger

* A hardware failure caused Oracle DB and OS crashed.


### Root Cause Analysis

Metadata

- Why Storage Service was down?
  - Oracle DB failed
- Why Oracle DB failure caused Storage Servcie offline?
  - Oracle DB stores the metadata (including location) of all files
- Why metadata unavailable caused Storage Service offline?
  - existing files could not be located
  - new files could not be assigned to a disk host


Oracle DB

- Why Oracle failed?
  - primary node had a hardware failure.
- Why failover didn't bring Oracle back online?
  - standby DB was overloaded by a retry storm, DB was not responding and eventually crashed.
  - we built in too much business logic into DB, with heavy PL/SQL Packages.
- Why restoring the DB takes 6 hours?
  - tables are big.
  - schema constraints are complicated.


## Lessons Learned

Metadata is the SPOF for Storeage Service.


### Things that went well

* Detection was fast. DBA oncall got alerts in 3 minutes.
* There was no major data loss.


### Things that went poorly

* DB failover was not tested in production, the failover was not successful.
* data processing jobs and offline jobs created a retry storm against Oracle DB.
* DB restoration took longer than expected.


### Where we got lucky

* We had daily backups for Oracle DB.
* DB restoration was eventually succeeded.


## Action Items

|     |Type | Owner | Priority | Task |
| --- | --- | ----- | -------- | ---- |
| Replace failed memory for Oracle node | mitigation | Adam | p0  | T111  |
| Restart Storage Service offline jobs | mitigation | Adam | p0  | T111  |
| Add retry backoff from client jobs | prevention | Adam | p0  | T111  |
| Implement file-based metadata storage | prevention | Adam | p0  | T111  |
| Add rate limitting and adminsion control on API gateway | prevention | Brain | p1   | T122 |


## Timeline

```
2023-03-01 (all times in US/Pacific)

	10:23 AM DB hardware failure, OS and DB went offline.

	10:25 AM DBA oncall got alerts.
	10:26 AM User reported Data Manager web UI show black page.
	10:26 AM User reported QC app cannot open new datasets.

	10:28 AM DB failover to standby.
	10:33 AM standby DB crashed.

	10:40 AM Storage Service major outage claimed, communication sent out to eng-team.
```


```
	10:47 AM standby reboot didn't work, DB tables were corrupted.
	11:00 AM DB restore started from backup

	 1:00 AM All processing jobs killed

	 5:00 PM DB restoration finished, DB back online
```


```
	 5:10 PM DB saw lock contentions on files table

	 5:15 PM All Storage Service offline jobs stopped.

	 6:45 PM Storage Service offline jobs started gradually.

	 7:00 PM Incident migrated, User start resubmitting processing jobs.
```


```
2023-03-02

    Started fixing minor missed files and data loss for users.
```