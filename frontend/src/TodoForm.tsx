import { useState } from "react";
import { Todo } from "./todoApi";



export interface TodoFormProps {
  todo: Todo
  onSubmit: (todo: Todo) => void;
}

const TodoForm = ({ todo, onSubmit }: TodoFormProps) => {
  const [title, setTitle] = useState(todo.title);
  const [body, setBody] = useState(todo.body);
  const { belongs_to, done } = todo;
  return (
    <div>
      <label> Title </label>
      <input data-testid="title-input" type="text" value={title} onChange={(e) => setTitle(e.target.value)} />
      <label> Body</label>
      <textarea data-testid='body-input' value={body} onChange={(e) => setBody(e.target.value)} />
      <button onClick={() => onSubmit({ title, body, belongs_to, done })}>Submit</button>
    </div>
  )
}

export default TodoForm;
