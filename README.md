LeetLog
=======

An IRC-logger for the channel *#scene.no*, where every day at 13:37 they post an empty post with the goal of hitting 13:37.000 (only been achieved once).   
This logger will log whether a post is valid (inside 13:37:00-13:37:59 and empty), or invalid by being outside or having multiple entries. It'll go active at 13:35 and be active to 13:40. The purpose is most for fun, but also having the ability to make statistics from de data collected.

You can check out the cool frontend which'll contain more stats as time goes [here](http://joms.github.io/LeetLog/)

---

**Log format**

    invalid/valid. A post will be invalid if it triggers any of the follow rules (Could be changed to an representing a reason)
    - Too early (before 13:37:00)
    - Too late (after 13:37:59)
    - Multiple entries inside 13:37. Only the first entry will count
    Date (yyyy/mm/dd)
    Time (hh:mm:ss:ms)
    Channel
    Status
    Nick

**Status**

    Valid/Invalid reasons:
    0 = Valid
    1 = Before 13:37
    2 = Text in 13:37
    3 = Already entered
    4 = After 13:37
    5 = Empty string before 13:37
    6 = Empty string after 13:37
    
**GROK**

    %{DATE_EU:date} %{TIME:time} %{INT:status} %{USERNAME:nick} %{GREEDYDATA:msg}

---

The statistics will be made by running the log through Logstash and into ElasticSearch, which then serves as an API for the website.   
Elastic should only be available locally by doing a block in the firewall.
