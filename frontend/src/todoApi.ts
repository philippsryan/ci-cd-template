export type Todo = {
  id?: number;
  belongs_to: string;
  title: string;
  body: string;
  done: boolean;
};

export const getTodos = async (user: string) => {
  try {
    const params = new URLSearchParams();
    params.append("user", user);
    const response = await fetch(
      `http://${window.location.hostname}:8000/todos?` + params,
      {
        method: "GET",
      },
    );

    return await response.json();
  } catch (e) {
    console.log(e);
    throw e;
  }
};

export const createTodo = async (todo: Todo) => {
  try {
    const response = await fetch(
      `http://${window.location.hostname}:8000/todos`,
      {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify(todo),
      },
    );

    return await response.json();
  } catch (e) {
    console.log(e);
    throw e;
  }
};
