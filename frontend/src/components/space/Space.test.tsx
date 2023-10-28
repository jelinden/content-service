import React from 'react';
import { render, screen } from '@testing-library/react';
import Space from './Space';

test('renders content service space page', () => {
  render(<Space />);
  const element = screen.getAllByText(/Space/);
  expect(element[0]).toBeInTheDocument();
});