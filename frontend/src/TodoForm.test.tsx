import { render, screen } from "@testing-library/react";
import TodoForm from './TodoForm';
import { describe, expect, it, vi } from "vitest";



describe("TodoForm", () => {
  it("can render a form with prefilled data", () => {
    const mockOnSubmit = vi.fn();
    const todo = {
      title: "hello",
      body: "world",
      done: true,
      belongs_to: "tester"
    }
    render(<TodoForm todo={todo} onSubmit={mockOnSubmit} />)

    const titleInput = screen.getByTestId("title-input");

    expect(titleInput).toBeInTheDocument();
  })
})
