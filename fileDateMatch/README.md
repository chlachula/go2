Program **fileDateMatch** has one text file argument. In filename is encoded date in form YYMMDD. Program is searching for line with structure like following example: 
 <date yyyy="2023" mm="11" dd="5"/>
Program with file argument ok231105.xml would end with exit code 0.
Program with file argument no231203.xml would end with exit code 1 and error message.

    Usage:

    go2 fileYYMMDD.xml

    Examples:

    go2 ok231105.xml

    go2 no231203.xml    
---
[Written with Markdown](https://www.markdownguide.org/basic-syntax/)