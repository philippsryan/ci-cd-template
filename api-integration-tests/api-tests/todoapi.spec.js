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


test('it can create and get a todo', async ({ request }) => {
  await request.post("/user", {
    form: {
      username: 'tester'
    }
  });


  const create_response = await request.post('/todos',
    {
      data: {
        belongs_to: "tester",
        title: "UT title",
        body: "UT body",
        done: false
      }
    })

  expect(create_response.status()).toBe(200);
  expect(await create_response.body()).not.toBeNull();


  const get_response = await request.get(`/todos`, {
    params: {
      "user": "tester"
    }
  });

  expect(get_response).toBeOK();

  const get_body = await get_response.json();

  expect(get_body).toHaveLength(1);

  expect(get_body[0].title).toBe("UT title");
  expect(get_body[0].body).toBe("UT body");
  expect(get_body[0].done).toBeFalsy();

})
