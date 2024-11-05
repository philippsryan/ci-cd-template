export const getUsers = async () => {
  try {
    const response = await fetch(
      `http://${window.location.hostname}:8000/user`,
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

export const createUser = async (username: string) => {
  try {
    const response = await fetch(
      `http://${window.location.hostname}:8000/user`,
      {
        method: "POST",
        body: new URLSearchParams({
          username: username,
        }),
      },
    );

    return await response.json();
  } catch (e) {
    console.log(e);
    throw e;
  }
};
