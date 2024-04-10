# Indeed Job Scraper

This Python script scrapes job listings from Indeed.com based on a specified keyword and location. It utilizes Selenium to automate the process of navigating through search results and extracting job details such as title, company, location, and posting time. The scraped data is then stored in a CSV file.

## Prerequisites

Before running the script, ensure you have the following dependencies installed:

- Python 3.x
- Selenium
- Chrome WebDriver (compatible with your Chrome browser version)

You can install the Python dependencies using pip:

```
pip install selenium webdriver-manager
```

## Usage

1. Clone the repository:

```
git clone https://github.com/your-username/indeed-job-scraper.git
```

2. Run the script:

```
python main.py
```

3. Enter search parameters:

You will be prompted to enter the search keyword (e.g., "Software Engineer") and the location (e.g., "San Jose"). The script will then scrape job listings from Indeed.com based on these parameters.

6. View the results:

Once the scraping process is complete, the scraped job data will be saved to a CSV file in the `results` directory.

[Example](https://github.com/jooyul-yoon/JobScrape/blob/master/results/software_engineer-california.csv)

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
