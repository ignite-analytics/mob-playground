import { test, expect } from '@playwright/test';

test('Sourcing has sourcing events in the list', async ({ page }) => {
  await page.goto('https://app.igniteprocurement.com/login/?referrer=/dashboard/');
  await page.click("role=button >> text=sign in", {})
  await page.fill('[aria-label="Email"]', "email")
  await page.fill('[aria-label="Password"]', "password")
  await page.screenshot({path: 'screenshot.png'}) 
  await page.click("role=button >> text=sign in", {})
  await page.screenshot({path: 'screenshot.png'}) 
  await page.click("text=Sourcing", {})
  await page.screenshot({path: 'screenshot2.png'}) 

  //console.log(register)

  // Expect a title "to contain" a substring.


  // create a locator
  const getStarted = page.locator('text=Sourcing');

  // Expect an attribute "to be strictly equal" to the value.
  await expect(getStarted).toHaveAttribute('href', '/docs/intro');

  // Click the get started link.
  await getStarted.click();

  // Expects the URL to contain intro.
  await expect(page).toHaveURL(/.*intro/);
});
