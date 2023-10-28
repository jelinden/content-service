import React from 'react';
import { render, screen } from '@testing-library/react';
import Login from './Login';

// mock router because it is not available in isolation for this test
const mockUsedNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
   ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockUsedNavigate,
}));

test('renders content service login page', () => {
  render(<Login />);
  const element = screen.getAllByText(/Login/);
  expect(element[0]).toBeInTheDocument();
});