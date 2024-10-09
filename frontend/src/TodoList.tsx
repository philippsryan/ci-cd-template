import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { createTodo, getTodos, Todo } from "./todoApi";
import { useMemo } from "react";
import TodoForm from "./TodoForm";



export interface TodoListProps {
  user: string
}

const TodoList = ({ user }: TodoListProps) => {
  const queryClient = useQueryClient();

  const { data: todos, status: fetchTodoStatus } = useQuery({
    queryKey: [`fetch-todos-${user}`],
    queryFn: () => getTodos(user)
  });

  const { mutate: makeTodo } = useMutation({
    mutationFn: (todo: Todo) => createTodo(todo),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [`fetch-todos-${user}`] })
    }
  })

  const todoItems = useMemo(() => {
    if (fetchTodoStatus === 'success') {
      console.log(fetchTodoStatus, todos);
      return todos.map((todo: Todo) => {
        return (
          <div key={todo.id}>
            <h4>{todo.title}</h4>
            <label>Done</label>
            <input type="checkbox" checked={todo.done} />
          </div>
        )
      });
    }
  }, [todos, fetchTodoStatus])

  return (
    <>
      <p>This is the todo list for {user}</p>
      <TodoForm todo={{ belongs_to: user, title: '', body: '', done: false }} onSubmit={(todo: Todo) => makeTodo(todo)} />
      {todoItems}
    </>
  )
}

export default TodoList;
