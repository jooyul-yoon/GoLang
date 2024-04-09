import csv
import time
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


def to_csv(data, filename):
    with open(filename+'.csv', 'w', newline='') as csvfile:
        writer = csv.writer(csvfile)
        writer.writerow(["Title", "Company", "Location", "Posted Time"])
        for row in data:
            writer.writerow(row)


def scrape_job_listings(driver, url, search_keyword, pages):
    job_data = []
    for page_num in range(1, pages + 1):
        driver.get(f"{url}?keywords={search_keyword}&pageNum={page_num}")
        WebDriverWait(driver, 10).until(EC.presence_of_element_located(
            (By.CLASS_NAME, "jobs-search__results-list")))
        job_listings = driver.find_elements(By.CLASS_NAME, "base-card")
        for listing in job_listings:
            try:
                title = listing.find_element(
                    By.CLASS_NAME, "base-search-card__title").get_attribute('innerText')
                company = listing.find_element(
                    By.CLASS_NAME, "base-search-card__subtitle").get_attribute('innerText')
                location = listing.find_element(
                    By.CLASS_NAME, "job-search-card__location").get_attribute('innerText')
                posted_time = listing.find_element(
                    By.CLASS_NAME, "job-search-card__listdate").get_attribute('innerText')
                job_data.append([title, company, location, posted_time])
            except:
                pass
    return job_data


driver = webdriver.Chrome()
PAGES = 1
search_keyword = "python developer".replace(" ", "%20")
url = "https://www.linkedin.com/jobs/search/"
filename = "jobs"

driver.get(url)

job_data = scrape_job_listings(driver, url, search_keyword, PAGES)
to_csv(job_data, filename)


driver.quit()
