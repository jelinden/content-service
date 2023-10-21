import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders content service', () => {
  render(<App />);
  const linkElement = screen.getAllByText(/Content service/i);
  expect(linkElement[0]).toBeInTheDocument();
});
