// @ts-check
const { test, expect } = require('@playwright/test');

test('can fetch users', async ({ request }) => {
  const response = await request.get("/user");

  await expect(response).toBeOK();
});

test('can create a user', async ({ request }) => {
  const response = await request.post("/user", {
    form: {
      username: 'unit test'
    }
  });

  await expect(response).toBeOK();
  expect(await response.text()).toBe("Added unit test");

})

test('throws a 400 when creating a todo for a user that does not exist', async ({ request }) => {
  const response = await request.post("/todos", {
    data: {
      belongs_to: "nobody",
      title: "This should fail",
      body: "This should fail"
    }
  });

  expect(response.status()).toBe(400);
});
