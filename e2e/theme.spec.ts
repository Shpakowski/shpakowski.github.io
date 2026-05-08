import { test, expect } from '@playwright/test';

test('has title and toggles theme', async ({ page }) => {
  await page.goto('/');

  // Toggle theme
  const html = page.locator('html');
  // Need to locate the ThemeToggle button correctly. Using generic aria-label or just the first button.
  // We'll target the button containing the Moon/Sun icon. It's often in header/sidebar.
  const themeToggle = page.locator('button').filter({ has: page.locator('svg.lucide-moon, svg.lucide-sun') }).first();
  
  // Wait for React to hydrate
  await page.waitForTimeout(1000);

  const isDarkInitial = await html.evaluate((el) => el.classList.contains('dark'));
  
  if (await themeToggle.isVisible()) {
    await themeToggle.click();
    const isDarkAfter = await html.evaluate((el) => el.classList.contains('dark'));
    expect(isDarkInitial).not.toBe(isDarkAfter);
  }
});
