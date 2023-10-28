import React from 'react';
import { render, screen } from '@testing-library/react';
import Register from './Register';

// mock router because it is not available in isolation for this test
const mockUsedNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
   ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockUsedNavigate,
}));

test('renders content service register page', () => {
  render(<Register />);
  const element = screen.getAllByText(/Register/);
  expect(element[0]).toBeInTheDocument();
});