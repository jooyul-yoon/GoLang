import csv
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


def to_csv(data, filename):
    with open('results/'+filename.replace('+', '_').lower()+'.csv', 'w', newline='') as csvfile:
        writer = csv.writer(csvfile)
        writer.writerow(["Title", "Company", "Location", "Posted Time"])
        for row in data:
            writer.writerow(row)


def scrape_job_listings(driver, url, search_keyword, origin, pages):
    job_data = []
    for page_num in range(pages):
        start = 10 * page_num
        driver.get(
            f"{url}q={search_keyword}&l={origin}&radius=50&start={start}")
        job_listings = driver.find_elements(
            By.CLASS_NAME, "cardOutline")
        for listing in job_listings:
            try:
                title = listing.find_element(
                    By.CLASS_NAME, "jobTitle").find_element(By.TAG_NAME, "span").text
                company = listing.find_element(
                    By.XPATH, "//*[@data-testid='company-name']").text
                location = listing.find_element(
                    By.XPATH, "//*[@data-testid='text-location']").text
                time = listing.find_element(
                    By.CSS_SELECTOR, '[data-testid="myJobsStateDate"]').text
                job_data.append([title, company, location, time])
                print(f"{time}")
            except:
                pass
    return job_data


driver = webdriver.Chrome()
PAGES = 1
KEYWORD = "software engineer".replace(" ", "+")
LOC = "San Jose".replace(" ", "+")
URL = f"https://www.indeed.com/jobs?"

driver.get(URL)

job_data = scrape_job_listings(driver, URL, KEYWORD, LOC, PAGES)
to_csv(job_data, KEYWORD+"-"+LOC)


driver.quit()
