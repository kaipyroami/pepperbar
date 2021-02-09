# pepperbar

---

A command line web scraper to RSS utility.
The mission of this project is to facilitate the cron-able generation of RSS feeds from whatever web content you want.



##Examples

---

###Example of twitter scraper usage

```bash
pepperbar twitter --count=5 --username=Twitter --format=atom --outfile=feed.xml
```

### Example output of ```feed.xml```
```bash
<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title>Feed generated with pepperbar</title>
  <id>https://github.com/kaipyroami/pepperbar</id>
  <updated>2021-02-08T20:20:08-08:00</updated>
  <subtitle>A command line web scraper to RSS utility.</subtitle>
  <link href="https://github.com/kaipyroami/pepperbar"></link>
  <author>
    <name>Kyle Crockett</name>
    <email>kyle@kyle-crockett.com</email>
  </author>
  <entry>
    <title></title>
    <updated>2021-02-08T11:30:27-08:00</updated>
    <id>1358860764732350465</id>
    <link href="https://twitter.com/Twitter/status/1358860764732350465" rel="alternate"></link>
    <summary type="html">memes have expiration dates</summary>
    <author>
      <name>Twitter</name>
    </author>
  </entry>
  <entry>
    <title></title>
    <updated>2021-02-04T11:58:07-08:00</updated>
    <id>1357418177941147648</id>
    <link href="https://twitter.com/Twitter/status/1357418177941147648" rel="alternate"></link>
    <summary type="html">thinking about oomf</summary>
    <author>
      <name>Twitter</name>
    </author>
  </entry>
  <entry>
    <title></title>
    <updated>2021-02-03T08:19:19-08:00</updated>
    <id>1357000724593442816</id>
    <link href="https://twitter.com/Twitter/status/1357000724593442816" rel="alternate"></link>
    <summary type="html">it&#39;s ok to Tweet before you text back</summary>
    <author>
      <name>Twitter</name>
    </author>
  </entry>
  <entry>
    <title></title>
    <updated>2021-02-02T07:46:00-08:00</updated>
    <id>1356629952398036995</id>
    <link href="https://twitter.com/Twitter/status/1356629952398036995" rel="alternate"></link>
    <summary type="html">5 likes is a lot</summary>
    <author>
      <name>Twitter</name>
    </author>
  </entry>
  <entry>
    <title></title>
    <updated>2021-01-02T10:19:52-08:00</updated>
    <id>1345434652891373571</id>
    <link href="https://twitter.com/Twitter/status/1345434652891373571" rel="alternate"></link>
    <summary type="html">RT @djarinculture: *me rereading my own tweet every time someone likes it* https://t.co/Q764BDICa9</summary>
    <author>
      <name>Twitter</name>
    </author>
  </entry>
</feed>
```