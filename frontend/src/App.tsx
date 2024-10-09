import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import './App.css'
import UserSelectionScreen from './UserSelectionScreen'
import { useState } from 'react';
import TodoList from './TodoList';

const queryClient = new QueryClient();


function App() {
  const [user, setUser] = useState('');
  return (
    <QueryClientProvider client={queryClient}>
      {
        user === '' ?

          <UserSelectionScreen onUserSelected={(selectedUser) => setUser(selectedUser)} />
          : <TodoList user={user} />
      }
    </QueryClientProvider>
  )
}

export default App
