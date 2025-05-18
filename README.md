# Wikipedia WikiScraper

This project is a simple and concurrent web scraper written in Go using the [Colly](https://github.com/gocolly/colly) and Mitchell (2018) frameworks. It scrapes a list of Wikipedia pages related to intelligent systems and robotics, extracts and summarizes their content, and prints a report.


## Features:

- Scrapes multiple Wikipedia pages concurrently
- Extracts and summarizes main content from each page
- Handles synchronization and concurrency
- Outputs summaries in a readable format (json)

## Wiki Pages Scraped:

- [Intelligent agent](https://en.wikipedia.org/wiki/Intelligent_agent)
- [Android (robot)](https://en.wikipedia.org/wiki/Android_(robot))
- [Reinforcement learning](https://en.wikipedia.org/wiki/Reinforcement_learning)
- [Robotics](https://en.wikipedia.org/wiki/Robotics)
- [Robotic process automation](https://en.wikipedia.org/wiki/Robotic_process_automation)
- [Robot Operating System](https://en.wikipedia.org/wiki/Robot_Operating_System)
- [Robot](https://en.wikipedia.org/wiki/Robot)
- [Software agent](https://en.wikipedia.org/wiki/Software_agent)
- [Chatbot](https://en.wikipedia.org/wiki/Chatbot)
- [Applications of artificial intelligence](https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence)

## Install and Run


1. Install Dependencies
    ```bash
    go mod tidy
    ```

2. To Run the Program:
    ```bash
    go run main.go
    ```



Reference:
Mitchell, Ryan. 2018. Web Scraping with Python: Collecting More Data from the Modern Web (second ed.). Sebastopol, Calif.: O'Reilly. [ISBN-13: 978-1491985571] Source code available at https://github.com/REMitchell/python-scraping  

ChatGPT was used in assisting in generating code and receiving errors. The main issues in coding this program were making sure the output json text was in a readable format. As seen in chatgpt_convo.text, chatgpt successfully assisted in making the program output easier to read.
