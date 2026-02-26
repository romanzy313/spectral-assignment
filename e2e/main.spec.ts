import { test, expect } from "@playwright/test";

test("page loads", async ({ page }) => {
  await page.goto("/");

  await expect(page).toHaveTitle(/Spectral Assignment/);
});

test("workflow", async ({ page }) => {
  await page.goto("/", { waitUntil: "load" });

  await expect(page.getByTestId("total-data-points")).toHaveText("0");

  await page.getByTestId("load-more-button").click();

  await expect(page.getByTestId("load-more-button")).toBeEnabled();

  await expect(page.getByTestId("total-data-points")).toHaveText("1000");

  await page.getByTestId("load-all-button").click();

  await expect(page.getByTestId("load-more-button")).toHaveText(/No more data/);

  await expect(page.getByTestId("load-more-button")).toBeDisabled();
  await expect(page.getByTestId("load-all-button")).toBeDisabled();
});

// test("get started link", async ({ page }) => {
//   await page.goto("https://playwright.dev/");

//   // Click the get started link.
//   await page.getByRole("link", { name: "Get started" }).click();

//   // Expects page to have a heading with the name of Installation.
//   await expect(
//     page.getByRole("heading", { name: "Installation" }),
//   ).toBeVisible();
// });
