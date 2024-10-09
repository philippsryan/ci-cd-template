import { useMutation, useQuery } from "@tanstack/react-query";
import { createUser, getUsers } from "./userApi";
import { useMemo, useState } from "react";



interface UserSelectionScreenProps {
  onUserSelected: (user: string) => void;
}

const UserSelectionScreen = ({ onUserSelected }: UserSelectionScreenProps) => {
  const [newUser, setNewUser] = useState<string>("");
  const { data: users, status: fetchUserStatus } = useQuery({
    queryKey: ['fetch-users'],
    queryFn: () => getUsers()
  });

  const { mutate: makeUser } = useMutation({
    mutationKey: ["fetch-users"],
    mutationFn: (username: string) => createUser(username),
  });


  const userElement = useMemo(() => {
    if (fetchUserStatus === 'success') {
      return users.map((user: string) => {
        return <button key={user} onClick={() => onUserSelected(user)}> {user} </button>;
      })
    }
  }, [fetchUserStatus, users])

  return (
    <>
      <h1> Select your user </h1>
      {userElement}

      <h3>Create a new user </h3>
      <input type="text" value={newUser} onChange={(e) => setNewUser(e.target.value)} />
      <button onClick={() => makeUser(newUser)}> Create User </button>
    </>
  );
}

export default UserSelectionScreen;
