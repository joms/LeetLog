LeetLog
=======

An IRC-logger for the channel *#scene.no*, where every day at 13:37 they post an empty post with the goal of hitting 13:37.000 (only been achieved once).   
This logger will log wether a post is valid (inside 13:37:00-13:37:59 and empty), or invalid by being outside or having multiple entries. It'll go active at 13:35 and be active to 13:40. The purpose is most for fun, but also having the ability to make statistics from de data collected.

---

**Log format**

    invalid/valid. A post will be invalid if it triggers any of the follow rules (Could be changed to an representing a reason)
    - Too early (before 13:37:00)
    - Too late (after 13:37:59)
    - Multiple entries inside 13:37. Only the first entry will count
    Date (dd-mm-yyyy)
    Time (hh:mm:ss)
    Milliseconds
    Nick
    Message
    - If not empty, this is what the user posted
    - If only spaces, this is the number of spaces

**Example**

    2014-03-06 13:37:01 124 valid JoMs 1
    2014-03-06 13:37:59 214 invalid JoMs google
    
**GROK**

    %{DATE_EU:date} %{TIME:time} %{WORD:status} %{USERNAME:nick} %{GREEDYDATA:msg}

---

The statistics will be made by runnin the log through Logstash and into ElasticSearch, wich then serves as an API for the website.
